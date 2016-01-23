package ri

/* SolidBegin */
func (ctx *Context) SolidBegin(operation RtToken) error {
	defer func() { ctx.depth++ }()
	return ctx.writef("SolidBegin", operation)
}

/* SolidEnd */
func (ctx *Context) SolidEnd() error {
	ctx.depth--
	return ctx.writef("SolidEnd")
}

/* ObjectBegin */
func (ctx *Context) ObjectBegin() (RtObjectHandle, error) {
	oh := RtObjectHandle(ctx.objects)
	ctx.objects++
	defer func() { ctx.depth++ }()
	return oh, ctx.writef("ObjectBegin", oh)
}

/* ObjectEnd */
func (ctx *Context) ObjectEnd() error {
	ctx.depth--
	return ctx.writef("ObjectEnd")
}

/* ObjectInstance */
func (ctx *Context) ObjectInstance(handle RtObjectHandle) error {
	if uint(handle) >= ctx.objects {
		return ErrBadHandle
	}
	return ctx.writef("ObjectInstance", handle)
}
