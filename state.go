/* rigo/state.go */
package ri

import (
	"sync"
	"fmt"
)

var (
	ErrInvalidContextHandle error = fmt.Errorf("Invalid Context Handle")
	ErrContextAlreadyExists error = fmt.Errorf("Context Already Exists")
)

type context struct {
	Name string
}

var internal struct {
	contexts map[string]*context	
	active *context
	sync.Mutex
}

func init() {

	internal.contexts = make(map[string]*context,0)
	internal.active = nil
}

/* Begin creates and initializes a new rendering context */
func Begin(name string) error {
	
	internal.Lock()
	defer internal.Unlock()

	if _,exists := internal.contexts[name]; exists {
		return ErrContextAlreadyExists
	}

	ctx := &context{Name:name}

	internal.contexts[name] = ctx
	internal.active = ctx

	return nil
}

/* End terminate the active rendering context */
func End() error {

	/* TODO: finish up stream */
	return nil
}

/* GetContext returns a handle for the current active rendering context */
func GetContext() RtContextHandle {
	
	internal.Lock()
	defer internal.Unlock()
	
	return internal.active
}

/* Context sets the current active rendering context */
func Context(handle RtContextHandle) error {

	if handle == nil {
		return ErrInvalidContextHandle
	}

	ctx,ok := handle.(*context)
	if !ok {
		return ErrInvalidContextHandle
	}

	internal.Lock()
	defer internal.Unlock()

	internal.active = ctx

	return nil
}






