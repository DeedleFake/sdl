package sdl

import (
	"errors"
)

// #cgo pkg-config: sdl2
//
// #include <SDL.h>
import "C"

const (
	INIT_TIMER       = 0x00000001
	INIT_AUDIO       = 0x00000010
	INIT_VIDEO       = 0x00000020
	INIT_JOYSTICK    = 0x00000200
	INIT_HAPTIC      = 0x00001000
	INIT_NOPARACHUTE = 0x00100000
	INIT_EVERYTHING  = 0x0000FFFF
)

func getError() error {
	err := C.GoString(C.SDL_GetError())
	if len(err) == 0 {
		panic("Blank error.")
	}

	return errors.New(err)
}

func Init(flags uint32) error {
	if C.SDL_Init(C.Uint32(flags)) != 0 {
		return getError()
	}

	return nil
}

func Quit() {
	C.SDL_Quit()
}
