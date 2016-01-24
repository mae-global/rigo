/* rigo/context.go */
package ri

import (
	"fmt"
	"sync"
)

var (
	ErrInvalidContextHandle = fmt.Errorf("Invalid Context Handle")
	ErrContextAlreadyExists = fmt.Errorf("Context Already Exists")
	ErrNoActiveContext      = fmt.Errorf("No Active Context")
	ErrNotImplemented				= fmt.Errorf("Not Implemented")
	ErrPipeDone							= fmt.Errorf("Pipe Done")	
	ErrEndOfLine					  = fmt.Errorf("End of Line")
)

type Result struct {
	Name RtName
	Args []Rter
	Info *Info	
	Err error
}

func Done() *Result {
	return &Result{"",nil,nil,ErrPipeDone}
}

func Next(name RtName,args []Rter,infom Info) *Result {
	info := &Info{Name:infom.Name,Depth:infom.Depth,Lights:infom.Lights,Objects:infom.Objects,Entity:infom.Entity}
	return &Result{name,args,info,nil}
}

func InError(err error) *Result {
	return &Result{"",nil,nil,err}
}

func Errored(message RtString) *Result {
	return &Result{"",nil,nil,fmt.Errorf(string(message))}
}

func EndOfLine() *Result {
	return &Result{"",nil,nil,ErrEndOfLine}
}


type Pipe struct {
	blocks []Piper
	sync.Mutex
}

func (p *Pipe) Append(block Piper) *Pipe {
	if block == nil {
		return p
	}	
	p.Lock()
	defer p.Unlock()
	p.blocks = append(p.blocks,block)
	return p
}

func (p *Pipe) Run(name RtName,list []Rter,info Info) error {
	p.Lock()
	defer p.Unlock()

	if len(p.blocks) == 0 {
		return nil
	}

	nblocks := make([]Piper,0)
	
	for _,b := range p.blocks {
		if b == nil {
			continue
		}

		r := b.Write(name,list,info)
		if r.Err != nil {
			if r.Err == ErrPipeDone {
				nblocks = append(nblocks,b)
				continue
			}

			if r.Err == ErrEndOfLine {
				/* then mark b ready to be removed */
				continue
			}
			
			return r.Err
		}

		nblocks = append(nblocks,b)
		/* TODO: take the output of last block */
	}

	p.blocks = nblocks
	return nil
}


func NewPipe() *Pipe {
	pipe := Pipe{}
	pipe.blocks = make([]Piper,0)
	return &pipe
}




type Piper interface {
	Write(RtName, []Rter, Info) *Result
}

type Info struct {
	Name    string
	Depth   int
	Lights  uint
	Objects uint
	Entity  bool
}



type Context struct {
	pipe *Pipe
	name string

	entity bool
	depth  int

	lights  uint
	objects uint
}

func (ctx *Context) info() Info {
	return Info{ctx.name, ctx.depth, ctx.lights, ctx.objects, ctx.entity}
}

func (ctx *Context) writef(name RtName, parameterlist ...Rter) error {
	return ctx.pipe.Run(name,parameterlist,ctx.info()) //ctx.pipe.Write(name, parameterlist, ctx.info())
}

func New(pipe *Pipe) *Context {
	if pipe == nil {
		pipe = DefaultFilePipe()
	}
	return &Context{name: "", pipe: pipe}
}

func NewEntity(pipe *Pipe) *Context {
	if pipe == nil {
		pipe = DefaultFilePipe()
	}
	return &Context{name: "", pipe: pipe, entity: true}
}
