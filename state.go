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

type Context struct {
	name string
	writer Filer
	depth int /* for pretty printing */
}

func (ctx *Context) write(out string) error {
	if ctx.writer == nil {
		return ErrNoActiveContext
	}
	return ctx.writer.Write(ctx.depth,out)
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
	return ctx.write(fmt.Sprintf("FrameBegin %d\n",frame))
}

/* FrameEnd mark the end of a single frame of an animated sequence */
func (ctx *Context) FrameEnd() error {
	ctx.depth-- 
	return ctx.write("FrameEnd\n")
}

/* When WorldBegin is invoked, all rendering options are frozen */
func (ctx *Context) WorldBegin() error {
	defer func() { ctx.depth++ }()
	return ctx.write("WorldBegin\n")
}

func (ctx *Context) WorldEnd() error {
	ctx.depth--
	return ctx.write("WorldEnd\n")
}

func (ctx *Context) Comment(comment string) error {
	return ctx.write("#" + comment + "\n")
}

/* set the horizontal and vertical resolution (in pixels) of the image to be rendered */
func (ctx *Context) Format(xresolution,yresolution RtInt,pixelaspectratio RtFloat) error {
	return ctx.write(fmt.Sprintf("Format %d %d %f\n",xresolution,yresolution,pixelaspectratio))
}

/* frameaspectratio is the ration of the width to the height of the desired image. */
func (ctx *Context) FrameAspectRatio(frameaspectratio RtFloat) error {
	return ctx.write(fmt.Sprintf("FrameAspectRatio %f\n",frameaspectratio))
}

/* ScreenWindow; this procedure defines a rectangle in the image plane. */
func (ctx *Context) ScreenWindow(left,right,bottom,top RtFloat) error {
	return ctx.write(fmt.Sprintf("ScreenWindow [%f %f %f %f]\n",left,right,bottom,top))
}

/* CropWindow; render only a subrectangle of the image. */
func (ctx *Context) CropWindow(xmin,xmax,ymin,ymax RtFloat) error {
	return ctx.write(fmt.Sprintf("CropWindow [%f %f %f %f]\n",xmin,xmax,ymin,ymax))
}

/* Projection; the project determines how camera coordinates are converted to screen coordinates */
func (ctx *Context) Projection(token RtToken, parameterlist ...interface{}) error {

	out := ""
	if len(parameterlist) > 0 {
		if len(parameterlist) % 2 != 0 {
			return ErrParameterlistMismatch
		}

		for _,p := range parameterlist {
			
				out += fmt.Sprintf("%s ",p)
			
			
		}
	}

	return ctx.write(fmt.Sprintf("Projection \"%s\" %s\n",token,out))
}

/* Clipping; sets the position of the near and far clipping planes along the direction of view. */
func (ctx *Context) Clipping(near,far RtFloat) error {
	return ctx.write(fmt.Sprintf("Clipping %s %s\n",near,far))
}

/* ClippingPlane; adds a user-specified clipping plane. */
func (ctx *Context) ClippingPlane(x,y,z,nx,ny,nz RtFloat) error {
	return ctx.write(fmt.Sprintf("ClippingPlane %s %s %s %s %s %s\n",x,y,z,nx,ny,nz))
}

/* DepthOfField; focaldistance sets the distance along the direction of view at which objects will be in focus. */
func (ctx *Context) DepthOfField(fstop,focallength,focaldistance RtFloat) error {
	return ctx.write(fmt.Sprintf("DepthOfField %s %s %s\n",fstop,focallength,focaldistance))
}

/* Shutter; sets the times at which the shutter opens and closes. */
func (ctx *Context) Shutter(min,max RtFloat) error {
	return ctx.write(fmt.Sprintf("Shutter %s %s\n",min,max))
}




