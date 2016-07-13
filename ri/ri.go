package ri

import (
	"fmt"
)

const USEDEBUG = false /* FIXME: fails tests when true */

type RiContexter interface {
	Write(RtName, []Rter, []Rter, []Rter) error
	OpenRaw(RtToken) (ArchiveWriter, error)
	CloseRaw(RtToken) error
	LightHandle() (RtLightHandle, error)
	CheckLightHandle(RtLightHandle) error
	ObjectHandle() (RtObjectHandle, error)
	CheckObjectHandle(RtObjectHandle) error
	ShaderHandle() (RtShaderHandle, error)
	CheckShaderHandle(RtShaderHandle) error

	Shader(RtShaderHandle) ShaderWriter
}


/* Ri is the main interface */
type Ri struct {
	RiContexter
	RiDictionarer
}

func (r *Ri) BloomFilter() *BloomFilter { return RiBloomFilter() }

/* User special func for client libraries to write to */
func (r *Ri) User(reader RterReader) error {
	if reader == nil {
		return ErrBadArgument
	}

	name, args, tokens, values := reader.ReadFrom()
	out := make([]Rter, len(args))
	copy(out, args)
	out = append(out, PARAMETERLIST)
	out = append(out, Mix(tokens, values)...)

	return r.writef(name, out...)
}

func (r *Ri) WriteTo(name RtName, args, tokens, values []Rter) error {

	out := make([]Rter, 0)
	out = append(out, args...)
	out = append(out, PARAMETERLIST)
	out = append(out, Mix(tokens, values)...)

	return r.writef(name, out...)
}

func (r *Ri) writef(name RtName, parameterlist ...Rter) error {
	if r.RiContexter == nil {
		return ErrProtocolBotch
	}

	para := -1
	/* find the actual parameterlist */
	for i, r := range parameterlist {
		if t, ok := r.(RtToken); ok {
			if t == PARAMETERLIST {
				para = i
				break
			}
		}
	}
	var args []Rter
	var list []Rter

	if para == -1 {
		args = parameterlist
	} else {
		args = make([]Rter, para)
		copy(args, parameterlist[:para])
		list = make([]Rter, len(parameterlist)-(para+1))
		copy(list, parameterlist[para+1:])
	}

	/* expand any RterArray present */
	nlist := make([]Rter,0)
	for _,ele := range list {
		if ele.Type() == "rter[]" {
			if attr,ok := ele.(RterArray); ok {
				nlist = append(nlist,attr...)
			}
		} else {
			nlist = append(nlist,ele)
		}
	}

	list = make([]Rter,len(nlist))
	copy(list,nlist)

	/* expand any Rtattr (attributes) present */
	nlist = make([]Rter,0)
	for _,ele := range list {
		if ele.Type() == "attribute" {
			if attr,ok := ele.(Rtattr); ok {
				name,value := attr.Break()
				nlist = append(nlist,name)
				nlist = append(nlist,value)
				continue
			}	else {
				return ErrBadArgument
			}
		}
		
		nlist = append(nlist,ele)
	}

	list = make([]Rter,len(nlist))
	copy(list,nlist)

	nlist = make([]Rter, len(list))

	
	var class RtToken
	var typeof RtToken
	var count RtInt
	var nameofparam RtToken


	params,values := Unmix(list)
	if len(params) != len(values) {
			return ErrBadArgument
	}

	nvalues := make([]Rter,0)
	params2 := make([]RtToken,len(params))

	var dict RiDictionarer
	if r.RiDictionarer != nil {
		dict = r.RiDictionarer
	}


	for i,param := range params {
			token,ok := param.(RtToken)
			if !ok {
				return ErrBadArgument
			}
	
			class,typeof,nameofparam,count = ClassTypeNameCount(token) /* TODO: need to look up the token class and type via global dictionary */
			/* if class,typeof are not defined then should be in the dictionary, if not then error
			 * otherwise if at least typeof then defined inline, class defaults to constant 
			 */
			params2[i] = token			

			if dict != nil {
				/* check if the token has been declared via ri.Declare(...)
         * if so then replace the presumed inline specification 
				 */
				spec := dict.Lookup(nameofparam)
				if spec != EmptyToken {
					class,typeof,_,count = ClassTypeNameCount(spec)
					params2[i] = spec				
				}
			}

			value := values[i]
			
			if count == 1 { /* singular */
	
				vtype := value.Type()
				/* if some form of handle switch the type to string */
				if vtype == "lighthandle" || vtype == "shaderhandle" || vtype == "objecthandle" {
					vtype = "string"
				}

				if string(typeof) != vtype {
					/* check for empty typeof information from the 
					 * token, if so then use the type from the actual value.
           * NOTE, this should be an error?!
					 */					
					if string(class) == "reference" {
						/* if the class is reference then the type is not
             * indictive of what the value is, instead we
						 * change it to the expected string type.
						 */
						typeof = RtToken("string")
					} else {
						return fmt.Errorf("%s -- %v",nameofparam,ErrBadArgument)
					}
				}

				var array Rter

				/* convert to array of that type */
				switch value.Type() {
					case "float":
						if v,ok := value.(RtFloat); ok {
							array = RtFloatArray{v}
						
						} else {
							return fmt.Errorf("%s -- %v, expecting float",nameofparam,ErrBadArgument)
						}
					break
					case "int":
						if v,ok := value.(RtInt); ok {
							array = RtIntArray{v}
					
						} else {
							return fmt.Errorf("%s -- %v, expecting int",nameofparam,ErrBadArgument)
						}
					break
					case "string":
						if v,ok := value.(RtString); ok {
							array = RtStringArray{v}
							
						} else {
							return fmt.Errorf("%s -- %v, expecting string",nameofparam,ErrBadArgument)
						}
					break
					case "point":
						if v,ok := value.(RtPoint); ok {
							array = RtPointArray{v}

						} else {
							return fmt.Errorf("%s -- %v, expecting point",nameofparam,ErrBadArgument)
						}
					break					
					default:
						array = value
					break
				}

				nvalues = append(nvalues,array)

			} else {
				if string(typeof) + "[]" != value.Type() {
					return fmt.Errorf("%s -- %v, expecting array of %s[] but was %s",nameofparam,ErrBadArgument,string(typeof),value.Type())
				}

				switch value.Type() {
					case "float[]":
						if array,ok := value.(RtFloatArray); ok {
							if int(count) != len(array) {
								return ErrBadArgument
							}
						} else {
							return ErrBadArgument
						}
					break
					case "int[]":
						if array,ok := value.(RtIntArray); ok {
							if int(count) != len(array) {
								return ErrBadArgument
							}
						} else {
							return ErrBadArgument
						}
					break
					case "string[]":
						if array,ok := value.(RtStringArray); ok {
							if int(count) != len(array) {
								return ErrBadArgument
							}
						} else {
							return ErrBadArgument
						}
					break
					case "point[]":
						if array,ok := value.(RtPointArray); ok {
							if int(count) != len(array) {
								return ErrBadArgument
							}
						} else {
							return ErrBadArgument
						}
					break
				}

				nvalues = append(nvalues,value)
			}
	}


	if USEDEBUG && len(params) > 0 {
		args = append(args, DEBUGBARRIER)
	}


	return r.Write(name, args, params, nvalues) /* TODO, do something with params2 */
}


