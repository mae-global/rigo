package rigo

import (
	"fmt"
	"os"
	"time"
	"sort"
	"strconv"

	. "github.com/mae-global/rigo/ri"
)

func DefaultFilePipe() *Pipe {
	pipe := NewPipe()
	return pipe.Append(&PipeTimer{}).Append(&PipeToStats{}).Append(&PipeToFile{})
}

func NullPipe() *Pipe {
	pipe := NewPipe()
	return pipe.Append(&PipeTimer{}).Append(&PipeToStats{})
}

/* Time from Begin to End */
type PipeTimer struct {
	start  time.Time
	finish time.Time
}

func (p *PipeTimer) ToRaw() ArchiveWriter {
	return nil
}

func (p PipeTimer) Name() string {
	return "default-pipe-timer"
}

func (p *PipeTimer) Pipe(name RtName,args,params,values []Rter, info Info) *Result {
	switch string(name) {
	case "Begin", "RiBegin":
		p.start = time.Now()
		p.finish = p.start
		break
	case "End", "RiEnd":
		p.finish = time.Now()
		break
	}
	return Done()
}

func (p *PipeTimer) String() string {
	return fmt.Sprintf("pipe took %s", p.finish.Sub(p.start))
}

func (p *PipeTimer) Took() time.Duration {
	return p.finish.Sub(p.start)
}

/* Pipe RI output to gathered states */
type PipeToStats struct {
	Stats map[RtName]int
}

func (p *PipeToStats) ToRaw() ArchiveWriter {
	return nil
}

func (p PipeToStats) Name() string {
	return "default-pipe-to-stats"
}

func (p *PipeToStats) Pipe(name RtName,args,params,values []Rter, info Info) *Result {
	if p.Stats == nil {
		p.Stats = make(map[RtName]int, 0)
	}
	if _, exists := p.Stats[name]; !exists {
		p.Stats[name] = 0
	}
	p.Stats[name]++

	return Done()
}

type record struct {
	Name RtName
	Count int
}

type byCount []record

func (bc byCount) Len() int { return len(bc) }
func (bc byCount) Less(i,j int) bool { return bc[i].Count < bc[j].Count }
func (bc byCount) Swap(i,j int) { bc[i],bc[j] = bc[j],bc[i] }


func (p *PipeToStats) String() string {
	if p.Stats == nil {
		return "stats [empty]"
	}

	if len(p.Stats) == 0 {
		return "stats [empty]"
	}

	max := 0
	records := make([]record,0)

	for n, v := range p.Stats {
		if v > max {
			max = v
		}
		
		records = append(records,record{n,v})
	}

	sort.Sort(byCount(records))

	dfmt := "\t%0" + fmt.Sprintf("%d", len(fmt.Sprintf("%d", max))) + "d"

	out := fmt.Sprintf("stats %d [\n", len(p.Stats))
	for _, r := range records {
		out += fmt.Sprintf(dfmt + " call(s).....%s\n", r.Count,r.Name)
	}
	return out + "]\n"
}

/* Pipe RI output to file */
type PipeToFile struct {
	file *os.File
}

func (p *PipeToFile) ToRaw() ArchiveWriter {
	if p.file == nil {
		return nil
	}
	return p.file
}

func (p PipeToFile) Name() string {
	return "default-pipe-to-file"
}

func (p *PipeToFile) Pipe(name RtName,args,params,values []Rter, info Info) *Result {

	list := Mix(params,values)

	if info.Formal {
		name = name.Trim("Ri")
	}

	if name == "Begin" {
		fmt.Printf("PipeToFile Write, name=Begin...\n")
		if p.file != nil {
			return InError(ErrProtocolBotch)
		}
		file := "out.rib"
		if len(args) > 0 {
			if t, ok := args[0].(RtString); ok {
				file = string(t)
			}
		}

		f, err := os.Create(file)
		fmt.Printf("\tfile %s created\n",file)
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
		fmt.Printf("FileToPipe Write, name=End, p.File=%v\n",p.file)
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
		if _, err := p.file.Write([]byte(Serialise(args) + " " + Serialise(list) + "\n")); err != nil {
			return InError(err)
		}
		return Done()
	}

	if name != "##" {
	  /* TODO: change this to a configurable scheme, N-spaces or \t character etc */
		prefix := ""
		if info.PrettyPrint {
			for i := 0; i < info.Depth; i++ {
				prefix += "\t"
			}
		}
		

		if _, err := p.file.Write([]byte(prefix + name.Serialise() + " " + Serialise(args) + " " + Serialise(list) + "\n")); err != nil {
			return InError(err)
		}
		return Done()
	}

	if _, err := p.file.Write([]byte("##" + Serialise(args) + " " + Serialise(list) + "\n")); err != nil {
		return InError(err)
	}
	return Done()
}


type FilterStringHandles struct {
	/* FIXME: this should actually be a filter */
}

func (p *FilterStringHandles) ToRaw() ArchiveWriter {
	return nil
}

func (p FilterStringHandles) Name() string {
	return "default-filter-string-handles"
}

func (p *FilterStringHandles) Pipe(name RtName,args,params,values []Rter, info Info) *Result {

	/* TODO: add filter to only those proceedures the include light and object handles */
		
	args1 := make([]Rter,len(args))
	
	for i := 0; i < len(args); i ++ {
		if lh,ok := args[i].(RtLightHandle); ok {
			id,err := strconv.Atoi(string(lh))
			if err != nil {
				return InError(err)
			}
			args1[i] = RtInt(id)
			continue
		} 
		if oh,ok := args[i].(RtObjectHandle); ok {
			id,err := strconv.Atoi(string(oh))
			if err != nil {
				return InError(err)
			}
			args1[i] = RtInt(id)
			continue
		}
		if sh,ok := args[i].(RtShaderHandle); ok {
			id,err := strconv.Atoi(string(sh))
			if err != nil {
				return InError(err)
			}
			args1[i] = RtInt(id)
			continue
		}
		args1[i] = args[i]
	}

	return Next(name, args1,params,values, info)
}










