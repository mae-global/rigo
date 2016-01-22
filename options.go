package ri

/* Option; render programs may have additional implementation-specific options. */
func (ctx *Context) Option(name RtToken,parameterlist ...Rter) error {
	out := make([]Rter,0)
	out = append(out,name)
	out = append(out,parameterlist...)

	return ctx.writef("Option",out...)
}

/* Attribute Set the paramters of the attribute name. Implementation-specific */
func (ctx *Context) Attribute(name RtToken,parameterlist ...Rter) error {
	
	var out = []Rter{name}
	out = append(out,parameterlist...)

	return ctx.writef("Attribute",out...)
}
