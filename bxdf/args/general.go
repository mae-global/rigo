package args

import (
	"sync"
	"fmt"

	. "github.com/mae-global/rigo/ri"
	. "github.com/mae-global/rigo/bxdf"
)

var (
	ErrTypeMismatch = fmt.Errorf("Type Mismatch")
)


type RtColorWidget struct {
	param *Param
	parent *GeneralBxdf
	
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

	if r.param.Value.Type() == value.Type() {
		return ErrTypeMismatch
	}

	r.param.Value = value
}

func (r *RtColorWidget) GetValue() Rter {
	r.param.RLock()
	defer r.param.RUnlock()
	return r.param.Value
}

func (r *RtColorWidget) Help() RtString {
	r.param.RLock()
	defer r.param.RUnlock()
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
	r.param.Lock()
	defer r.param.Unlock()
	r.param.Value = r.param.Default
	return nil /* FIXME, remove error */
}

func (r *RtColorWidget) Value() RtColor {
	r.param.RLock()
	defer r.param.RUnlock()
	return r.param.Value.(RtColor)
}

func (r *RtColorWidget) Set(color RtColor) error {
	/* TODO: set to min/max */
	r.param.Lock()
	defer r.param.Unlock()
	r.param.Value = color
	return nil /* FIXME, remove error */
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


type RtNormalWidget struct {
	param *Param
	parent *GeneralBxdf

	next RtToken
	prev RtToken
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
	val := r.parent.Value(r.param.Name)
	return val.(RtNormal)
}

func (r *RtNormalWidget) Set(value RtNormal) error {
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
	Value Rter

	sync.RWMutex

	/* TODO: add option hints dictionary here */
}


type GeneralBxdf struct {
	nodeid RtToken
	name RtToken
	classification RtString

	params []Param /* inorder of definition */

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

	var next,prev RtToken
	var p *Param
	found := -1	

	for i,param := range g.params {
		if param.Name == name {
			param.RLock()
			defer param.RUnlock()
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
			w = &RtColorWidget{param:p,parent:g,next:next,prev:prev}
		break
		case "float":
			w = &RtFloatWidget{param:p,parent:g,next:next,prev:prev}
		break
		case "int":
			w = &RtIntWidget{param:p,parent:g,next:next,prev:prev}
		break
		case "normal":
			w = &RtNormalWidget{param:p,parent:g,next:next,prev:prev}
		break
		/* TODO: add normal etc.. here*/
	}

	return w
}

func (g *GeneralBxdf) Names() []RtToken {
	names := make([]RtToken,len(g.params))
	for i,param := range g.params {
		names[i] = param.Name
	}
	return names
}

func (g *GeneralBxdf)	NamesSpec() []RtToken {
	names := make([]RtToken,len(g.params))
	for i,param := range g.params {
		names[i] = RtToken(string(param.Type) + " " + string(param.Name)) /* FIXME, this is not a complete spec : missing [n] */
	}
	return names
}

func (g *GeneralBxdf)	SetValue(name RtToken,value Rter) error {
	
	var val Rter
	for _,param := range g.params {
		if param.Name == name {
			param.Lock()
			defer param.Unlock()

			if param.Value.Type() != value.Type() {
				return fmt.Errorf("Type mismatch, setting with \"%s\", wants \"%s\"",value.Type(),param.Value.Type())
			}

			param.Value = value
			return nil
		}
	}
	return fmt.Errorf("Unknown parameter %s",name)
}

func (g *GeneralBxdf) Value(name RtToken) Rter {

	for _,param := range g.params {
		if param.Name == name {
			param.RLock()
			defer param.RUnlock()

			return para.Value
		}
	}
	return nil
}


