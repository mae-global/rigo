package ri

func (r *Ri) Begin(name RtString) error {
	if err := r.writef("Begin", name); err != nil {
		return err
	}
	return r.writef("version",Version)
}

func (r *Ri) End() error {
	return r.writef("End")
}

/* FrameBegin mark the beginning of a single frame of an animated sequenece */
func (r *Ri) FrameBegin(frame RtInt) error {
	defer func() { r.Depth(1) }()
	return r.writef("FrameBegin", frame)
}

/* FrameEnd mark the end of a single frame of an animated sequence */
func (r *Ri) FrameEnd() error {
	r.Depth(-1)
	return r.writef("FrameEnd")
}

/* WorldBegin is invoked, all rendering options are frozen */
func (r *Ri) WorldBegin(args ...RtAnnotation) error {
	defer func() { r.Depth(1) }()
	if len(args) > 1 {
		return ErrBadParamlist
	}
	return r.writef("WorldBegin", parseAnnotations(args...)...)
}

/* WorldEnd */
func (r *Ri) WorldEnd(args ...RtAnnotation) error {
	r.Depth(-1)
	if len(args) > 1 {
		return ErrBadParamlist
	}
	return r.writef("WorldEnd", parseAnnotations(args...)...)
}

/* Comment */
func (r *Ri) Comment(comment RtName) error {
	return r.writef("#", comment)
}

/* Format set the horizontal and vertical resolution (in pixels) of the image to be rendered */
func (r *Ri) Format(xresolution, yresolution RtInt, pixelaspectratio RtFloat) error {
	return r.writef("Format", xresolution, yresolution, pixelaspectratio)
}

/* FrameAspectRatio is the ration of the width to the height of the desired image. */
func (r *Ri) FrameAspectRatio(frameaspectratio RtFloat) error {
	return r.writef("FrameAspectRatio", frameaspectratio)
}

/* ScreenWindow this procedure defines a rectangle in the image plane. */
func (r *Ri) ScreenWindow(left, right, bottom, top RtFloat) error {
	return r.writef("ScreenWindow", left, right, bottom, top)
}

/* CropWindow render only a subrectangle of the image. */
func (r *Ri) CropWindow(xmin, xmax, ymin, ymax RtFloat) error {
	return r.writef("CropWindow", RtFloatArray{xmin, xmax, ymin, ymax})
}

/* Projection the project determines how camera coordinates are converted to screen coordinates */
func (r *Ri) Projection(token RtToken, parameterlist ...Rter) error {
	var out = []Rter{token,PARAMETERLIST}
	out = append(out, parameterlist...)
	return r.writef("Projection", out...)
}

/* Clipping sets the position of the near and far clipping planes along the direction of view. */
func (r *Ri) Clipping(near, far RtFloat) error {
	return r.writef("Clipping", near, far)
}

/* ClippingPlane adds a user-specified clipping plane. */
func (r *Ri) ClippingPlane(x, y, z, nx, ny, nz RtFloat) error {
	return r.writef("ClippingPlane", x, y, z, nx, ny, nz)
}

/* DepthOfField focaldistance sets the distance along the direction of view at which objects will be in focus. */
func (r *Ri) DepthOfField(fstop, focallength, focaldistance RtFloat) error {
	return r.writef("DepthOfField", fstop, focallength, focaldistance)
}

/* Shutter sets the times at which the shutter opens and closes. */
func (r *Ri) Shutter(min, max RtFloat) error {
	return r.writef("Shutter", min, max)
}
