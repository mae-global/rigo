package ri

/* Points */
func (ctx *Context) Points(npoints RtInt,parameterlist ...Rter) error {
	return ctx.writef("Points",parameterlist...)
}

/* Curves */
func (ctx *Context) Curves(typeof RtToken,ncurves RtInt,nvertices RtIntArray,wrap RtToken,parameterlist ...Rter) error {

	var out = []Rter{typeof,nvertices,wrap}
	out = append(out,parameterlist...)

	return ctx.writef("Curves",out...)
}
