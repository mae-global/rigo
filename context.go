/* rigo/context.go */
package rigo

import (
	"sync"

	. "github.com/mae-global/rigo/ri"
	. "github.com/mae-global/rigo/ris"
)

type Configuration struct {
	Entity             bool
	PrettyPrint        bool
	PrettyPrintSpacing string
}

type Context struct {
	mux  sync.RWMutex
	pipe *Pipe
	/* TODO: replace individual handlers with a manager */
	objects ObjectHandler
	lights  LightHandler
	shaders ShaderHandler

	files map[RtToken]bool

	cache map[RtShaderHandle]Shader

	Info
}

func (ctx *Context) Write(name RtName, args, params, values []Rter) error {
	ctx.mux.Lock()
	defer ctx.mux.Unlock()
	/* FIXME: add pipe check here */
	return ctx.pipe.Run(name, args, Mix(params, values), ctx.Info)
}

func (ctx *Context) OpenRaw(id RtToken) (ArchiveWriter, error) {
	ctx.mux.Lock()
	defer ctx.mux.Unlock()

	if ctx.files == nil {
		ctx.files = make(map[RtToken]bool, 0)
	}

	if _, exists := ctx.files[id]; exists {
		return nil, ErrNotSupported
	}

	for _, r := range ctx.files {
		if r {
			return nil, ErrNotSupported
		}
	}

	ctx.files[id] = true

	return ctx.pipe.ToRaw(), nil
}

func (ctx *Context) CloseRaw(id RtToken) error {
	ctx.mux.Lock()
	defer ctx.mux.Unlock()

	if ctx.files == nil {
		return ErrNotSupported
	}

	if _, exists := ctx.files[id]; !exists {
		return ErrNotSupported
	}

	delete(ctx.files, id)

	return nil
}

func (ctx *Context) ShaderHandle() (RtShaderHandle, error) {
	ctx.mux.Lock()
	defer ctx.mux.Unlock()
	return ctx.shaders.Generate()
}

func (ctx *Context) CheckShaderHandle(sh RtShaderHandle) error {
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

func (ctx *Context) SetShader(sh RtShaderHandle, s Shader) {
	ctx.mux.Lock()
	defer ctx.mux.Unlock()
	ctx.cache[sh] = s
}

func (ctx *Context) GetShader(sh RtShaderHandle) Shader {
	ctx.mux.Lock()
	defer ctx.mux.Unlock()
	if s, ok := ctx.cache[sh]; ok {
		return s
	}
	return nil
}

func (ctx *Context) Shader(sh RtShaderHandle) ShaderWriter {
	return ctx.GetShader(sh)
}

func NewContext(pipe *Pipe, lights LightHandler, objects ObjectHandler, shaders ShaderHandler, config *Configuration) *Context {
	if pipe == nil {
		pipe = DefaultFilePipe()
	}
	if config == nil {
		config = &Configuration{Entity: false, PrettyPrint: false, PrettyPrintSpacing: "\t"}
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

	info := Info{Name: "", Entity: config.Entity, PrettyPrint: config.PrettyPrint, PrettyPrintSpacing: config.PrettyPrintSpacing}

	ctx := &Context{pipe: pipe, lights: lights, objects: objects, shaders: shaders, Info: info}
	ctx.cache = make(map[RtShaderHandle]Shader, 0)
	return ctx
}

func RIS(ctx RisContexter) *Ris {
	return &Ris{ctx}
}

func RI(ctx RiContexter) *Ri {
	return &Ri{ctx}
}
