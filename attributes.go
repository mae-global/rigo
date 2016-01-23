package ri


/* AttributeBegin push the current set of attributes, you can add a single annotation */
func (ctx *Context) AttributeBegin(args ...RtAnnotation) error {
	defer func() {ctx.depth++}()
	if len(args) > 1 {
		return ErrBadParamlist
	}
	return ctx.writef("AttributeBegin",parseAnnotations(args...)...)
}

/* AttributeEnd pop the current set of attributes */
func (ctx *Context) AttributeEnd(args ...RtAnnotation) error {
	ctx.depth--
	if len(args) > 1 {
		return ErrBadParamlist
	}
	return ctx.writef("AttributeEnd",parseAnnotations(args...)...)
}

/* Color set the current color */
func (ctx *Context) Color(color RtColor) error {
	return ctx.writef("Color",color)
}

/* Opacity set the current opacity */
func (ctx *Context) Opacity(color RtColor) error {
	return ctx.writef("Opacity",color)
}

/* TextureCoordinates set the current set of texture coordinates to the values */
func (ctx *Context) TextureCoordinates(s1,t1,s2,t2,s3,t3,s4,t4 RtFloat) error {
	return ctx.writef("TextureCoordinates",RtFloatArray{s1,t1,s2,t2,s3,t3,s4,t4})
}

/* LightSource create a non-area light */
func (ctx *Context) LightSource(shadername RtToken,parameterlist ...Rter) (RtLightHandle,error) {
	lh := RtLightHandle(ctx.lights)
	ctx.lights++
	
	var out = []Rter{shadername,lh}
	out = append(out,parameterlist...)

	return lh,ctx.writef("LightSource",out...)
}

/* AreaLightSource */
func (ctx *Context) AreaLightSource(shadername RtToken,parameterlist ...Rter) (RtLightHandle,error) {
	lh := RtLightHandle(ctx.lights)
	ctx.lights++

	var out = []Rter{shadername,lh}
	out = append(out,parameterlist...)

	return lh,ctx.writef("AreaLightSource",out...)
}
	
/* Illuminate If onoff is true and the light source referred to by the light is not
 * currently in the4 current light source list, then add it to that list */
func (ctx *Context) Illuminate(light RtLightHandle,onoff RtBoolean) error {

	/* TODO: could actuall check that the light exists */
	return ctx.writef("Illuminate",light,onoff)
}

/* Surface shadername is the name of a surface shader */
func (ctx *Context) Surface(shadername RtToken, parameterlist ...Rter) error {
	
	var out = []Rter{shadername}
	out = append(out,parameterlist...)
	
	return ctx.writef("Surface",out...)
}

/* Displacement set the current displacement shader to the named shader. */
func (ctx *Context) Displacement(shadername RtToken,parameterlist ...Rter) error {

	var out = []Rter{shadername}
	out = append(out,parameterlist...)

	return ctx.writef("Displacement",out...)
}

/* Atmosphere this procedure sets the current atmosphere shader. */
func (ctx *Context) Atmosphere(shadername RtToken,parameterlist ...Rter) error {

	var out = []Rter{shadername}
	out = append(out,parameterlist...)

	return ctx.writef("Atmosphere",out...)
}

/* Interior this procedure sets the current interior volume shader. */
func (ctx *Context) Interior(shadername RtToken,parameterlist ...Rter) error {

	var out = []Rter{shadername}
	out = append(out,parameterlist...)

	return ctx.writef("Interior",out...)
}

/* Exterior this procedure sets the curent exterior volume shader. */
func (ctx *Context) Exterior(shadername RtToken,parameterlist ...Rter) error {

	var out = []Rter{shadername}
	out = append(out,parameterlist...)
	
	return ctx.writef("Exterior",out...)
}

/* ShadingRate set the current shading rate to size */
func (ctx *Context) ShadingRate(size RtFloat) error {
	return ctx.writef("ShadingRate",size)
}

/* ShadingInterpolation this function controls how values are interpolated between shading samples */
func (ctx *Context) ShadingInterpolation(typeof RtToken) error {
	return ctx.writef("ShadingInterpolation",typeof)
}

/* Matte indicates whether subsequent primitives are matte objects */
func (ctx *Context) Matte(onoff RtBoolean) error {
	return ctx.writef("Matte",onoff)
}

/* Bound This procedure sets the current bound to bound. The bounding box. */
func (ctx *Context) Bound(bound RtBound) error {
	return ctx.writef("Bound",bound)
}
	
/* Detail set the current bound to bound */
func (ctx *Context) Detail(bound RtBound) error {
	return ctx.writef("Detail",bound)
}

/* DetailRange set the current detail range */
func (ctx *Context) DetailRange(minvisible,lowertransition,uppertransition,maxvisible RtFloat) error {
	return ctx.writef("DetailRange",RtFloatArray{minvisible,lowertransition,uppertransition,maxvisible})
}

/* GeometricApproximation */
func (ctx *Context) GeometricApproximation(typeof RtToken,value RtFloat) error {
	return ctx.writef("GeometricApproximation",typeof,value)
}

/* Orientation This procedure sets the current orientation to be either "outside", "inside","lh","rh" */
func (ctx *Context) Orientation(orientation RtToken) error {
	return ctx.writef("Orientation",orientation)
}

/* ReverseOrientation causes the current orientation to be toggled */
func (ctx *Context) ReverseOrientation() error {
	return ctx.writef("ReverseOrientation")
}

/* Sides if sides is 2, subsequent surfaces are considered two-sided */
func (ctx *Context) Sides(sides RtInt) error {
	return ctx.writef("Sides",sides)
}




