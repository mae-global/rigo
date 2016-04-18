package args

import (
	"sync"
	"fmt"

	. "github.com/mae-global/rigo/ri"
	. "github.com/mae-global/rigo/bxdf"
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

func (r *RtColorWidget) Set(color RtColor) error {
	/* TODO: set to min/max */
	return r.parent.SetValue(r.param.Name,color)
}

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

type RtFloatWidget struct {
	param *Param
	parent *GeneralBxdf

	next RtToken
	prev RtToken
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
	val := r.parent.Value(r.param.Name)
	return val.(RtFloat)
}

func (r *RtFloatWidget) Set(value RtFloat) error {
	/* TODO: check the min and max */
	//if r.param.Min != r.param.Max {
		
	return r.parent.SetValue(r.param.Name,value)
} 



type Param struct {
	Label RtString
	Name RtToken
	Type RtToken
	Default Rter
	Min Rter
	Max Rter
	Widget RtToken
	Help RtString

	/* TODO: add option hints dictionary here */
}


type GeneralBxdf struct {
	nodeid RtToken
	name RtToken
	classification RtString

	params []Param /* inorder of definition */
	values map[RtToken] Rter

	mux sync.RWMutex
}

func (g *GeneralBxdf) Write() []Rter {
	return nil
}

func (g *GeneralBxdf) Name() RtToken {
	return g.name
}

func (g *GeneralBxdf) NodeId() RtToken {
	return g.nodeid 
}

func (g *GeneralBxdf) Classifiation() RtString {
	return g.classification
}

func (g *GeneralBxdf) Widget(name RtToken) Widget {
	g.mux.Lock()
	defer g.mux.Unlock()

	var next,prev RtToken
	var p *Param
	found := -1	

	for i,param := range g.params {
		if param.Name == name {
			p = &param
			found = i
			break
		}
	}

	if p == nil {
		return nil
	}
	
	if found - 1 < 0 {
		prev = g.params[len(g.params) - 1].Name
	} else {
		prev = g.params[found - 1].Name
	}

	if found + 1 >= len(g.params) {
		next = g.params[0].Name
	} else {
		next = g.params[found + 1].Name
	}
	
	var w Widget

	switch p.Type {
		case "color":
			cw := &RtColorWidget{param:p,parent:g,next:next,prev:prev}

			w = cw
		break
		case "float":
			fw := &RtFloatWidget{param:p,parent:g,next:next,prev:prev}

			w = fw
		break
		case "int":
			iw := &RtIntWidget{param:p,parent:g,next:next,prev:prev}

			w = iw
		break
	}

	return w
}

func (g *GeneralBxdf) Names() []RtToken {
	g.mux.RLock()
	defer g.mux.RUnlock()
	names := make([]RtToken,len(g.params))
	for i,param := range g.params {
		names[i] = param.Name
	}
	return names
}

func (g *GeneralBxdf)	NamesSpec() []RtToken {
	g.mux.RLock()
	defer g.mux.RUnlock()
	names := make([]RtToken,len(g.params))
	for i,param := range g.params {
		names[i] = RtToken(string(param.Type) + " " + string(param.Name)) /* FIXME, this is not a complete spec : missing [n] */
	}
	return names
}


func (g *GeneralBxdf)	SetValue(name RtToken,value Rter) error {
	g.mux.Lock()
	defer g.mux.Unlock()

	if r,ok := g.values[name]; ok {
		if r.Type() != value.Type() {
			return fmt.Errorf("Type mismatch, setting with \"%s\", wants \"%s\"",value.Type(),r.Type())
		}

		g.values[name] = value
		return nil
	}
	return fmt.Errorf("Unknown parameter %s",name)
}

func (g *GeneralBxdf) Value(name RtToken) Rter {
	g.mux.RLock()
	defer g.mux.RUnlock()

	if r,ok := g.values[name]; ok {
		return r
	}	

	return nil
}

