package ttf

import (
	"github.com/DeedleFake/sdl"
	"unsafe"
)

// #include <SDL_version.h>
import "C"

func cVersion(v *sdl.Version) *C.SDL_version {
	return (*C.SDL_version)(unsafe.Pointer(v))
}
