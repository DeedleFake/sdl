package sdl

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
