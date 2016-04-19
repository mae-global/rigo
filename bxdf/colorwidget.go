package bxdf

import (
	. "github.com/mae-global/rigo/ri"
)

type RtColorWidget struct {
	param *Param
	parent *GeneralBxdf

	next,prev RtToken
}

func (r *RtColorWidget) Name() RtToken {
	return r.param.Name
}

func (r *RtColorWidget) NameSpec() RtToken {
	return RtToken(string(r.param.Type) + " " + string(r.param.Name))
}

func (r *RtColorWidget) Label() RtString {
	return r.param.Label
}

func (r *RtColorWidget) SetValue(value Rter) error {
	return r.parent.SetValue(r.param.Name,value)
}

func (r *RtColorWidget) GetValue() Rter {
	return r.parent.Value(r.param.Name)
}

func (r *RtColorWidget) Help() RtString {
	return r.param.Help
}

func (r *RtColorWidget) Bounds() (Rter,Rter) {
	return nil,nil
}

func (r *RtColorWidget) Next() Widget {
	return r.parent.Widget(r.next)
}

func (r *RtColorWidget) Prev() Widget {
	return r.parent.Widget(r.prev)
}

func (r *RtColorWidget) Default() error {
	return r.parent.SetValue(r.param.Name,r.param.Default)
}

func (r *RtColorWidget) Value() RtColor {
	return r.parent.Value(r.param.Name).(RtColor)
}

func (r *RtColorWidget) Set(value RtColor) error {
	/* TODO: check min/max */
	return r.parent.SetValue(r.param.Name,value)
}
