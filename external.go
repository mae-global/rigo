package ri

/* MakeTexture Convert an image in a standard picture file whise name is picturename */
func (ctx *Context) MakeTexture(picturename,texturename RtString,swrap,twrap RtToken,filterfunc RtFilterFunc,swidth,twidth RtFloat,parameterlist ...Rter) error {

	var out = []Rter{picturename,texturename,swrap,twrap,filterfunc,swidth,twidth}
	out = append(out,parameterlist...)

	return ctx.writef("MakeTexture",out...)
}

/* MakeLatLongEnvironment */
func (ctx *Context) MakeLatLongEnvironment(picturename,texturename RtString,filterfunc RtFilterFunc,swidth,twidth RtFloat,parameterlist ...Rter) error {

	var out = []Rter{picturename,texturename,filterfunc,swidth,twidth}
	out = append(out,parameterlist...)

	return ctx.writef("MakeLatLongEnvironment",out...)
}


