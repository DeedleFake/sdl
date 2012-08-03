package sdl

// #include <SDL.h>
import "C"

type GestureID C.SDL_GestureID

func (id GestureID) c() C.SDL_GestureID {
	return C.SDL_GestureID(id)
}

func RecordGesture() error {
	if C.SDL_RecordGesture(-1) != 0 {
		return getError()
	}

	return nil
}

func (id TouchID) RecordGesture() error {
	if C.SDL_RecordGesture(id.c()) != 0 {
		return getError()
	}

	return nil
}

func SaveAllDollarTemplates(rw *RWops) error {
	if C.SDL_SaveAllDollarTemplates(rw.c) != 0 {
		return getError()
	}

	return nil
}

func SaveDollarTemplate(rw *RWops, id GestureID) error {
	if C.SDL_SaveDollarTemplate(id.c(), rw.c) != 0 {
		return getError()
	}

	return nil
}

func LoadDollarTemplates(rw *RWops, id TouchID) error {
	if C.SDL_LoadDollarTemplates(id.c(), rw.c) != 0 {
		return getError()
	}

	return nil
}
