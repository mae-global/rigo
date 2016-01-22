package ri

/* Basis Set the current u-basis to ubasis and the current v-basis to vbasis. */
func (ctx *Context) Basis(ubasis RtBasis,ustep RtInt,vbasis RtBasis,vstep RtInt) error {

	/* TODO: see spec pg.68, required uname for known basis's */ 

	return ctx.writef("Basis",ubasis,ustep,vbasis,vstep)
}
	
/* Patch define a single patch */
func (ctx *Context) Patch(typeof RtToken,parameterlist ...Rter) error {

	var out = []Rter{typeof}
	out = append(out,parameterlist...)
	
	return ctx.writef("Patch",out...)
}

/* PatchMesh This primitive is a compact way of specifying a quadrilateral mesh of patches */
func (ctx *Context) PatchMesh(typeof RtToken,nu RtInt,uwrap RtToken,nv RtInt,vwrap RtToken,parameterlist ...Rter) error {

	var out = []Rter{typeof,nu,uwrap,nv,vwrap}
	out = append(out,parameterlist...)

	return ctx.writef("PatchMesh",out...)
}

/* NuPatch */
func (ctx *Context) NuPatch(nu,uorder RtInt,uknot RtFloatArray,umin,umax RtFloat,nv,vorder RtInt,vknot RtFloatArray,vmin,vmax RtFloat,parameterlist ...Rter) error {
	
	var out = []Rter{nu,uorder,uknot,umin,umax,nv,vorder,vknot,vmin,vmax}
	out = append(out,parameterlist...)

	return ctx.writef("NuPatch",out...)
}

/* TrimCurve */
func (ctx *Context) TrimCurve(nloops RtInt,ncurves RtIntArray,order RtIntArray,knot RtFloatArray,min,max,RtFloat,n RtIntArray,u,v,w RtFloatArray) error {

	return ctx.writef("TrimCurve",ncurves,order,knot,min,max,n,u,v,w)
}

