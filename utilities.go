package ri


func serialise(parameterlist ...Rter) (string,error) {
	out := ""
	if len(parameterlist) > 0 {
		if len(parameterlist) % 2 != 0 {
			return "",ErrParameterlistMismatch
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

func serialise_s(parameterlist ...Rter) string {
	out := ""
	for i,p := range parameterlist {
		out += p.Serialise()
		if i < (len(parameterlist) - 1) {
			out += " "
		}
	}
	return out
} 
