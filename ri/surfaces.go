package ri

/* SubdivisionMesh */
func (r *Ri) SubdivisionMesh(scheme RtToken, nfaces RtInt, nvertices, vertices RtIntArray, ntags RtInt, tags RtTokenArray, nargs, intargs RtIntArray, floatargs RtFloatArray, parameterlist ...Rter) error {

	var out = []Rter{scheme, nvertices, vertices, tags, nargs, intargs, floatargs, PARAMETERLIST}
	out = append(out, parameterlist...)

	return r.writef("SubdivisionMesh", out...)
}

/* Points */
func (r *Ri) Points(npoints RtInt, parameterlist ...Rter) error {
	var out = []Rter{PARAMETERLIST}
	out = append(out, parameterlist...)
	return r.writef("Points", out...)
}

/* Curves */
func (r *Ri) Curves(typeof RtToken, ncurves RtInt, nvertices RtIntArray, wrap RtToken, parameterlist ...Rter) error {

	var out = []Rter{typeof, nvertices, wrap, PARAMETERLIST}
	out = append(out, parameterlist...)

	return r.writef("Curves", out...)
}

/* Blobby */
func (r *Ri) Blobby(nleaf, ncode RtInt, code RtIntArray, nfloats RtInt, floats RtFloatArray, nstrings RtInt, strings RtStringArray, parameterlist ...Rter) error {

	var out = []Rter{nleaf, floats, strings, PARAMETERLIST}
	out = append(out, parameterlist...)

	return r.writef("Blobby", out...)
}
