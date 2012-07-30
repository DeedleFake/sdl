package sdl

import (
	"unsafe"
)

// #cgo pkg-config: sdl2
//
// #include <SDL.h>
import "C"

type Point struct {
	X int32
	Y int32
}

func (p *Point) c() *C.SDL_Point {
	return (*C.SDL_Point)(unsafe.Pointer(p))
}

type Rect struct {
	X, Y int32
	W, H int32
}

func (r *Rect) c() *C.SDL_Rect {
	return (*C.SDL_Rect)(unsafe.Pointer(r))
}

func (r *Rect) Empty() bool {
	return (r == nil) || (r.W <= 0) || (r.H <= 0)
}

func (r *Rect) Equals(r2 *Rect) bool {
	return (r != nil) && (r2 != nil) && (r.X == r2.X) && (r.Y == r2.Y) && (r.W == r2.W) && (r.H == r2.H)
}

func (r *Rect) HasIntersection(r2 *Rect) bool {
	return C.SDL_HasIntersection(r.c(), r2.c()) == C.SDL_TRUE
}

func (r *Rect) IntersectRect(r2 *Rect) *Rect {
	var rect Rect
	if C.SDL_IntersectRect(r.c(), r2.c(), rect.c()) == C.SDL_FALSE {
		return nil
	}

	return &rect
}

func (r *Rect) UnionRect(r2 *Rect) *Rect {
	var rect Rect
	C.SDL_UnionRect(r.c(), r2.c(), rect.c())

	return &rect
}

func EnclosePoints(points []Point, clip *Rect) (*Rect, bool) {
	var rect Rect
	return &rect, C.SDL_EnclosePoints(
		(*C.SDL_Point)(unsafe.Pointer(&points[0])),
		C.int(len(points)),
		clip.c(),
		rect.c(),
	) == C.SDL_TRUE
}

func (r *Rect) IntersectRectAndLine(x1, y1, x2, y2 *int) bool {
	cx1 := (*C.int)(unsafe.Pointer(x1))
	cy1 := (*C.int)(unsafe.Pointer(y1))
	cx2 := (*C.int)(unsafe.Pointer(x2))
	cy2 := (*C.int)(unsafe.Pointer(y2))

	return C.SDL_IntersectRectAndLine(r.c(), cx1, cy1, cx2, cy2) == C.SDL_TRUE
}
