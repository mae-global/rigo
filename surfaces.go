package ri

/* SubdivisionMesh */
func (ctx *Context) SubdivisionMesh(scheme RtToken, nfaces RtInt, nvertices, vertices RtIntArray, ntags RtInt, tags RtTokenArray, nargs, intargs RtIntArray, floatargs RtFloatArray, parameterlist ...Rter) error {

	var out = []Rter{scheme, nvertices, vertices, tags, nargs, intargs, floatargs}
	out = append(out, parameterlist...)

	return ctx.writef("SubdivisionMesh", out...)
}

/* Points */
func (ctx *Context) Points(npoints RtInt, parameterlist ...Rter) error {
	return ctx.writef("Points", parameterlist...)
}

/* Curves */
func (ctx *Context) Curves(typeof RtToken, ncurves RtInt, nvertices RtIntArray, wrap RtToken, parameterlist ...Rter) error {

	var out = []Rter{typeof, nvertices, wrap}
	out = append(out, parameterlist...)

	return ctx.writef("Curves", out...)
}

/* Blobby */
func (ctx *Context) Blobby(nleaf, ncode RtInt, code RtIntArray, nfloats RtInt, floats RtFloatArray, nstrings RtInt, strings RtStringArray, parameterlist ...Rter) error {

	var out = []Rter{nleaf, floats, strings}
	out = append(out, parameterlist...)

	return ctx.writef("Blobby", out...)
}
