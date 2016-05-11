package rigo

import (
	"sync"

	. "github.com/mae-global/rigo/ri"
)


/* HandleManager -- Collects all the handle management */
type HandleManagerer interface {
	
	/* Light... */
	LightHandle() (RtLightHandle,error)
	CheckLightHandle(RtLightHandle) error 

	/* Object... */
	ObjectHandle() (RtObjectHandle,error)
	CheckObjectHandle(RtObjectHandle) error

	/* Shader... */
	ShaderHandle() (RtShaderHandle,error)
	CheckShaderHandle(RtShaderHandle) error	
}


type HandleManager struct {
	mux sync.RWMutex

	objects 	ObjectHandler
	lights 		LightHandler
	shaders 	ShaderHandler
}


func (mgr *HandleManager) LightHandle() (RtLightHandle,error) {
	mgr.mux.Lock()
	defer mgr.mux.Unlock()
	return mgr.lights.Generate()	
}

func (mgr *HandleManager) CheckLightHandle(h RtLightHandle) error {
	mgr.mux.RLock()
	defer mgr.mux.RUnlock()
	return mgr.lights.Check(h)
}


func (mgr *HandleManager) ObjectHandle() (RtObjectHandle,error) {
	mgr.mux.Lock()
	defer mgr.mux.Unlock()
	return mgr.objects.Generate()
}

func (mgr *HandleManager) CheckObjectHandle(h RtObjectHandle) error {
	mgr.mux.RLock()
	defer mgr.mux.RUnlock()
	return mgr.objects.Check(h)
}


func (mgr *HandleManager) ShaderHandle() (RtShaderHandle,error) {
	mgr.mux.Lock()
	defer mgr.mux.Unlock()
	return mgr.shaders.Generate()
}

func (mgr *HandleManager) CheckShaderHandle(h RtShaderHandle) error {
	mgr.mux.RLock()
	defer mgr.mux.RUnlock()
	return mgr.shaders.Check(h)
}

func NewHandleManager(object ObjectHandler,light LightHandler,shader ShaderHandler) *HandleManager {
	
	mgr := new(HandleManager)
	if object == nil {
		object = NewObjectNumberGenerator()
	}
	if light == nil {
		light = NewLightNumberGenerator()
	}
	if shader == nil {
		shader = NewShaderNumberGenerator()
	}

	mgr.objects = object
	mgr.lights  = light
	mgr.shaders = shader


	return mgr
}





