package rigo

import (
	"sync"

	. "github.com/mae-global/rigo/ri"
	. "github.com/mae-global/rigo/ris"
)

type Configuration struct {
	Entity	bool
	PrettyPrint bool
	PrettyPrintSpacing string
}


type Context struct {
	mux sync.RWMutex
	pipe *Pipe

	HandleManagerer

	files map[RtToken]bool
	cache map[RtShaderHandle]Shader

	Info
}

/* Write */
func (ctx *Context) Write(name RtName,args,params,values []Rter) error {
	ctx.mux.Lock()
	defer ctx.mux.Unlock()
	return ctx.pipe.Run(name,args,Mix(params,values),ctx.Info)
}

/* OpenRaw */
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

/* CloseRaw */
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

/* SetShader */
func (ctx *Context) SetShader(sh RtShaderHandle,s Shader) {
	ctx.mux.Lock()
	defer ctx.mux.Unlock()
	ctx.cache[sh] = s
}

/* GetShader */
func (ctx *Context) GetShader(sh RtShaderHandle) Shader {
	ctx.mux.Lock()
	defer ctx.mux.Unlock()
	if s,ok := ctx.cache[sh]; ok {
		return s
	}
	return nil
}

/* Shader */
func (ctx *Context) Shader(sh RtShaderHandle) ShaderWriter {
	return ctx.GetShader(sh)
}


func NewContext(pipe *Pipe,mgr HandleManagerer,config *Configuration) *Context {

	if pipe == nil {
		pipe = DefaultFilePipe()
	}

	if config == nil {
		config = &Configuration{Entity:false,PrettyPrint:false,PrettyPrintSpacing:"\t"}
	} else {
		if config.PrettyPrint && len(config.PrettyPrintSpacing) == 0 {
			config.PrettyPrintSpacing = "\t"
		}
	}

	if mgr == nil {
		mgr = NewHandleManager(nil,nil,nil)
	}

	info := Info{Name:"",Entity:config.Entity,PrettyPrint:config.PrettyPrint,
							 PrettyPrintSpacing:config.PrettyPrintSpacing}

	ctx := &Context{pipe:pipe,Info:info}
	ctx.HandleManagerer = mgr
	ctx.cache = make(map[RtShaderHandle]Shader,0)
	return ctx
}

func RIS(ctx RisContexter) *Ris {
	return &Ris{ctx}
}

func RI(ctx RiContexter) *Ri {
	return &Ri{ctx}
}





