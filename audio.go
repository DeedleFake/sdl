// TODO: Large parts of this are probably broken. Find out what they are and fix them.

package sdl

import (
	"errors"
	"reflect"
	"unsafe"
)

// #include <SDL.h>
//
// #include "audio.h"
import "C"

type AudioFormat C.SDL_AudioFormat

func (af *AudioFormat) c() *C.SDL_AudioFormat {
	return (*C.SDL_AudioFormat)(unsafe.Pointer(af))
}

const (
	AUDIO_MASK_BITSIZE  = C.SDL_AUDIO_MASK_BITSIZE
	AUDIO_MASK_DATATYPE = C.SDL_AUDIO_MASK_DATATYPE
	AUDIO_MASK_ENDIAN   = C.SDL_AUDIO_MASK_ENDIAN
	AUDIO_MASK_SIGNED   = C.SDL_AUDIO_MASK_SIGNED
)

func AUDIO_BITSIZE(x uint32) uint32 {
	return uint32(C.AUDIO_BITSIZE(C.Uint32(x)))
}

func AUDIO_ISFLOAT(x uint32) uint32 {
	return uint32(C.AUDIO_ISFLOAT(C.Uint32(x)))
}

func AUDIO_ISBIGENDIAN(x uint32) uint32 {
	return uint32(C.AUDIO_ISBIGENDIAN(C.Uint32(x)))
}

func AUDIO_ISSIGNED(x uint32) uint32 {
	return uint32(C.AUDIO_ISSIGNED(C.Uint32(x)))
}

func AUDIO_ISINT(x uint32) uint32 {
	return uint32(C.AUDIO_ISINT(C.Uint32(x)))
}

func AUDIO_ISLITTLEENDIAN(x uint32) uint32 {
	return uint32(C.AUDIO_ISLITTLEENDIAN(C.Uint32(x)))
}

func AUDIO_ISUNSIGNED(x uint32) uint32 {
	return uint32(C.AUDIO_ISUNSIGNED(C.Uint32(x)))
}

const (
	AUDIO_U8     = C.AUDIO_U8
	AUDIO_S8     = C.AUDIO_S8
	AUDIO_U16LSB = C.AUDIO_U16LSB
	AUDIO_S16LSB = C.AUDIO_S16LSB
	AUDIO_U16MSB = C.AUDIO_U16MSB
	AUDIO_S16MSB = C.AUDIO_S16MSB
	AUDIO_U16    = C.AUDIO_U16
	AUDIO_S16    = C.AUDIO_S16
)

const (
	AUDIO_S32LSB = C.AUDIO_S32LSB
	AUDIO_S32MSB = C.AUDIO_S32MSB
	AUDIO_S32    = C.AUDIO_S32
)

const (
	AUDIO_F32LSB = C.AUDIO_F32LSB
	AUDIO_F32MSB = C.AUDIO_F32MSB
	AUDIO_F32    = C.AUDIO_F32
)

const (
	AUDIO_U16SYS = C.AUDIO_U16SYS
	AUDIO_S16SYS = C.AUDIO_S16SYS
	AUDIO_S32SYS = C.AUDIO_S32SYS
	AUDIO_F32SYS = C.AUDIO_F32SYS
)

const (
	AUDIO_ALLOW_FREQUENCY_CHANGE = C.SDL_AUDIO_ALLOW_FREQUENCY_CHANGE
	AUDIO_ALLOW_FORMAT_CHANGE    = C.SDL_AUDIO_ALLOW_FORMAT_CHANGE
	AUDIO_ALLOW_CHANNELS_CHANGE  = C.SDL_AUDIO_ALLOW_CHANNELS_CHANGE
	AUDIO_ALLOW_ANY_CHANGE       = C.SDL_AUDIO_ALLOW_ANY_CHANGE
)

type AudioCallback func(interface{}, []uint8)

type audioCallbackCtx struct {
	f AudioCallback
	d interface{}
}

//export audioCallback
func audioCallback(data unsafe.Pointer, stream *C.Uint8, length C.int) {
	ctx := (*audioCallbackCtx)(data)
	ss := *(*[]uint8)(unsafe.Pointer(&reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(stream)),
		Len:  int(length),
		Cap:  int(length),
	}))

	ctx.f(ctx.d, ss)
}

type AudioSpec struct {
	Freq     int32
	Format   AudioFormat
	Channels uint8
	Silence  uint8
	Samples  uint16
	padding  uint16
	Size     uint32
	callback C.SDL_AudioCallback
	data     unsafe.Pointer
}

