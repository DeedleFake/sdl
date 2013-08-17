package sdl

import (
	"math"
	"strconv"
	"sync"
	"time"
	"unsafe"
)

// #include <SDL.h>
//
// #include "events.h"
import "C"

type EventType uint32

const (
	FIRSTEVENT EventType = C.SDL_FIRSTEVENT

	QUIT EventType = C.SDL_QUIT

	APP_TERMINATING         EventType = C.SDL_APP_TERMINATING
	APP_LOWMEMORY           EventType = C.SDL_APP_LOWMEMORY
	APP_WILLENTERBACKGROUND EventType = C.SDL_APP_WILLENTERBACKGROUND
	APP_DIDENTERBACKGROUND  EventType = C.SDL_APP_DIDENTERBACKGROUND
	APP_WILLENTERFOREGROUND EventType = C.SDL_APP_WILLENTERFOREGROUND
	APP_DIDENTERFOREGROUND  EventType = C.SDL_APP_DIDENTERFOREGROUND

	WINDOWEVENT EventType = C.SDL_WINDOWEVENT
	SYSWMEVENT  EventType = C.SDL_SYSWMEVENT

	KEYDOWN     EventType = C.SDL_KEYDOWN
	KEYUP       EventType = C.SDL_KEYUP
	TEXTEDITING EventType = C.SDL_TEXTEDITING
	TEXTINPUT   EventType = C.SDL_TEXTINPUT

	MOUSEMOTION     EventType = C.SDL_MOUSEMOTION
	MOUSEBUTTONDOWN EventType = C.SDL_MOUSEBUTTONDOWN
	MOUSEBUTTONUP   EventType = C.SDL_MOUSEBUTTONUP
	MOUSEWHEEL      EventType = C.SDL_MOUSEWHEEL

	JOYAXISMOTION    EventType = C.SDL_JOYAXISMOTION
	JOYBALLMOTION    EventType = C.SDL_JOYBALLMOTION
	JOYHATMOTION     EventType = C.SDL_JOYHATMOTION
	JOYBUTTONDOWN    EventType = C.SDL_JOYBUTTONDOWN
	JOYBUTTONUP      EventType = C.SDL_JOYBUTTONUP
	JOYDEVICEADDED   EventType = C.SDL_JOYDEVICEADDED
	JOYDEVICEREMOVED EventType = C.SDL_JOYDEVICEREMOVED

	CONTROLLERAXISMOTION     EventType = C.SDL_CONTROLLERAXISMOTION
	CONTROLLERBUTTONDOWN     EventType = C.SDL_CONTROLLERBUTTONDOWN
	CONTROLLERBUTTONUP       EventType = C.SDL_CONTROLLERBUTTONUP
	CONTROLLERDEVICEADDED    EventType = C.SDL_CONTROLLERDEVICEADDED
	CONTROLLERDEVICEREMOVED  EventType = C.SDL_CONTROLLERDEVICEREMOVED
	CONTROLLERDEVICEREMAPPED EventType = C.SDL_CONTROLLERDEVICEREMAPPED

	FINGERDOWN   EventType = C.SDL_FINGERDOWN
	FINGERUP     EventType = C.SDL_FINGERUP
	FINGERMOTION EventType = C.SDL_FINGERMOTION

	DOLLARGESTURE EventType = C.SDL_DOLLARGESTURE
	DOLLARRECORD  EventType = C.SDL_DOLLARRECORD
	MULTIGESTURE  EventType = C.SDL_MULTIGESTURE

	CLIPBOARDUPDATE EventType = C.SDL_CLIPBOARDUPDATE

	DROPFILE EventType = C.SDL_DROPFILE

	USEREVENT EventType = C.SDL_USEREVENT

	LASTEVENT EventType = C.SDL_LASTEVENT
)

type WindowEvent struct {
	Type                         EventType
	Timestamp                    uint32
	WindowID                     uint32
	Event                        uint8
	padding1, padding2, padding3 uint8
	Data1, Data2                 int32
}

