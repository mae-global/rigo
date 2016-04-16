package ri

import (
	"fmt"
	"strings"
)

func serialise(parameterlist ...Rter) (string, error) {
	out := ""
	if len(parameterlist) > 0 {
		if len(parameterlist)%2 != 0 {
			return "", ErrBadParamlist
		}
		for i, p := range parameterlist {
			out += p.Serialise()
			if i < (len(parameterlist) - 1) {
				out += " "
			}
		}
	}

	return out, nil
}

func serialiseToString(parameterlist ...Rter) string {
	out := ""
	for i, p := range parameterlist {
		out += p.Serialise()
		if i < (len(parameterlist) - 1) {
			out += " "
		}
	}
	return out
}

func parseAnnotations(annotations ...RtAnnotation) []Rter {
	out := make([]Rter, 0)
	for i, a := range annotations {
		if i > 0 {
			out = append(out, RtName(string(a)))
			continue
		}
		out = append(out, a)
	}
	return out
}

func Serialise(list []Rter) string {
	if len(list) == 0 {
		return ""
	}
	out := ""
	for i, p := range list {
		out += p.Serialise()
		if i < len(list)-1 {
			out += " "
		}
	}
	return strings.TrimSpace(out)
}

func reduce(f RtFloat) string {

	str := fmt.Sprintf("%f", f)
	s := 0
	neg := false
	for i, c := range str {
		if c != '0' {
			if c == '-' {
				neg = true
				continue
			}
			s = i
			break
		}
		if c == '.' {
			break
		}
	}

	e := 0
	for i := len(str) - 1; i >= 0; i-- {
		if str[i] != '0' {
			e = i + 1
			break
		}
		if str[i] == '.' {
			break
		}
	}

	str = str[s:e]
	if str == "." {
		str = "0"
	}
	if neg {
		str = "-" + str
	}
	if str[len(str)-1] == '.' {
		str = str[:len(str)-1]
	}

	return str
}

func reducev(fv []RtFloat) string {
	out := ""
	for i, f := range fv {
		out += reduce(f)
		if i < len(fv)-1 {
			out += " "
		}
	}
	return out
}
