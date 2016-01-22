package ri

/* SubdivisionMesh */
func (ctx *Context) SubdivisionMesh(scheme RtToken,nfaces RtInt,nvertices,vertices RtIntArray,ntags RtInt,tags RtTokenArray,nargs,intargs RtIntArray,floatargs RtFloatArray,parameterlist ...Rter) error {

	var out = []Rter{scheme,nvertices,vertices,tags,nargs,intargs,floatargs}
	out = append(out,parameterlist...)

	return ctx.writef("SubdivisionMesh",out...)
}


