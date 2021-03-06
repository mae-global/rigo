package ris

import (
	"fmt"
	"sync"

	. "github.com/mae-global/rigo/ri"
)

var (
	ErrInvalidParam       = fmt.Errorf("Invalid Param")
	ErrParamAlreadyExists = fmt.Errorf("Param already exists")
)

type Param struct {
	Label   RtString
	Name    RtToken
	Type    RtToken
	Default Rter
	Min     Rter
	Max     Rter
	Widget  RtToken
	Help    RtString
	Value   Rter
	

	ref RtString
	sync.RWMutex

	/* TODO: add option hints dictionary here */
}

type Output struct {
	Name RtToken
	Types []RtToken
}

/* IsDefault is the param value same as the default */
func (param *Param) IsDefault() bool {
	param.RLock()
	defer param.RUnlock()

	return (param.Value.Equal(param.Default))
}

type GeneralShader struct {
	shadertype     RtName
	nodeid         RtToken
	name           RtToken
	classification RtString
	handle         RtShaderHandle

	params []*Param
	outputs []*Output
}

func NewGeneralShader(shadertype RtName, name, nodeid RtToken, classification RtString, handle RtShaderHandle) *GeneralShader {
	g := &GeneralShader{shadertype: shadertype, name: name, nodeid: nodeid, classification: classification, handle: handle}
	g.params = make([]*Param, 0)
	g.outputs = make([]*Output,0)	
	return g
}

func (g *GeneralShader) AddParam(p *Param) error {
	if p == nil {
		return ErrInvalidParam
	}
	for _, param := range g.params {
		if param.Name == p.Name {
			return ErrParamAlreadyExists
		}
	}
	g.params = append(g.params, p)
	return nil
}

func (g *GeneralShader) Handle() RtShaderHandle {
	return g.handle
}

func (g *GeneralShader) ShaderType() RtName {
	return g.shadertype
}

func namespec(name, typeof RtToken) RtToken {
	return RtToken(string(typeof) + " " + string(name))
}

func refnamespec(name,typeof RtToken) RtToken {
	return RtToken(fmt.Sprintf("reference %s %s",string(typeof),string(name)))
}

func (g *GeneralShader) Write() (RtName, RtShaderHandle, []Rter, []Rter, []Rter) {

	args := make([]Rter, 0)
	params := make([]Rter, 0)
	values := make([]Rter, 0)

	args = append(args, RtToken(g.name))

	for _, param := range g.params {
		/* if the value is equal to the default value then we don't need to
		 * write it out TODO: add a flag to control this */
		if param.IsDefault() && param.ref == "" {
			continue
		}

		param.RLock()
		/* override with reference */
		if param.ref != "" {
			params = append(params, refnamespec(param.Name,param.Type))
			values = append(values, param.ref)
		} else {
			params = append(params, namespec(param.Name, param.Type))
			values = append(values, param.Value)
		}
		param.RUnlock()
	}

	return g.shadertype, g.handle, args, params, values
}

func (g *GeneralShader) Name() RtToken {
	return g.name
}

func (g *GeneralShader) NodeId() RtToken {
	return g.nodeid
}

func (g *GeneralShader) Classifiation() RtString {
	return g.classification
}

func (g *GeneralShader) Widget(name RtToken) Widget {

	var next, prev RtToken
	var p *Param
	found := -1

	for i, param := range g.params {
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

	if found-1 < 0 {
		prev = g.params[len(g.params)-1].Name
	} else {
		prev = g.params[found-1].Name
	}

	if found+1 >= len(g.params) {
		next = g.params[0].Name
	} else {
		next = g.params[found+1].Name
	}

	var w Widget

	switch p.Type {
	case "color":
		w = &RtColorWidget{param: p, parent: g, next: next, prev: prev}
		break
	case "float":
		w = &RtFloatWidget{param: p, parent: g, next: next, prev: prev}
		break
	case "int":
		w = &RtIntWidget{param: p, parent: g, next: next, prev: prev}
		break
	case "normal":
		w = &RtNormalWidget{param: p, parent: g, next: next, prev: prev}
		break
		/* TODO: add normal etc.. here*/
	}

	return w
}

func (g *GeneralShader) FirstWidget() Widget {
	if len(g.params) == 0 {
		return nil
	}
	return g.Widget(g.params[0].Name)
}

func (g *GeneralShader) LastWidget() Widget {
	if len(g.params) == 0 {
		return nil
	}
	return g.Widget(g.params[len(g.params)-1].Name)
}

func (g *GeneralShader) Names() []RtToken {
	names := make([]RtToken, len(g.params))
	for i, param := range g.params {
		names[i] = param.Name
	}
	return names
}

func (g *GeneralShader) Info() (RtInt,RtInt) {
	return RtInt(len(g.params)),RtInt(len(g.outputs))
}

func (g *GeneralShader) NamesSpec() []RtToken {
	names := make([]RtToken,0)
	for _, param := range g.params {
		names = append(names,RtToken(string(param.Type) + " " + string(param.Name))) /* FIXME, this is not a complete spec : missing [n] */
	}
	for _,param := range g.outputs {
		names = append(names,RtToken("output " + string(param.Name) + " " + RtTokenArray(param.Types).String()))
	}
	return names
}

func (g *GeneralShader) SetValue(name RtToken, value Rter) error {

	for _, param := range g.params {
		if param.Name == name {
			param.Lock()
			defer param.Unlock()

			if param.Value.Type() != value.Type() {
				return fmt.Errorf("Type mismatch, setting with \"%s\", wants \"%s\"", value.Type(), param.Value.Type())
			}

			param.Value = value
			return nil
		}
	}
	return fmt.Errorf("Unknown parameter %s", name)
}

func (g *GeneralShader) SetReferencedValue(name RtToken, value RtString) error {

	for _, param := range g.params {
		if param.Name == name {
			param.Lock()
			defer param.Unlock()

			param.ref = value
			return nil
		}
	}
	return fmt.Errorf("Unknown parameter %s",name)
}

func (g *GeneralShader) Value(name RtToken) Rter {

	for _, param := range g.params {
		if param.Name == name {
			param.RLock()
			defer param.RUnlock()

			return param.Value
		}
	}
	return nil
}

func (g *GeneralShader) ReferenceOutput(name RtToken) RtString {
	
	for _, output := range g.outputs {
		if output.Name == name {
			return RtString(fmt.Sprintf("%s:%s",string(g.handle),string(name)))
		}
	}
	return ""
}
		