func (ev *WindowEvent) c() *C.SDL_WindowEvent {
	return (*C.SDL_WindowEvent)(unsafe.Pointer(ev))
}

func (ev *WindowEvent) sdlEvent() *cEvent {
	return (*cEvent)(unsafe.Pointer(ev))
}

type KeyboardEvent struct {
	Type               EventType
	Timestamp          uint32
	WindowID           uint32
	State              uint8
	Repeat             uint8
	padding2, padding3 uint8
	Keysym             Keysym
}

func (ev *KeyboardEvent) c() *C.SDL_KeyboardEvent {
	return (*C.SDL_KeyboardEvent)(unsafe.Pointer(ev))
}

func (ev *KeyboardEvent) sdlEvent() *cEvent {
	return (*cEvent)(unsafe.Pointer(ev))
}

const TEXTEDITINGEVENT_TEXT_SIZE = C.SDL_TEXTEDITINGEVENT_TEXT_SIZE

type TextEditingEvent struct {
	Type      EventType
	Timestamp uint32
	WindowID  uint32
	Text      [TEXTEDITINGEVENT_TEXT_SIZE]byte
	Start     int32
	Length    int32
}

func (ev *TextEditingEvent) c() *C.SDL_TextEditingEvent {
	return (*C.SDL_TextEditingEvent)(unsafe.Pointer(ev))
}

func (ev *TextEditingEvent) sdlEvent() *cEvent {
	return (*cEvent)(unsafe.Pointer(ev))
}

const TEXTINPUTEVENT_TEXT_SIZE = C.SDL_TEXTINPUTEVENT_TEXT_SIZE

type TextInputEvent struct {
	Type      EventType
	Timestamp uint32
	WindowID  uint32
	Text      [TEXTINPUTEVENT_TEXT_SIZE]byte
}

func (ev *TextInputEvent) c() *C.SDL_TextInputEvent {
	return (*C.SDL_TextInputEvent)(unsafe.Pointer(ev))
}

func (ev *TextInputEvent) sdlEvent() *cEvent {
	return (*cEvent)(unsafe.Pointer(ev))
}

type MouseMotionEvent struct {
	Type                         EventType
	Timestamp                    uint32
	WindowID                     uint32
	State                        uint8
	padding1, padding2, padding3 uint8
	X, Y                         int32
	Xrel, Yrel                   int32
}

func (ev *MouseMotionEvent) c() *C.SDL_MouseMotionEvent {
	return (*C.SDL_MouseMotionEvent)(unsafe.Pointer(ev))
}

func (ev *MouseMotionEvent) sdlEvent() *cEvent {
	return (*cEvent)(unsafe.Pointer(ev))
}

type MouseButtonEvent struct {
	Type               EventType
	Timestamp          uint32
	WindowID           uint32
	Button             uint8
	State              uint8
	padding1, padding2 uint8
	X, Y               int32
}

func (ev *MouseButtonEvent) c() *C.SDL_MouseButtonEvent {
	return (*C.SDL_MouseButtonEvent)(unsafe.Pointer(ev))
}

func (ev *MouseButtonEvent) sdlEvent() *cEvent {
	return (*cEvent)(unsafe.Pointer(ev))
}

type MouseWheelEvent struct {
	Type      EventType
	Timestamp uint32
	WindowID  uint32
	X, Y      int32
}

func (ev *MouseWheelEvent) c() *C.SDL_MouseWheelEvent {
	return (*C.SDL_MouseWheelEvent)(unsafe.Pointer(ev))
}

func (ev *MouseWheelEvent) sdlEvent() *cEvent {
	return (*cEvent)(unsafe.Pointer(ev))
}

type JoyAxisEvent struct {
	Type               EventType
	Timestamp          uint32
	Which              uint8
	Axis               uint8
	padding1, padding2 uint8
	Value              int32
}

func (ev *JoyAxisEvent) c() *C.SDL_JoyAxisEvent {
	return (*C.SDL_JoyAxisEvent)(unsafe.Pointer(ev))
}

func (ev *JoyAxisEvent) sdlEvent() *cEvent {
	return (*cEvent)(unsafe.Pointer(ev))
}

