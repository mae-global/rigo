/* rigo/presets.go */
package rigo

import (
	. "github.com/mae-global/rigo/ri"
)

func DefaultPipeline(config *Configuration) (*Ri,*Pipe) {

	pipe := DefaultFilePipe()

	return RI(NewContext(pipe,nil,config)),pipe
}


func StrictPipeline() (*Ri, *Pipe) {

	pipe := NewPipe()
	pipe.Append(&PipeTimer{}).Append(&PipeToStats{}).Append(&FilterStringHandles{}).Append(&PipeToPrettyPrint{}).Append(&PipeToFile{})

	ctx := NewContext(pipe, nil, nil)

	return RI(ctx), pipe
}

func EntityPipeline() (*Ri, *Pipe) {

	pipe := DefaultFilePipe()

	ctx := NewContext(pipe, nil, &Configuration{Entity: true, PrettyPrint: true})

	return RI(ctx), pipe
}

func CustomEntityPipeline(pipe *Pipe) *Ri {

	return RI(NewContext(pipe, nil, &Configuration{Entity: true, PrettyPrint: true}))
}
