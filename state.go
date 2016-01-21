/* rigo/state.go */
package ri

import (

	"fmt"
	"io"
	"os"
)

var (
	ErrInvalidContextHandle error = fmt.Errorf("Invalid Context Handle")
	ErrContextAlreadyExists error = fmt.Errorf("Context Already Exists")
	ErrNoActiveContext      error = fmt.Errorf("No Active Context")
)

type Filer interface {
	io.Writer
	io.Closer
}

type Context struct {
	name string
	writer Filer
}

func (ctx *Context) write(out string) error {
	if ctx.writer == nil {
		return ErrNoActiveContext
	}
	ctx.writer.Write([]byte(out))
	return nil
}


func New(writer Filer) *Context {
	return &Context{name:"",writer:writer}
}

func (ctx *Context) Begin(name string) error {
	if ctx.writer == nil {
		if f,err := os.Create(name); err != nil {
			return err
		} else {
			ctx.writer = f
		}
	}
	return nil
}

func (ctx *Context) End() error {
	if ctx.writer == nil {
		return ErrNoActiveContext
	}
	if err := ctx.writer.Close(); err != nil {
		return err
	}
	ctx.writer = nil
	return nil
}	

/* FrameBegin mark the beginning of a single frame of an animated sequenece */
func (ctx *Context) FrameBegin(frame RtInt) error { 
	return ctx.write(fmt.Sprintf("FrameBegin %d\n",frame))
}

/* FrameEnd mark the end of a single frame of an animated sequence */
func (ctx *Context) FrameEnd() error {
	return ctx.write("FrameEnd\n")
}

/* When WorldBegin is invoked, all rendering options are frozen */
func (ctx *Context) WorldBegin() error {
	return ctx.write("WorldBegin\n")
}

func (ctx *Context) WorldEnd() error {
	return ctx.write("WorldEnd\n")
}

/* set the horizontal and vertical resolution (in pixels) of the image to be rendered */
func (ctx *Context) Format(xresolution,yresolution RtInt,pixelaspectratio RtFloat) error {
	return ctx.write(fmt.Sprintf("Format %d %d %f\n",xresolution,yresolution,pixelaspectratio))
}


