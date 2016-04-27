package ri

/* see Siggraph 2003, Course 9 'The Evolution of RIB' by Byron Bashforth, pg.5 */

/* IfBegin */
func (r *Ri) IfBegin(operation RtString) error {
	return r.writef("IfBegin", operation, PARAMETERLIST)
}

/* ElseIf */
func (r *Ri) ElseIf(operation RtString) error {
	return r.writef("ElseIf", operation, PARAMETERLIST)
}

/* Else */
func (r *Ri) Else() error {
	return r.writef("Else")
}

/* IfEnd */
func (r *Ri) IfEnd() error {
	return r.writef("IfEnd")
}
