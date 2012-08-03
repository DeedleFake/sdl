package sdl

import (
	"errors"
)

// #include <SDL.h>
import "C"

func getError() error {
	err := C.GoString(C.SDL_GetError())
	if len(err) == 0 {
		panic("Blank error.")
	}

	return errors.New(err)
}
