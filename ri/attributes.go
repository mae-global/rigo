package ri

/* AttributeBegin push the current set of attributes, you can add a single annotation */
func (r *Ri) AttributeBegin(args ...RtAnnotation) error {

	if len(args) > 1 {
		return ErrBadParamlist
	}
	return r.writef("AttributeBegin", parseAnnotations(args...)...)
}

func (r *Ri) AttributeBeginV(args, tokens, values []Rter) error {
	return r.writef("AttributeBegin", args...)
}

/* AttributeEnd pop the current set of attributes */
func (r *Ri) AttributeEnd(args ...RtAnnotation) error {

	if len(args) > 1 {
		return ErrBadParamlist
	}
	return r.writef("AttributeEnd", parseAnnotations(args...)...)
}

func (r *Ri) AttributeEndV(args, tokens, values []Rter) error {
	return r.writef("AttributeEnd", args...)
}

/* Color set the current color */
func (r *Ri) Color(color RtColor) error {
	return r.writef("Color", color)
}

func (r *Ri) ColorV(args, tokens, values []Rter) error {
	return r.writef("Color", args...)
}

/* Opacity set the current opacity */
func (r *Ri) Opacity(color RtColor) error {
	return r.writef("Opacity", color)
}

func (r *Ri) OpacityV(args, tokens, values []Rter) error {
	return r.writef("Opacity", args...)
}

/* TextureCoordinates set the current set of texture coordinates to the values */
func (r *Ri) TextureCoordinates(s1, t1, s2, t2, s3, t3, s4, t4 RtFloat) error {
	return r.writef("TextureCoordinates", RtFloatArray{s1, t1, s2, t2, s3, t3, s4, t4})
}

func (r *Ri) TextureCoordinatesV(args, tokens, values []Rter) error {
	return r.writef("TextureCoordinates", args...)
}

/* LightSource create a non-area light */
func (r *Ri) LightSource(shadername RtToken, parameterlist ...Rter) (RtLightHandle, error) {
	lh, err := r.LightHandle()
	if err != nil {
		return lh, err
	}

	var out = []Rter{shadername, lh, PARAMETERLIST}
	out = append(out, parameterlist...)

	return lh, r.writef("LightSource", out...)
}

func (r *Ri) LightSourceV(args, tokens, values []Rter) error {

	out := make([]Rter, 0)
	out = append(out, args...)
	out = append(out, PARAMETERLIST)
	out = append(out, Mix(tokens, values)...)

	return r.writef("LightSource", out...)
}

/* AreaLightSource */
func (r *Ri) AreaLightSource(shadername RtToken, parameterlist ...Rter) (RtLightHandle, error) {
	lh, err := r.LightHandle()
	if err != nil {
		return lh, err
	}
	var out = []Rter{shadername, lh, PARAMETERLIST}
	out = append(out, parameterlist...)

	return lh, r.writef("AreaLightSource", out...)
}

func (r *Ri) AreaLightSourceV(args, tokens, values []Rter) error {

	out := make([]Rter, 0)
	out = append(out, args...)
	out = append(out, PARAMETERLIST)
	out = append(out, Mix(tokens, values)...)

	return r.writef("AreaLightSource", out...)
}

/* Illuminate If onoff is true and the light source referred to by the light is not
 * currently in the4 current light source list, then add it to that list */
func (r *Ri) Illuminate(light RtLightHandle, onoff RtBoolean) error {

	/* TODO: could actuall check that the light exists */
	return r.writef("Illuminate", light, onoff)
}

func (r *Ri) IlluminateV(args, tokens, values []Rter) error {
	return r.writef("Illuminate", args...)
}

/* Surface shadername is the name of a surface shader */
func (r *Ri) Surface(shadername RtToken, parameterlist ...Rter) error {

	var out = []Rter{shadername, PARAMETERLIST}
	out = append(out, parameterlist...)

	return r.writef("Surface", out...)
}

func (r *Ri) SurfaceV(args, tokens, values []Rter) error {

	out := make([]Rter, 0)
	out = append(out, args...)
	out = append(out, PARAMETERLIST)
	out = append(out, Mix(tokens, values)...)

	return r.writef("Surface", out...)
}

/* Displacement set the current displacement shader to the named shader. */
func (r *Ri) Displacement(shadername RtToken, parameterlist ...Rter) error {

	var out = []Rter{shadername, PARAMETERLIST}
	out = append(out, parameterlist...)

	return r.writef("Displacement", out...)
}

func (r *Ri) DisplacementV(args, tokens, values []Rter) error {

	out := make([]Rter, 0)
	out = append(out, args...)
	out = append(out, PARAMETERLIST)
	out = append(out, Mix(tokens, values)...)

	return r.writef("Displacement", out...)
}

