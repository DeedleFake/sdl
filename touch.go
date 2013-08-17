package sdl
/*
import (
	"reflect"
	"unsafe"
)*/

// #include <SDL.h>
import "C"

type TouchID C.SDL_TouchID

func (id TouchID) c() C.SDL_TouchID {
	return C.SDL_TouchID(id)
}

type FingerID C.SDL_FingerID

func (id FingerID) c() C.SDL_FingerID {
	return C.SDL_FingerID(id)
}

type Finger struct {
	Id                         FingerID
	X, Y                       float32
	Pressure                   float32
}
/*
func (f *Finger) c() *C.SDL_Finger {
	return (*C.SDL_Finger)(unsafe.Pointer(f))
}

func (f *Finger) Down() bool {
	return f.down == C.SDL_TRUE
}

type Touch struct {
	freeTouch uintptr

	PressureMax, PressureMin                  float32
	XMax, XMin                                float32
	YMax, YMin                                float32
	Xres, Yres, Pressureres                   uint16
	NativeXres, NativeYres, NativePressureres float32
	TiltX, TiltY                              float32
	Rotation                                  float32

	Id    TouchID
	Focus *Window

	name         *C.char
	Buttonstate  uint8
	relativeMode C.SDL_bool
	flushMotion  C.SDL_bool

	NumFingers int32
	MaxFingers int32
	fingers    uintptr

	Driverdata uintptr
}

func (t *Touch) c() *C.SDL_Touch {
	return (*C.SDL_Touch)(unsafe.Pointer(t))
}

func (t *Touch) Name() string {
	return C.GoString(t.name)
}

func (t *Touch) RelativeMode() bool {
	return t.relativeMode == C.SDL_TRUE
}

func (t *Touch) FlushMotion() bool {
	return t.flushMotion == C.SDL_TRUE
}

func (t *Touch) Fingers() []*Finger {
	return *(*[]*Finger)(unsafe.Pointer(&reflect.SliceHeader{
		Data: t.fingers,
		Len:  int(t.NumFingers),
		Cap:  int(t.NumFingers),
	}))
}

func (t TouchID) GetTouchDevice() TouchID {
	return (*Touch)(unsafe.Pointer(C.SDL_GetTouchDevice(t.c())))
}

func (t *Touch) GetFingerDevice(f FingerID) *Finger {
	return (*Finger)(unsafe.Pointer(C.SDL_GetFingerDevice(t.c(), f.c())))
}*/
