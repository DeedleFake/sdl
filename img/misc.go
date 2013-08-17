package img

import (
	"github.com/salviati/sdl"
	"unsafe"
)

// #include <SDL_version.h>
// #include <SDL_surface.h>
// #include <SDL_rwops.h>
// #include <SDL_render.h>
import "C"

func cVersion(v *sdl.Version) *C.SDL_version {
	return (*C.SDL_version)(unsafe.Pointer(v))
}

func cSurface(s *sdl.Surface) *C.SDL_Surface {
	return (*C.SDL_Surface)(unsafe.Pointer(s))
}

func cRWops(rw *sdl.RWops) *C.SDL_RWops {
	return *(**C.SDL_RWops)(unsafe.Pointer(rw))
}

func goRWops(rw *C.SDL_RWops) *sdl.RWops {
	return (*sdl.RWops)(unsafe.Pointer(&rw))
}

func cRenderer(r *sdl.Renderer) *C.SDL_Renderer {
	return *(**C.SDL_Renderer)(unsafe.Pointer(r))
}

func goRenderer(r *C.SDL_Renderer) *sdl.Renderer {
	return (*sdl.Renderer)(unsafe.Pointer(&r))
}

func cTexture(tex *sdl.Texture) *C.SDL_Texture {
	return *(**C.SDL_Texture)(unsafe.Pointer(tex))
}

func goTexture(tex *C.SDL_Texture) *sdl.Texture {
	return (*sdl.Texture)(unsafe.Pointer(&tex))
}
