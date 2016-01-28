package ri

/* SubdivisionMesh */
func (r *Ri) SubdivisionMesh(scheme RtToken, nfaces RtInt, nvertices, vertices RtIntArray, ntags RtInt, tags RtTokenArray, nargs, intargs RtIntArray, floatargs RtFloatArray, parameterlist ...Rter) error {

	var out = []Rter{scheme, nvertices, vertices, tags, nargs, intargs, floatargs}
	out = append(out, parameterlist...)

	return r.writef("SubdivisionMesh", out...)
}

/* Points */
func (r *Ri) Points(npoints RtInt, parameterlist ...Rter) error {
	return r.writef("Points", parameterlist...)
}

/* Curves */
func (r *Ri) Curves(typeof RtToken, ncurves RtInt, nvertices RtIntArray, wrap RtToken, parameterlist ...Rter) error {

	var out = []Rter{typeof, nvertices, wrap}
	out = append(out, parameterlist...)

	return r.writef("Curves", out...)
}

/* Blobby */
func (r *Ri) Blobby(nleaf, ncode RtInt, code RtIntArray, nfloats RtInt, floats RtFloatArray, nstrings RtInt, strings RtStringArray, parameterlist ...Rter) error {

	var out = []Rter{nleaf, floats, strings}
	out = append(out, parameterlist...)

	return r.writef("Blobby", out...)
}
