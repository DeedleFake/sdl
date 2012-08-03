package sdl

// #include <SDL.h>
import "C"

const (
	CACHELINE_SIZE = C.SDL_CACHELINE_SIZE
)

func GetCPUCount() int {
	return int(C.SDL_GetCPUCount())
}

func GetCPUCacheLineSize() int {
	return int(C.SDL_GetCPUCacheLineSize())
}

func HasRDTSC() bool {
	return C.SDL_HasRDTSC() == C.SDL_TRUE
}

func HasAltiVec() bool {
	return C.SDL_HasAltiVec() == C.SDL_TRUE
}

func HasMMX() bool {
	return C.SDL_HasMMX() == C.SDL_TRUE
}

func Has3DNow() bool {
	return C.SDL_Has3DNow() == C.SDL_TRUE
}

func HasSSE() bool {
	return C.SDL_HasSSE() == C.SDL_TRUE
}

func HasSSE2() bool {
	return C.SDL_HasSSE2() == C.SDL_TRUE
}

func HasSSE3() bool {
	return C.SDL_HasSSE3() == C.SDL_TRUE
}

func HasSSE41() bool {
	return C.SDL_HasSSE41() == C.SDL_TRUE
}

func HasSSE42() bool {
	return C.SDL_HasSSE42() == C.SDL_TRUE
}