func (a *AudioSpec) c() *C.SDL_AudioSpec {
	return (*C.SDL_AudioSpec)(unsafe.Pointer(a))
}

func (a *AudioSpec) GetData() interface{} {
	return (*interface{})(a.data)
}

func (a *AudioSpec) SetData(v interface{}) {
	*(*interface{})(a.data) = v
}

//type AudioFilter func(*AudioCVT, AudioFormat)
//
//type audioFilterCtx struct {
//	f AudioFilter
//	d interface{}
//}
//
//// TODO: Fix AudioFilter related stuff.
//
////export audioFilter
//func audioFilter(ccvt *C.SDL_AudioCVT, cformat C.SDL_AudioFormat) {
//	cvt := (*AudioCVT)(unsafe.Pointer(ccvt))
//	format := AudioFormat(cformat)
//}

type AudioCVT struct {
	needed      int32
	SrcFormat   AudioFormat
	DstFormat   AudioFormat
	RateIncr    float64
	buf         uintptr
	length      int32
	lenCvt      int32
	lenMult     int32
	lenRatio    float64
	filters     uintptr
	filterIndex int32
}

func (a *AudioCVT) c() *C.SDL_AudioCVT {
	return (*C.SDL_AudioCVT)(unsafe.Pointer(a))
}

func (a *AudioCVT) Buf() []byte {
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: a.buf,
		Len:  int(a.length),
		Cap:  int(a.length),
	}))
}

func (a *AudioCVT) Filters() []uintptr {
	return *(*[]uintptr)(unsafe.Pointer(&reflect.SliceHeader{
		Data: a.filters,
		Len:  10,
		Cap:  10,
	}))
}

func GetNumAudioDrivers() int {
	return int(C.SDL_GetNumAudioDrivers())
}

func GetAudioDriver(i int) (string, error) {
	name := C.SDL_GetAudioDriver(C.int(i))
	if name == nil {
		return "", getError()
	}

	return C.GoString(name), nil
}

func AudioInit(name string) error {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	if C.SDL_AudioInit(cname) != 0 {
		return getError()
	}

	return nil
}

func AudioQuit() {
	C.SDL_AudioQuit()
}

func GetCurrentAudioDriver() string {
	return C.GoString(C.SDL_GetCurrentAudioDriver())
}

func OpenAudio(des *AudioSpec) (*AudioSpec, error) {
	var o *AudioSpec
	if C.SDL_OpenAudio(des.c(), o.c()) != 0 {
		return nil, getError()
	}

	return o, nil
}

type AudioDeviceID C.SDL_AudioDeviceID

func (id *AudioDeviceID) c() *C.SDL_AudioDeviceID {
	return (*C.SDL_AudioDeviceID)(unsafe.Pointer(id))
}

func GetNumAudioDevices(ic bool) int {
	var cic C.int
	if ic {
		cic = 1
	}

	return int(C.SDL_GetNumAudioDevices(cic))
}

func GetAudioDeviceName(i int, ic bool) (string, error) {
	var cic C.int
	if ic {
		cic = 1
	}

	name := C.SDL_GetAudioDeviceName(C.int(i), cic)
	if name == nil {
		return "", getError()
	}

	return C.GoString(name), nil
}

func OpenAudioDevice(dev string, ic bool, des *AudioSpec, ac bool) (AudioDeviceID, *AudioSpec, error) {
	cdev := C.CString(dev)
	defer C.free(unsafe.Pointer(cdev))

	var cic C.int
	if ic {
		cic = 1
	}

	var cac C.int
	if ac {
		cac = 1
	}

	var o AudioSpec
	id := C.SDL_OpenAudioDevice(cdev, cic, des.c(), o.c(), cac)
	if id == 0 {
		return 0, nil, getError()
	}

	return AudioDeviceID(id), &o, nil
}

type AudioStatus C.SDL_AudioStatus

func (a *AudioStatus) c() *C.SDL_AudioStatus {
	return (*C.SDL_AudioStatus)(unsafe.Pointer(a))
}

const (
	AUDIO_STOPPED AudioStatus = C.SDL_AUDIO_STOPPED
	AUDIO_PLAYING AudioStatus = C.SDL_AUDIO_PLAYING
	AUDIO_PAUSED  AudioStatus = C.SDL_AUDIO_PAUSED
)

func GetAudioStatus() AudioStatus {
	return AudioStatus(C.SDL_GetAudioStatus())
}

