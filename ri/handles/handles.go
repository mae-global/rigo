/* rigo/handles/handles.go */
package handles

import (
	"fmt"
	"sync"
	"crypto/rand"
	"io"
	"encoding/hex"
)


/* RtLightHandle 
 * note -- updated to string handles [Siggraph 2009 course 9; Byron Bashforth] */
type RtLightHandle string

func (l RtLightHandle) Type() string {
	return "lighthandle"
}

func (l RtLightHandle) String() string {
	return l.Serialise()
}

func (l RtLightHandle) Serialise() string {
	return fmt.Sprintf("\"%s\"", string(l))
}

type LightHandler interface {
	Generate() (RtLightHandle,error)
	Check(RtLightHandle) error
	Example() RtLightHandle
}

/* RtObjectHandle 
 * note -- updated to string handles [Siggraph 2009 course 9; Byron Bashforth] */
type RtObjectHandle string

func (l RtObjectHandle) Type() string {
	return "objecthandle"
}

func (l RtObjectHandle) String() string {
	return l.Serialise()
}

func (l RtObjectHandle) Serialise() string {
	return fmt.Sprintf("\"%s\"", string(l))
}

/* ObjectHandler */
type ObjectHandler interface {
	Generate() (RtObjectHandle,error)
	Check(RtObjectHandle) error
	Example() RtObjectHandle
}


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

	b := make([]byte,g.size)
	n,err := io.ReadFull(rand.Reader,b)
	if err != nil {
		return "",err
	}
	h := fmt.Sprintf(g.format,hex.EncodeToString(b[:n]))
	return RtShaderHandle(h),nil
}

func (g *ShaderUniqueGenerator) Check(h RtShaderHandle) error {
	return nil /* FIXME */
}

func (g *ShaderUniqueGenerator) Example() RtShaderHandle {
	g.mux.RLock()
	defer g.mux.RUnlock()
	example := []byte("abcdefghijklnmopqrstuvw123456789") /* FIXME */
	return RtShaderHandle(fmt.Sprintf(g.format,hex.EncodeToString(example[:g.size])))
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
 






/* LightNumberGenerator implements the old style of int handles */
type LightNumberGenerator struct {
	current uint
	mux sync.RWMutex
	format string
}

func (g *LightNumberGenerator) Generate() (RtLightHandle,error) {
	g.mux.Lock()
	defer g.mux.Unlock()
	
	h := fmt.Sprintf(g.format,g.current)
	g.current ++
	return RtLightHandle(h),nil
}

func (g *LightNumberGenerator) Check(h RtLightHandle) error {
	return nil /* FIXME */
}

func (g *LightNumberGenerator) Example() RtLightHandle {
	g.mux.RLock()
	defer g.mux.RUnlock()
	return RtLightHandle(fmt.Sprintf(g.format,0))
}

/* NewLightNumberGenerator */
func NewLightNumberGenerator() *LightNumberGenerator {
	return &LightNumberGenerator{format:"%d"}
}

/* NewPrefixLightNumberGenerator */
func NewPrefixLightNumberGenerator(prefix string) *LightNumberGenerator {
	if len(prefix) == 0 {
		return NewLightNumberGenerator()
	}
	return &LightNumberGenerator{format:prefix + "%d"}
}

/* ObjectNumberGenerator implements the old style of int handles */
type ObjectNumberGenerator struct {
	current uint
	mux sync.RWMutex
	format string
}

func (g *ObjectNumberGenerator) Generate() (RtObjectHandle,error) {
	g.mux.Lock()
	defer g.mux.Unlock()

	h := fmt.Sprintf(g.format,g.current)
	g.current ++
	return RtObjectHandle(h),nil
}

func (g *ObjectNumberGenerator) Check(h RtObjectHandle) error {
	return nil /* FIXME */
}

func (g *ObjectNumberGenerator) Example() RtObjectHandle {
	g.mux.RLock()
	defer g.mux.RUnlock()
	return RtObjectHandle(fmt.Sprintf(g.format,0))
}

/* NewObjectNumberGenerator */
func NewObjectNumberGenerator() *ObjectNumberGenerator {
	return &ObjectNumberGenerator{format:"%d"}
}

/* NewPrefixObjectNumberGenerator */
func NewPrefixObjectNumberGenerator(prefix string) *ObjectNumberGenerator {
	if len(prefix) == 0 {
		return NewObjectNumberGenerator()
	}
	return &ObjectNumberGenerator{format:prefix + "%d"}
}


/* LightUniqueGenerator */
type LightUniqueGenerator struct {
	mux sync.RWMutex
	size   int
	format string
}

func (g *LightUniqueGenerator) Generate() (RtLightHandle,error) {
	g.mux.Lock()
	defer g.mux.Unlock()

	b := make([]byte,g.size)
	n,err := io.ReadFull(rand.Reader,b)
	if err != nil {
		return "",err
	}
	h := fmt.Sprintf(g.format,hex.EncodeToString(b[:n]))
	return RtLightHandle(h),nil
}

func (g *LightUniqueGenerator) Check(h RtLightHandle) error {
	return nil /* FIXME */
}

func (g *LightUniqueGenerator) Example() RtLightHandle {
	g.mux.RLock()
	defer g.mux.RUnlock()
	example := []byte("abcdefghijklnmopqrstuvw123456789") /* FIXME */
	return RtLightHandle(fmt.Sprintf(g.format,hex.EncodeToString(example[:g.size])))
}	

/* NewLightUniqueGenerator */
func NewLightUniqueGenerator() *LightUniqueGenerator {
	return &LightUniqueGenerator{format:"%s",size:4}
}

/* NewPrefixLightUniqueGenerator */
func NewPrefixLightUniqueGenerator(prefix string) *LightUniqueGenerator {
	if len(prefix) == 0 {
		return NewLightUniqueGenerator()
	}
	return &LightUniqueGenerator{format:prefix + "%s",size:4}
}
 

/* ObjectUniqueGenerator */
type ObjectUniqueGenerator struct {
	mux sync.RWMutex
	size   int
	format string
}

func (g *ObjectUniqueGenerator) Generate() (RtObjectHandle,error) {
	g.mux.Lock()
	defer g.mux.Unlock()

	b := make([]byte,g.size)
	n,err := io.ReadFull(rand.Reader,b)
	if err != nil {
		return "",err
	}
	h := fmt.Sprintf(g.format,hex.EncodeToString(b[:n]))
	return RtObjectHandle(h),nil
}

func (g *ObjectUniqueGenerator) Check(h RtObjectHandle) error {
	return nil /* FIXME */
}

func (g *ObjectUniqueGenerator) Example() RtObjectHandle {
	g.mux.RLock()
	defer g.mux.RUnlock()
	example := []byte("abcdefghijklnmopqrstuvw123456789") /* FIXME */
	return RtObjectHandle(fmt.Sprintf(g.format,hex.EncodeToString(example[:g.size])))
}	

/* NewObjectUniqueGenerator */
func NewObjectUniqueGenerator() *ObjectUniqueGenerator {
	return &ObjectUniqueGenerator{format:"%s",size:4}
}

/* NewPrefixObjectUniqueGenerator */
func NewPrefixObjectUniqueGenerator(prefix string) *ObjectUniqueGenerator {
	if len(prefix) == 0 {
		return NewObjectUniqueGenerator()
	}
	return &ObjectUniqueGenerator{format:prefix + "%s",size:4}
}
 






	


