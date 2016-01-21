package ri

import (
	"fmt"
)

/* PixelVariance; the color of a pixel computed by the rendering program */
func (ctx *Context) PixelVariance(variation RtFloat) error {
	return ctx.write(fmt.Sprintf("PixelVariance %s\n",variation))
}

/* PixelSamples; set the effective hider sampling rate in the horizontal and vertial directions */
func (ctx *Context) PixelSamples(xsamples,ysamples RtFloat) error {
	return ctx.write(fmt.Sprintf("PixelSamples %s %s\n",xsamples,ysamples))
}

/* PixelFilter; antialiasing is performed by filtering the geometry (or supersampling) and then
 * sampling at pixel locations. */
func (ctx *Context) PixelFilter(filterfunc RtToken,xwidth,ywidth RtFloat) error {
	return ctx.write(fmt.Sprintf("PixelFilter %s %s %s\n",filterfunc,xwidth,ywidth))
}

/* Exposure; this functions controls the sensitivity and nonlinearity of the exposure process. */
func (ctx *Context) Exposure(gain,gamma RtFloat) error {
	return ctx.write(fmt.Sprintf("Exposure %s %s\n",gain,gamma))
}

/* Imager; select an imager function programmed in the Shading Language. */
func (ctx *Context) Imager(name RtToken,parameterlist ...interface{}) error {
	return nil
}
	
