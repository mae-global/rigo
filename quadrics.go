package ri

/* Sphere */
func (ctx *Context) Sphere(radius,zmin,zmax,thetamax RtFloat,parameterlist ...Rter) error {
	
	var out = []Rter{radius,zmin,zmax,thetamax}
	out = append(out,parameterlist...)

	return ctx.writef("Sphere",out...)
}

/* Cone */
func (ctx *Context) Cone(height,radius,thetamax RtFloat,parameterlist ...Rter) error {
	
	var out = []Rter{height,radius,thetamax}
	out = append(out,parameterlist...)

	return ctx.writef("Cone",out...)
}
	
/* Cylinder */
func (ctx *Context) Cylinder(radius,zmin,zmax,thetamax RtFloat,parameterlist ...Rter) error {

	var out = []Rter{radius,zmin,zmax,thetamax}
	out = append(out,parameterlist...)

	return ctx.writef("Cylinder",out...)
}

/* Hyperboloid */
func (ctx *Context) Hyperboloid(point1,point2 RtPoint,thetamax RtFloat,parameterlist ...Rter) error {

	var out = []Rter{point1,point2,thetamax}
	out = append(out,parameterlist...)

	return ctx.writef("Hyperboloid",out...)
}

/* Disk */
func (ctx *Context) Disk(height,radius,thetamax RtFloat,parameterlist ...Rter) error {

	var out = []Rter{height,radius,thetamax}
	out = append(out,parameterlist...)

	return ctx.writef("Disk",out...)
}

/* Torus */
func (ctx *Context) Torus(majorradius,minorradius,phimin,phimax,thetamax RtFloat,parameterlist ...Rter) error {

	var out = []Rter{majorradius,minorradius,phimin,phimax,thetamax}
	out = append(out,parameterlist...)

	return ctx.writef("Torus",out...)
}


