package ris

import (
	. "github.com/mae-global/rigo/ri"
)

type RtNormalWidget struct {
	param *Param
	parent Shader

	next,prev RtToken
}

func (r *RtNormalWidget) Name() RtToken {
	r.param.RLock()
	defer r.param.RUnlock()
	return r.param.Name
}

func (r *RtNormalWidget) NameSpec() RtToken {
	r.param.RLock()
	defer r.param.RUnlock()
	return RtToken(string(r.param.Type) + " " + string(r.param.Name))
}

func (r *RtNormalWidget) Label() RtString {
	r.param.RLock()
	defer r.param.RUnlock()
	return r.param.Label
}

func (r *RtNormalWidget) SetValue(value Rter) error {
	r.param.Lock()
	defer r.param.Unlock()
	if _,ok := value.(RtNormal); !ok {
		return ErrInvalidType
	}
	r.param.Value = value 
	return  nil	
}

func (r *RtNormalWidget) GetValue() Rter {
	r.param.RLock()
	defer r.param.RUnlock()
	return r.param.Name
}

func (r *RtNormalWidget) Help() RtString {
	r.param.RLock()
	defer r.param.RUnlock()
	return r.param.Help
}

func (r *RtNormalWidget) Bounds() (Rter,Rter) {
	r.param.RLock()
	defer r.param.RUnlock()
	return r.param.Min,r.param.Max
}

func (r *RtNormalWidget) Next() Widget {
	return r.parent.Widget(r.next)
}

func (r *RtNormalWidget) Prev() Widget {
	return r.parent.Widget(r.prev)
}

func (r *RtNormalWidget) Default() error {
	r.param.Lock()
	defer r.param.Unlock()
	r.param.Value = r.param.Default
	return nil
}

func (r *RtNormalWidget) Value() RtNormal {
	r.param.RLock()
	defer r.param.RUnlock()
	return r.param.Value.(RtNormal)
}

func (r *RtNormalWidget) Set(value RtNormal) error {
	/* TODO: check min/max */
	r.param.Lock()
	defer r.param.Unlock()
	r.param.Value = value
	return nil
}
