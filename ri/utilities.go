package ri

import (
	"fmt"
	"strconv"
	"strings"
)

/* FIXME */
func ClassTypeNameCount(t RtToken) (RtToken, RtToken, RtToken, RtInt) {
	p := strings.Split(strings.TrimSpace(string(t)), " ")
	if len(p) == 1 {
		return "", "", t, 1
	}

	count := 1

	ty := RtToken("")

	if len(p) == 2 { /* should be just type and name */

		ty = RtToken(p[0])

		if strings.Contains(p[0], "[") && strings.Contains(p[0], "]") {

			parts := strings.Split(p[0], "[")
			out := strings.Replace(parts[1], "]", "", -1)
			if c, err := strconv.Atoi(out); err == nil {
				count = c
			}
			ty = RtToken(parts[0])
		}
		return "", ty, RtToken(p[1]), RtInt(count)
	}

	ty = RtToken(p[1])

	if strings.Contains(p[1], "[") && strings.Contains(p[1], "]") {
		parts := strings.Split(p[1], "[")
		out := strings.Replace(parts[1], "]", "", -1)
		if c, err := strconv.Atoi(out); err == nil {
			count = c
		}
		ty = RtToken(parts[0])
	}
	return RtToken(p[0]), ty, RtToken(p[2]), RtInt(count)
}

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

func mix(tokens, values []Rter) []Rter {
	out := make([]Rter, 0)
	for i := 0; i < len(tokens); i++ {
		out = append(out, tokens[i])
		out = append(out, values[i])
	}
	return out
}

func unmix(params []Rter) ([]Rter, []Rter) {
	tokens := make([]Rter, 0)
	values := make([]Rter, 0)

	for i := 0; i < len(params); i++ {
		if i%2 == 0 {
			tokens = append(tokens, params[i])
		} else {
			values = append(values, params[i])
		}
	}
	return tokens, values
}
