package bxdf

import (
	. "github.com/mae-global/rigo/ri"
)

type RtIntWidget struct {
	param *Param
	parent *GeneralBxdf

	next,prev RtToken
}

func (r *RtIntWidget) Name() RtToken {
	return r.param.Name
}

func (r *RtIntWidget) NameSpec() RtToken {
	return RtToken(string(r.param.Type) + " " + string(r.param.Name))
}

func (r *RtIntWidget) Label() RtString {
	return r.param.Label
}

func (r *RtIntWidget) SetValue(value Rter) error {
	return r.parent.SetValue(r.param.Name,value)
}

func (r *RtIntWidget) GetValue() Rter {
	return r.parent.Value(r.param.Name)
}

func (r *RtIntWidget) Help() RtString {
	return r.param.Help
}

func (r *RtIntWidget) Bounds() (Rter,Rter) {
	return nil,nil
}

func (r *RtIntWidget) Next() Widget {
	return r.parent.Widget(r.next)
}

func (r *RtIntWidget) Prev() Widget {
	return r.parent.Widget(r.prev)
}

func (r *RtIntWidget) Default() error {
	return r.parent.SetValue(r.param.Name,r.param.Default)
}

func (r *RtIntWidget) Value() RtInt {
	return r.parent.Value(r.param.Name).(RtInt)
}

func (r *RtIntWidget) Set(value RtInt) error {
	/* TODO: check min/max */
	return r.parent.SetValue(r.param.Name,value)
}
