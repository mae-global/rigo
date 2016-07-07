package ri

import (
	"fmt"
)

/* RIS Procedures -- https://renderman.pixar.com/resources/current/RenderMan/risProcedures.html */

var dashed = RtShaderHandle("-")

type ShaderWriter interface {
	Write() (RtName, RtShaderHandle, []Rter, []Rter, []Rter)
}

/* Integrator procedure is used to specify an integrator. RIS-Mode only. */
func (r *Ri) Integrator(name RtToken, handle RtShaderHandle, parameterlist ...Rter) error {

	list := sort(r.RiContexter,name,handle,parameterlist...)

	return r.writef("Integrator", list...)
}

/* Bxdf is used to assign a Bxdf to a surface. Dxdfs take precedence over Surface when an integratir is
 * active but are ignored whren no itegrator has been specified. */
func (r *Ri) Bxdf(name RtToken, handle RtShaderHandle, parameterlist ...Rter) error {

	/* We take handle, and if we find the shader in the cache of the contexter then we use it as the
	 * basis for the output. All the parameterlist are thus applied ontop of the basis.
	 */
	list := sort(r.RiContexter,name,handle,parameterlist...)
	
	return r.writef("Bxdf", list...)
}

/* Pattern is used to wire in textures and patterns. */
func (r *Ri) Pattern(name RtToken, handle RtShaderHandle, parameterlist ...Rter) error {

	list := sort(r.RiContexter,name,handle,parameterlist...)

	return r.writef("Pattern", list...)
}

func sort(ctx RiContexter,name RtToken, handle RtShaderHandle, parameterlist ...Rter) []Rter {

	shader := ctx.Shader(handle)

	if shader == nil {
		list := []Rter{name, handle, PARAMETERLIST}
		list = append(list, parameterlist...)
		return list
	}

	_, h, args, params, values := shader.Write()

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
		fmt.Printf("looking for %s to override\n\t%d params to search through\n",param,len(params))
		found := false
		for j, p := range params {
			if param == p {
				fmt.Printf("\tfound it -- overidden value as %s from %s\n",ovalues[i],values[j])
				values[j] = ovalues[i]
				found = true
				break
			}
		}

		if !found {
			fmt.Printf("\tadded\n")
			params = append(params,param)
			values = append(values,ovalues[i])
		}
	}

	for i, param := range params {
		list = append(list, param)
		list = append(list, values[i])
	}

	return list
}



