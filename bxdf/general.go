package bxdf

import (
	"sync"
	"fmt"
	"strings"

	. "github.com/mae-global/rigo/ri"
)

var (
	ErrInvalidParam = fmt.Errorf("Invalid Param")
	ErrParamAlreadyExists = fmt.Errorf("Param already exists")
)

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
	
	params []*Param
}

func NewGeneralBxdf(name,nodeid RtToken,classification RtString) *GeneralBxdf {
	g := &GeneralBxdf{name:name,nodeid:nodeid,classification:classification}
	g.params = make([]*Param,0)
	return g
}

func (g *GeneralBxdf) AddParam(p *Param) error {
	if p == nil {
		return ErrInvalidParam
	}
	for _,param := range g.params {
		if param.Name == p.Name {
			return ErrParamAlreadyExists 
		}
	}
	g.params = append(g.params,p)
	return nil
}		


func (g *GeneralBxdf) Write() (RtName,[]Rter,[]Rter) {

	args := make([]Rter,0)
	params := make([]Rter,0)

	n := strings.ToLower(string(g.name))
	name := string(n[0]) + string(g.name[1:])
	
	args = append(args,RtToken(name))

	for _,param := range g.params {
		param.RLock()

		params = append(params,param.Name)
		params = append(params,param.Value)

		param.RUnlock()
	}

	return RtName("Bxdf"),args,params
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
			p = param
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

func (g *GeneralBxdf) FirstWidget() Widget {
	if len(g.params) == 0 {
		return nil
	}
	return g.Widget(g.params[0].Name)
}

func (g *GeneralBxdf) LastWidget() Widget {
	if len(g.params) == 0 {
		return nil
	}
	return g.Widget(g.params[len(g.params) - 1].Name)
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

			return param.Value
		}
	}
	return nil
}


