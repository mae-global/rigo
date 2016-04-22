package ris

import (
	. "github.com/mae-global/rigo/ri"
)

type RtColorWidget struct {
	param *Param
	parent Shader

	next,prev RtToken
}

func (r *RtColorWidget) Name() RtToken {
	r.param.RLock()
	defer r.param.RUnlock()
	return r.param.Name
}

func (r *RtColorWidget) NameSpec() RtToken {
	r.param.RLock()
	defer r.param.RUnlock()
	return RtToken(string(r.param.Type) + " " + string(r.param.Name))
}

func (r *RtColorWidget) Label() RtString {
	r.param.RLock()
	defer r.param.RUnlock()
	return r.param.Label
}

func (r *RtColorWidget) SetValue(value Rter) error {
	r.param.Lock()
	defer r.param.Unlock()
	if _,ok := value.(RtColor); !ok {
		return ErrInvalidType
	}
	r.param.Value = value 
	return  nil	
}

func (r *RtColorWidget) GetValue() Rter {
	r.param.RLock()
	defer r.param.RUnlock()
	return r.param.Name
}

func (r *RtColorWidget) Help() RtString {
	r.param.RLock()
	defer r.param.RUnlock()
	return r.param.Help
}

func (r *RtColorWidget) Bounds() (Rter,Rter) {
	r.param.RLock()
	defer r.param.RUnlock()
	return r.param.Min,r.param.Max
}

func (r *RtColorWidget) Next() Widget {
	return r.parent.Widget(r.next)
}

func (r *RtColorWidget) Prev() Widget {
	return r.parent.Widget(r.prev)
}

func (r *RtColorWidget) Default() error {
	r.param.Lock()
	defer r.param.Unlock()
	r.param.Value = r.param.Default
	return nil
}

func (r *RtColorWidget) Value() RtColor {
	r.param.RLock()
	defer r.param.RUnlock()
	return r.param.Value.(RtColor)
}

func (r *RtColorWidget) Set(value RtColor) error {
	/* TODO: check min/max */
	r.param.Lock()
	defer r.param.Unlock()
	r.param.Value = value
	return nil
}
