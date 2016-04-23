package rigo

import (
	. "github.com/mae-global/rigo/ri"
)

func Mix(params,values []Rter) []Rter {

	out := make([]Rter,0)

	for i := 0; i < len(params); i++ {

		out = append(out,params[i])
		out = append(out,values[i])
	}
	return out
}

func Unmix(list []Rter) ([]Rter,[]Rter) {

	params := make([]Rter,0)
	values := make([]Rter,0)

	flipflop := false

	for _,param := range list {
		if !flipflop {
			params = append(params,param)
		} else {
			values = append(values,param)
		}
		flipflop = !flipflop
	}

	return params,values
}

