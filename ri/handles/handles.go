/* rigo/handles/handles.go */
package handles

import (
	"fmt"
	"sync"
)

/* RtLightHandle 
 * note -- updated to string handles [Siggraph 2009 course 9; Byron Bashforth] */
type RtLightHandle string

func (l RtLightHandle) String() string {
	return l.Serialise()
}

func (l RtLightHandle) Serialise() string {
	return fmt.Sprintf("\"%s\"", string(l))
}

type LightHandler interface {
	Generate() (RtLightHandle,error)
	Check(RtLightHandle) error
}

/* RtObjectHandle 
 * note -- updated to string handles [Siggraph 2009 course 9; Byron Bashforth] */
type RtObjectHandle string

func (l RtObjectHandle) String() string {
	return l.Serialise()
}

func (l RtObjectHandle) Serialise() string {
	return fmt.Sprintf("\"%s\"", string(l))
}


type ObjectHandler interface {
	Generate() (RtObjectHandle,error)
	Check(RtObjectHandle) error
}



/* basic generators */
type LightNumberGenerator struct {
	current uint
	mux sync.RWMutex
}

func (lng *LightNumberGenerator) Generate() (RtLightHandle,error) {
	lng.mux.Lock()
	defer lng.mux.Unlock()

	h := fmt.Sprintf("%d",lng.current)
	lng.current++
	return RtLightHandle(h),nil
}

func (lng *LightNumberGenerator) Check(h RtLightHandle) error {
	return nil
}

type ObjectNumberGenerator struct {
	current uint
	mux sync.RWMutex
}

func (ong *ObjectNumberGenerator) Generate() (RtObjectHandle,error) {
	ong.mux.Lock()
	defer ong.mux.Unlock()

	h := fmt.Sprintf("%d",ong.current)
	ong.current++
	return RtObjectHandle(h),nil
}

func (ong *ObjectNumberGenerator) Check(h RtObjectHandle) error {
	return nil
}





	


