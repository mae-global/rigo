package ri

import (
	"os"
)

func DefaultFilePipe() *Pipe {
	pipe := NewPipe()
	return pipe.Append(&PipeToFile{})
}


type PipeToFile struct {
	file *os.File
}

func (p *PipeToFile) Write(name RtName, list []Rter, info Info) *Result {
	if name == "Begin" {
		if p.file != nil {
			return InError(ErrProtocolBotch)
		}
		file := "out.rib"
		if len(list) > 0 {
			if t, ok := list[0].(RtString); ok {
				file = string(t)
			}
		}

		f, err := os.Create(file)
		if err != nil {
			return InError(err)
		}
		p.file = f

		postfix := "\n"
		if info.Entity {
			postfix = " Entity\n"
		}
		if _, err = p.file.Write([]byte("##RenderMan RIB-Structure 1.1" + postfix)); err != nil {
			return InError(err)
		}
		return Done()
	}

	if name == "End" {
		if p.file == nil {
			return InError(ErrProtocolBotch)
		}
		if err := p.file.Close(); err != nil {
			return InError(err)
		}
		return Done()
	}

	if p.file == nil {
		return InError(ErrNoActiveContext)
	}

	if name == "Verbatim" {
		if _,err := p.file.Write([]byte(Serialise(list) + "\n")); err != nil {
			return InError(err)
		}
		return Done()
	}

	if name != "##"  {

		prefix := ""
		for i := 0; i < info.Depth; i++ {
			prefix += "\t"
		}

		if _, err := p.file.Write([]byte(prefix + name.Serialise() + " " + Serialise(list) + "\n")); err != nil {
			return InError(err)
		}
		return Done()
	}

	if _, err := p.file.Write([]byte("##" + Serialise(list) + "\n")); err != nil {
		return InError(err)
	}
	return Done()
}







