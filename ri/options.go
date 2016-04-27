package ri

/* Declare Declare the name and type of a variable */
func (r *Ri) Declare(name, declaration RtString) error {
	return r.writef("Declare", name, declaration)
}

/* Option render programs may have additional implementation-specific options. */
func (r *Ri) Option(name RtToken, parameterlist ...Rter) error {
	out := []Rter{name, PARAMETERLIST}
	out = append(out, parameterlist...)

	return r.writef("Option", out...)
}

/* Attribute Set the paramters of the attribute name. Implementation-specific */
func (r *Ri) Attribute(name RtToken, parameterlist ...Rter) error {

	var out = []Rter{name, PARAMETERLIST}
	out = append(out, parameterlist...)

	return r.writef("Attribute", out...)
}

/* Geometry */
func (r *Ri) Geometry(typeof RtToken, parameterlist ...Rter) error {

	var out = []Rter{typeof, PARAMETERLIST}
	out = append(out, parameterlist...)

	return r.writef("Geometry", out...)
}

/* MotionBegin */
func (r *Ri) MotionBegin(n RtInt, t ...RtFloat) error {
	defer func() { r.Depth(1) }()
	return r.writef("MotionBegin", RtFloatArray(t))
}

/* MotionEnd */
func (r *Ri) MotionEnd() error {
	r.Depth(-1)
	return r.writef("MotionEnd")
}
