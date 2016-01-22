package ri

import (
	"fmt"
)

/* PixelVariance; the color of a pixel computed by the rendering program */
func (ctx *Context) PixelVariance(variation RtFloat) error {
	return ctx.write(fmt.Sprintf("PixelVariance %s\n",variation.Serialise()))
}

/* PixelSamples; set the effective hider sampling rate in the horizontal and vertial directions */
func (ctx *Context) PixelSamples(xsamples,ysamples RtFloat) error {
	return ctx.write(fmt.Sprintf("PixelSamples %s\n",serialise_s(xsamples,ysamples)))
}

/* PixelFilter; antialiasing is performed by filtering the geometry (or supersampling) and then
 * sampling at pixel locations. */
func (ctx *Context) PixelFilter(filterfunc RtToken,xwidth,ywidth RtFloat) error {
	return ctx.write(fmt.Sprintf("PixelFilter %s\n",serialise_s(filterfunc,xwidth,ywidth)))
}

/* Exposure; this functions controls the sensitivity and nonlinearity of the exposure process. */
func (ctx *Context) Exposure(gain,gamma RtFloat) error {
	return ctx.write(fmt.Sprintf("Exposure %s\n",serialise_s(gain,gamma)))
}

/* Imager; select an imager function programmed in the Shading Language. */
func (ctx *Context) Imager(name RtToken,parameterlist ...Rter) error {

	out,err := serialise(parameterlist...)
	if err != nil {
		return err
	}
	return ctx.write(fmt.Sprintf("Imager %s %s\n",name.Serialise(),out))
}


/* Quantize; set the quantization parameters for colors or depth. */
func (ctx *Context) Quantize(typeof RtToken,one,min,max RtInt,ditheramplitude RtFloat) error {
	return ctx.write(fmt.Sprintf("Quantize %s\n",serialise_s(typeof,one,min,max,ditheramplitude)))
}
	
/* Display; choose a display by name and set the type of output being generated. */
func (ctx *Context) Display(name,typeof,mode RtToken,parameterlist ...Rter) error {

	out,err := serialise(parameterlist...)
	if err != nil {
		return err
	}
	return ctx.write(fmt.Sprintf("Display %s %s\n",serialise_s(name,typeof,mode),out))
}

/* Hider; */
func (ctx *Context) Hider(typeof RtToken, parameterlist ...Rter) error {
	
	out,err := serialise(parameterlist...)
	if err != nil {
		return err
	}
	return ctx.write(fmt.Sprintf("Hider %s %s\n",typeof.Serialise(),out))
}

/* ColorSamples; controls the number of color components or samples to be used in specifying colors */
func (ctx *Context) ColorSamples(n RtInt, nRGB RtFloatArray,RGBn RtFloatArray) error {
	return ctx.write(fmt.Sprintf("ColorSamples %s\n",serialise_s(n,nRGB,RGBn)))
}

/* RelativeDetail; the relative level of detail scales the results of all level of detail calculations. */
func (ctx *Context) RelativeDetail(relativedetail RtFloat) error {
	return ctx.write(fmt.Sprintf("RelativeDetail %s\n",relativedetail))
}

	



