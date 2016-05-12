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

func Mix(params, values []Rter) []Rter {

	out := make([]Rter, 0)

	for i := 0; i < len(params); i++ {

		out = append(out, params[i])
		out = append(out, values[i])
	}
	return out
}

func Unmix(list []Rter) ([]Rter, []Rter) {

	params := make([]Rter, 0)
	values := make([]Rter, 0)

	flipflop := false

	for _, param := range list {
		if !flipflop {
			params = append(params, param)
		} else {
			values = append(values, param)
		}
		flipflop = !flipflop
	}

	return params, values
}

type PrototypeArgument struct {
	Type string
	Example Rter
	Name string
}

func (p *PrototypeArgument) String() string {
	return fmt.Sprintf("name=\"%s\", type=\"%s\", example=%s",p.Name,p.Type,p.Example)
} 

type PrototypeInformation struct {
	Name RtName 
	Arguments []*PrototypeArgument
	Parameterlist bool
}

func (p *PrototypeInformation) String() string {
	out := fmt.Sprintf("PrototypeInformation \"%s\", parameterlist=%v\n",p.Name,p.Parameterlist)
	if len(p.Arguments) > 0 {
		out += fmt.Sprintf("\t%d arguments\n",len(p.Arguments))
		for i,arg := range p.Arguments {
			out += fmt.Sprintf("\t\t[%03d] \"%s\" -- type=\"%s\", example=%s\n",i,arg.Name,arg.Type,arg.Example)
		}
	}
	return out
}
	

func ParsePrototype(stream string) *PrototypeInformation {

	list := strings.Split(stream," ")
	if len(list) == 0 {
		return &PrototypeInformation{Name:RtName(stream)}
	}

	proto := &PrototypeInformation{}
	
	proto.Name = RtName(list[0])
	proto.Arguments = make([]*PrototypeArgument,0)	

	var arg *PrototypeArgument

	/* example :- "Shader token name token handle ..." */	
	var r Rter
	for i := 1; i < len(list); i++ {
		r = nil

		switch list[i] {
			case "string":
				r = RtString("string")
			break
			case "string[]":
				r = RtStringArray{}
			break
			case "float":
				r = RtFloat(1)
			break
			case "float[]":
				r = RtFloatArray{1,2,3}
			break
			case "int":
				r = RtInt(1)
			break
			case "int[]":
				r = RtIntArray{1,2,3}
			break
			case "token":
				r = RtToken("name")
			break
			case "lighthandle":
				r = RtLightHandle("light")
			break
			case "objecthandle":
				r = RtObjectHandle("object")
			break
			case "filterfunc":
				r = BesselFilter
			break
			case "boolean":
				r = RtBoolean(true)
			break
			case "color":
				r = RtColor{1,1,1}
			break
			case "point":
				r = RtPoint{1,1,1}
			break
			case "point[]":
				r = RtPointArray{}
			break
			case "basis":
				r = RtBasis{}
			break
			case "bound":
				r = RtBound{}
			break
			case "matrix":
				r = RtMatrix{}
			break
			case "pointer":
				r = RtStringArray{}
			break
			case "procsubdivfunc":
				r = ProcDelayedReadArchive
			break
			case "procfreefunc":
				r = ProcFree
			break
			case "proc2subdivfunc":
				r = Proc2DelayedReadArchive
			break
			case "...":
				r = PARAMETERLIST
			break
		}

		if r == nil {
			/* then it is a name */
			fmt.Printf("list[i]:proto.Name=\"%s\" = %s\n",proto.Name,list[i])
			arg.Name = list[i]
		} else {

			if arg != nil && len(arg.Type) > 0 {
				proto.Arguments = append(proto.Arguments,arg)
				arg = new(PrototypeArgument)
			}

			if r == PARAMETERLIST {
				
				proto.Parameterlist = true
				continue
			}

			arg = new(PrototypeArgument)
			arg.Type = list[i]
			arg.Example = r
		}
	}
	if arg != nil && len(arg.Type) > 0 {
		proto.Arguments = append(proto.Arguments,arg)
	}	
	
	return proto
}




