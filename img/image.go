package img

import (
	"github.com/DeedleFake/sdl"
	"unsafe"
)

// #cgo pkg-config: SDL2_image
// #include <SDL_image.h>
//
// #include "image.h"
import "C"

const (
	MAJOR_VERSION = C.SDL_IMAGE_MAJOR_VERSION
	MINOR_VERSION = C.SDL_IMAGE_MINOR_VERSION
	PATCHLEVEL    = C.SDL_IMAGE_PATCHLEVEL
)

func VERSION() *sdl.Version {
	var v sdl.Version
	C.IMAGE_VERSION(cVersion(&v))

	return &v
}

func Linked_Version() *sdl.Version {
	return (*sdl.Version)(unsafe.Pointer(C.IMG_Linked_Version()))
}

type InitFlags C.IMG_InitFlags

const (
	INIT_JPG  InitFlags = C.IMG_INIT_JPG
	INIT_PNG  InitFlags = C.IMG_INIT_PNG
	INIT_TIF  InitFlags = C.IMG_INIT_TIF
	INIT_WEBP InitFlags = C.IMG_INIT_WEBP
)

func Init(flags InitFlags) error {
	if C.IMG_Init(C.int(flags)) != 0 {
		return getError()
	}

	return nil
}

func Quit() {
	C.IMG_Quit()
}

func Load(name string) (*sdl.Surface, error) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	s := C.IMG_Load(cname)
	if s == nil {
		return nil, getError()
	}
	return (*sdl.Surface)(unsafe.Pointer(s)), nil
}

func Load_RW(rw *sdl.RWops, freesrc bool) (*sdl.Surface, error) {
	cfreesrc := C.int(0)
	if freesrc {
		cfreesrc = 1
	}

	s := C.IMG_Load_RW(cRWops(rw), cfreesrc)
	if s == nil {
		return nil, getError()
	}

	return (*sdl.Surface)(unsafe.Pointer(s)), nil
}

func LoadTexture(r *sdl.Renderer, name string) (*sdl.Texture, error) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	s := C.IMG_LoadTexture(cRenderer(r), cname)
	if s == nil {
		return nil, getError()
	}

	return goTexture(s), nil
}

func LoadTexture_RW(r *sdl.Renderer, rw *sdl.RWops, freesrc bool) (*sdl.Texture, error) {
	cfreesrc := C.int(0)
	if freesrc {
		cfreesrc = 1
	}

	s := C.IMG_LoadTexture_RW(cRenderer(r), cRWops(rw), cfreesrc)
	if s == nil {
		return nil, getError()
	}

	return goTexture(s), nil
}
