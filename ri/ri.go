package ri



const USEDEBUG = false

type RiContexter interface {
	Write(RtName, []Rter, []Rter, []Rter) error
	OpenRaw(RtToken) (ArchiveWriter, error)
	CloseRaw(RtToken) error
	//Depth(int)
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

	nlist := make([]Rter, len(list))

	for i, r := range list {
		ar := r
		if i%2 == 0 {
			_, ok := r.(RtToken)
			if !ok {
				return ErrBadArgument
			}

			//	cl,ty,nam,n := ClassTypeNameCount(t)
			/* TODO: parse token, lookup class and type then check inputs of values */
			//	fmt.Printf("Debug,Ri.writef token, class=%s,type=%s,name=%s,count=%d\n",cl,ty,nam,n)

		} else {
			/* convert all outputs of the parameterlist in
			 * array types as per RIB standard */
			if a, ok := r.(RtString); ok {
				ar = RtStringArray{a}
			}
			if a, ok := r.(RtFloat); ok {
				ar = RtFloatArray{a}
			}
			if a, ok := r.(RtInt); ok {
				ar = RtIntArray{a}
			}
			if a, ok := r.(RtToken); ok {
				ar = RtTokenArray{a}
			}
			if a, ok := r.(RtPoint); ok {
				ar = RtPointArray{a}
			}
		}

		nlist[i] = ar
	}

	/* tokens to values is unbalanced */
	if len(nlist)%2 != 0 {
		return ErrBadArgument
	}

	if USEDEBUG && len(nlist) > 0 {
		args = append(args, DEBUGBARRIER)
	}


	params,values := Unmix(nlist)

	return r.Write(name, args, params,values)
}


