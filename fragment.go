/* rigo/fragment.go */
package rigo

import (
	"sync"

	. "github.com/mae-global/rigo/ri"
	. "github.com/mae-global/rigo/ri/handles"
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
	lights LightHandler
	objects ObjectHandler
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

func (f *Fragment) OpenRaw(id RtToken) (ArchiveWriter,error) {
	return nil,ErrNotSupported
}

func (f *Fragment) CloseRaw(id RtToken) error {
	return ErrNotSupported
}


func (f *Fragment) Depth(d int) {
	f.mux.Lock()
	defer f.mux.Unlock()
	f.statements = append(f.statements,FragmentStatement{Name:"_depth",Args:[]Rter{RtInt(d)}})
}

func (f *Fragment) LightHandle() (RtLightHandle,error) {
	f.mux.Lock()
	defer f.mux.Unlock()
	return f.lights.Generate()
}

func (f *Fragment) CheckLightHandle(lh RtLightHandle) error {
	f.mux.RLock()
	defer f.mux.RUnlock()
	return f.lights.Check(lh)
}

func (f *Fragment) ObjectHandle() (RtObjectHandle, error) {
	f.mux.Lock()
	defer f.mux.RLock()
	return f.objects.Generate()
}

func (f *Fragment) CheckObjectHandle(oh RtObjectHandle) error {
	f.mux.RLock()
	defer f.mux.RUnlock()
	return f.objects.Check(oh)
}

func NewFragment(name RtName) *Fragment {
	return NewCustomFragment(name,nil,nil)
}

func NewCustomFragment(name RtName,lights LightHandler,objects ObjectHandler) *Fragment {
	if lights == nil {
		lights = NewLightNumberGenerator()
	}
	if objects == nil {
		objects = NewObjectNumberGenerator()
	}
	f := &Fragment{Name:name,lights:lights,objects:objects}
	f.statements = make([]FragmentStatement,0)
	return f
}




