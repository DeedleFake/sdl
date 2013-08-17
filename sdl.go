package sdl

// #cgo pkg-config: sdl2
//
// #include <SDL.h>
import "C"

const (
	INIT_TIMER       = C.SDL_INIT_TIMER
	INIT_AUDIO       = C.SDL_INIT_AUDIO
	INIT_VIDEO       = C.SDL_INIT_VIDEO
	INIT_JOYSTICK    = C.SDL_INIT_JOYSTICK
	INIT_HAPTIC      = C.SDL_INIT_HAPTIC
	INIT_GAMECONTROLLER = C.SDL_INIT_GAMECONTROLLER
	INIT_EVENTS      = C.SDL_INIT_EVENTS
	INIT_NOPARACHUTE = C.SDL_INIT_NOPARACHUTE
	INIT_EVERYTHING  = C.SDL_INIT_EVERYTHING
)

func Init(flags uint32) error {
	if C.SDL_Init(C.Uint32(flags)) != 0 {
		return getError()
	}

	return nil
}

func InitSubSystem(flags uint32) error {
	if C.SDL_InitSubSystem(C.Uint32(flags)) != 0 {
		return getError()
	}

	return nil
}

func QuitSubSystem(flags uint32) {
	C.SDL_QuitSubSystem(C.Uint32(flags))
}

func WasInit(flags uint32) uint32 {
	return uint32(C.SDL_WasInit(C.Uint32(flags)))
}

func Quit() {
	C.SDL_Quit()
}
