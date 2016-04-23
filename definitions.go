package rigo

import (
	"fmt"
	"sync"

	. "github.com/mae-global/rigo/ri"
)


const (
	Author RtToken = "RiGO;ver 0"
)


var (
	ErrInvalidContextHandle = fmt.Errorf("Invalid Context Handle")
	ErrContextAlreadyExists = fmt.Errorf("Context Already Exists")
	ErrNoActiveContext      = fmt.Errorf("No Active Context")
	ErrNotImplemented       = fmt.Errorf("Not Implemented")
	ErrPipeDone             = fmt.Errorf("Pipe Done")
	ErrEndOfLine            = fmt.Errorf("End of Line")
)


type Pipe struct {
	blocks []Piper
	sync.Mutex
}

func (p *Pipe) Last() Piper {
	p.Lock()
	defer p.Unlock()
	if len(p.blocks) == 0 {
		return nil
	}
	return p.blocks[len(p.blocks)-1]
}

func (p *Pipe) Append(block Piper) *Pipe {
	if block == nil {
		return p
	}
	p.Lock()
	defer p.Unlock()
	p.blocks = append(p.blocks, block)
	return p
}

/* Len get the length of the pipe */
func (p *Pipe) Len() int {
	p.Lock()
	defer p.Unlock()
	return len(p.blocks)
}

/* Get get a Piper object via index */
func (p *Pipe) Get(idx int) Piper {
	p.Lock()
	defer p.Unlock()
	if idx < 0 || idx >= len(p.blocks) {
		return nil
	}
	return p.blocks[idx]
}

/* GetByName get the first Piper object by name */
func (p *Pipe) GetByName(name string) Piper {
	p.Lock()
	defer p.Unlock()
	for _, b := range p.blocks {
		if b.Name() == name {
			return b
		}
	}
	return nil
}

func (p *Pipe) Run(name RtName, args,list []Rter, info Info) error {
	p.Lock()
	defer p.Unlock()

	if len(p.blocks) == 0 {
		return nil
	}

	nblocks := make([]Piper, 0)
	
	params,values := Unmix(list)

	for _, b := range p.blocks {
		if b == nil {
			continue
		}

		r := b.Pipe(name, args,params,values, info)
		if r.Err != nil {
			if r.Err == ErrPipeDone {
				nblocks = append(nblocks, b)
				continue
			}

			if r.Err == ErrEndOfLine {
				/* then mark b ready to be removed */
				continue
			}

			return r.Err
		}

		nblocks = append(nblocks, b)
		if r.Args != nil {
			args = r.Args
		}
		if r.Params != nil {
			params = make([]Rter,len(r.Params))
			copy(params,r.Params)
		}
		if r.Values != nil {
			values = make([]Rter,len(r.Values))
			copy(values,r.Values)
		}
	}

	p.blocks = nblocks
	return nil
}

func (p *Pipe) ToRaw() ArchiveWriter {
	p.Lock()
	defer p.Unlock()

	if len(p.blocks) == 0 {
		return nil
	}

	for _,b := range p.blocks {
		if b == nil {
			continue
		}
		if aw := b.ToRaw(); aw != nil {
			return aw
		}
	}
	return nil
}
	

func NewPipe() *Pipe {
	pipe := Pipe{}
	pipe.blocks = make([]Piper, 0)
	return &pipe
}


type Result struct {
	Name RtName
	Args []Rter
	Params []Rter
	Values []Rter
	Info *Info
	Err  error
}

func Done() *Result {
	return &Result{"", nil, nil,nil,nil, ErrPipeDone}
}

func Next(name RtName, args,params,values []Rter, info Info) *Result {
	return &Result{name, args,params,values, info.Copy(), nil}
}

func InError(err error) *Result {
	return &Result{"", nil, nil,nil,nil, err}
}

func Errored(message RtString) *Result {
	return &Result{"", nil, nil,nil,nil, fmt.Errorf(string(message))}
}

func EndOfLine() *Result {
	return &Result{"", nil, nil,nil,nil, ErrEndOfLine}
}


type Info struct {
	Name        string
	Depth       int
	Lights      uint
	Objects     uint
	Entity      bool
	Formal      bool
	PrettyPrint bool
}

func (info Info) Copy() *Info {
	n := Info{}
	n.Name = info.Name
	n.Depth = info.Depth
	n.Lights = info.Lights
	n.Objects = info.Objects
	n.Entity = info.Entity
	n.Formal = info.Formal
	n.PrettyPrint = info.PrettyPrint
	return &n
}


type Piper interface {
	/* name, []args,[]params,[]values,info */
	Pipe(RtName, []Rter,[]Rter,[]Rter, Info) *Result
	Name() string
	ToRaw() ArchiveWriter
}



