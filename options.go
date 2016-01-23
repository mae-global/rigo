package ri

/* Declare Declare the name and type of a variable */
func (ctx *Context) Declare(name, declaration RtString) error {
	return ctx.writef("Declare", name, declaration)
}

/* Option render programs may have additional implementation-specific options. */
func (ctx *Context) Option(name RtToken, parameterlist ...Rter) error {
	out := make([]Rter, 0)
	out = append(out, name)
	out = append(out, parameterlist...)

	return ctx.writef("Option", out...)
}

/* Attribute Set the paramters of the attribute name. Implementation-specific */
func (ctx *Context) Attribute(name RtToken, parameterlist ...Rter) error {

	var out = []Rter{name}
	out = append(out, parameterlist...)

	return ctx.writef("Attribute", out...)
}

/* Geometry */
func (ctx *Context) Geometry(typeof RtToken, parameterlist ...Rter) error {

	var out = []Rter{typeof}
	out = append(out, parameterlist...)

	return ctx.writef("Geometry", out...)
}

/* MotionBegin */
func (ctx *Context) MotionBegin(n RtInt, t ...RtFloat) error {
	defer func() { ctx.depth++ }()
	return ctx.writef("MotionBegin", RtFloatArray(t))
}

/* MotionEnd */
func (ctx *Context) MotionEnd() error {
	ctx.depth--
	return ctx.writef("MotionEnd")
}
