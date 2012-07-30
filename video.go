package sdl

import (
	"unsafe"
)

// #cgo pkg-config: sdl2
//
// #include <SDL.h>
import "C"

type DisplayMode struct {
	Format      uint32
	W           int32
	H           int32
	RefreshRate int32
	DriverData  uintptr
}

func (dm *DisplayMode) c() *C.SDL_DisplayMode {
	return (*C.SDL_DisplayMode)(unsafe.Pointer(dm))
}

type Window struct {
	c *C.SDL_Window
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
	WINDOWEVENT_NONE = iota
	WINDOWEVENT_SHOWN
	WINDOWEVENT_HIDDEN
	WINDOWEVENT_EXPOSED
	WINDOWEVENT_MOVED
	WINDOWEVENT_RESIZED
	WINDOWEVENT_SIZE_CHANGED
	WINDOWEVENT_MINIMIZED
	WINDOWEVENT_MAXIMIZED
	WINDOWEVENT_RESTORED
	WINDOWEVENT_ENTER
	WINDOWEVENT_LEAVE
	WINDOWEVENT_FOCUS_GAINED
	WINDOWEVENT_FOCUS_LOST
	WINDOWEVENT_CLOSE
)

type GLContext struct {
	c *C.SDL_GLContext
}

const (
	GL_RED_SIZE = iota
	GL_GREEN_SIZE
	GL_BLUE_SIZE
	GL_ALPHA_SIZE
	GL_BUFFER_SIZE
	GL_DOUBLEBUFFER
	GL_DEPTH_SIZE
	GL_STENCIL_SIZE
	GL_ACCUM_RED_SIZE
	GL_ACCUM_GREEN_SIZE
	GL_ACCUM_BLUE_SIZE
	GL_ACCUM_ALPHA_SIZE
	GL_STEREO
	GL_MULTISAMPLEBUFFERS
	GL_MULTISAMPLESAMPLES
	GL_ACCELERATED_VISUAL
	GL_RETAINED_BACKING
	GL_CONTEXT_MAJOR_VERSION
	GL_CONTEXT_MINOR_VERSION
	GL_CONTEXT_EGL
	GL_CONTEXT_FLAGS
	GL_CONTEXT_PROFILE_MASK
)

const (
	GL_CONTEXT_PROFILE_CORE          = 0x0001
	GL_CONTEXT_PROFILE_COMPATIBILITY = 0x0002
	GL_CONTEXT_PROFILE_ES2           = 0x0004
)

const (
	GL_CONTEXT_DEBUG_FLAG              = 0x0001
	GL_CONTEXT_FORWARD_COMPATIBLE_FLAG = 0x0002
	GL_CONTEXT_ROBUST_ACCESS_FLAG      = 0x0004
)

func GetNumVideoDrivers() int {
	return int(C.SDL_GetNumVideoDrivers())
}

func GetVideoDriver(i int) string {
	return C.GoString(C.SDL_GetVideoDriver(C.int(i)))
}

func VideoInit(driver string) error {
	cdriver := C.CString(driver)
	defer C.free(unsafe.Pointer(cdriver))

	if C.SDL_VideoInit(cdriver) != 0 {
		return getError()
	}

	return nil
}

func VideoQuit() {
	C.SDL_VideoQuit()
}

func GetCurrentVideoDriver() string {
	return C.GoString(C.SDL_GetCurrentVideoDriver())
}

func GetNumVideoDisplays() int {
	return int(C.SDL_GetNumVideoDisplays())
}

func GetDisplayBounds(i int) (*Rect, error) {
	var rect Rect
	if C.SDL_GetDisplayBounds(C.int(i), rect.c()) != 0 {
		return nil, getError()
	}

	return &rect, nil
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
