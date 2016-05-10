package ri

/* SolidBegin */
func (r *Ri) SolidBegin(operation RtToken) error {
	return r.writef("SolidBegin", operation)
}

/* SolidEnd */
func (r *Ri) SolidEnd() error {
	return r.writef("SolidEnd")
}

/* ObjectBegin */
func (r *Ri) ObjectBegin() (RtObjectHandle, error) {
	oh, err := r.ObjectHandle()
	if err != nil {
		return oh, err
	}
	return oh, r.writef("ObjectBegin", oh)
}

/* ObjectEnd */
func (r *Ri) ObjectEnd() error {
	return r.writef("ObjectEnd")
}

/* ObjectInstance */
func (r *Ri) ObjectInstance(handle RtObjectHandle) error {
	if err := r.CheckObjectHandle(handle); err != nil {
		return err
	}
	return r.writef("ObjectInstance", handle)
}
