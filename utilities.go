package ri


func serialise(parameterlist ...Rter) (string,error) {
	out := ""
	if len(parameterlist) > 0 {
		if len(parameterlist) % 2 != 0 {
			return "",ErrBadParamlist
		}
		for i,p := range parameterlist {
			out += p.Serialise()
			if i < (len(parameterlist) - 1) {
				out += " "
			}
		}
	}

	return out,nil
}

func serialiseToString(parameterlist ...Rter) string {
	out := ""
	for i,p := range parameterlist {
		out += p.Serialise()
		if i < (len(parameterlist) - 1) {
			out += " "
		}
	}
	return out
} 

func parseAnnotations(annotations ...RtAnnotation) []Rter {
	out := make([]Rter,0)
	for i,a := range annotations {
		if i > 0 {
			out = append(out,RtName(string(a)))
			continue
		}
		out = append(out,a)
	}
	return out
}
