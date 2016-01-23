package ri

import (
	"os"
)

func (ctx *Context) Begin(name string) error {
	if ctx.writer == nil {
		if f,err := os.Create(name); err != nil {
			return err
		} else {
			ctx.writer = &File{f,"",false}
		}
	}

	var out = []Rter{RtName("##RenderMan RIB-Structure 1.1")}
	if ctx.entity {
		out = append(out,RtName("Entity"))
	}

	return ctx.write(out...)
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

/* WorldBegin is invoked, all rendering options are frozen */
func (ctx *Context) WorldBegin(args ...RtAnnotation) error {
	defer func() { ctx.depth++ }()
	if len(args) > 1 {
		return ErrBadParamlist
	}
	return ctx.writef("WorldBegin",parseAnnotations(args...)...)
}

/* WorldEnd */
func (ctx *Context) WorldEnd(args ...RtAnnotation) error {
	ctx.depth--
	if len(args) > 1 {
		return ErrBadParamlist
	}	
	return ctx.writef("WorldEnd",parseAnnotations(args...)...)
}

/* Comment */
func (ctx *Context) Comment(comment RtName) error {
	return ctx.writef("#",comment)
}

/* Format set the horizontal and vertical resolution (in pixels) of the image to be rendered */
func (ctx *Context) Format(xresolution,yresolution RtInt,pixelaspectratio RtFloat) error {
	return ctx.writef("Format",xresolution,yresolution,pixelaspectratio)
}

/* FrameAspectRatio is the ration of the width to the height of the desired image. */
func (ctx *Context) FrameAspectRatio(frameaspectratio RtFloat) error {
	return ctx.writef("FrameAspectRatio",frameaspectratio)
}

/* ScreenWindow this procedure defines a rectangle in the image plane. */
func (ctx *Context) ScreenWindow(left,right,bottom,top RtFloat) error {
	return ctx.writef("ScreenWindow",left,right,bottom,top)
}

/* CropWindow render only a subrectangle of the image. */
func (ctx *Context) CropWindow(xmin,xmax,ymin,ymax RtFloat) error {
	return ctx.writef("CropWindow",RtFloatArray{xmin,xmax,ymin,ymax})
}

/* Projection the project determines how camera coordinates are converted to screen coordinates */
func (ctx *Context) Projection(token RtToken, parameterlist ...Rter) error {
	var out = []Rter{token}
	out = append(out,parameterlist...)
	return ctx.writef("Projection",out...)
}

/* Clipping sets the position of the near and far clipping planes along the direction of view. */
func (ctx *Context) Clipping(near,far RtFloat) error {
	return ctx.writef("Clipping",near,far)
}

/* ClippingPlane adds a user-specified clipping plane. */
func (ctx *Context) ClippingPlane(x,y,z,nx,ny,nz RtFloat) error {
	return ctx.writef("ClippingPlane",x,y,z,nx,ny,nz)
}

/* DepthOfField focaldistance sets the distance along the direction of view at which objects will be in focus. */
func (ctx *Context) DepthOfField(fstop,focallength,focaldistance RtFloat) error {
	return ctx.writef("DepthOfField",fstop,focallength,focaldistance)
}

/* Shutter sets the times at which the shutter opens and closes. */
func (ctx *Context) Shutter(min,max RtFloat) error {
	return ctx.writef("Shutter",min,max)
}




