package ris

import (
	. "github.com/mae-global/rigo/ri"
)

type RtFloatWidget struct {
	param  *Param
	parent Shader

	next, prev RtToken
}

func (r *RtFloatWidget) Name() RtToken {
	r.param.RLock()
	defer r.param.RUnlock()
	return r.param.Name
}

func (r *RtFloatWidget) NameSpec() RtToken {
	r.param.RLock()
	defer r.param.RUnlock()
	return RtToken(string(r.param.Type) + " " + string(r.param.Name))
}

func (r *RtFloatWidget) Label() RtString {
	r.param.RLock()
	defer r.param.RUnlock()
	return r.param.Label
}

func (r *RtFloatWidget) SetValue(value Rter) error {
	r.param.Lock()
	defer r.param.Unlock()
	if _, ok := value.(RtInt); !ok {
		return ErrInvalidType
	}
	r.param.Value = value
	return nil
}

func (r *RtFloatWidget) GetValue() Rter {
	r.param.RLock()
	defer r.param.RUnlock()
	return r.param.Value
}

func (r *RtFloatWidget) Help() RtString {
	r.param.RLock()
	defer r.param.RUnlock()
	return r.param.Help
}

func (r *RtFloatWidget) Bounds() (Rter, Rter) {
	r.param.RLock()
	defer r.param.RUnlock()
	return r.param.Min, r.param.Max
}

func (r *RtFloatWidget) Next() Widget {
	return r.parent.Widget(r.next)
}

func (r *RtFloatWidget) Prev() Widget {
	return r.parent.Widget(r.prev)
}

func (r *RtFloatWidget) Default() error {
	r.param.Lock()
	defer r.param.Unlock()
	r.param.Value = r.param.Default
	return nil
}

func (r *RtFloatWidget) Value() RtFloat {
	r.param.RLock()
	defer r.param.RUnlock()
	return r.param.Value.(RtFloat)
}

func (r *RtFloatWidget) Set(value RtFloat) error {
	/* TODO: check min/max */
	r.param.Lock()
	defer r.param.Unlock()
	r.param.Value = value
	return nil
}
