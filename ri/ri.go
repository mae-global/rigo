package ri

import (
	"fmt"
	"strconv"

	. "github.com/mae-global/rigo/ri/handles"
)

type Contexter interface {
	Write(RtName, []Rter) error
	Depth(int)
	LightHandle() (RtLightHandle, error)
	CheckLightHandle(RtLightHandle) error
	ObjectHandle() (RtObjectHandle, error)
	CheckObjectHandle(RtObjectHandle) error
}

type TestContext struct {
	lights  uint
	objects uint
	depth   int
}

func (b *TestContext) Write(name RtName, list []Rter) error {
	str := Serialise(list)
	if len(str) == 0 {
		return fmt.Errorf("%s", name)
	}

	return fmt.Errorf("%s %s", name, str)
}

func (b *TestContext) Depth(d int) {
	b.depth += d
}

func (b *TestContext) LightHandle() (RtLightHandle, error) {
	h := RtLightHandle(fmt.Sprintf("%d",b.lights))
	b.lights++

	return h, nil
}

func (b *TestContext) CheckLightHandle(h RtLightHandle) error {
	if i,err := strconv.Atoi(string(h)); err != nil || uint(i) >= b.lights {
		return ErrBadHandle
	} 
	return nil
}

func (b *TestContext) ObjectHandle() (RtObjectHandle, error) {
	h := RtObjectHandle(fmt.Sprintf("%d",b.objects))
	b.objects++
	return h, nil
}

func (b *TestContext) CheckObjectHandle(h RtObjectHandle) error {
	if i,err := strconv.Atoi(string(h)); err != nil || uint(i) >= b.objects {
		return ErrBadHandle
	}
	return nil
}

func NewTest() *Ri {
	return &Ri{&TestContext{0, 0, 0}}
}

/* Ri is the main interface */
type Ri struct {
	Contexter
}

func (r *Ri) writef(name RtName, parameterlist ...Rter) error {
	if r.Contexter == nil {
		return ErrProtocolBotch
	}
	return r.Write(name, parameterlist)
}
