package rigo

import (
	"fmt"
	"os"
	"time"

	. "github.com/mae-global/rigo/ri"
)


func DefaultFilePipe() *Pipe {
	pipe := NewPipe()
	return pipe.Append(&PipeTimer{}).Append(&PipeToStats{}).Append(&PipeToFile{})
}

/* Time from Begin to End */
type PipeTimer struct {
	start time.Time
	finish time.Time
}

func (p PipeTimer) Name() string {
	return "default-pipe-timer"
}

func (p *PipeTimer) Write(name RtName,list []Rter,info Info) *Result {
	switch string(name) {
		case "Begin","RiBegin":
			p.start = time.Now()
			p.finish = p.start
			break
		case "End","RiEnd":
			p.finish = time.Now()
			break
	}
	return Done()
}

func (p *PipeTimer) String() string {
	return fmt.Sprintf("pipe took %s",p.finish.Sub(p.start))
}

func (p *PipeTimer) Took() time.Duration {
	return p.finish.Sub(p.start)
}

/* Pipe RI output to gathered states */
type PipeToStats struct {
	Stats map[RtName]int
}

func (p PipeToStats) Name() string {
	return "default-pipe-to-stats"
}

func (p *PipeToStats) Write(name RtName, list []Rter, info Info) *Result {
	if p.Stats == nil {
		p.Stats = make(map[RtName]int, 0)
	}
	if _, exists := p.Stats[name]; !exists {
		p.Stats[name] = 0
	}
	p.Stats[name]++

	return Done()
}

func (p *PipeToStats) String() string {
	if p.Stats == nil {
		return "stats [empty]"
	}

	if len(p.Stats) == 0 {
		return "stats [empty]"
	}

	max := 0
	for _, v := range p.Stats {
		if v > max {
			max = v
		}
	}

	dfmt := "\t%0" + fmt.Sprintf("%d", len(fmt.Sprintf("%d", max))) + "d"

	out := fmt.Sprintf("stats %d [\n", len(p.Stats))
	for n, v := range p.Stats {
		out += fmt.Sprintf(dfmt+" call(s).....%s\n", v, n)
	}
	return out + "]\n"
}

/* Pipe RI output to file */
type PipeToFile struct {
	file *os.File
}

func (p PipeToFile) Name() string {
	return "default-pipe-to-file"
}

func (p *PipeToFile) Write(name RtName, list []Rter, info Info) *Result {
	if info.Formal {
		name = name.Trim("Ri")
	}

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
		if _, err = p.file.Write([]byte(string(RIBStructure) + postfix)); err != nil {
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
		if _, err := p.file.Write([]byte(Serialise(list) + "\n")); err != nil {
			return InError(err)
		}
		return Done()
	}

	if name != "##" {

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
