/* rigo/fragment.go */
package rigo

import (
	"sync"

	. "github.com/mae-global/rigo/ri"
)

type Fragmenter interface {
	Replay() error
	ReplayPartial(start,finish uint) error
	Statements() uint
	Ri() *Ri
}
	


type FragmentStatement struct {
	Name RtName
	Args []Rter
}
	

type Fragment struct {
	Name RtName
	mux sync.RWMutex
	statements []FragmentStatement	
}

func (f *Fragment) Ri() *Ri {
	return &Ri{f}
}

func (f *Fragment) Replay(ctx Contexter) error {
	f.mux.RLock()
	defer f.mux.RUnlock()

	for _,s := range f.statements {
		if s.Name == "_depth" {
			if r,ok := s.Args[0].(RtInt); ok {
				ctx.Depth(int(r))
			}
			continue
		}

		if err := ctx.Write(s.Name,s.Args); err != nil {
			return err
		}
	}
	return nil
}

func (f *Fragment) ReplayPartial(ctx Contexter,start,finish uint) error {
	f.mux.RLock()
	defer f.mux.RUnlock()

	var count uint
	for _,s := range f.statements {
		if count >= start && count < finish {

			if s.Name == "_depth" {
				if r,ok := s.Args[0].(RtInt); ok {
					ctx.Depth(int(r))
				}
				continue
			}

			if err := ctx.Write(s.Name,s.Args); err != nil {
				return err
			}

			count ++
		}
		if count >= finish {
			break
		}
	}			

	return nil
}

func (f *Fragment) Statements() uint {
	f.mux.RLock()
	defer f.mux.RUnlock()

	var count uint
	for _,s := range f.statements {
		/* skip all the _depth statements */
		if s.Name == "_depth" {
			continue
		}
		count++
	}

	return count
}
	



func (f *Fragment) Write(name RtName, list []Rter) error {
	f.mux.Lock()
	defer f.mux.Unlock()

	f.statements = append(f.statements,FragmentStatement{Name:name,Args:list})
	return nil
}

func (f *Fragment) Depth(d int) {
	f.mux.Lock()
	defer f.mux.Unlock()
	f.statements = append(f.statements,FragmentStatement{Name:"_depth",Args:[]Rter{RtInt(d)}})
}

func (f *Fragment) LightHandle() (RtLightHandle,error) {
	return 0,ErrNotImplemented
}

func (f *Fragment) CheckLightHandle(lh RtLightHandle) error {
	return ErrNotImplemented
}

func (f *Fragment) ObjectHandle() (RtObjectHandle, error) {
	return 0,ErrNotImplemented
}

func (f *Fragment) CheckObjectHandle(oh RtObjectHandle) error {
	return ErrNotImplemented
}

func NewFragment(name RtName) *Fragment {
	f := &Fragment{Name:name}
	f.statements = make([]FragmentStatement,0)
	return f
}
