package ri

/* Option; render programs may have additional implementation-specific options. */
func (ctx *Context) Option(name RtToken,parameterlist ...Rter) error {
	out := make([]Rter,0)
	out = append(out,name)
	out = append(out,parameterlist...)

	return ctx.writef("Option",out...)
}
