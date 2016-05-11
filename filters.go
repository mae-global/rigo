package rigo

import (
	"strconv"

	. "github.com/mae-global/rigo/ri"
)


type FilterStringHandles struct{}

/* FIXME: this should actually be a filter */

func (p *FilterStringHandles) ToRaw() ArchiveWriter {
	return nil
}

func (p FilterStringHandles) Name() string {
	return "default-filter-string-handles"
}

func (p *FilterStringHandles) Pipe(name RtName, args, params, values []Rter, info Info) *Result {

	/* TODO: add filter to only those proceedures the include light and object handles */

	args1 := make([]Rter, len(args))

	for i := 0; i < len(args); i++ {
		if lh, ok := args[i].(RtLightHandle); ok {
			id, err := strconv.Atoi(string(lh))
			if err != nil {
				return InError(err)
			}
			args1[i] = RtInt(id)
			continue
		}
		if oh, ok := args[i].(RtObjectHandle); ok {
			id, err := strconv.Atoi(string(oh))
			if err != nil {
				return InError(err)
			}
			args1[i] = RtInt(id)
			continue
		}
		if sh, ok := args[i].(RtShaderHandle); ok {
			id, err := strconv.Atoi(string(sh))
			if err != nil {
				return InError(err)
			}
			args1[i] = RtInt(id)
			continue
		}
		args1[i] = args[i]
	}

	return Next(name, args1, params, values, info)
}

