package ris

import (
	"os"
	"fmt"
	"io/ioutil"
	"strings"

	. "github.com/mae-global/rigo/ri"
)

var dashed = RtShaderHandle("-")

type RisContexter interface {
	ShaderHandle() (RtShaderHandle,error)

	SetShader(RtShaderHandle,Shader) 
	GetShader(RtShaderHandle) Shader
}

type Ris struct {
	RisContexter
}

func (ris *Ris) Bxdf(name string,sh RtShaderHandle) (Shader,error) {
	
	if s := ris.GetShader(sh); s != nil {
		return s,nil
	}
	
	if len(sh) == 0 || sh == dashed {
		if h,err := ris.ShaderHandle(); err != nil {
			return nil,err
		} else {
			sh = h
		}
	}
		
	bxdf,err := Bxdf(name,sh)
	if err != nil {
		return nil,err
	}

	ris.SetShader(sh,bxdf)

	return bxdf,nil
}

/* Load a bxdf shader from RMANTREE */
func Bxdf(name string,sh RtShaderHandle) (Shader,error) {

	rmantree := os.Getenv("RMANTREE")
	if len(rmantree) == 0 {
		return nil,fmt.Errorf("is RMANTREE set?")
	}

	debug := os.Getenv("DEBUG")
	if debug == "testing" {
		name = strings.Replace(name,"Pxr","Test",-1)
	}

	filepath := rmantree + "/lib/RIS/bxdf/Args/" + name + ".args"

	file,err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil,err
	}

	return Parse(name,sh,file)
}

func (ris *Ris) Integrator(name string,sh RtShaderHandle) (Shader,error) {

	if s := ris.GetShader(sh); s != nil {
		return s,nil
	}

	if len(sh) == 0 || sh == dashed {
		if h,err := ris.ShaderHandle(); err != nil {
			return nil,err
		} else {
			sh = h
		}
	}

	integrator,err := Integrator(name,sh)
	if err != nil {
		return nil,err
	}

	ris.SetShader(sh,integrator)
	
	return integrator,nil
}

/* Load a integrator shader from RMANTREE */
func Integrator(name string,sh RtShaderHandle) (Shader,error) {

	rmantree := os.Getenv("RMANTREE")
	if len(rmantree) == 0 {
		return nil,fmt.Errorf("is RMANTREE set?")
	}

	debug := os.Getenv("DEBUG")
	if debug == "testing" {
		name = strings.Replace(name,"Pxr","Test",-1)
	}


	filepath := rmantree + "/lib/RIS/integrator/Args/" + name + ".args"
	
	file,err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil,err
	}

	return Parse(name,sh,file)
}

func (ris *Ris) LightFilter(name string,sh RtShaderHandle) (Shader,error) {

	if s := ris.GetShader(sh); s != nil {
		return s,nil
	}

	if len(sh) == 0 || sh == dashed {
		if h,err := ris.ShaderHandle(); err != nil {
			return nil,err
		} else {
			sh = h
		}
	}

	lightfilter,err := LightFilter(name,sh)
	if err != nil {
		return nil,err
	}

	ris.SetShader(sh,lightfilter)
	
	return lightfilter,nil
}

/* Load a light shader from RMANTREE */
func LightFilter(name string,sh RtShaderHandle) (Shader,error) {

	rmantree := os.Getenv("RMANTREE")
	if len(rmantree) == 0 {
		return nil,fmt.Errorf("is RMANTREE set?")
	}

	debug := os.Getenv("DEBUG")
	if debug == "testing" {
		name = strings.Replace(name,"Pxr","Test",-1)
	}

	filepath := rmantree + "/lib/RIS/light/Args/" + name + ".args"

	file,err := ioutil.ReadFile(filepath) 
	if err != nil {
		return nil,err 
	}

	return Parse(name,sh,file)
}

func (ris *Ris) Projection(name string,sh RtShaderHandle) (Shader,error) {

	if s := ris.GetShader(sh); s != nil {
		return s,nil
	}

	if len(sh) == 0 || sh == dashed {
		if h,err := ris.ShaderHandle(); err != nil {
			return nil,err
		} else {
			sh = h
		}
	}

	projection,err := Projection(name,sh)
	if err != nil {
		return nil,err
	}

	ris.SetShader(sh,projection)

	return projection,nil
}

/* Load a projection shader from RMANTREE */
func Projection(name string,sh RtShaderHandle) (Shader,error) {

	rmantree := os.Getenv("RMANTREE")
	if len(rmantree) == 0 {
		return nil,fmt.Errorf("is RMANTREE set?")
	}

	debug := os.Getenv("DEBUG")
	if debug == "testing" {
		name = strings.Replace(name,"Pxr","Test",-1)
	}

	filepath := rmantree + "/lib/RIS/projection/Args/" + name + ".args"

	file,err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil,err
	}

	return Parse(name,sh,file)
}

func (ris *Ris) Pattern(name string,sh RtShaderHandle) (Shader,error) {
	
	if s := ris.GetShader(sh); s != nil {
		return s,nil
	}

	if len(sh) == 0 || sh == dashed {
		if h,err := ris.ShaderHandle(); err != nil {
			return nil,err
		} else {
			sh = h
		}
	}

	pattern,err := Pattern(name,sh)
	if err != nil {
		return nil,err
	}

	ris.SetShader(sh,pattern)
	
	return pattern,nil
}

/* Load a pattern shader from RMANTREE */
func Pattern(name string,sh RtShaderHandle) (Shader,error) {

	rmantree := os.Getenv("RMANTREE")
	if len(rmantree) == 0 {
		return nil,fmt.Errorf("is RMANTREE set?")
	}

	debug := os.Getenv("DEBUG")
	if debug == "testing" {
		name = strings.Replace(name,"Pxr","Test",-1)
	}

	filepath := rmantree + "/lib/RIS/pattern/Args/" + name + ".args"

	file,err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil,err
	}
	
	return Parse(name,sh,file)
}



