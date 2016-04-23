/* rigo/fragment.go */
package rigo

import (
	"sync"

	. "github.com/mae-global/rigo/ri"
	. "github.com/mae-global/rigo/ri/handles"
	. "github.com/mae-global/rigo/ris"
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
	List []Rter
}
	

type Fragment struct {
	Name RtName
	mux sync.RWMutex
	lights LightHandler
	objects ObjectHandler
	shaders ShaderHandler
	statements []FragmentStatement	
	cache map[RtShaderHandle]Shader
}


func (f *Fragment) Replay(ctx RiContexter) error {
	f.mux.RLock()
	defer f.mux.RUnlock()

	for _,s := range f.statements {
		if s.Name == DEPTH {
			if r,ok := s.Args[0].(RtInt); ok {
				ctx.Depth(int(r))
			}
			continue
		}

		if err := ctx.Write(s.Name,s.Args,s.List); err != nil {
			return err
		}
	}
	return nil
}

func (f *Fragment) ReplayPartial(ctx RiContexter,start,finish uint) error {
	f.mux.RLock()
	defer f.mux.RUnlock()

	var count uint
	for _,s := range f.statements {
		if count >= start && count < finish {

			if s.Name == DEPTH {
				if r,ok := s.Args[0].(RtInt); ok {
					ctx.Depth(int(r))
				}
				continue
			}

			if err := ctx.Write(s.Name,s.Args,s.List); err != nil {
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
		if s.Name == DEPTH {
			continue
		}
		count++
	}

	return count
}
	



func (f *Fragment) Write(name RtName,args, list []Rter) error {
	f.mux.Lock()
	defer f.mux.Unlock()

	f.statements = append(f.statements,FragmentStatement{Name:name,Args:args,List:list})
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
	f.statements = append(f.statements,FragmentStatement{Name:DEPTH,Args:[]Rter{RtInt(d)}})
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

func (f *Fragment) ShaderHandle() (RtShaderHandle,error) {
	f.mux.Lock()
	defer f.mux.Unlock()
	return f.shaders.Generate()
}

func (f *Fragment) CheckShaderHandle(h RtShaderHandle) error {
	f.mux.RLock()
	defer f.mux.Unlock()
	return f.shaders.Check(h)
}

func (f *Fragment) GetShader(h RtShaderHandle) Shader {
	f.mux.RLock()
	defer f.mux.RUnlock()
	if s,ok := f.cache[h]; ok {
		return s
	}
	return nil
}

func (f *Fragment) SetShader(h RtShaderHandle,s Shader)  {
	f.mux.Lock()
	defer f.mux.Unlock()
	f.cache[h] = s
}

func (f *Fragment) Shader(h RtShaderHandle) ShaderWriter {
	return f.GetShader(h)
}
	 

func NewFragment(name RtName,lights LightHandler,objects ObjectHandler,shaders ShaderHandler) *Fragment {
	if lights == nil {
		lights = NewLightNumberGenerator()
	}
	if objects == nil {
		objects = NewObjectNumberGenerator()
	}
	if shaders == nil {
		shaders = NewShaderNumberGenerator()
	}
	f := &Fragment{Name:name,lights:lights,objects:objects,shaders:shaders}
	f.statements = make([]FragmentStatement,0)
	f.cache = make(map[RtShaderHandle]Shader,0)
	return f
}




