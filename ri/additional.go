package ri

/* Camera */
func (r *Ri) Camera(name RtToken, parameterlist ...Rter) error {

	list := []Rter{name,PARAMETERLIST}
	list = append(list,parameterlist...)

	return r.writef("Camera", list...)
}

/* Deformation */
func (r *Ri) Deformation(name RtToken, parameterlist ...Rter) error {

	list := []Rter{name, PARAMETERLIST}
	list = append(list, parameterlist...)

	return r.writef("Deformation", list...)
}

/* DisplayChannel */
func (r *Ri) DisplayChannel(channel RtToken, parameterlist ...Rter) error {

	list := []Rter{channel, PARAMETERLIST}
	list = append(list, parameterlist...)

	return r.writef("DisplayChannel", list...)
}

/* EditAttributeBegin */
func (r *Ri) EditAttributeBegin(name RtToken) error {
	return r.writef("EditAttributeBegin", name)
}

/* EditAttributeEnd */
func (r *Ri) EditAttributeEnd() error {
	return r.writef("EditAttributeEnd")
}

/* EditBegin */
func (r *Ri) EditBegin(name RtToken, parameterlist ...Rter) error {

	list := []Rter{name, PARAMETERLIST}
	list = append(list, parameterlist...)

	return r.writef("EditBegin", list...)
}

/* EditEnd */
func (r *Ri) EditEnd() error {
	return r.writef("EditEnd")
}

func (r *Ri) EditWorldBegin(name RtToken, parameterlist ...Rter) error {

	list := []Rter{name, PARAMETERLIST}
	list = append(list, parameterlist...)

	return r.writef("EditWorldBegin", list...)
}

/* EditWorldEnd */
func (r *Ri) EditWorldEnd() error {

	return r.writef("EditWorldEnd")
}

/* EnableLightFilter */
func (r *Ri) EnableLightFilter(light RtLightHandle, filter RtToken, onoff RtBoolean) error {

	list := []Rter{light, filter, onoff, PARAMETERLIST}

	return r.writef("EnableLightFilter", list...)
}



/* HierarchicalSubdivisionMesh */
func (r *Ri) HierarchicalSubdivisionMesh(mask RtToken, nf RtInt, nverts, verts RtIntArray, nt RtInt, tags RtTokenArray,
	nargs, intargs RtIntArray, floatargs RtFloatArray, stringargs RtTokenArray, parameterlist ...Rter) error {

	list := []Rter{mask, nf, nverts, verts, nt, tags, nargs, intargs, floatargs, stringargs, PARAMETERLIST}
	list = append(list, parameterlist...)

	return r.writef("HierarchicalSubdivisionMesh", list...)
}

/* LightFilter */
func (r *Ri) LightFilter(name RtToken, handle RtToken, parameterlist ...Rter) error {

	list := []Rter{name, handle, PARAMETERLIST}
	list = append(list, parameterlist...)

	return r.writef("LightFilter", list...)
}

/* MakeBrickMap */
func (r *Ri) MakeBrickMap(nptcs RtInt, ptcs RtStringArray, bkm RtString, parameterlist ...Rter) error {

	list := []Rter{nptcs, ptcs, bkm, PARAMETERLIST}
	list = append(list, parameterlist...)

	return r.writef("MakeBrickMap", list...)
}

/* MakeBump */
func (r *Ri) MakeBump(pic, text RtString, swrap, twrap RtToken, filterfunc RtFilterFunc, swidth, twidth RtFloat, parameterlist ...Rter) error {

	list := []Rter{pic, text, swrap, twrap, filterfunc, swidth, twidth, PARAMETERLIST}
	list = append(list, parameterlist...)

	return r.writef("MakeBump", list...)
}

/* Paraboloid */
func (r *Ri) Paraboloid(radius, zmin, zmax, tmax RtFloat, parameterlist ...Rter) error {

	list := []Rter{radius, zmin, zmax, tmax, PARAMETERLIST}
	list = append(list, parameterlist...)

	return r.writef("Paraboloid", list...)
}

/* Procedural2 */
func (r *Ri) Procedural2(subdividefunc RtProc2SubdivFunc, boundfunc RtProc2BoundFunc, parameterlist ...Rter) error {

	list := []Rter{subdividefunc, boundfunc, PARAMETERLIST}
	list = append(list, parameterlist...)

	return r.writef("Procedural2", list...)
}

/* ResourceBegin */
func (r *Ri) ResourceBegin() error {
	return r.writef("ResourceBegin")
}

/* ResourceEnd */
func (r *Ri) ResourceEnd() error {
	return r.writef("ResourceEnd")
}

/* ScopedCoordinateSystem */
func (r *Ri) ScopedCoordinateSystem(name RtString) error {

	list := []Rter{name, PARAMETERLIST}
	return r.writef("ScopedCoordinateSystem", list...)
}

/* Shader */
func (r *Ri) Shader(name RtToken, handle RtToken, parameterlist ...Rter) error {

	list := []Rter{name, handle, PARAMETERLIST}
	list = append(list, parameterlist...)

	return r.writef("Shader", list...)
}

/* System */
func (r *Ri) System(name RtString) error {

	list := []Rter{name, PARAMETERLIST}
	return r.writef("System", list...)
}

/* FIXME
func (r *Ri) RiTransformPoints(RtToken fromspace,
                               RtToken tospace, RtInt n, RtPoint * points) ([]RtPoint,error) {

}
*/

/* VPAtmosphere */
func (r *Ri) VPAtmosphere(name RtToken, parameterlist ...Rter) error {

	list := []Rter{name, PARAMETERLIST}
	list = append(list, parameterlist...)

	return r.writef("VPAtmosphere", list...)
}

/* VPInterior */
func (r *Ri) VPInterior(name RtToken, parameterlist ...Rter) error {

	list := []Rter{name, PARAMETERLIST}
	list = append(list, parameterlist...)

	return r.writef("VPInterior", list...)
}

/* VPSurface */
func (r *Ri) VPSurface(name RtToken, parameterlist ...Rter) error {

	list := []Rter{name, PARAMETERLIST}
	list = append(list, parameterlist...)

	return r.writef("VPSurface", list...)
}

/* Volume */
func (r *Ri) Volume(typeof RtToken, bound RtBound, dimensions RtIntArray, parameterlist ...Rter) error {

	list := []Rter{typeof, bound, dimensions, PARAMETERLIST}
	list = append(list, parameterlist...)

	return r.writef("Volume", list...)
}

/* VolumePixelSamples */
func (r *Ri) VolumePixelSamples(x RtFloat, y RtFloat) error {

	list := []Rter{x, y, PARAMETERLIST}
	return r.writef("VolumePixelSamples", list...)
}