type JoyBallEvent struct {
	Type               EventType
	Timestamp          uint32
	Which              uint8
	Ball               uint8
	padding1, padding2 uint8
	Xrel, Yrel         int32
}

func (ev *JoyBallEvent) c() *C.SDL_JoyBallEvent {
	return (*C.SDL_JoyBallEvent)(unsafe.Pointer(ev))
}

func (ev *JoyBallEvent) sdlEvent() *cEvent {
	return (*cEvent)(unsafe.Pointer(ev))
}

type JoyHatEvent struct {
	Type      EventType
	Timestamp uint32
	Which     uint8
	Hat       uint8
	Value     uint8
	padding1  uint8
}

func (ev *JoyHatEvent) c() *C.SDL_JoyHatEvent {
	return (*C.SDL_JoyHatEvent)(unsafe.Pointer(ev))
}

func (ev *JoyHatEvent) sdlEvent() *cEvent {
	return (*cEvent)(unsafe.Pointer(ev))
}

type JoyButtonEvent struct {
	Type      EventType
	Timestamp uint32
	Which     uint8
	Hat       uint8
	Value     uint8
	padding1  uint8
}

func (ev *JoyButtonEvent) c() *C.SDL_JoyButtonEvent {
	return (*C.SDL_JoyButtonEvent)(unsafe.Pointer(ev))
}

func (ev *JoyButtonEvent) sdlEvent() *cEvent {
	return (*cEvent)(unsafe.Pointer(ev))
}

type TouchFingerEvent struct {
	Type                         EventType
	Timestamp                    uint32
	WindowID                     uint32
	TouchId                      TouchID
	FingerId                     FingerID
	State                        uint8
	padding1, padding2, padding3 uint8
	X, Y                         uint16
	Dx, Dy                       int16
	Pressure                     uint16
}

func (ev *TouchFingerEvent) c() *C.SDL_TouchFingerEvent {
	return (*C.SDL_TouchFingerEvent)(unsafe.Pointer(ev))
}

func (ev *TouchFingerEvent) sdlEvent() *cEvent {
	return (*cEvent)(unsafe.Pointer(ev))
}

type MultiGestureEvent struct {
	Type       EventType
	Timestamp  uint32
	WindowID   uint32
	TouchId    TouchID
	DTheta     float32
	DDist      float32
	X, Y       float32
	NumFingers uint16
	padding    uint16
}

func (ev *MultiGestureEvent) c() *C.SDL_MultiGestureEvent {
	return (*C.SDL_MultiGestureEvent)(unsafe.Pointer(ev))
}

func (ev *MultiGestureEvent) sdlEvent() *cEvent {
	return (*cEvent)(unsafe.Pointer(ev))
}

type DollarGestureEvent struct {
	Type       EventType
	Timestamp  uint32
	WindowID   uint32
	TouchId    TouchID
	GestureId  GestureID
	NumFingers uint32
	Error      float32
}

func (ev *DollarGestureEvent) c() *C.SDL_DollarGestureEvent {
	return (*C.SDL_DollarGestureEvent)(unsafe.Pointer(ev))
}

func (ev *DollarGestureEvent) sdlEvent() *cEvent {
	return (*cEvent)(unsafe.Pointer(ev))
}

type DropEvent struct {
	Type      EventType
	Timestamp uint32
	file      *C.char
}

func (ev *DropEvent) c() *C.SDL_DropEvent {
	return (*C.SDL_DropEvent)(unsafe.Pointer(ev))
}

func (ev *DropEvent) sdlEvent() *cEvent {
	return (*cEvent)(unsafe.Pointer(ev))
}

func (ev *DropEvent) File() string {
	return C.GoString(ev.file)
}

type QuitEvent struct {
	Type      EventType
	Timestamp uint32
}

func (ev *QuitEvent) c() *C.SDL_QuitEvent {
	return (*C.SDL_QuitEvent)(unsafe.Pointer(ev))
}

func (ev *QuitEvent) sdlEvent() *cEvent {
	return (*cEvent)(unsafe.Pointer(ev))
}

