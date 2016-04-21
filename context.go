/* rigo/context.go */
package rigo

import (
	"fmt"
	"sync"

	. "github.com/mae-global/rigo/ri"
	. "github.com/mae-global/rigo/ri/handles"
	. "github.com/mae-global/rigo/ris"
)

const (
	Author RtToken = "RiGO;ver 0"
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
	List []Rter
	Info *Info
	Err  error
}

func Done() *Result {
	return &Result{"", nil, nil,nil, ErrPipeDone}
}

func Next(name RtName, args,list []Rter, info Info) *Result {
	return &Result{name, args,list, info.Copy(), nil}
}

func InError(err error) *Result {
	return &Result{"", nil, nil,nil, err}
}

func Errored(message RtString) *Result {
	return &Result{"", nil, nil,nil, fmt.Errorf(string(message))}
}

func EndOfLine() *Result {
	return &Result{"", nil, nil,nil, ErrEndOfLine}
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

func (p *Pipe) Run(name RtName, args,list []Rter, info Info) error {
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

		r := b.Write(name, args,list, info)
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
		if r.Args != nil {
			args = r.Args
		}
		if r.List != nil {
			list = r.List
		}
	}

	p.blocks = nblocks
	return nil
}

func (p *Pipe) ToRaw() ArchiveWriter {
	p.Lock()
	defer p.Unlock()

	if len(p.blocks) == 0 {
		return nil
	}

	for _,b := range p.blocks {
		if b == nil {
			continue
		}
		if aw := b.ToRaw(); aw != nil {
			return aw
		}
	}
	return nil
}
	

func NewPipe() *Pipe {
	pipe := Pipe{}
	pipe.blocks = make([]Piper, 0)
	return &pipe
}

type Piper interface {
	Write(RtName, []Rter,[]Rter, Info) *Result
	Name() string
	ToRaw() ArchiveWriter
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
	objects ObjectHandler
	lights LightHandler
	shaders ShaderHandler

	files map[RtToken]bool

	cache map[RtShaderHandle] Shader

	Info
}

func (ctx *Context) Write(name RtName, args,list []Rter) error {
	ctx.mux.Lock()
	defer ctx.mux.Unlock()
	if ctx.Formal {
		name = name.Prefix("Ri")
	}
	return ctx.pipe.Run(name, args,list, ctx.Info)
}

func (ctx *Context) OpenRaw(id RtToken) (ArchiveWriter,error) {
	ctx.mux.Lock()
	defer ctx.mux.Unlock()

	if ctx.files == nil {
		ctx.files = make(map[RtToken]bool,0)
	}

	if _,exists := ctx.files[id]; exists {
		return nil,ErrNotSupported
	}

	for _,r := range ctx.files {
		if r {
			return nil,ErrNotSupported
		}
	}

	ctx.files[id] = true

	return ctx.pipe.ToRaw(),nil
}

func (ctx *Context) CloseRaw(id RtToken) error {
	ctx.mux.Lock()
	defer ctx.mux.Unlock()
	
	if ctx.files == nil {
		return ErrNotSupported
	}


	if _,exists := ctx.files[id]; !exists {
		return ErrNotSupported
	}

	delete(ctx.files,id)

	return nil
}

func (ctx *Context) Depth(d int) {
	ctx.mux.Lock()
	defer ctx.mux.Unlock()
	ctx.Info.Depth += d /* TODO: this is a little clunky */
}

func (ctx *Context) ShaderHandle() (RtShaderHandle,error) {
	ctx.mux.Lock()
	defer ctx.mux.Unlock()
	return ctx.shaders.Generate()
}

func (ctx *Context) CheckShaderHandler(sh RtShaderHandle) error {
	ctx.mux.Lock()
	defer ctx.mux.Unlock()
	return ctx.shaders.Check(sh)
}

func (ctx *Context) LightHandle() (RtLightHandle, error) {
	ctx.mux.Lock()
	defer ctx.mux.Unlock()
	return ctx.lights.Generate()
}

func (ctx *Context) CheckLightHandle(lh RtLightHandle) error {
	ctx.mux.RLock()
	defer ctx.mux.RUnlock()
	return ctx.lights.Check(lh)
}

func (ctx *Context) ObjectHandle() (RtObjectHandle, error) {
	ctx.mux.Lock()
	defer ctx.mux.Unlock()
	return ctx.objects.Generate()
}

func (ctx *Context) CheckObjectHandle(oh RtObjectHandle) error {
	ctx.mux.RLock()
	defer ctx.mux.RUnlock()
	return ctx.objects.Check(oh)
}
	

func (ctx *Context) SetShader(sh RtShaderHandle,s Shader) {
	ctx.mux.Lock() 
	defer ctx.mux.Unlock()
	ctx.cache[sh] = s
}

func (ctx *Context) GetShader(sh RtShaderHandle) Shader {
	ctx.mux.Lock()
	defer ctx.mux.Unlock()
	if s,ok := ctx.cache[sh]; ok {
		return s
	}
	return nil
}

func New(pipe *Pipe, config *Configuration) *Ri {
	return NewCustom(pipe,nil,nil,nil,config)
}

func NewCustom(pipe *Pipe,lights LightHandler,objects ObjectHandler,shaders ShaderHandler,config *Configuration) *Ri {
	if pipe == nil {
		pipe = DefaultFilePipe()
	}
	if config == nil {
		config = &Configuration{Entity: false,Formal: false,PrettyPrint: false}
	}
	
	if lights == nil {
		lights = NewLightNumberGenerator()
	}
	if objects == nil {
		objects = NewObjectNumberGenerator()
	}
	if shaders == nil {
		shaders = NewShaderNumberGenerator()
	}	

	ctx := &Context{pipe:pipe,lights:lights,objects:objects,shaders:shaders,Info:Info{Name: "", Entity: config.Entity, Formal: config.Formal,PrettyPrint: config.PrettyPrint}}

	/* cache for the shaders */
	ctx.cache = make(map[RtShaderHandle]Shader,0)
	return &Ri{ctx}
}

func NewEntity(pipe *Pipe) *Ri {
	return New(pipe, &Configuration{Entity: true, Formal: false})
}

func RIS(ctx *Context) *Ris {
	return &Ris{ctx}
}



