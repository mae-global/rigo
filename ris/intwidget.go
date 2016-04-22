package ris

import (
	. "github.com/mae-global/rigo/ri"
)

type RtIntWidget struct {
	param *Param
	parent Shader

	next,prev RtToken
}

func (r *RtIntWidget) Name() RtToken {
	r.param.RLock()
	defer r.param.RUnlock()
	return r.param.Name
}

func (r *RtIntWidget) NameSpec() RtToken {
	r.param.RLock()
	defer r.param.RUnlock()
	return RtToken(string(r.param.Type) + " " + string(r.param.Name))
}

func (r *RtIntWidget) Label() RtString {
	r.param.RLock()
	defer r.param.RUnlock()
	return r.param.Label
}

func (r *RtIntWidget) SetValue(value Rter) error {
	r.param.Lock()
	defer r.param.Unlock()
	if _,ok := value.(RtInt); !ok {
		return ErrInvalidType
	}
	r.param.Value = value
	return nil
}

func (r *RtIntWidget) GetValue() Rter {
	r.param.RLock()
	defer r.param.RUnlock()
	return r.param.Value
}

func (r *RtIntWidget) Help() RtString {
	r.param.RLock()
	defer r.param.RUnlock()
	return r.param.Help
}

func (r *RtIntWidget) Bounds() (Rter,Rter) {
	r.param.RLock()
	defer r.param.RUnlock()
	return r.param.Min,r.param.Max
}

func (r *RtIntWidget) Next() Widget {
	return r.parent.Widget(r.next)
}

func (r *RtIntWidget) Prev() Widget {
	return r.parent.Widget(r.prev)
}

func (r *RtIntWidget) Default() error {
	r.param.Lock()
	defer r.param.Unlock()
	r.param.Value = r.param.Default
	return nil
}

func (r *RtIntWidget) Value() RtInt {
	r.param.RLock()
	defer r.param.RUnlock()	
	return r.param.Value.(RtInt)
}

func (r *RtIntWidget) Set(value RtInt) error {
	/* TODO: check min/max */
	r.param.Lock()
	defer r.param.Unlock()	
	r.param.Value = value
	return nil
}


