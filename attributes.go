package ri


/* AttributeBegin push the current set of attributes. */
func (ctx *Context) AttributeBegin() error {
	defer func() {ctx.depth++}()
	return ctx.writef("AttributeBegin")
}

/* AttributeEnd pop the current set of attributes */
func (ctx *Context) AttributeEnd() error {
	ctx.depth--
	return ctx.writef("AttributeEnd")
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
	
