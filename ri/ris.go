package ri

/* RIS Procedures -- https://renderman.pixar.com/resources/current/RenderMan/risProcedures.html */

var dashed = RtShaderHandle("-")

type ShaderWriter interface {
	Write() (RtName, RtShaderHandle, []Rter, []Rter, []Rter)
}

/* Integrator procedure is used to specify an integrator. RIS-Mode only. */
func (r *Ri) Integrator(name RtToken, handle RtShaderHandle, parameterlist ...Rter) error {

	integrator := r.RiContexter.Shader(handle)

	if integrator == nil {

		list := []Rter{name, handle, PARAMETERLIST}
		list = append(list, parameterlist...)

		return r.writef("Integrator", list...)
	}

	n, h, args, params, values := integrator.Write()
	list := []Rter{n, h}
	list = append(list, args...)
	list = append(list, PARAMETERLIST)

	oparams := make([]Rter, 0)
	ovalues := make([]Rter, 0)

	/* FIXME: convert to unmix() */
	flipflop := false
	for _, param := range parameterlist {
		if !flipflop {
			oparams = append(oparams, param)
		} else {
			ovalues = append(ovalues, param)
		}
		flipflop = !flipflop
	}

	for i, param := range oparams {
		for j, p := range params {
			if param == p {
				values[j] = ovalues[i]
			}
		}
	}

	for i, param := range params {
		list = append(list, param)
		list = append(list, values[i])
	}

	return r.writef("Integrator", list...)
}

/* Bxdf is used to assign a Bxdf to a surface. Dxdfs take precedence over Surface when an integratir is
 * active but are ignored whren no itegrator has been specified. */
func (r *Ri) Bxdf(name RtToken, handle RtShaderHandle, parameterlist ...Rter) error {

	/* We take handle, and if we find the shader in the cache of the contexter then we use it as the
	 * basis for the output. All the parameterlist are thus applied ontop of the basis.
	 */

	bxdf := r.RiContexter.Shader(handle)

	if bxdf == nil {

		/* otherwise we parse the attributes as normal */
		list := []Rter{name, handle, PARAMETERLIST}
		list = append(list, parameterlist...)

		return r.writef("Bxdf", list...)
	}

	_, h, args, params, values := bxdf.Write()

	list := []Rter{}
	list = append(list, args...)
	list = append(list, h)
	list = append(list, PARAMETERLIST)

	oparams := make([]Rter, 0)
	ovalues := make([]Rter, 0)

	/* FIXME: convert to unmix() */
	flipflop := false
	for _, param := range parameterlist {
		if !flipflop {
			oparams = append(oparams, param)
		} else {
			ovalues = append(ovalues, param)
		}
		flipflop = !flipflop
	}

	for i, param := range oparams {
		for j, p := range params {
			if param == p {
				values[j] = ovalues[i]
			}
		}
	}

	for i, param := range params {
		list = append(list, param)
		list = append(list, values[i])
	}

	return r.writef("Bxdf", list...)
}

/* Pattern is used to wire in textures and patterns. */
func (r *Ri) Pattern(name RtToken, handle RtShaderHandle, parameterlist ...Rter) error {

	pattern := r.RiContexter.Shader(handle)

	if pattern == nil {
		list := []Rter{name, handle, PARAMETERLIST}
		list = append(list, parameterlist...)
		return r.writef("Pattern", list...)
	}

	n, h, args, params, values := pattern.Write()
	list := []Rter{n, h}
	list = append(list, args...)
	list = append(list, PARAMETERLIST)

	oparams := make([]Rter, 0)
	ovalues := make([]Rter, 0)

	/* FIXME: convert to unmix() */
	flipflop := false
	for _, param := range parameterlist {
		if !flipflop {
			oparams = append(oparams, param)
		} else {
			ovalues = append(ovalues, param)
		}
		flipflop = !flipflop
	}

	for i, param := range oparams {
		for j, p := range params {
			if param == p {
				values[j] = ovalues[i]
			}
		}
	}

	for i, param := range params {
		list = append(list, param)
		list = append(list, values[i])
	}

	return r.writef("Pattern", list...)
}
