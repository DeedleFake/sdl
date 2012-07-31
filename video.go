package sdl

import (
	"errors"
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
	WINDOW_FULLSCREEN    = C.SDL_WINDOW_FULLSCREEN
	WINDOW_OPENGL        = C.SDL_WINDOW_OPENGL
	WINDOW_SHOWN         = C.SDL_WINDOW_SHOWN
	WINDOW_HIDDEN        = C.SDL_WINDOW_HIDDEN
	WINDOW_BORDERLESS    = C.SDL_WINDOW_BORDERLESS
	WINDOW_RESIZABLE     = C.SDL_WINDOW_RESIZABLE
	WINDOW_MINIMIZED     = C.SDL_WINDOW_MINIMIZED
	WINDOW_MAXIMIZED     = C.SDL_WINDOW_MAXIMIZED
	WINDOW_INPUT_GRABBED = C.SDL_WINDOW_INPUT_GRABBED
	WINDOW_INPUT_FOCUS   = C.SDL_WINDOW_INPUT_FOCUS
	WINDOW_MOUSE_FOCUS   = C.SDL_WINDOW_MOUSE_FOCUS
	WINDOW_FOREIGN       = C.SDL_WINDOW_FOREIGN
)

const (
	WINDOWPOS_UNDEFINED_MASK = C.SDL_WINDOWPOS_UNDEFINED_MASK
	WINDOWPOS_UNDEFINED      = C.SDL_WINDOWPOS_UNDEFINED

	WINDOWPOS_CENTERED_MASK = C.SDL_WINDOWPOS_CENTERED_MASK
	WINDOWPOS_CENTERED      = C.SDL_WINDOWPOS_CENTERED
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
	WINDOWEVENT_NONE         = C.SDL_WINDOWEVENT_NONE
	WINDOWEVENT_SHOWN        = C.SDL_WINDOWEVENT_SHOWN
	WINDOWEVENT_HIDDEN       = C.SDL_WINDOWEVENT_HIDDEN
	WINDOWEVENT_EXPOSED      = C.SDL_WINDOWEVENT_EXPOSED
	WINDOWEVENT_MOVED        = C.SDL_WINDOWEVENT_MOVED
	WINDOWEVENT_RESIZED      = C.SDL_WINDOWEVENT_RESIZED
	WINDOWEVENT_SIZE_CHANGED = C.SDL_WINDOWEVENT_SIZE_CHANGED
	WINDOWEVENT_MINIMIZED    = C.SDL_WINDOWEVENT_MINIMIZED
	WINDOWEVENT_MAXIMIZED    = C.SDL_WINDOWEVENT_MAXIMIZED
	WINDOWEVENT_RESTORED     = C.SDL_WINDOWEVENT_RESTORED
	WINDOWEVENT_ENTER        = C.SDL_WINDOWEVENT_ENTER
	WINDOWEVENT_LEAVE        = C.SDL_WINDOWEVENT_LEAVE
	WINDOWEVENT_FOCUS_GAINED = C.SDL_WINDOWEVENT_FOCUS_GAINED
	WINDOWEVENT_FOCUS_LOST   = C.SDL_WINDOWEVENT_FOCUS_LOST
	WINDOWEVENT_CLOSE        = C.SDL_WINDOWEVENT_CLOSE
)

type GLContext uintptr

func (ctx GLContext) c() C.SDL_GLContext {
	return C.SDL_GLContext(ctx)
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
	GL_CONTEXT_PROFILE_CORE          = C.SDL_GL_CONTEXT_PROFILE_CORE
	GL_CONTEXT_PROFILE_COMPATIBILITY = C.SDL_GL_CONTEXT_PROFILE_COMPATIBILITY
	GL_CONTEXT_PROFILE_ES2           = C.SDL_GL_CONTEXT_PROFILE_ES2
)

