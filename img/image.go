package img

import (
	"github.com/DeedleFake/sdl"
	"unsafe"
)

// #cgo pkg-config: SDL2_image
// #include <SDL_image.h>
import "C"

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
