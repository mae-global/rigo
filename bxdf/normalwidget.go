package bxdf

import (
	. "github.com/mae-global/rigo/ri"
)

type RtNormalWidget struct {
	param *Param
	parent *GeneralBxdf

	next,prev RtToken
}

func (r *RtNormalWidget) Name() RtToken {
	return r.param.Name
}

func (r *RtNormalWidget) NameSpec() RtToken {
	return RtToken(string(r.param.Type) + " " + string(r.param.Name))
}

func (r *RtNormalWidget) Label() RtString {
	return r.param.Label
}

func (r *RtNormalWidget) SetValue(value Rter) error {
	return r.parent.SetValue(r.param.Name,value)
}

func (r *RtNormalWidget) GetValue() Rter {
	return r.parent.Value(r.param.Name)
}

func (r *RtNormalWidget) Help() RtString {
	return r.param.Help
}

func (r *RtNormalWidget) Bounds() (Rter,Rter) {
	return nil,nil
}

func (r *RtNormalWidget) Next() Widget {
	return r.parent.Widget(r.next)
}

func (r *RtNormalWidget) Prev() Widget {
	return r.parent.Widget(r.prev)
}

func (r *RtNormalWidget) Default() error {
	return r.parent.SetValue(r.param.Name,r.param.Default)
}

func (r *RtNormalWidget) Value() RtNormal {
	return r.parent.Value(r.param.Name).(RtNormal)
}

func (r *RtNormalWidget) Set(value RtNormal) error {
	/* TODO: check min/max */
	return r.parent.SetValue(r.param.Name,value)
}
