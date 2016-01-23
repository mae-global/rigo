/* rigo/context.go */
package ri

import (

	"fmt"
	"io"
	"os"
)

var (
	ErrInvalidContextHandle  = fmt.Errorf("Invalid Context Handle")
	ErrContextAlreadyExists  = fmt.Errorf("Context Already Exists")
	ErrNoActiveContext       = fmt.Errorf("No Active Context")
)

type Writer interface {
	Write(int,string) error
}


type Filer interface {
	Writer
	io.Closer
}

type File struct {
	file *os.File
}

func (f *File) Write(depth int,out string) error {
	if f.file == nil {
		return ErrNoActiveContext
	}

	content := "" /* TODO: replace with compact version */
	for i := 0; i < depth; i++ {
		content += "\t"
	}

	_,err := f.file.Write([]byte(content + out))
	return err
}

func (f *File) Close() error {
	return f.file.Close()
}

type Filterer interface {
	Filter(name RtName,parameterlist ...Rter) error
}



type Context struct {
	name string
	writer Filer
	entity bool
	depth int /* for pretty printing */

	lights uint
	objects uint

	filters map[RtName]Filterer
}

func (ctx *Context) IsEntity() bool {
	return ctx.entity
}

func (ctx *Context) filter(name RtName,parameterlist ...Rter) error {
	if f,exists := ctx.filters[name]; exists {
			return f.Filter(name,parameterlist...)
	}
	return nil
}

func (ctx *Context) write(parameterlist ...Rter) error {
	/* TODO: add general filter here */
	if ctx.writer == nil {
		return ErrNoActiveContext
	}

	return ctx.writer.Write(ctx.depth,fmt.Sprintf("%s\n",serialiseToString(parameterlist...)))
}

func (ctx *Context) writef(name RtName,parameterlist ...Rter) error {
	if f,exists := ctx.filters[name]; exists {
		if err := f.Filter(name,parameterlist...); err != nil {
			return err
		}
	}
	if ctx.writer == nil {
		return ErrNoActiveContext
	}
	
	
	if len(parameterlist) == 0 {
		return ctx.writer.Write(ctx.depth,fmt.Sprintf("%s\n",name))
	}

	return ctx.writer.Write(ctx.depth,fmt.Sprintf("%s %s\n",name,serialiseToString(parameterlist...)))
}

func (ctx *Context) Filter(name RtName,filter Filterer) {
	ctx.filters[name] = filter
}


func New(writer Filer) *Context {
	return &Context{name:"",writer:writer}
}

func NewEntity(writer Filer) *Context {
	return &Context{name:"",writer:writer,entity:true}
}

