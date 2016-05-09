package ri

/* see Siggraph 2003, Course 9 'The Evolution of RIB' by Byron Bashforth, pg.5 */

/* IfBegin */
func (r *Ri) IfBegin(operation RtString) error {
	return r.writef("IfBegin", operation, PARAMETERLIST)
}

func (r *Ri) IfBeginV(args,tokens,values []Rter) error {
	return r.writef("IfBegin",args...)
}

/* ElseIf */
func (r *Ri) ElseIf(operation RtString) error {
	return r.writef("ElseIf", operation, PARAMETERLIST)
}

func (r *Ri) ElseIfV(args,tokens,values []Rter) error {
	return r.writef("ElseIf",args...)
}

/* Else */
func (r *Ri) Else() error {
	return r.writef("Else")
}

func (r *Ri) ElseV(args,tokens,values []Rter) error {
	return r.writef("Else")
}

/* IfEnd */
func (r *Ri) IfEnd() error {
	return r.writef("IfEnd")
}

func (r *Ri) IfEndV(args,tokens,values []Rter) error {
	return r.writef("IfEnd")
}
