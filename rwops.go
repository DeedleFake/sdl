package sdl

import (
	"bytes"
	"io"
	"reflect"
	"unsafe"
)

// #cgo pkg-config: sdl2
//
// #include <SDL.h>
//
// #include "rwops.h"
import "C"

type RWops struct {
	c *C.SDL_RWops
}

func (rw *RWops) Seek(offset int64, whence int) (int64, error) {
	return int64(C.seek(rw.c, C.long(offset), C.int(whence))), nil
}

func (rw *RWops) Tell() (int64, error) {
	return int64(C.tell(rw.c)), nil
}

func (rw *RWops) Read(data []byte) (int, error) {
	n := C.read(rw.c, unsafe.Pointer(&data[0]), 1, C.size_t(len(data)))
	switch {
	case n == 0:
		return 0, io.EOF
	case n < 0:
		return 0, getError()
	}

	return int(n), nil
}

func (rw *RWops) Write(data []byte) (int, error) {
	n := C.write(rw.c, unsafe.Pointer(&data[0]), 1, C.size_t(len(data)))
	switch {
	case n == 0:
		return 0, io.EOF
	case n < 0:
		return 0, getError()
	}

	return int(n), nil
}

func (rw *RWops) Close() error {
	if C.close(rw.c) != 0 {
		return getError()
	}

	return nil
}

func RWFromFile(file string, mode string) (*RWops, error) {
	cfile := C.CString(file)
	defer C.free(unsafe.Pointer(cfile))

	cmode := C.CString(mode)
	defer C.free(unsafe.Pointer(cmode))

	rw := C.SDL_RWFromFile(cfile, cmode)
	if rw == nil {
		return nil, getError()
	}

	return &RWops{rw}, nil
}

func RWFromMem(data []byte) (*RWops, error) {
	rw := C.SDL_RWFromMem(unsafe.Pointer(&data[0]), C.int(len(data)))
	if rw == nil {
		return nil, getError()
	}

	return &RWops{rw}, nil
}

func RWFromConstMem(data []byte) (*RWops, error) {
	rw := C.SDL_RWFromConstMem(unsafe.Pointer(&data[0]), C.int(len(data)))
	if rw == nil {
		return nil, getError()
	}

	return &RWops{rw}, nil
}

//export seekReadSeeker
func seekReadSeeker(ctx unsafe.Pointer, off int64, wh int) (int64, string) {
	r := *(*io.ReadSeeker)(ctx)

	n, err := r.Seek(off, wh)
	if err != nil {
		return -1, err.Error()
	}

	return n, ""
}

//export readReadSeeker
func readReadSeeker(ctx unsafe.Pointer, data unsafe.Pointer, size int, num int) (int, string) {
	r := *(*io.ReadSeeker)(ctx)
	buf := *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: uintptr(data),
		Len:  size * num,
		Cap:  size * num,
	}))

	n, err := r.Read(buf)
	if err != nil {
		if err == io.EOF {
			return 0, ""
		}

		return -1, err.Error()
	}

	return n, ""
}

func RWFromReadSeeker(r io.ReadSeeker) *RWops {
	rw := C.RWFromReadSeeker(unsafe.Pointer(&r))

	return &RWops{rw}
}

func RWFromReader(r io.Reader) (*RWops, error) {
	var buf bytes.Buffer
	_, err := io.Copy(&buf, r)
	if err != nil {
		return nil, err
	}

	return RWFromReadSeeker(bytes.NewReader(buf.Bytes())), nil
}
