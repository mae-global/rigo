package ri

import (
	"fmt"
	"strconv"

	. "github.com/mae-global/rigo/ri/handles"
)

const USEDEBUG = false


type Contexter interface {
	Write(RtName, []Rter,[]Rter) error
	OpenRaw(RtToken) (ArchiveWriter,error)
	CloseRaw(RtToken) error
	Depth(int)
	LightHandle() (RtLightHandle, error)
	CheckLightHandle(RtLightHandle) error
	ObjectHandle() (RtObjectHandle, error)
	CheckObjectHandle(RtObjectHandle) error
	ShaderHandle() (RtShaderHandle, error)
	CheckShaderHandle(RtShaderHandle) error

	Shader(RtShaderHandle) ShaderWriter
}

type TestContext struct {
	lights  uint
	objects uint
	shaders uint
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



func (b *TestContext) Write(name RtName,args,list []Rter) error {
	if name != "ArchiveBegin" && b.raw != nil {
		return ErrNotSupported
	}
	/* combine args and parameterlist */
	args = append(args,list...)	
	astr := Serialise(args)
	if len(args) == 0 {
		return fmt.Errorf("%s",name)
	}
	return fmt.Errorf("%s %s", name,astr)
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

func (b *TestContext) ShaderHandle() (RtShaderHandle,error) {
	h := RtShaderHandle(fmt.Sprintf("%s",b.shaders))
	b.shaders++

	return h,nil
}

func (b *TestContext) CheckShaderHandle(h RtShaderHandle) error {
	if i,err := strconv.Atoi(string(h)); err != nil || uint(i) >= b.shaders {
		return ErrBadHandle
	}
	return nil
}

func (b *TestContext) Shader(h RtShaderHandle) ShaderWriter {
	return nil
}


func NewTest() *Ri {
	return &Ri{&TestContext{0, 0, 0, 0,nil}}
}

/* Ri is the main interface */
type Ri struct {
	Contexter
}


/* User special func for client libraries to write to */
func (r *Ri) User(w RterWriter) error {
	if w == nil {
		return ErrBadArgument
	}

	name,args,params := w.Write() 
	out := make([]Rter,len(args))
	copy(out,args)
	out = append(out,PARAMETERLIST)
	out = append(out,params...)

	return r.writef(name,out...)
}
	

func (r *Ri) writef(name RtName, parameterlist ...Rter) error {
	if r.Contexter == nil {
		return ErrProtocolBotch
	}

	para := -1
	/* find the actual parameterlist */
	for i,r := range parameterlist {
		if t,ok := r.(RtToken); ok {
			if t == PARAMETERLIST {
				para = i
				break
			}
		}
	}
	var args []Rter
	var list []Rter

	if para == -1 {
		args = parameterlist
	} else {
		args = make([]Rter,para)
		copy(args,parameterlist[:para])
		list = make([]Rter,len(parameterlist) - (para + 1))
		copy(list,parameterlist[para + 1:])
	}

	nlist := make([]Rter,len(list))		

	for i,r := range list {
		ar := r		
		if i % 2 == 0 {
			t,ok := r.(RtToken)
			if !ok {
				return ErrBadArgument  			
			}
		
			cl,ty,nam,n := ClassTypeNameCount(t)
			/* TODO: parse token, lookup class and type then check inputs of values */
			fmt.Printf("Debug,Ri.writef token, class=%s,type=%s,name=%s,count=%d\n",cl,ty,nam,n)

		} else {
			if a,ok := r.(RtString); ok {
				ar = RtStringArray{a}
				
			}
			if a,ok := r.(RtFloat); ok {
				ar = RtFloatArray{a}
			}
			if a,ok := r.(RtInt); ok {
				ar = RtIntArray{a}
			}
			if a,ok := r.(RtToken); ok {
				ar = RtTokenArray{a}
			}
			if a,ok := r.(RtPoint); ok {
				ar = RtPointArray{a}
			}
		}

		nlist[i] = ar
	}

	/* tokens to values is unbalanced */
	if len(nlist) % 2 != 0 {
		return ErrBadArgument 
	}

	if USEDEBUG && len(nlist) > 0  {
		args = append(args,DEBUGBARRIER)
	}

	return r.Write(name, args,nlist)
}



