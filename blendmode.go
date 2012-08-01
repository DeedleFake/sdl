package sdl

// #cgo pkg-config: sdl2
//
// #include <SDL.h>
import "C"

type BlendMode C.SDL_BlendMode

func (m BlendMode) c() C.SDL_BlendMode {
	return C.SDL_BlendMode(m)
}

const (
	BLENDMODE_NONE  BlendMode = C.SDL_BLENDMODE_NONE
	BLENDMODE_BLEND BlendMode = C.SDL_BLENDMODE_BLEND
	BLENDMODE_ADD   BlendMode = C.SDL_BLENDMODE_ADD
	BLENDMODE_MOD   BlendMode = C.SDL_BLENDMODE_MOD
)
