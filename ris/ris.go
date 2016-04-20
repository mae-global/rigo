package ris

import (
	"os"
	"fmt"
	"io/ioutil"

)

/* Load a bxdf shader from RMANTREE */
func Bxdf(name string) (Shader,error) {

	rmantree := os.Getenv("RMANTREE")
	if len(rmantree) == 0 {
		return nil,fmt.Errorf("is RMANTREE set?")
	}

	filepath := rmantree + "/lib/RIS/bxdf/Args/" + name + ".args"

	file,err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil,err
	}

	return Parse(name,file)
}

/* Load a integrator shader from RMANTREE */
func Integrator(name string) (Shader,error) {

	rmantree := os.Getenv("RMANTREE")
	if len(rmantree) == 0 {
		return nil,fmt.Errorf("is RMANTREE set?")
	}

	filepath := rmantree + "/lib/RIS/integrator/Args/" + name + ".args"
	
	file,err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil,err
	}

	return Parse(name,file)
}

/* Load a light shader from RMANTREE */
func LightFilter(name string) (Shader,error) {

	rmantree := os.Getenv("RMANTREE")
	if len(rmantree) == 0 {
		return nil,fmt.Errorf("is RMANTREE set?")
	}

	filepath := rmantree + "/lib/RIS/light/Args/" + name + ".args"

	file,err := ioutil.ReadFile(filepath) 
	if err != nil {
		return nil,err 
	}

	return Parse(name,file)
}

/* Load a projection shader from RMANTREE */
func Projection(name string) (Shader,error) {

	rmantree := os.Getenv("RMANTREE")
	if len(rmantree) == 0 {
		return nil,fmt.Errorf("is RMANTREE set?")
	}

	filepath := rmantree + "/lib/RIS/projection/Args/" + name + ".args"

	file,err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil,err
	}

	return Parse(name,file)
}

/* Load a pattern shader from RMANTREE */
func Pattern(name string) (Shader,error) {

	rmantree := os.Getenv("RMANTREE")
	if len(rmantree) == 0 {
		return nil,fmt.Errorf("is RMANTREE set?")
	}

	filepath := rmantree + "/lib/RIS/pattern/Args/" + name + ".args"

	file,err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil,err
	}
	
	return Parse(name,file)
}



