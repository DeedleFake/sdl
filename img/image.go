package img

import (
	"github.com/salviati/sdl"
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

func LoadTyped_RW(rw *sdl.RWops, freesrc bool, t string) (*sdl.Surface, error) {
	cfreesrc := C.int(0)
	if freesrc {
		cfreesrc = 1
	}

	ct := C.CString(t)
	defer C.free(unsafe.Pointer(ct))

	s := C.IMG_LoadTyped_RW(cRWops(rw), cfreesrc, ct)
	if s == nil {
		return nil, getError()
	}

	return (*sdl.Surface)(unsafe.Pointer(s)), nil
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

func LoadTextureTyped_RW(r *sdl.Renderer, rw *sdl.RWops, freesrc bool, t string) (*sdl.Texture, error) {
	cfreesrc := C.int(0)
	if freesrc {
		cfreesrc = 1
	}

	ct := C.CString(t)
	defer C.free(unsafe.Pointer(ct))

	s := C.IMG_LoadTextureTyped_RW(cRenderer(r), cRWops(rw), cfreesrc, ct)
	if s == nil {
		return nil, getError()
	}

	return goTexture(s), nil
}

// NOTE: Purposefully not implementing IMG_InvertAlpha().

func IsICO(rw *sdl.RWops) bool {
	return C.IMG_isICO(cRWops(rw)) != 0
}

func IsCUR(rw *sdl.RWops) bool {
	return C.IMG_isCUR(cRWops(rw)) != 0
}

func IsBMP(rw *sdl.RWops) bool {
	return C.IMG_isBMP(cRWops(rw)) != 0
}

func IsGIF(rw *sdl.RWops) bool {
	return C.IMG_isGIF(cRWops(rw)) != 0
}

func IsJPG(rw *sdl.RWops) bool {
	return C.IMG_isJPG(cRWops(rw)) != 0
}

func IsLBM(rw *sdl.RWops) bool {
	return C.IMG_isLBM(cRWops(rw)) != 0
}

func IsPCX(rw *sdl.RWops) bool {
	return C.IMG_isPCX(cRWops(rw)) != 0
}

func IsPNG(rw *sdl.RWops) bool {
	return C.IMG_isPNG(cRWops(rw)) != 0
}

func IsPNM(rw *sdl.RWops) bool {
	return C.IMG_isPNM(cRWops(rw)) != 0
}

func IsTIF(rw *sdl.RWops) bool {
	return C.IMG_isTIF(cRWops(rw)) != 0
}

func IsXCF(rw *sdl.RWops) bool {
	return C.IMG_isXCF(cRWops(rw)) != 0
}

func IsXPM(rw *sdl.RWops) bool {
	return C.IMG_isXPM(cRWops(rw)) != 0
}

func IsXV(rw *sdl.RWops) bool {
	return C.IMG_isXV(cRWops(rw)) != 0
}

func IsWEBP(rw *sdl.RWops) bool {
	return C.IMG_isWEBP(cRWops(rw)) != 0
}

func LoadICO_RW(rw *sdl.RWops) (*sdl.Surface, error) {
	s := C.IMG_LoadICO_RW(cRWops(rw))
	if s == nil {
		return nil, getError()
	}

	return (*sdl.Surface)(unsafe.Pointer(s)), nil
}

func LoadCUR_RW(rw *sdl.RWops) (*sdl.Surface, error) {
	s := C.IMG_LoadCUR_RW(cRWops(rw))
	if s == nil {
		return nil, getError()
	}

	return (*sdl.Surface)(unsafe.Pointer(s)), nil
}

func LoadBMP_RW(rw *sdl.RWops) (*sdl.Surface, error) {
	s := C.IMG_LoadBMP_RW(cRWops(rw))
	if s == nil {
		return nil, getError()
	}

	return (*sdl.Surface)(unsafe.Pointer(s)), nil
}

func LoadGIF_RW(rw *sdl.RWops) (*sdl.Surface, error) {
	s := C.IMG_LoadGIF_RW(cRWops(rw))
	if s == nil {
		return nil, getError()
	}

	return (*sdl.Surface)(unsafe.Pointer(s)), nil
}

func LoadJPG_RW(rw *sdl.RWops) (*sdl.Surface, error) {
	s := C.IMG_LoadJPG_RW(cRWops(rw))
	if s == nil {
		return nil, getError()
	}

	return (*sdl.Surface)(unsafe.Pointer(s)), nil
}

func LoadLBM_RW(rw *sdl.RWops) (*sdl.Surface, error) {
	s := C.IMG_LoadLBM_RW(cRWops(rw))
	if s == nil {
		return nil, getError()
	}

	return (*sdl.Surface)(unsafe.Pointer(s)), nil
}

func LoadPCX_RW(rw *sdl.RWops) (*sdl.Surface, error) {
	s := C.IMG_LoadPCX_RW(cRWops(rw))
	if s == nil {
		return nil, getError()
	}

	return (*sdl.Surface)(unsafe.Pointer(s)), nil
}

func LoadPNG_RW(rw *sdl.RWops) (*sdl.Surface, error) {
	s := C.IMG_LoadPNG_RW(cRWops(rw))
	if s == nil {
		return nil, getError()
	}

	return (*sdl.Surface)(unsafe.Pointer(s)), nil
}

func LoadPNM_RW(rw *sdl.RWops) (*sdl.Surface, error) {
	s := C.IMG_LoadPNM_RW(cRWops(rw))
	if s == nil {
		return nil, getError()
	}

	return (*sdl.Surface)(unsafe.Pointer(s)), nil
}

func LoadTGA_RW(rw *sdl.RWops) (*sdl.Surface, error) {
	s := C.IMG_LoadTGA_RW(cRWops(rw))
	if s == nil {
		return nil, getError()
	}

	return (*sdl.Surface)(unsafe.Pointer(s)), nil
}

func LoadTIF_RW(rw *sdl.RWops) (*sdl.Surface, error) {
	s := C.IMG_LoadTIF_RW(cRWops(rw))
	if s == nil {
		return nil, getError()
	}

	return (*sdl.Surface)(unsafe.Pointer(s)), nil
}

func LoadXCF_RW(rw *sdl.RWops) (*sdl.Surface, error) {
	s := C.IMG_LoadXCF_RW(cRWops(rw))
	if s == nil {
		return nil, getError()
	}

	return (*sdl.Surface)(unsafe.Pointer(s)), nil
}

func LoadXPM_RW(rw *sdl.RWops) (*sdl.Surface, error) {
	s := C.IMG_LoadXPM_RW(cRWops(rw))
	if s == nil {
		return nil, getError()
	}

	return (*sdl.Surface)(unsafe.Pointer(s)), nil
}

func LoadXV_RW(rw *sdl.RWops) (*sdl.Surface, error) {
	s := C.IMG_LoadXV_RW(cRWops(rw))
	if s == nil {
		return nil, getError()
	}

	return (*sdl.Surface)(unsafe.Pointer(s)), nil
}

func LoadWEBP_RW(rw *sdl.RWops) (*sdl.Surface, error) {
	s := C.IMG_LoadWEBP_RW(cRWops(rw))
	if s == nil {
		return nil, getError()
	}

	return (*sdl.Surface)(unsafe.Pointer(s)), nil
}

// XPM's are C source files, so IMG_ReadXPMFromArray() is somewhat
// pointless in Go...
