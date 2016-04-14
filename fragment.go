/* rigo/fragment.go */
package rigo

import (
	"sync"

	. "github.com/mae-global/rigo/ri"
)

type Fragmenter interface {
	Replay() error
	ReplayPartial(start,finish uint) error
	Statements() (uint,uint)
}
	


type FragmentStatement struct {
	Name RtName
	Args []Rter
}
	

type Fragment struct {
	mux sync.RWMutex
	statements []FragmentStatement	

	Info
}

func (f *Fragment) Replay() error {

	return ErrNotImplemented
}

func (f *Fragment) ReplayPartial(start,finish uint) error {

	return ErrNotImplemented
}




func (f *Fragment) Write(name RtName, list []Rter) error {
	f.mux.Lock()
	defer f.mux.Unlock()
	if f.Formal {
		name = name.Prefix("Ri")
	}
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


