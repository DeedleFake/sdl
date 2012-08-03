package sdl

import (
	"errors"
	"unsafe"
)

// #include <SDL.h>
//
// #include "mouse.h"
import "C"

type Cursor struct {
	c *C.SDL_Cursor
}

func GetMouseFocus() *Window {
	return &Window{C.SDL_GetMouseFocus()}
}

func GetMouseState() (state uint8, x, y int) {
	var cx, cy C.int
	cstate := C.SDL_GetMouseState(&cx, &cy)

	return uint8(cstate), int(cx), int(cy)
}

func GetRelativeMouseState() (state uint8, x, y int) {
	var cx, cy C.int
	cstate := C.SDL_GetRelativeMouseState(&cx, &cy)

	return uint8(cstate), int(cx), int(cy)
}

func WarpMouse(x, y int) {
	C.SDL_WarpMouseInWindow(nil, C.int(x), C.int(y))
}

func (w *Window) WarpMouse(x, y int) {
	C.SDL_WarpMouseInWindow(w.c, C.int(x), C.int(y))
}

func SetRelativeMouseMode(e bool) error {
	ce := C.SDL_bool(C.SDL_FALSE)
	if e {
		ce = C.SDL_TRUE
	}

	if C.SDL_SetRelativeMouseMode(ce) != 0 {
		return getError()
	}

	return nil
}

func GetRelativeMouseMode() bool {
	return C.SDL_GetRelativeMouseMode() == C.SDL_TRUE
}

func CreateCursor(data []uint8, mask []uint8, w, h, hx, hy int) (*Cursor, error) {
	wh := w * h
	if len(data) < wh {
		return nil, errors.New("len(data) < w * h")
	}
	if (len(mask) < wh) && (len(mask) > 0) {
		return nil, errors.New("len(mask) < w * h and > 0")
	}

	var mp *C.Uint8
	if len(mask) > 0 {
		mp = (*C.Uint8)(unsafe.Pointer(&mask[0]))
	}

	c := C.SDL_CreateCursor(
		(*C.Uint8)(unsafe.Pointer(&data[0])),
		mp,
		C.int(w),
		C.int(h),
		C.int(hx),
		C.int(hy),
	)
	if c == nil {
		return nil, getError()
	}

	return &Cursor{c}, nil
}

func (s *Surface) CreateColorCursor(hx, hy int) (*Cursor, error) {
	c := C.SDL_CreateColorCursor(s.c(), C.int(hx), C.int(hy))
	if c == nil {
		return nil, getError()
	}

	return &Cursor{c}, nil
}

func (c *Cursor) Set() {
	C.SDL_SetCursor(c.c)
}

func GetCursor() (*Cursor, error) {
	c := C.SDL_GetCursor()
	if c == nil {
		return nil, getError()
	}

	return &Cursor{c}, nil
}

func (c *Cursor) Free() {
	C.SDL_FreeCursor(c.c)
}

func ShowCursor(t int) bool {
	return C.SDL_ShowCursor(C.int(t)) != 0
}

func BUTTON(x uint8) uint8 {
	return uint8(C.BUTTON(C.Uint8(x)))
}

const (
	BUTTON_LEFT   = C.SDL_BUTTON_LEFT
	BUTTON_MIDDLE = C.SDL_BUTTON_MIDDLE
	BUTTON_RIGHT  = C.SDL_BUTTON_RIGHT
	BUTTON_X1     = C.SDL_BUTTON_X1
	BUTTON_X2     = C.SDL_BUTTON_X2
	BUTTON_LMASK  = C.SDL_BUTTON_LMASK
	BUTTON_MMASK  = C.SDL_BUTTON_MMASK
	BUTTON_RMASK  = C.SDL_BUTTON_RMASK
	BUTTON_X1MASK = C.SDL_BUTTON_X1MASK
	BUTTON_X2MASK = C.SDL_BUTTON_X2MASK
)
