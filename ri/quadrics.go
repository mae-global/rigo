package ri

/* Sphere */
func (r *Ri) Sphere(radius, zmin, zmax, thetamax RtFloat, parameterlist ...Rter) error {

	var out = []Rter{radius, zmin, zmax, thetamax, PARAMETERLIST}
	out = append(out, parameterlist...)

	return r.writef("Sphere", out...)
}

/* Cone */
func (r *Ri) Cone(height, radius, thetamax RtFloat, parameterlist ...Rter) error {

	var out = []Rter{height, radius, thetamax, PARAMETERLIST}
	out = append(out, parameterlist...)

	return r.writef("Cone", out...)
}

/* Cylinder */
func (r *Ri) Cylinder(radius, zmin, zmax, thetamax RtFloat, parameterlist ...Rter) error {

	var out = []Rter{radius, zmin, zmax, thetamax, PARAMETERLIST}
	out = append(out, parameterlist...)

	return r.writef("Cylinder", out...)
}

/* Hyperboloid */
func (r *Ri) Hyperboloid(point1, point2 RtPoint, thetamax RtFloat, parameterlist ...Rter) error {

	var out = []Rter{point1, point2, thetamax, PARAMETERLIST}
	out = append(out, parameterlist...)

	return r.writef("Hyperboloid", out...)
}

/* Disk */
func (r *Ri) Disk(height, radius, thetamax RtFloat, parameterlist ...Rter) error {

	var out = []Rter{height, radius, thetamax, PARAMETERLIST}
	out = append(out, parameterlist...)

	return r.writef("Disk", out...)
}

/* Torus */
func (r *Ri) Torus(majorradius, minorradius, phimin, phimax, thetamax RtFloat, parameterlist ...Rter) error {

	var out = []Rter{majorradius, minorradius, phimin, phimax, thetamax, PARAMETERLIST}
	out = append(out, parameterlist...)

	return r.writef("Torus", out...)
}
