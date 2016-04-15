package ri

import (
	"fmt"
	"strconv"

	. "github.com/mae-global/rigo/ri/handles"
)

type Contexter interface {
	Write(RtName, []Rter) error
	OpenRaw(RtToken) (ArchiveWriter,error)
	CloseRaw(RtToken) error
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

	raw *testArchiveWriter /* single instance only */
}

type testArchiveWriter struct {
	id RtToken
	content []byte
}

func (w *testArchiveWriter) Write(c []byte) (int,error) {
	if w.content == nil {
		w.content = make([]byte,len(c))
		copy(w.content,c)
		return len(c),nil
	}
	ncontent := make([]byte,len(w.content) + len(c))
	copy(ncontent[:len(w.content)],w.content)
	copy(ncontent[len(w.content):],c)
	return len(c),nil
}



func (b *TestContext) Write(name RtName, list []Rter) error {
	if name != "ArchiveBegin" && b.raw != nil {
		return ErrNotSupported
	}

	str := Serialise(list)
	if len(str) == 0 {
		return fmt.Errorf("%s", name)
	}

	return fmt.Errorf("%s %s", name, str)
}

func (b *TestContext) OpenRaw(id RtToken) (ArchiveWriter,error) {
	if b.raw != nil {
		return nil,ErrNotSupported
	}

	b.raw = &testArchiveWriter{id:id}

	return b.raw,nil
}

func (b *TestContext) CloseRaw(id RtToken) error {
	if b.raw == nil {
		return ErrNotSupported
	}

	if b.raw.id != id {
		return ErrNotSupported
	}

	b.raw = nil
	return nil
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
	return &Ri{&TestContext{0, 0, 0,nil}}
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
