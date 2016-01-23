/* rigo/context.go */
package ri

import (

	"fmt"
	"os"
)

var (
	ErrInvalidContextHandle  = fmt.Errorf("Invalid Context Handle")
	ErrContextAlreadyExists  = fmt.Errorf("Context Already Exists")
	ErrNoActiveContext       = fmt.Errorf("No Active Context")
)

type Piper interface {
	Write(RtName,[]Rter,Info) error
}

type Info struct {
	Name string
	Depth int
	Lights uint
	Objects uint
	Entity bool
}

type DefaultFilePipe struct {
	file *os.File
}

func (p *DefaultFilePipe) Write(name RtName,list []Rter,info Info) error {
	if name == "Begin" {
		if p.file != nil {
			return ErrProtocolBotch
		}
		file := "out.rib"
		if len(list) > 0 {
			if t,ok := list[0].(RtString); ok {
				file = string(t)
			}
		}

		f,err := os.Create(file)
		if err != nil {
			return err
		}
		p.file = f

		postfix := "\n"
		if info.Entity {
			postfix = " Entity\n"
		}
		_,err = p.file.Write([]byte("##RenderMan RIB-Structure 1.1" + postfix))
		return err
	}

	if name == "End" {
		if p.file == nil {
			return ErrProtocolBotch
		}
		return p.file.Close()
	}

	if p.file == nil {
		return ErrNoActiveContext
	}

	if name != "##" { 

		prefix := "" 
		for i := 0; i < info.Depth; i++ {
			prefix += "\t"
		}

		_,err := p.file.Write([]byte(prefix + name.Serialise() + " " + Serialise(list) + "\n"))
		return err
	}
	
	/* FIXME: structural needs work */

	_,err := p.file.Write([]byte("##" + Serialise(list) + "\n"))
	return err
}

type Context struct {
	pipe Piper
	name string
	
	entity bool
	depth int 

	lights uint
	objects uint
}

func (ctx *Context) info() Info {
	return Info{ctx.name,ctx.depth,ctx.lights,ctx.objects,ctx.entity}
}

func (ctx *Context) writef(name RtName,parameterlist ...Rter) error {
	return ctx.pipe.Write(name,parameterlist,ctx.info())
}

func New(writer Piper) *Context {
	if writer == nil {
		writer = &DefaultFilePipe{}
	}

	return &Context{name:"",pipe:writer}
}

func NewEntity(writer Piper) *Context {
	if writer == nil {
		writer = &DefaultFilePipe{}
	}
	return &Context{name:"",pipe:writer,entity:true}
}