/* Atmosphere this procedure sets the current atmosphere shader. */
func (r *Ri) Atmosphere(shadername RtToken, parameterlist ...Rter) error {

	var out = []Rter{shadername, PARAMETERLIST}
	out = append(out, parameterlist...)

	return r.writef("Atmosphere", out...)
}

func (r *Ri) AtmosphereV(args, tokens, values []Rter) error {

	out := make([]Rter, 0)
	out = append(out, args...)
	out = append(out, PARAMETERLIST)
	out = append(out, Mix(tokens, values)...)

	return r.writef("Atmosphere", out...)
}

/* Interior this procedure sets the current interior volume shader. */
func (r *Ri) Interior(shadername RtToken, parameterlist ...Rter) error {

	var out = []Rter{shadername, PARAMETERLIST}
	out = append(out, parameterlist...)

	return r.writef("Interior", out...)
}

func (r *Ri) InteriorV(args, tokens, values []Rter) error {

	out := make([]Rter, 0)
	out = append(out, args...)
	out = append(out, PARAMETERLIST)
	out = append(out, Mix(tokens, values)...)

	return r.writef("Interior", out...)
}

/* Exterior this procedure sets the curent exterior volume shader. */
func (r *Ri) Exterior(shadername RtToken, parameterlist ...Rter) error {

	var out = []Rter{shadername, PARAMETERLIST}
	out = append(out, parameterlist...)

	return r.writef("Exterior", out...)
}

func (r *Ri) ExteriorV(args, tokens, values []Rter) error {

	out := make([]Rter, 0)
	out = append(out, args...)
	out = append(out, PARAMETERLIST)
	out = append(out, Mix(tokens, values)...)

	return r.writef("Exterior", out...)
}

/* ShadingRate set the current shading rate to size */
func (r *Ri) ShadingRate(size RtFloat) error {
	return r.writef("ShadingRate", size)
}

func (r *Ri) ShadingRateV(args, tokens, values []Rter) error {
	return r.writef("ShadingRate", args...)
}

/* ShadingInterpolation this function controls how values are interpolated between shading samples */
func (r *Ri) ShadingInterpolation(typeof RtToken) error {
	return r.writef("ShadingInterpolation", typeof)
}

func (r *Ri) ShadingInterpolationV(args, tokens, values []Rter) error {
	return r.writef("ShadingInterpolation", args...)
}

/* Matte indicates whether subsequent primitives are matte objects */
func (r *Ri) Matte(onoff RtBoolean) error {
	return r.writef("Matte", onoff)
}

func (r *Ri) MatteV(args, tokens, values []Rter) error {
	return r.writef("Matte", args...)
}

/* Bound This procedure sets the current bound to bound. The bounding box. */
func (r *Ri) Bound(bound RtBound) error {
	return r.writef("Bound", bound)
}

func (r *Ri) BoundV(args, tokens, values []Rter) error {
	return r.writef("Bound", args...)
}

/* Detail set the current bound to bound */
func (r *Ri) Detail(bound RtBound) error {
	return r.writef("Detail", bound)
}

func (r *Ri) DetailV(args, tokens, values []Rter) error {
	return r.writef("Detail", args...)
}

/* DetailRange set the current detail range */
func (r *Ri) DetailRange(minvisible, lowertransition, uppertransition, maxvisible RtFloat) error {
	return r.writef("DetailRange", RtFloatArray{minvisible, lowertransition, uppertransition, maxvisible})
}

func (r *Ri) DetailRangeV(args, tokens, values []Rter) error {
	return r.writef("DetailRange", args...)
}

/* GeometricApproximation */
func (r *Ri) GeometricApproximation(typeof RtToken, value RtFloat) error {
	return r.writef("GeometricApproximation", typeof, value)
}

func (r *Ri) GeometricApproximationV(args, tokens, values []Rter) error {
	return r.writef("GeometricApproximation", args...)
}

/* Orientation This procedure sets the current orientation to be either "outside", "inside","lh","rh" */
func (r *Ri) Orientation(orientation RtToken) error {
	return r.writef("Orientation", orientation)
}

func (r *Ri) OrientationV(args, tokens, values []Rter) error {
	return r.writef("Orientation", args...)
}

/* ReverseOrientation causes the current orientation to be toggled */
func (r *Ri) ReverseOrientation() error {
	return r.writef("ReverseOrientation")
}

func (r *Ri) ReverseOrientationV(args, tokens, values []Rter) error {
	return r.writef("ReverseOrientation")
}

/* Sides if sides is 2, subsequent surfaces are considered two-sided */
func (r *Ri) Sides(sides RtInt) error {
	return r.writef("Sides", sides)
}


func (r *Ri) SidesV(args, tokens, values []Rter) error {
	return r.writef("Sides", args...)
}

