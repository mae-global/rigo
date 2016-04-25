/* rigo/handles/handles.go */
package handles

import (
	"fmt"
	"sync"
)

/* RtShaderHandler - used in RIS */
type RtShaderHandle string

func (l RtShaderHandle) Type() string {
	return "shaderhandle"
}

func (l RtShaderHandle) String() string {
	return l.Serialise()
}

func (l RtShaderHandle) Serialise() string {
	return fmt.Sprintf("\"%s\"",string(l))
}


/* ShaderHandler */
type ShaderHandler interface {
	Generate() (RtShaderHandle,error)
	Check(RtShaderHandle) error 
	Example() RtShaderHandle
}
/* ShaderNumberGenerator */
type ShaderNumberGenerator struct {
	current uint
	mux sync.RWMutex
	format string
}

func (g *ShaderNumberGenerator) Generate() (RtShaderHandle,error) {
	g.mux.Lock()
	defer g.mux.Unlock()
	
	h := fmt.Sprintf(g.format,g.current)
	g.current ++
	return RtShaderHandle(h),nil
}

func (g *ShaderNumberGenerator) Check(h RtShaderHandle) error {
	return nil /* FIXME */
}

func (g *ShaderNumberGenerator) Example() RtShaderHandle {
	g.mux.RLock()
	defer g.mux.RUnlock()
	return RtShaderHandle(fmt.Sprintf(g.format,0))
}

/* NewShaderNumberGenerator */
func NewShaderNumberGenerator() *ShaderNumberGenerator {
	return &ShaderNumberGenerator{format:"%d"}
}

/* NewPrefixShaderNumberGenerator */
func NewPrefixShaderNumberGenerator(prefix string) *ShaderNumberGenerator {
	if len(prefix) == 0 {
		return NewShaderNumberGenerator()
	}
	return &ShaderNumberGenerator{format:prefix + "%d"}
}

/* ShaderUniqueGenerator */
type ShaderUniqueGenerator struct {
	mux sync.RWMutex
	size   int
	format string
}

func (g *ShaderUniqueGenerator) Generate() (RtShaderHandle,error) {
	g.mux.Lock()
	defer g.mux.Unlock()

	un,err := read(g.size)
	if err != nil {
		return "",err
	}

	h := fmt.Sprintf(g.format,un)
	return RtShaderHandle(h),nil
}

func (g *ShaderUniqueGenerator) Check(h RtShaderHandle) error {
	return nil /* FIXME */
}

func (g *ShaderUniqueGenerator) Example() RtShaderHandle {
	g.mux.RLock()
	defer g.mux.RUnlock()
	return RtShaderHandle(fmt.Sprintf(g.format,readExample(g.size)))
}	

/* NewShaderUniqueGenerator */
func NewShaderUniqueGenerator() *ShaderUniqueGenerator {
	return &ShaderUniqueGenerator{format:"%s",size:4}
}

/* NewPrefixShaderUniqueGenerator */
func NewPrefixShaderUniqueGenerator(prefix string) *ShaderUniqueGenerator {
	if len(prefix) == 0 {
		return NewShaderUniqueGenerator()
	}
	return &ShaderUniqueGenerator{format:prefix + "%s",size:4}
}
 


