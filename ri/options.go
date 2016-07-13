package ri

/* Declare Declare the name and type of a variable */
func (r *Ri) Declare(name, declaration RtString) error {
	/* declare to the dictionary */
	if r.RiDictionarer != nil {
		if err := r.RiDictionarer.Declare(RtToken(string(declaration) + " " + string(name))); err != nil {
			return err
		}
	}

	return r.writef("Declare", name, declaration)
}

func (r *Ri) DeclareV(args, tokens, values []Rter) error {
	return r.writef("Declare", args...)
}

/* Option render programs may have additional implementation-specific options. */
func (r *Ri) Option(name RtToken, parameterlist ...Rter) error {
	out := []Rter{name, PARAMETERLIST}
	out = append(out, parameterlist...)

	return r.writef("Option", out...)
}

func (r *Ri) OptionV(args, tokens, values []Rter) error {

	out := make([]Rter, 0)
	out = append(out, args...)
	out = append(out, PARAMETERLIST)
	out = append(out, Mix(tokens, values)...)
	return r.writef("Option", out...)
}

/* Attribute Set the paramters of the attribute name. Implementation-specific */
func (r *Ri) Attribute(name RtToken, parameterlist ...Rter) error {

	var out = []Rter{name, PARAMETERLIST}
	out = append(out, parameterlist...)

	return r.writef("Attribute", out...)
}

func (r *Ri) AttributeV(args, tokens, values []Rter) error {

	out := make([]Rter, 0)
	out = append(out, args...)
	out = append(out, PARAMETERLIST)
	out = append(out, Mix(tokens, values)...)
	return r.writef("Attribute", out...)
}

/* Geometry */
func (r *Ri) Geometry(typeof RtToken, parameterlist ...Rter) error {

	var out = []Rter{typeof, PARAMETERLIST}
	out = append(out, parameterlist...)

	return r.writef("Geometry", out...)
}

func (r *Ri) GeometryV(args, tokens, values []Rter) error {

	out := make([]Rter, 0)
	out = append(out, args...)
	out = append(out, PARAMETERLIST)
	out = append(out, Mix(tokens, values)...)
	return r.writef("Geometry", out...)
}

/* MotionBegin */
func (r *Ri) MotionBegin(n RtInt, t ...RtFloat) error {
	return r.writef("MotionBegin", RtFloatArray(t))
}

func (r *Ri) MotionBeginV(args, tokens, values []Rter) error {
	return r.writef("MotionBegin", args...)
}

/* MotionEnd */
func (r *Ri) MotionEnd() error {
	return r.writef("MotionEnd")
}

func (r *Ri) MotionEndV(args, tokens, values []Rter) error {
	return r.writef("MotionEnd")
}
