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
	ErrParameterlistMismatch error = fmt.Errorf("Parameterlist Mismatch")
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
	depth int /* for pretty printing */

	lights uint

	filters map[RtName]Filterer
}

func (ctx *Context) filter(name RtName,parameterlist ...Rter) error {
	if f,exists := ctx.filters[name]; exists {
			return f.Filter(name,parameterlist...)
	}
	return nil
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
	return ctx.writer.Write(ctx.depth,fmt.Sprintf("%s %s\n",name,serialise_s(parameterlist...)))
}

func (ctx *Context) Filter(name RtName,filter Filterer) {
	ctx.filters[name] = filter
}


func New(writer Filer) *Context {
	return &Context{name:"",writer:writer}
}

func (ctx *Context) Begin(name string) error {
	if ctx.writer == nil {
		if f,err := os.Create(name); err != nil {
			return err
		} else {
			ctx.writer = &File{f}
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
	defer func() { ctx.depth++ }()
	return ctx.writef("FrameBegin",frame)
}

/* FrameEnd mark the end of a single frame of an animated sequence */
func (ctx *Context) FrameEnd() error {
	ctx.depth-- 
	return ctx.writef("FrameEnd")
}

/* When WorldBegin is invoked, all rendering options are frozen */
func (ctx *Context) WorldBegin() error {
	defer func() { ctx.depth++ }()
	return ctx.writef("WorldBegin")
}

func (ctx *Context) WorldEnd() error {
	ctx.depth--	
	return ctx.writef("WorldEnd")
}

func (ctx *Context) Comment(comment RtName) error {
	return ctx.writef("#",comment)
}

/* set the horizontal and vertical resolution (in pixels) of the image to be rendered */
func (ctx *Context) Format(xresolution,yresolution RtInt,pixelaspectratio RtFloat) error {
	return ctx.writef("Format",xresolution,yresolution,pixelaspectratio)
}

/* frameaspectratio is the ration of the width to the height of the desired image. */
func (ctx *Context) FrameAspectRatio(frameaspectratio RtFloat) error {
	return ctx.writef("FrameAspectRatio",frameaspectratio)
}

/* ScreenWindow; this procedure defines a rectangle in the image plane. */
func (ctx *Context) ScreenWindow(left,right,bottom,top RtFloat) error {
	return ctx.writef("ScreenWindow",left,right,bottom,top)
}

/* CropWindow; render only a subrectangle of the image. */
func (ctx *Context) CropWindow(xmin,xmax,ymin,ymax RtFloat) error {
	return ctx.writef("CropWindow",RtFloatArray{xmin,xmax,ymin,ymax})
}

/* Projection; the project determines how camera coordinates are converted to screen coordinates */
func (ctx *Context) Projection(token RtToken, parameterlist ...Rter) error {
	outs := make([]Rter,0)
	outs = append(outs,token)
	outs = append(outs,parameterlist...)
	return ctx.writef("Projection",outs...)
}

/* Clipping; sets the position of the near and far clipping planes along the direction of view. */
func (ctx *Context) Clipping(near,far RtFloat) error {
	return ctx.writef("Clipping",near,far)
}

/* ClippingPlane; adds a user-specified clipping plane. */
func (ctx *Context) ClippingPlane(x,y,z,nx,ny,nz RtFloat) error {
	return ctx.writef("ClippingPlane",x,y,z,nx,ny,nz)
}

/* DepthOfField; focaldistance sets the distance along the direction of view at which objects will be in focus. */
func (ctx *Context) DepthOfField(fstop,focallength,focaldistance RtFloat) error {
	return ctx.writef("DepthOfField",fstop,focallength,focaldistance)
}

/* Shutter; sets the times at which the shutter opens and closes. */
func (ctx *Context) Shutter(min,max RtFloat) error {
	return ctx.writef("Shutter",min,max)
}




