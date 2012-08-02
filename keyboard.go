package sdl

import (
	"reflect"
	"unsafe"
)

// #cgo pkg-config: sdl2
//
// #include <SDL.h>
import "C"

type Keysym struct {
	Scancode Scancode
	Sym      Keycode
	Mod      uint16
	Unicode  uint32
}

func (ks *Keysym) c() *C.SDL_Keysym {
	return (*C.SDL_Keysym)(unsafe.Pointer(ks))
}

func GetKeyboardFocus() *Window {
	return &Window{C.SDL_GetKeyboardFocus()}
}

func GetKeyboardState() []uint8 {
	var n C.int
	s := C.SDL_GetKeyboardState(&n)

	return *(*[]uint8)(unsafe.Pointer(&reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(s)),
		Len:  int(n),
		Cap:  int(n),
	}))
}

func GetModState() Keymod {
	return Keymod(C.SDL_GetModState())
}

func (s Scancode) GetKey() Keycode {
	return Keycode(C.SDL_GetKeyFromScancode(s.c()))
}

func (k Keycode) GetScancode() Scancode {
	return Scancode(C.SDL_GetScancodeFromKey(k.c()))
}

func (s Scancode) GetName() string {
	return C.GoString(C.SDL_GetScancodeName(s.c()))
}

func GetScancodeFromName(name string) Scancode {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	return Scancode(C.SDL_GetScancodeFromName(cname))
}

func (k Keycode) GetName() string {
	return C.GoString(C.SDL_GetKeyName(k.c()))
}

func GetKeyFromName(name string) Keycode {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	return Keycode(C.SDL_GetKeyFromName(cname))
}

func StartTextInput() {
	C.SDL_StartTextInput()
}

func StopTextInput() {
	C.SDL_StopTextInput()
}

func SetTextInputRect(r *Rect) {
	C.SDL_SetTextInputRect(r.c())
}
