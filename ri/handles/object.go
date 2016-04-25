/* rigo/handles/handles.go */
package handles

import (
	"fmt"
	"sync"
)

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

/* ObjectUniqueGenerator */
type ObjectUniqueGenerator struct {
	mux sync.RWMutex
	size   int
	format string
}

func (g *ObjectUniqueGenerator) Generate() (RtObjectHandle,error) {
	g.mux.Lock()
	defer g.mux.Unlock()

	un,err := read(g.size)
	if err != nil {
		return "",err
	}
	return RtObjectHandle(fmt.Sprintf(g.format,un)),nil
}

func (g *ObjectUniqueGenerator) Check(h RtObjectHandle) error {
	return nil /* FIXME */
}

func (g *ObjectUniqueGenerator) Example() RtObjectHandle {
	g.mux.RLock()
	defer g.mux.RUnlock()
	return RtObjectHandle(fmt.Sprintf(g.format,readExample(g.size)))
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
 


