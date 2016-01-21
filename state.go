/* rigo/state.go */
package ri

import (
	"sync"
	"fmt"
	"io"
	"os"
)

var (
	ErrInvalidContextHandle error = fmt.Errorf("Invalid Context Handle")
	ErrContextAlreadyExists error = fmt.Errorf("Context Already Exists")
	ErrNoActiveContext      error = fmt.Errorf("No Active Context")
)

type context struct {
	Name string
	writer io.Writer
}

func (ctx *context) Write(out string) {
	ctx.writer.Write([]byte(out))
}

var internal struct {
	contexts map[string]*context	
	active *context
	sync.Mutex
}

func write(out string) error {
	internal.Lock()
	defer internal.Unlock()
	if internal.active == nil {
		return ErrNoActiveContext
	}
	internal.active.Write(out)
	return nil
}

func init() {

	internal.contexts = make(map[string]*context,0)
	internal.active = nil
}

/* Begin creates and initializes a new rendering context */
func Begin(name string) error {
	
	internal.Lock()
	defer internal.Unlock()

	if _,exists := internal.contexts[name]; exists {
		return ErrContextAlreadyExists
	}

	file,err := os.Create(name)
	if err != nil {
		return err
	}

	ctx := &context{Name:name,writer:file}
	ctx.writer.Write([]byte("## RIB /* rigo */\n"))
	
	internal.contexts[name] = ctx
	internal.active = ctx
	return nil
}

func BeginWriter(name string,writer io.Writer) error {

	internal.Lock()
	defer internal.Unlock()

	if _,exists := internal.contexts[name]; exists {
		return ErrContextAlreadyExists 
	}

	ctx := &context{Name:name,writer:writer}
	ctx.writer.Write([]byte("## RIB /* rigo */\n"))

	internal.contexts[name] = ctx
	internal.active = ctx
	return nil
}

/* End terminate the active rendering context */
func End() error {

	/* TODO: finish up stream */

	internal.Lock()
	defer internal.Unlock()

	if internal.active == nil {
		return ErrNoActiveContext
	}

	delete(internal.contexts,internal.active.Name)
	internal.active = nil
	return nil
}

/* GetContext returns a handle for the current active rendering context */
func GetContext() RtContextHandle {
	
	internal.Lock()
	defer internal.Unlock()
	
	return internal.active
}

/* Context sets the current active rendering context */
func Context(handle RtContextHandle) error {

	if handle == nil {
		return ErrInvalidContextHandle
	}

	ctx,ok := handle.(*context)
	if !ok {
		return ErrInvalidContextHandle
	}

	internal.Lock()
	defer internal.Unlock()

	internal.active = ctx

	return nil
}

/* FrameBegin mark the beginning of a single frame of an animated sequenece */
func FrameBegin(frame RtInt) error { 
	return write(fmt.Sprintf("FrameBegin %d\n",frame))
}

/* FrameEnd mark the end of a single frame of an animated sequence */
func FrameEnd() error {
	return write("FrameEnd\n")
}

/* When WorldBegin is invoked, all rendering options are frozen */
func WorldBegin() error {
	return write("WorldBegin\n")
}

func WorldEnd() error {
	return write("WorldEnd\n")
}

/* set the horizontal and vertical resolution (in pixels) of the image to be rendered */
func Format(xresolution,yresolution RtInt,pixelaspectratio RtFloat) error {
	return write(fmt.Sprintf("Format %d %d %f\n",xresolution,yresolution,pixelaspectratio))
}


