package ri

/* PixelVariance; the color of a pixel computed by the rendering program */
func (ctx *Context) PixelVariance(variation RtFloat) error {
	return ctx.writef("PixelVariance",variation)
}

/* PixelSamples; set the effective hider sampling rate in the horizontal and vertial directions */
func (ctx *Context) PixelSamples(xsamples,ysamples RtFloat) error {
	return ctx.writef("PixelSamples",xsamples,ysamples)
}

/* PixelFilter; antialiasing is performed by filtering the geometry (or supersampling) and then
 * sampling at pixel locations. */
func (ctx *Context) PixelFilter(filterfunc RtToken,xwidth,ywidth RtFloat) error {
	return ctx.writef("PixelFilter",filterfunc,xwidth,ywidth)
}

/* Exposure; this functions controls the sensitivity and nonlinearity of the exposure process. */
func (ctx *Context) Exposure(gain,gamma RtFloat) error {
	return ctx.writef("Exposure",gain,gamma)
}

/* Imager; select an imager function programmed in the Shading Language. */
func (ctx *Context) Imager(name RtToken,parameterlist ...Rter) error {

	out := make([]Rter,0)
	out = append(out,name)
	out = append(out,parameterlist...)

	return ctx.writef("Imager",out...)
}


/* Quantize; set the quantization parameters for colors or depth. */
func (ctx *Context) Quantize(typeof RtToken,one,min,max RtInt,ditheramplitude RtFloat) error {
	return ctx.writef("Quantize",typeof,one,min,max,ditheramplitude)
}
	
/* Display; choose a display by name and set the type of output being generated. */
func (ctx *Context) Display(name,typeof,mode RtToken,parameterlist ...Rter) error {

	out := make([]Rter,0)
	out = append(out,name)
	out = append(out,typeof)
	out = append(out,mode)
	out = append(out,parameterlist...)

	return ctx.writef("Display",out...)
}

/* Hider; */
func (ctx *Context) Hider(typeof RtToken, parameterlist ...Rter) error {
	
	out := make([]Rter,0)
	out = append(out,typeof)
	out = append(out,parameterlist...)

	return ctx.writef("Hider",out...)
}

/* ColorSamples; controls the number of color components or samples to be used in specifying colors */
func (ctx *Context) ColorSamples(n RtInt, nRGB RtFloatArray,RGBn RtFloatArray) error {
	return ctx.writef("ColorSamples",n,nRGB,RGBn)
}

/* RelativeDetail; the relative level of detail scales the results of all level of detail calculations. */
func (ctx *Context) RelativeDetail(relativedetail RtFloat) error {
	return ctx.writef("RelativeDetail",relativedetail)
}

	