func (id AudioDeviceID) GetStatus() AudioStatus {
	return AudioStatus(C.SDL_GetAudioDeviceStatus(C.SDL_AudioDeviceID(id)))
}

func PauseAudio(p bool) {
	var cp C.int
	if p {
		cp = 1
	}

	C.SDL_PauseAudio(cp)
}

func (id AudioDeviceID) PauseAudio(p bool) {
	var cp C.int
	if p {
		cp = 1
	}

	C.SDL_PauseAudioDevice(C.SDL_AudioDeviceID(id), cp)
}

func LoadWAV_RW(rw *RWops, free bool) (*AudioSpec, []byte, error) {
	var cfree C.int
	if free {
		cfree = 1
	}

	var a AudioSpec
	var h reflect.SliceHeader
	var l C.Uint32
	if C.SDL_LoadWAV_RW(rw.c, cfree, a.c(), (**C.Uint8)(unsafe.Pointer(&h.Data)), &l) == nil {
		return nil, nil, getError()
	}
	h.Len = int(l)
	h.Cap = int(l)

	return &a, *(*[]byte)(unsafe.Pointer(&h)), nil
}

func LoadWAV(file string) (*AudioSpec, []byte, error) {
	cfile := C.CString(file)
	defer C.free(unsafe.Pointer(cfile))

	var a AudioSpec
	var h reflect.SliceHeader
	var l C.Uint32
	if C.LoadWAV(cfile, a.c(), (**C.Uint8)(unsafe.Pointer(&h.Data)), &l) == nil {
		return nil, nil, getError()
	}
	h.Len = int(l)
	h.Cap = int(l)

	return &a, *(*[]byte)(unsafe.Pointer(&h)), nil
}

func FreeWAV(buf []byte) {
	C.SDL_FreeWAV((*C.Uint8)(unsafe.Pointer(&buf[0])))
}

func BuildAudioCVT(sf AudioFormat, sc uint8, sr int, df AudioFormat, dc uint8, dr int) (*AudioCVT, bool, error) {
	var cvt AudioCVT
	r := C.SDL_BuildAudioCVT(cvt.c(),
		C.SDL_AudioFormat(sf),
		C.Uint8(sc),
		C.int(sr),
		C.SDL_AudioFormat(df),
		C.Uint8(dc),
		C.int(dr),
	)
	switch r {
	case 0:
		return &cvt, false, nil
	case 1:
		return &cvt, true, nil
	}

	return nil, false, getError()
}

func (cvt *AudioCVT) ConvertAudio() error {
	if C.SDL_ConvertAudio(cvt.c()) != 0 {
		return getError()
	}

	return nil
}

const (
	MIX_MAXVOLUME = C.SDL_MIX_MAXVOLUME
)

func MixAudio(dst, src []byte, vol int) error {
	if len(dst) != len(src) {
		return errors.New("len(dst) != len(src)")
	}

	C.SDL_MixAudio(
		(*C.Uint8)(unsafe.Pointer(&dst[0])),
		(*C.Uint8)(unsafe.Pointer(&src[0])),
		C.Uint32(len(dst)),
		C.int(vol),
	)

	return nil
}

func MixAudioFormat(dst, src []byte, f AudioFormat, vol int) error {
	if len(dst) != len(src) {
		return errors.New("len(dst) != len(src)")
	}

	C.SDL_MixAudioFormat(
		(*C.Uint8)(unsafe.Pointer(&dst[0])),
		(*C.Uint8)(unsafe.Pointer(&src[0])),
		C.SDL_AudioFormat(f),
		C.Uint32(len(dst)),
		C.int(vol),
	)

	return nil
}

func LockAudio() {
	C.SDL_LockAudio()
}

func (dev AudioDeviceID) Lock() {
	C.SDL_LockAudioDevice(C.SDL_AudioDeviceID(dev))
}

func UnlockAudio() {
	C.SDL_UnlockAudio()
}

func (dev AudioDeviceID) Unlock() {
	C.SDL_UnlockAudioDevice(C.SDL_AudioDeviceID(dev))
}

func CloseAudio() {
	C.SDL_CloseAudio()
}

func (dev AudioDeviceID) Close() {
	C.SDL_CloseAudioDevice(C.SDL_AudioDeviceID(dev))
}

// TODO: Is this not implemented in SDL?
//func (dev AudioDeviceID) Connected() (bool, error) {
//	r := C.SDL_AudioDeviceConnected(dev.c())
//	switch r {
//	case 0:
//		return false, nil
//	case 1:
//		return true, nil
//	}
//
//	return false, getError()
//}
