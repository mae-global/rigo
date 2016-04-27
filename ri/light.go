/* rigo/ri/handles.go */
package ri

import (
	"fmt"
	"sync"
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

func (l RtLightHandle) Equal(o Rter) bool {
	if other, ok := o.(RtLightHandle); ok {
		return (other == l)
	}
	return false
}

/* LightHandler */
type LightHandler interface {
	/* Generate or provide a RtLightHandle */
	Generate() (RtLightHandle, error)
	/* Check an existing RtLightHandle */
	Check(RtLightHandle) error
	/* Example output of this generator */
	Example() RtLightHandle
}

/* LightNumberGenerator implements the old style of int handles */
type LightNumberGenerator struct {
	current uint
	mux     sync.RWMutex
	format  string
}

func (g *LightNumberGenerator) Generate() (RtLightHandle, error) {
	g.mux.Lock()
	defer g.mux.Unlock()

	h := fmt.Sprintf(g.format, g.current)
	g.current++
	return RtLightHandle(h), nil
}

func (g *LightNumberGenerator) Check(h RtLightHandle) error {
	return nil /* FIXME */
}

func (g *LightNumberGenerator) Example() RtLightHandle {
	g.mux.RLock()
	defer g.mux.RUnlock()
	return RtLightHandle(fmt.Sprintf(g.format, 0))
}

/* NewLightNumberGenerator */
func NewLightNumberGenerator() *LightNumberGenerator {
	return &LightNumberGenerator{format: "%d"}
}

/* NewPrefixLightNumberGenerator */
func NewPrefixLightNumberGenerator(prefix string) *LightNumberGenerator {
	if len(prefix) == 0 {
		return NewLightNumberGenerator()
	}
	return &LightNumberGenerator{format: prefix + "%d"}
}

/* LightUniqueGenerator */
type LightUniqueGenerator struct {
	mux    sync.RWMutex
	size   int
	format string
}

func (g *LightUniqueGenerator) Generate() (RtLightHandle, error) {
	g.mux.Lock()
	defer g.mux.Unlock()

	un, err := read(g.size)
	if err != nil {
		return "", err
	}
	return RtLightHandle(fmt.Sprintf(g.format, un)), nil
}

func (g *LightUniqueGenerator) Check(h RtLightHandle) error {
	return nil /* FIXME */
}

func (g *LightUniqueGenerator) Example() RtLightHandle {
	g.mux.RLock()
	defer g.mux.RUnlock()
	return RtLightHandle(fmt.Sprintf(g.format, readExample(g.size)))
}

/* NewLightUniqueGenerator */
func NewLightUniqueGenerator() *LightUniqueGenerator {
	return &LightUniqueGenerator{format: "%s", size: 4}
}

/* NewPrefixLightUniqueGenerator */
func NewPrefixLightUniqueGenerator(prefix string) *LightUniqueGenerator {
	if len(prefix) == 0 {
		return NewLightUniqueGenerator()
	}
	return &LightUniqueGenerator{format: prefix + "%s", size: 4}
}