type UserEvent struct {
	Type      EventType
	Timestamp uint32
	WindowID  uint32
	Code      int32
	Data1     uintptr
	Data2     uintptr
}

func (ev *UserEvent) c() *C.SDL_UserEvent {
	return (*C.SDL_UserEvent)(unsafe.Pointer(ev))
}

func (ev *UserEvent) sdlEvent() *cEvent {
	return (*cEvent)(unsafe.Pointer(ev))
}

type SysWMmsg C.SDL_SysWMmsg

type SysWMEvent struct {
	Type      EventType
	Timestamp uint32
	Msg       *SysWMmsg
}

func (ev *SysWMEvent) c() *C.SDL_SysWMEvent {
	return (*C.SDL_SysWMEvent)(unsafe.Pointer(ev))
}

func (ev *SysWMEvent) sdlEvent() *cEvent {
	return (*cEvent)(unsafe.Pointer(ev))
}

type cEvent C.SDL_Event

func (ev *cEvent) c() *C.SDL_Event {
	return (*C.SDL_Event)(unsafe.Pointer(ev))
}

func (ev *cEvent) getType() EventType {
	return *(*EventType)(unsafe.Pointer(ev))
}

func (ev *cEvent) toGo() Event {
	switch ev.getType() {
	case QUIT:
		return (*QuitEvent)(unsafe.Pointer(ev))
	case WINDOWEVENT:
		return (*WindowEvent)(unsafe.Pointer(ev))
	case SYSWMEVENT:
		return (*SysWMEvent)(unsafe.Pointer(ev))
	case KEYDOWN, KEYUP:
		return (*KeyboardEvent)(unsafe.Pointer(ev))
	case TEXTEDITING:
		return (*TextEditingEvent)(unsafe.Pointer(ev))
	case TEXTINPUT:
		return (*TextInputEvent)(unsafe.Pointer(ev))
	case MOUSEMOTION:
		return (*MouseMotionEvent)(unsafe.Pointer(ev))
	case MOUSEBUTTONDOWN, MOUSEBUTTONUP:
		return (*MouseButtonEvent)(unsafe.Pointer(ev))
	case MOUSEWHEEL:
		return (*MouseWheelEvent)(unsafe.Pointer(ev))
	case JOYAXISMOTION:
		return (*JoyAxisEvent)(unsafe.Pointer(ev))
	case JOYBALLMOTION:
		return (*JoyBallEvent)(unsafe.Pointer(ev))
	case JOYHATMOTION:
		return (*JoyHatEvent)(unsafe.Pointer(ev))
	case JOYBUTTONDOWN, JOYBUTTONUP:
		return (*JoyButtonEvent)(unsafe.Pointer(ev))
	case FINGERDOWN, FINGERUP, FINGERMOTION:
		return (*TouchFingerEvent)(unsafe.Pointer(ev))
	case DOLLARGESTURE:
		return (*DollarGestureEvent)(unsafe.Pointer(ev))
	case MULTIGESTURE:
		return (*MultiGestureEvent)(unsafe.Pointer(ev))
	case DROPFILE:
		return (*DropEvent)(unsafe.Pointer(ev))
	case USEREVENT:
		return (*UserEvent)(unsafe.Pointer(ev))
	}

	panic("Can't convert event type: " + strconv.FormatUint(uint64(ev.getType()), 16))
}

type Event interface {
	sdlEvent() *cEvent
}

func PumpEvents() {
	C.SDL_PumpEvents()
}

type Eventaction C.SDL_eventaction

const (
	ADDEVENT  Eventaction = C.SDL_ADDEVENT
	PEEKEVENT Eventaction = C.SDL_PEEKEVENT
	GETEVENT  Eventaction = C.SDL_GETEVENT
)

func PeepEvents(events []Event, action Eventaction, min, max uint32) (int, error) {
	cevents := make([]cEvent, len(events))
	n := C.SDL_PeepEvents(
		(*C.SDL_Event)(unsafe.Pointer(&cevents[0])),
		C.int(len(cevents)),
		C.SDL_eventaction(action),
		C.Uint32(min),
		C.Uint32(max),
	)
	if n < 0 {
		return 0, getError()
	}

	for i := 0; i < int(n); i++ {
		events[i] = cevents[i].toGo()
	}

	return int(n), nil
}

