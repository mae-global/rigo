package ri

/* Procedural */
func (ctx *Context) Procedural(args RtStringArray,bound RtBound, subdividefunc RtProcSubdivFunc,freefunc RtProcFreeFunc) error {
	return ctx.writef("Procedural",subdividefunc,args,bound)
}


