package sdl

import (
	"fmt"
	"unsafe"
)

// #include <SDL.h>
//
// #include "log.h"
import "C"

const (
	MAX_LOG_MESSAGE = C.SDL_MAX_LOG_MESSAGE
)

const (
	LOG_CATEGORY_APPLICATION = C.SDL_LOG_CATEGORY_APPLICATION
	LOG_CATEGORY_ERROR       = C.SDL_LOG_CATEGORY_ERROR
	LOG_CATEGORY_SYSTEM      = C.SDL_LOG_CATEGORY_SYSTEM
	LOG_CATEGORY_AUDIO       = C.SDL_LOG_CATEGORY_AUDIO
	LOG_CATEGORY_VIDEO       = C.SDL_LOG_CATEGORY_VIDEO
	LOG_CATEGORY_RENDER      = C.SDL_LOG_CATEGORY_RENDER
	LOG_CATEGORY_INPUT       = C.SDL_LOG_CATEGORY_INPUT

	LOG_CATEGORY_RESERVED1  = C.SDL_LOG_CATEGORY_RESERVED1
	LOG_CATEGORY_RESERVED2  = C.SDL_LOG_CATEGORY_RESERVED2
	LOG_CATEGORY_RESERVED3  = C.SDL_LOG_CATEGORY_RESERVED3
	LOG_CATEGORY_RESERVED4  = C.SDL_LOG_CATEGORY_RESERVED4
	LOG_CATEGORY_RESERVED5  = C.SDL_LOG_CATEGORY_RESERVED5
	LOG_CATEGORY_RESERVED6  = C.SDL_LOG_CATEGORY_RESERVED6
	LOG_CATEGORY_RESERVED7  = C.SDL_LOG_CATEGORY_RESERVED7
	LOG_CATEGORY_RESERVED8  = C.SDL_LOG_CATEGORY_RESERVED8
	LOG_CATEGORY_RESERVED9  = C.SDL_LOG_CATEGORY_RESERVED9
	LOG_CATEGORY_RESERVED10 = C.SDL_LOG_CATEGORY_RESERVED10

	LOG_CATEGORY_CUSTOM = C.SDL_LOG_CATEGORY_CUSTOM
)

type LogPriority C.SDL_LogPriority

const (
	LOG_PRIORITY_VERBOSE  LogPriority = C.SDL_LOG_PRIORITY_VERBOSE
	LOG_PRIORITY_DEBUG    LogPriority = C.SDL_LOG_PRIORITY_DEBUG
	LOG_PRIORITY_INFO     LogPriority = C.SDL_LOG_PRIORITY_INFO
	LOG_PRIORITY_WARN     LogPriority = C.SDL_LOG_PRIORITY_WARN
	LOG_PRIORITY_ERROR    LogPriority = C.SDL_LOG_PRIORITY_ERROR
	LOG_PRIORITY_CRITICAL LogPriority = C.SDL_LOG_PRIORITY_CRITICAL
	NUM_LOG_PRIORITIES    LogPriority = C.SDL_NUM_LOG_PRIORITIES
)

func LogSetAllPriority(p LogPriority) {
	C.SDL_LogSetAllPriority(C.SDL_LogPriority(p))
}

func LogSetPriority(cat int, p LogPriority) {
	C.SDL_LogSetPriority(C.int(cat), C.SDL_LogPriority(p))
}

func LogGetPriority(cat int) LogPriority {
	return LogPriority(C.SDL_LogGetPriority(C.int(cat)))
}

func LogResetPriorities() {
	C.SDL_LogResetPriorities()
}

func Log(str string, args ...interface{}) {
	str = fmt.Sprintf(str, args...)

	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))

	C.Log(cstr)
}

func LogVerbose(cat int, str string, args ...interface{}) {
	str = fmt.Sprintf(str, args...)

	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))

	C.LogVerbose(C.int(cat), cstr)
}

func LogDebug(cat int, str string, args ...interface{}) {
	str = fmt.Sprintf(str, args...)

	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))

	C.LogDebug(C.int(cat), cstr)
}

func LogInfo(cat int, str string, args ...interface{}) {
	str = fmt.Sprintf(str, args...)

	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))

	C.LogInfo(C.int(cat), cstr)
}

func LogWarn(cat int, str string, args ...interface{}) {
	str = fmt.Sprintf(str, args...)

	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))

	C.LogWarn(C.int(cat), cstr)
}

func LogError(cat int, str string, args ...interface{}) {
	str = fmt.Sprintf(str, args...)

	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))

	C.LogError(C.int(cat), cstr)
}

func LogCritical(cat int, str string, args ...interface{}) {
	str = fmt.Sprintf(str, args...)

	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))

	C.LogCritical(C.int(cat), cstr)
}

func LogMessage(cat int, pri LogPriority, str string, args ...interface{}) {
	str = fmt.Sprintf(str, args...)

	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))

	C.LogMessage(C.int(cat), C.SDL_LogPriority(pri), cstr)
}

type LogOutputFunction func(data interface{}, cat int, pri LogPriority, message string)

type logOutputFunctionCtx struct {
	f LogOutputFunction
	d interface{}
}

//export logOutputFunction
func logOutputFunction(data unsafe.Pointer, cat C.int, pri C.SDL_LogPriority, message *C.char) {
	ctx := (*logOutputFunctionCtx)(data)

	ctx.f(ctx.d, int(cat), LogPriority(pri), C.GoString(message))
}

var (
	logOutputFunctionCache LogOutputFunction
	logOutputDataCache     interface{}
)

func LogGetOutputFunction() (LogOutputFunction, interface{}) {
	return logOutputFunctionCache, logOutputDataCache
}

func LogSetOutputFunction(f LogOutputFunction, data interface{}) {
	ctx := &logOutputFunctionCtx{
		f: f,
		d: data,
	}

	C.LogSetOutputFunction(unsafe.Pointer(ctx))

	logOutputFunctionCache = f
	logOutputDataCache = data
}
