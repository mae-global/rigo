package ri

/* Identity set the current transformation to the identity */
func (r *Ri) Identity() error {
	return r.writef("Identity")
}

/* Transform set the current transformation to the transformation transform */
func (r *Ri) Transform(transform RtMatrix) error {
	return r.writef("Transform", transform)
}

/* ConcatTransform concatenate the transformation transform onto the current transformation. */
func (r *Ri) ConcatTransform(transform RtMatrix) error {
	return r.writef("ConcatTransform", transform)
}

/* Perspective Concatenate a perspective transformation onto the current transformation. */
func (r *Ri) Perspective(fov RtFloat) error {
	return r.writef("Perspective", fov)
}

/* Translate Concatenate a translation onto the current transformation */
func (r *Ri) Translate(dx, dy, dz RtFloat) error {
	return r.writef("Translate", dx, dy, dz)
}

/* Rotate Concatenate a rotation of angle degrees about the given axis onto the current transformation */
func (r *Ri) Rotate(angle, dx, dy, dz RtFloat) error {
	return r.writef("Rotate", angle, dx, dy, dz)
}

/* Scale Concatenate a scaling onto the current transformation */
func (r *Ri) Scale(sx, sy, sz RtFloat) error {
	return r.writef("Scale", sx, sy, sz)
}

/* Skew Concatenate a skew onto the current transformation */
func (r *Ri) Skew(angle, dx1, dy1, dz1, dx2, dy2, dz2 RtFloat) error {
	return r.writef("Skew", RtFloatArray{angle, dx1, dy1, dz1, dx2, dy2, dz2})
}

/* CoordinateSystem This function marks the coordinate system defined by the current transformation
 * with the name space and saves it */
func (r *Ri) CoordinateSystem(name RtToken) error {
	return r.writef("CoordinateSystem", name)
}

/* CoordSysTransform This function replaces the current transformation matrix with the matrix that
 * forms the name coordinate system.
 */
func (r *Ri) CoordSysTransform(name RtToken) error {
	return r.writef("CoordSysTransform", name)
}

/* RtPoint * RiTransformPoints --- TODO */

/* TransformBegin Push the current transformation */
func (r *Ri) TransformBegin() error {
	defer func() { r.Depth(1) }()
	return r.writef("TransformBegin")
}

/* TransformEnd Pop the current transformation */
func (r *Ri) TransformEnd() error {
	r.Depth(-1)
	return r.writef("TransformEnd")
}
