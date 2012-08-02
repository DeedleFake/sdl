package sdl

import (
	"fmt"
	"unsafe"
)

// #cgo pkg-config: sdl2
//
// #include <SDL.h>
//
// #include "version.h"
import "C"

type Version struct {
	Major uint8
	Minor uint8
	Patch uint8
}

func (v *Version) c() *C.SDL_version {
	return (*C.SDL_version)(unsafe.Pointer(v))
}

func (v *Version) String() string {
	return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
}

const (
	MAJOR_VERSION = C.SDL_MAJOR_VERSION
	MINOR_VERSION = C.SDL_MINOR_VERSION
	PATCHLEVEL    = C.SDL_PATCHLEVEL
)

func VERSION() *Version {
	var v Version
	C.VERSION(v.c())

	return &v
}

func VERSTIONNUM(x, y, z uint8) uint8 {
	return uint8(C.VERSIONNUM(C.Uint8(x), C.Uint8(y), C.Uint8(z)))
}

const COMPILEDVERSION = C.SDL_COMPILEDVERSION

func VERSION_ATLEAST(x, y, z uint8) bool {
	return C.VERSION_ATLEAST(C.Uint8(x), C.Uint8(y), C.Uint8(z)) == 1
}

func GetVersion() *Version {
	var v Version
	C.SDL_GetVersion(v.c())

	return &v
}

func GetRevision() string {
	return C.GoString(C.SDL_GetRevision())
}

func GetRevisionNumber() int {
	return int(C.SDL_GetRevisionNumber())
}
