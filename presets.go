/* rigo/presets.go */
package rigo

import (
	. "github.com/mae-global/rigo/ri"
	. "github.com/mae-global/rigo/ri/handles"
)


func StrictPipeline() (*Ri,*Pipe) {

	pipe := NewPipe()
	pipe.Append(&PipeTimer{}).Append(&PipeToStats{}).Append(&FilterStringHandles{}).Append(&PipeToFile{})

	return NewCustom(pipe,NewLightNumberGenerator(),NewObjectNumberGenerator(),nil,nil),pipe
}
