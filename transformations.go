package ri

/* Identity set the current transformation to the identity */
func (ctx *Context) Identity() error {
	return ctx.writef("Identity")
}

/* Transform set the current transformation to the transformation transform */
func (ctx *Context) Transform(transform RtMatrix) error {
	return ctx.writef("Transform", transform)
}

/* ConcatTransform concatenate the transformation transform onto the current transformation. */
func (ctx *Context) ConcatTransform(transform RtMatrix) error {
	return ctx.writef("ConcatTransform", transform)
}

/* Perspective Concatenate a perspective transformation onto the current transformation. */
func (ctx *Context) Perspective(fov RtFloat) error {
	return ctx.writef("Perspecitve", fov)
}

/* Translate Concatenate a translation onto the current transformation */
func (ctx *Context) Translate(dx, dy, dz RtFloat) error {
	return ctx.writef("Translate", dx, dy, dz)
}

/* Rotate Concatenate a rotation of angle degrees about the given axis onto the current transformation */
func (ctx *Context) Rotate(angle, dx, dy, dz RtFloat) error {
	return ctx.writef("Rotate", angle, dx, dy, dz)
}

/* Scale Concatenate a scaling onto the current transformation */
func (ctx *Context) Scale(sx, sy, sz RtFloat) error {
	return ctx.writef("Scale", sx, sy, sz)
}

/* Skew Concatenate a skew onto the current transformation */
func (ctx *Context) Skew(angle, dx1, dy1, dz1, dx2, dy2, dz2 RtFloat) error {
	return ctx.writef("Skew", RtFloatArray{angle, dx1, dy1, dz1, dx2, dy2, dz2})
}

/* CoordinateSystem This function marks the coordinate system defined by the current transformation
 * with the name space and saves it */
func (ctx *Context) CoordinateSystem(name RtToken) error {
	return ctx.writef("CoordinateSystem", name)
}

/* CoordSysTransform This function replaces the current transformation matrix with the matrix that
 * forms the name coordinate system.
 */
func (ctx *Context) CoordSysTransform(name RtToken) error {
	return ctx.writef("CoordSysTransform", name)
}

/* RtPoint * RiTransformPoints --- TODO */

/* TransformBegin Push the current transformation */
func (ctx *Context) TransformBegin() error {
	defer func() { ctx.depth++ }()
	return ctx.writef("TransformBegin")
}

/* TransformEnd Pop the current transformation */
func (ctx *Context) TransformEnd() error {
	ctx.depth--
	return ctx.writef("TransformEnd")
}
