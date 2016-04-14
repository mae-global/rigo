/* rigo/context.go */
package rigo

import (
	"fmt"
	"sync"

	. "github.com/mae-global/rigo/ri"
)
const (
	Version = "0"
	Author  = "rigo"
)

var (
	ErrInvalidContextHandle = fmt.Errorf("Invalid Context Handle")
	ErrContextAlreadyExists = fmt.Errorf("Context Already Exists")
	ErrNoActiveContext      = fmt.Errorf("No Active Context")
	ErrNotImplemented       = fmt.Errorf("Not Implemented")
	ErrPipeDone             = fmt.Errorf("Pipe Done")
	ErrEndOfLine            = fmt.Errorf("End of Line")
)

type Result struct {
	Name RtName
	Args []Rter
	Info *Info
	Err  error
}

func Done() *Result {
	return &Result{"", nil, nil, ErrPipeDone}
}

func Next(name RtName, args []Rter, info Info) *Result {
	return &Result{name, args, info.Copy(), nil}
}

func InError(err error) *Result {
	return &Result{"", nil, nil, err}
}

func Errored(message RtString) *Result {
	return &Result{"", nil, nil, fmt.Errorf(string(message))}
}

func EndOfLine() *Result {
	return &Result{"", nil, nil, ErrEndOfLine}
}

type Pipe struct {
	blocks []Piper
	sync.Mutex
}

func (p *Pipe) Last() Piper {
	p.Lock()
	defer p.Unlock()
	if len(p.blocks) == 0 {
		return nil
	}
	return p.blocks[len(p.blocks)-1]
}

func (p *Pipe) Append(block Piper) *Pipe {
	if block == nil {
		return p
	}
	p.Lock()
	defer p.Unlock()
	p.blocks = append(p.blocks, block)
	return p
}

/* Len get the length of the pipe */
func (p *Pipe) Len() int {
	p.Lock()
	defer p.Unlock()
	return len(p.blocks)
}

/* Get get a Piper object via index */
func (p *Pipe) Get(idx int) Piper {
	p.Lock()
	defer p.Unlock()
	if idx < 0 || idx >= len(p.blocks) {
		return nil
	}
	return p.blocks[idx]
}

/* GetByName get the first Piper object by name */
func (p *Pipe) GetByName(name string) Piper {
	p.Lock()
	defer p.Unlock()
	for _, b := range p.blocks {
		if b.Name() == name {
			return b
		}
	}
	return nil
}

func (p *Pipe) Run(name RtName, list []Rter, info Info) error {
	p.Lock()
	defer p.Unlock()

	if len(p.blocks) == 0 {
		return nil
	}

	nblocks := make([]Piper, 0)

	for _, b := range p.blocks {
		if b == nil {
			continue
		}

		r := b.Write(name, list, info)
		if r.Err != nil {
			if r.Err == ErrPipeDone {
				nblocks = append(nblocks, b)
				continue
			}

			if r.Err == ErrEndOfLine {
				/* then mark b ready to be removed */
				continue
			}

			return r.Err
		}

		nblocks = append(nblocks, b)
		/* TODO: take the output of last block */
	}

	p.blocks = nblocks
	return nil
}

func NewPipe() *Pipe {
	pipe := Pipe{}
	pipe.blocks = make([]Piper, 0)
	return &pipe
}

type Piper interface {
	Write(RtName, []Rter, Info) *Result
	Name() string
}

type Info struct {
	Name        string
	Depth       int
	Lights      uint
	Objects     uint
	Entity      bool
	Formal      bool
	PrettyPrint bool
}

func (info Info) Copy() *Info {
	n := Info{}
	n.Name = info.Name
	n.Depth = info.Depth
	n.Lights = info.Lights
	n.Objects = info.Objects
	n.Entity = info.Entity
	n.Formal = info.Formal
	n.PrettyPrint = info.PrettyPrint
	return &n
}

type Configuration struct {
	Entity bool
	Formal bool
	PrettyPrint bool
}

type Context struct {
	mux sync.RWMutex
	pipe *Pipe

	Info
}

func (ctx *Context) Write(name RtName, list []Rter) error {
	ctx.mux.Lock()
	defer ctx.mux.Unlock()
	if ctx.Formal {
		name = name.Prefix("Ri")
	}
	return ctx.pipe.Run(name, list, ctx.Info)
}

func (ctx *Context) Depth(d int) {
	ctx.mux.Lock()
	defer ctx.mux.Unlock()
	ctx.Info.Depth += d /* TODO: this is a little clunky */
}

func (ctx *Context) LightHandle() (RtLightHandle, error) {
	ctx.mux.Lock()
	defer ctx.mux.Unlock()
	lh := RtLightHandle(ctx.Lights)
	ctx.Lights++
	return lh, nil
}

func (ctx *Context) CheckLightHandle(lh RtLightHandle) error {
	ctx.mux.RLock()
	defer ctx.mux.RUnlock()
	if uint(lh) >= ctx.Lights {
		return ErrBadHandle
	}
	return nil
}

func (ctx *Context) ObjectHandle() (RtObjectHandle, error) {
	ctx.mux.Lock()
	defer ctx.mux.Unlock()
	oh := RtObjectHandle(ctx.Objects)
	ctx.Objects++
	return oh, nil
}

func (ctx *Context) CheckObjectHandle(oh RtObjectHandle) error {
	ctx.mux.RLock()
	defer ctx.mux.RUnlock()
	if uint(oh) >= ctx.Objects {
		return ErrBadHandle
	}
	return nil
}


func New(pipe *Pipe, config *Configuration) *Ri {
	if pipe == nil {
		pipe = DefaultFilePipe()
	}
	if config == nil {
		config = &Configuration{Entity: false, Formal: false,PrettyPrint: false}
	}
	ctx := &Context{pipe:pipe,Info:Info{Name: "", Entity: config.Entity, Formal: config.Formal,PrettyPrint: config.PrettyPrint}}
	return &Ri{ctx}
}

func NewEntity(pipe *Pipe) *Ri {
	return New(pipe, &Configuration{Entity: true, Formal: false})
}