const (
	GL_CONTEXT_DEBUG_FLAG              = C.SDL_GL_CONTEXT_DEBUG_FLAG
	GL_CONTEXT_FORWARD_COMPATIBLE_FLAG = C.SDL_GL_CONTEXT_FORWARD_COMPATIBLE_FLAG
	GL_CONTEXT_ROBUST_ACCESS_FLAG      = C.SDL_GL_CONTEXT_ROBUST_ACCESS_FLAG
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

func GetNumDisplayModes(i int) (int, error) {
	nm := C.SDL_GetNumDisplayModes(C.int(i))
	if nm < 0 {
		return 0, getError()
	}

	return int(nm), nil
}

func GetDisplayMode(i int, m int) (*DisplayMode, error) {
	var dm DisplayMode
	if C.SDL_GetDisplayMode(C.int(i), C.int(m), dm.c()) != 0 {
		return nil, getError()
	}

	return &dm, nil
}

func GetDesktopDisplayMode(i int) (*DisplayMode, error) {
	var dm DisplayMode
	if C.SDL_GetDesktopDisplayMode(C.int(i), dm.c()) != 0 {
		return nil, getError()
	}

	return &dm, nil
}

func GetCurrentDisplayMode(i int) (*DisplayMode, error) {
	var dm DisplayMode
	if C.SDL_GetCurrentDisplayMode(C.int(i), dm.c()) != 0 {
		return nil, getError()
	}

	return &dm, nil
}

func GetClosestDisplayMode(i int, m *DisplayMode) (*DisplayMode, error) {
	var dm DisplayMode
	if C.SDL_GetClosestDisplayMode(C.int(i), m.c(), dm.c()) == nil {
		return nil, errors.New("Couldn't find close match")
	}

	return &dm, nil
}

func (win *Window) GetDisplay() (int, error) {
	if d := C.SDL_GetWindowDisplay(win.c); d >= 0 {
		return int(d), nil
	}

	return 0, getError()
}

func (win *Window) SetDisplayMode(dm *DisplayMode) error {
	if C.SDL_SetWindowDisplayMode(win.c, dm.c()) != 0 {
		return getError()
	}

	return nil
}

func (win *Window) GetDisplayMode() (*DisplayMode, error) {
	var dm DisplayMode
	if C.SDL_GetWindowDisplayMode(win.c, dm.c()) != 0 {
		return nil, getError()
	}

	return &dm, nil
}

func (win *Window) GetPixelFormat() uint32 {
	return uint32(C.SDL_GetWindowPixelFormat(win.c))
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

func CreateWindowFrom(data uintptr) (*Window, error) {
	win := C.SDL_CreateWindowFrom(unsafe.Pointer(data))
	if win == nil {
		return nil, getError()
	}

	return &Window{win}, nil
}

func (win *Window) GetID() uint32 {
	return uint32(C.SDL_GetWindowID(win.c))
}

func GetWindowFromID(id uint32) (*Window, error) {
	win := C.SDL_GetWindowFromID(C.Uint32(id))
	if win == nil {
		return nil, getError()
	}

	return &Window{win}, nil
}

func (win *Window) GetFlags() uint32 {
	return uint32(C.SDL_GetWindowFlags(win.c))
}

func (win *Window) SetTitle(title string) {
	ctitle := C.CString(title)
	defer C.free(unsafe.Pointer(ctitle))

	C.SDL_SetWindowTitle(win.c, ctitle)
}

func (win *Window) GetTitle() string {
	return C.GoString(C.SDL_GetWindowTitle(win.c))
}

func (win *Window) SetIcon(icon *Surface) {
	C.SDL_SetWindowIcon(win.c, icon.c())
}

func (win *Window) SetData(name string, data interface{}) interface{} {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	pd := C.SDL_SetWindowData(win.c, cname, unsafe.Pointer(&data))
	if pd == nil {
		return nil
	}

	return *(*interface{})(pd)
}

func (win *Window) GetData(name string) interface{} {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	data := C.SDL_GetWindowData(win.c, cname)
	if data == nil {
		return nil
	}

	return *(*interface{})(data)
}

func (win *Window) SetPosition(x, y int) {
	C.SDL_SetWindowPosition(win.c, C.int(x), C.int(y))
}

func (win *Window) GetPosition() (x, y int) {
	C.SDL_GetWindowPosition(win.c,
		(*C.int)(unsafe.Pointer(&x)),
		(*C.int)(unsafe.Pointer(&y)),
	)

	return
}

func (win *Window) SetSize(w, h int) {
	C.SDL_SetWindowSize(win.c, C.int(w), C.int(h))
}

func (win *Window) GetSize() (w, h int) {
	C.SDL_GetWindowSize(win.c,
		(*C.int)(unsafe.Pointer(&w)),
		(*C.int)(unsafe.Pointer(&h)),
	)

	return
}

func (win *Window) Show() {
	C.SDL_ShowWindow(win.c)
}

func (win *Window) Hide() {
	C.SDL_HideWindow(win.c)
}

func (win *Window) Raise() {
	C.SDL_RaiseWindow(win.c)
}

func (win *Window) Maximize() {
	C.SDL_MaximizeWindow(win.c)
}

func (win *Window) Minimize() {
	C.SDL_MinimizeWindow(win.c)
}

func (win *Window) Restore() {
	C.SDL_RestoreWindow(win.c)
}

func (win *Window) SetFullscreen(fs bool) error {
	cfs := C.SDL_bool(C.SDL_FALSE)
	if fs {
		cfs = C.SDL_TRUE
	}

	if C.SDL_SetWindowFullscreen(win.c, cfs) != 0 {
		return getError()
	}

	return nil
}

func (win *Window) GetSurface() (*Surface, error) {
	s := C.SDL_GetWindowSurface(win.c)
	if s == nil {
		return nil, getError()
	}

	return (*Surface)(unsafe.Pointer(s)), nil
}

func (win *Window) UpdateSurface() error {
	if C.SDL_UpdateWindowSurface(win.c) != 0 {
		return getError()
	}

	return nil
}

func (win *Window) UpdateSurfaceRects(rects []Rect) error {
	if C.SDL_UpdateWindowSurfaceRects(win.c, (*C.SDL_Rect)(unsafe.Pointer(&rects[0])), C.int(len(rects))) != 0 {
		return getError()
	}

	return nil
}

func (win *Window) SetGrab(grab bool) {
	cgrab := C.SDL_bool(C.SDL_FALSE)
	if grab {
		cgrab = C.SDL_TRUE
	}

	C.SDL_SetWindowGrab(win.c, cgrab)
}

func (win *Window) GetGrab() bool {
	return C.SDL_GetWindowGrab(win.c) == C.SDL_TRUE
}

func (win *Window) SetBrightness(b float32) error {
	if C.SDL_SetWindowBrightness(win.c, C.float(b)) != 0 {
		return getError()
	}

	return nil
}

func (win *Window) GetBrightness() float32 {
	return float32(C.SDL_GetWindowBrightness(win.c))
}

func (win *Window) SetGammaRamp(r, g, b uint16) error {
	cr := (*C.Uint16)(unsafe.Pointer(&r))
	cg := (*C.Uint16)(unsafe.Pointer(&g))
	cb := (*C.Uint16)(unsafe.Pointer(&b))

	if C.SDL_SetWindowGammaRamp(win.c, cr, cg, cb) != 0 {
		return getError()
	}

	return nil
}

func (win *Window) GetGammaRamp() (r, g, b uint16, err error) {
	var cr, cg, cb C.Uint16
	if C.SDL_GetWindowGammaRamp(win.c, &cr, &cg, &cb) != 0 {
		return 0, 0, 0, getError()
	}

	return uint16(cr), uint16(cg), uint16(cb), nil
}

func (win *Window) Destroy() {
	C.SDL_DestroyWindow(win.c)
}

func IsScreenSaverEnabled() bool {
	return C.SDL_IsScreenSaverEnabled() == C.SDL_TRUE
}

func EnableScreenSaver() {
	C.SDL_EnableScreenSaver()
}

func DisableScreenSaver() {
	C.SDL_DisableScreenSaver()
}

func GL_LoadLibrary(path string) error {
	cpath := C.CString(path)
	defer C.free(unsafe.Pointer(cpath))

	if C.SDL_GL_LoadLibrary(cpath) != 0 {
		return getError()
	}

	return nil
}

func GL_GetProcAddress(proc string) uintptr {
	cproc := C.CString(proc)
	defer C.free(unsafe.Pointer(cproc))

	return uintptr(C.SDL_GL_GetProcAddress(cproc))
}

func GL_UnloadLibrary() {
	C.SDL_GL_UnloadLibrary()
}

func GL_ExtensionSupported(ext string) bool {
	cext := C.CString(ext)
	defer C.free(unsafe.Pointer(cext))

	return C.SDL_GL_ExtensionSupported(cext) == C.SDL_TRUE
}

func GL_SetAttribute(attr uint32, val int) error {
	if C.SDL_GL_SetAttribute(C.SDL_GLattr(attr), C.int(val)) != 0 {
		return getError()
	}

	return nil
}

func GL_GetAttribute(attr uint32) (int, error) {
	var val C.int
	if C.SDL_GL_GetAttribute(C.SDL_GLattr(attr), &val) != 0 {
		return 0, getError()
	}

	return int(val), nil
}

func (win *Window) GL_CreateContext() (GLContext, error) {
	ctx := C.SDL_GL_CreateContext(win.c)
	if ctx == nil {
		return 0, getError()
	}

	return GLContext(ctx), nil
}

func (win *Window) GL_MakeCurrent(ctx GLContext) error {
	if C.SDL_GL_MakeCurrent(win.c, ctx.c()) != 0 {
		return getError()
	}

	return nil
}

func GL_SetSwapInterval(vsync bool) error {
	var interval C.int
	if vsync {
		interval = 1
	}

	if C.SDL_GL_SetSwapInterval(interval) != 0 {
		return getError()
	}

	return nil
}

func GL_GetSwapInterval() (bool, error) {
	interval := C.SDL_GL_GetSwapInterval()
	if interval < 0 {
		return false, getError()
	}

	return interval == 1, nil
}

func (win *Window) GL_Swap() {
	C.SDL_GL_SwapWindow(win.c)
}

func (ctx GLContext) Delete() {
	C.SDL_GL_DeleteContext(ctx.c())
}
