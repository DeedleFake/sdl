package sdl

import (
	"unsafe"
)

// #include <SDL.h>
import "C"

func SetClipboardText(text string) error {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))

	if C.SDL_SetClipboardText(ctext) != 0 {
		return getError()
	}

	return nil
}

func GetClipboardText() string {
	ctext := C.SDL_GetClipboardText()
	defer C.SDL_free(unsafe.Pointer(ctext))

	return C.GoString(ctext)
}

func HasClipboardText() bool {
	return C.SDL_HasClipboardText() == C.SDL_TRUE
}
