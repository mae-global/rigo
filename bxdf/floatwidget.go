package bxdf

import (
	. "github.com/mae-global/rigo/ri"
)

type RtFloatWidget struct {
	param *Param
	parent *GeneralBxdf

	next,prev RtToken
}

func (r *RtFloatWidget) Name() RtToken {
	return r.param.Name
}

func (r *RtFloatWidget) NameSpec() RtToken {
	return RtToken(string(r.param.Type) + " " + string(r.param.Name))
}

func (r *RtFloatWidget) Label() RtString {
	return r.param.Label
}

func (r *RtFloatWidget) SetValue(value Rter) error {
	return r.parent.SetValue(r.param.Name,value)
}

func (r *RtFloatWidget) GetValue() Rter {
	return r.parent.Value(r.param.Name)
}

func (r *RtFloatWidget) Help() RtString {
	return r.param.Help
}

func (r *RtFloatWidget) Bounds() (Rter,Rter) {
	return nil,nil
}

func (r *RtFloatWidget) Next() Widget {
	return r.parent.Widget(r.next)
}

func (r *RtFloatWidget) Prev() Widget {
	return r.parent.Widget(r.prev)
}

func (r *RtFloatWidget) Default() error {
	return r.parent.SetValue(r.param.Name,r.param.Default)
}

func (r *RtFloatWidget) Value() RtFloat {
	return r.parent.Value(r.param.Name).(RtFloat)
}

func (r *RtFloatWidget) Set(value RtFloat) error {
	/* TODO: check min/max */
	return r.parent.SetValue(r.param.Name,value)
}