func HasEvent(t uint32) bool {
	return C.SDL_HasEvent(C.Uint32(t)) == C.SDL_TRUE
}

func HasEvents(min, max uint32) bool {
	return C.SDL_HasEvents(C.Uint32(min), C.Uint32(max)) == C.SDL_TRUE
}

// TODO: Find a way to make this more multi-threading friendly.

var (
	evCache cEvent
	evLock  sync.Mutex
)

func PollEvent(ev *Event) bool {
	evLock.Lock()
	defer evLock.Unlock()

	if C.SDL_PollEvent(evCache.c()) == 0 {
		return false
	}

	*ev = evCache.toGo()

	return true
}

func WaitEvent(ev *Event) error {
	evLock.Lock()
	defer evLock.Unlock()

	if C.SDL_WaitEvent(evCache.c()) == 0 {
		return getError()
	}

	*ev = evCache.toGo()

	return nil
}

func WaitEventTimeout(ev *Event, timeout time.Duration) error {
	evLock.Lock()
	defer evLock.Unlock()

	if C.SDL_WaitEventTimeout(evCache.c(), C.int(timeout/time.Millisecond)) == 0 {
		return getError()
	}

	return nil
}

func PushEvent(ev Event) (bool, error) {
	switch C.SDL_PushEvent(ev.sdlEvent().c()) {
	case 1:
		return true, nil
	case 0:
		return false, nil
	}

	return false, getError()
}

type EventFilter func(data interface{}, ev Event) bool

type eventFilterCtx struct {
	f EventFilter
	d interface{}
}

//export eventFilter
func eventFilter(data unsafe.Pointer, ev *cEvent) C.int {
	ctx := (*eventFilterCtx)(data)
	r := ctx.f(ctx.d, ev.toGo())

	if r {
		return 1
	}

	return 0
}

func SetEventFilter(filter EventFilter, userdata interface{}) {
	ctx := &eventFilterCtx{
		f: filter,
		d: userdata,
	}

	C.SetEventFilter(unsafe.Pointer(ctx))
}

func GetEventFilter() (EventFilter, interface{}) {
	var filter C.SDL_EventFilter
	var data unsafe.Pointer
	if C.SDL_GetEventFilter(&filter, &data) == C.SDL_FALSE {
		return nil, nil
	}

	ctx := (*eventFilterCtx)(data)

	return ctx.f, ctx.d
}

func AddEventWatch(filter EventFilter, data interface{}) {
	ctx := &eventFilterCtx{
		f: filter,
		d: data,
	}

	C.AddEventWatch(unsafe.Pointer(ctx))
}

// TODO: Find a way to make this work.
//func DelEventWatch(filter EventFilter, data interface{}) {
//	ctx := &eventFilterCtx{
//		f: filter,
//		d: data,
//	}
//
//	C.DelEventWatch(unsafe.Pointer(&eventFilter), ctx)
//}

func FilterEvents(filter EventFilter, data unsafe.Pointer) {
	ctx := &eventFilterCtx{
		f: filter,
		d: data,
	}

	C.FilterEvents(unsafe.Pointer(ctx))
}

const (
	QUERY   = C.SDL_QUERY
	IGNORE  = C.SDL_IGNORE
	DISABLE = C.SDL_DISABLE
	ENABLE  = C.SDL_ENABLE
)

func EventState(t uint32, s int) uint8 {
	return uint8(C.SDL_EventState(C.Uint32(t), C.int(s)))
}

func GetEventState(t uint32) uint8 {
	return uint8(C.GetEventState(C.Uint32(t)))
}

func RegisterEvents(num int) (uint32, error) {
	n := C.SDL_RegisterEvents(C.int(num))
	if uint32(n) == math.MaxUint32 {
		return 0, getError()
	}

	return uint32(n), nil
}
