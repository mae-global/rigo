package ri

/* Procedural */
func (r *Ri) Procedural(args RtStringArray, bound RtBound, subdividefunc RtProcSubdivFunc, freefunc RtProcFreeFunc) error {
	return r.writef("Procedural", subdividefunc, args, bound)
}
