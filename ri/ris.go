package ri

/* RIS Procedures -- https://renderman.pixar.com/resources/current/RenderMan/risProcedures.html */

/* Integrator procedure is used to specify an integrator. RIS-Mode only. */
func (r *Ri) Integrator(name,handle RtToken,parameterlist... Rter) error {
	list := []Rter{name,handle,PARAMETERLIST}
	list = append(list,parameterlist...)
	return r.writef("Integrator",list...)
}

/* Bxdf is used to assign a Bxdf to a surface. Dxdfs take precedence over Surface when an integratir is
 * active but are ignored whren no itegrator has been specified. */
func (r *Ri) Bxdf(name,handle RtToken,parameterlist... Rter) error {

	/* We take handle, and if we find the shader in the cache of the contexter then we use it as the 
	 * basis for the output. All the parameterlist are thus applied ontop of the basis.
	 */



	/* otherwise we parse the attributes as normal */

	list := []Rter{name,handle,PARAMETERLIST}
	list = append(list,parameterlist...)
	return r.writef("Bxdf",list...)

}

/* Pattern is used to wire in textures and patterns. */
func (r *Ri) Pattern(name,handle RtToken,parameterlist... Rter) error {

	list := []Rter{name,handle,PARAMETERLIST}
	list = append(list,parameterlist...)
	return r.writef("Pattern",list...)
}



