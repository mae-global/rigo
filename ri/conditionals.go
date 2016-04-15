package ri

/* IfBegin */
func (r *Ri) IfBegin(operation RtString) error {
	return r.writef("IfBegin",operation)
}

/* ElseIf */
func (r *Ri) ElseIf(operation RtString) error {
	return r.writef("ElseIf",operation)
}

/* Else */
func (r *Ri) Else() error {
	return r.writef("Else")
}

/* IfEnd */
func (r *Ri) IfEnd() error {
	return r.writef("IfEnd")
}
