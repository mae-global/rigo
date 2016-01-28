package ri

type Contexter interface {
	Write(RtName,[]Rter) error
	Depth(int)
	LightHandle() (RtLightHandle,error)
	CheckLightHandle(RtLightHandle) error
	ObjectHandle() (RtObjectHandle,error)
	CheckObjectHandle(RtObjectHandle) error
}


/* Ri is the main interface */
type Ri struct {
	Contexter
}

func (r *Ri) writef(name RtName,parameterlist ...Rter) error {
	if r.Contexter == nil {
		return ErrNoActiveContext
	}
	return r.Write(name,parameterlist)
}
