package sdl

// #include <SDL.h>
import "C"

import "unsafe"

const (
	HINT_FRAMEBUFFER_ACCELERATION = string(C.SDL_HINT_FRAMEBUFFER_ACCELERATION)
	IDLE_TIMER_DISABLED           = string(C.SDL_HINT_IDLE_TIMER_DISABLED)
	HINT_ORIENTATIONS             = string(C.SDL_HINT_ORIENTATIONS)
	HINT_RENDER_DRIVER            = string(C.SDL_HINT_RENDER_DRIVER)
	HINT_RENDER_OPENGL_SHADERS    = string(C.SDL_HINT_RENDER_OPENGL_SHADERS)
	HINT_RENDER_SCALE_QUALITY     = string(C.SDL_HINT_RENDER_SCALE_QUALITY)
	HINT_RENDER_VSYNC             = string(C.SDL_HINT_RENDER_VSYNC)
)

func SetHint(name, value string) bool {
	cname, cvalue := C.CString(name), C.CString(value)
	defer C.free(unsafe.Pointer(cname))
	defer C.free(unsafe.Pointer(cvalue))

	return C.SDL_SetHint(cname, cvalue) == C.SDL_TRUE
}

func GetHint(name string) string {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	cvalue := C.SDL_GetHint(cname)
	if cvalue == nil {
		return ""
	}
	return C.GoString(cvalue)
}
