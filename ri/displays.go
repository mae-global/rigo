package ri

/* PixelVariance the color of a pixel computed by the rendering program */
func (r *Ri) PixelVariance(variation RtFloat) error {
	return r.writef("PixelVariance", variation)
}

/* PixelSamples set the effective hider sampling rate in the horizontal and vertial directions */
func (r *Ri) PixelSamples(xsamples, ysamples RtFloat) error {
	return r.writef("PixelSamples", xsamples, ysamples)
}

/* PixelFilter antialiasing is performed by filtering the geometry (or supersampling) and then
 * sampling at pixel locations. */
func (r *Ri) PixelFilter(filterfunc RtFilterFunc, xwidth, ywidth RtFloat) error {
	return r.writef("PixelFilter", filterfunc, xwidth, ywidth)
}

/* Exposure this functions controls the sensitivity and nonlinearity of the exposure process. */
func (r *Ri) Exposure(gain, gamma RtFloat) error {
	return r.writef("Exposure", gain, gamma)
}

/* Imager; select an imager function programmed in the Shading Language. */
func (r *Ri) Imager(name RtToken, parameterlist ...Rter) error {

	var out = []Rter{name,PARAMETERLIST}
	out = append(out, parameterlist...)

	return r.writef("Imager", out...)
}

/* Quantize set the quantization parameters for colors or depth. */
func (r *Ri) Quantize(typeof RtToken, one, min, max RtInt, ditheramplitude RtFloat) error {
	return r.writef("Quantize", typeof, one, min, max, ditheramplitude)
}

/* Display choose a display by name and set the type of output being generated. */
func (r *Ri) Display(name, typeof, mode RtToken, parameterlist ...Rter) error {

	var out = []Rter{name, typeof, mode,PARAMETERLIST}
	out = append(out, parameterlist...)

	return r.writef("Display", out...)
}

/* Hider */
func (r *Ri) Hider(typeof RtToken, parameterlist ...Rter) error {

	var out = []Rter{typeof,PARAMETERLIST}
	out = append(out, parameterlist...)

	return r.writef("Hider", out...)
}

/* ColorSamples controls the number of color components or samples to be used in specifying colors */
func (r *Ri) ColorSamples(n RtInt, nRGB RtFloatArray, RGBn RtFloatArray) error {
	return r.writef("ColorSamples", n, nRGB, RGBn)
}

/* RelativeDetail the relative level of detail scales the results of all level of detail calculations. */
func (r *Ri) RelativeDetail(relativedetail RtFloat) error {
	return r.writef("RelativeDetail", relativedetail)
}
