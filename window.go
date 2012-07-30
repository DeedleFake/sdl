package sdl

import (
	"unsafe"
)

// #cgo pkg-config: sdl2
//
// #include <SDL.h>
import "C"

const (
	WINDOWPOS_UNDEFINED_MASK = 0x1FFF0000
	WINDOWPOS_UNDEFINED      = WINDOWPOS_UNDEFINED_MASK | 0

	WINDOWPOS_CENTERED_MASK = 0x2FFF0000
	WINDOWPOS_CENTERED      = WINDOWPOS_CENTERED_MASK | 0
)

func WINDOWPOS_UNDEFINED_DISPLAY(x uint32) uint32 {
	return WINDOWPOS_UNDEFINED_MASK | x
}

func WINDOWPOS_ISUNDEFINED(x uint32) bool {
	return (x & 0xFFFF0000) == WINDOWPOS_UNDEFINED_MASK
}

func WINDOWPOS_CENTERED_DISPLAY(x uint32) uint32 {
	return WINDOWPOS_CENTERED_MASK | x
}

func WINDOWPOS_CENTERED_ISCENTERED(x uint32) bool {
	return (x & 0xFFFF0000) == WINDOWPOS_CENTERED_MASK
}

const (
	WINDOW_FULLSCREEN    = 0x00000001
	WINDOW_OPENGL        = 0x00000002
	WINDOW_SHOWN         = 0x00000004
	WINDOW_HIDDEN        = 0x00000008
	WINDOW_BORDERLESS    = 0x00000010
	WINDOW_RESIZABLE     = 0x00000020
	WINDOW_MINIMIZED     = 0x00000040
	WINDOW_MAXIMIZED     = 0x00000080
	WINDOW_INPUT_GRABBED = 0x00000100
	WINDOW_INPUT_FOCUS   = 0x00000200
	WINDOW_MOUSE_FOCUS   = 0x00000400
	WINDOW_FOREIGN       = 0x00000800
)

type Window struct {
	c *C.SDL_Window
}

func CreateWindow(title string, x, y, w, h int, flags uint32) (*Window, error) {
	ctitle := C.CString(title)
	defer C.free(unsafe.Pointer(ctitle))

	win := C.SDL_CreateWindow(
		ctitle,
		C.int(x),
		C.int(y),
		C.int(w),
		C.int(h),
		C.Uint32(flags),
	)
	if win == nil {
		return nil, getError()
	}

	return &Window{win}, nil
}

func (win *Window) Destroy() {
	C.SDL_DestroyWindow(win.c)
}
