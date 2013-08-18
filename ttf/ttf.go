package ttf

import (
	"github.com/DeedleFake/sdl"
	"unsafe"
)

// #cgo pkg-config: SDL2_ttf
// #include <SDL_ttf.h>
//
// #include "ttf.h"
import "C"

const (
	MAJOR_VERSION = C.SDL_TTF_MAJOR_VERSION
	MINOR_VERSION = C.SDL_TTF_MINOR_VERSION
	PATCHLEVEL    = C.SDL_TTF_PATCHLEVEL
)

const (
	UNICODE_BOM_NATIVE  = C.UNICODE_BOM_NATIVE
	UNICODE_BOM_SWAPPED = C.UNICODE_BOM_SWAPPED
)

const (
	STYLE_NORMAL        = C.TTF_STYLE_NORMAL
	STYLE_BOLD          = C.TTF_STYLE_BOLD
	STYLE_ITALIC        = C.TTF_STYLE_ITALIC
	STYLE_UNDERLINE     = C.TTF_STYLE_UNDERLINE
	STYLE_STRIKETHROUGH = C.TTF_STYLE_STRIKETHROUGH
)

const (
	HINTING_NORMAL = C.TTF_HINTING_NORMAL
	HINTING_LIGHT  = C.TTF_HINTING_LIGHT
	HINTING_MONO   = C.TTF_HINTING_MONO
	HINTING_NONE   = C.TTF_HINTING_NONE
)

type Font struct {
	c *C.TTF_Font
}

func (f *Font) GetStyle() int {
	return int(C.TTF_GetFontStyle(f.c))
}

func (f *Font) SetStyle(style int) {
	C.TTF_SetFontStyle(f.c, C.int(style))
}

func (f *Font) GetOutline() int {
	return int(C.TTF_GetFontOutline(f.c))
}

func (f *Font) SetOutline(outline int) {
	C.TTF_SetFontOutline(f.c, C.int(outline))
}

func (f *Font) GetHinting() int {
	return int(C.TTF_GetFontHinting(f.c))
}

func (f *Font) SetHinting(hinting int) {
	C.TTF_SetFontHinting(f.c, C.int(hinting))
}

func (f *Font) Height() int {
	return int(C.TTF_FontHeight(f.c))
}

func (f *Font) Ascent() int {
	return int(C.TTF_FontAscent(f.c))
}

func (f *Font) Descent() int {
	return int(C.TTF_FontDescent(f.c))
}

func (f *Font) LineSkip() int {
	return int(C.TTF_FontLineSkip(f.c))
}

func (f *Font) GetKerning() int {
	return int(C.TTF_GetFontKerning(f.c))
}

func (f *Font) SetKerning(kerning int) {
	C.TTF_SetFontKerning(f.c, C.int(kerning))
}

func (f *Font) Faces() int64 {
	return int64(C.TTF_FontFaces(f.c))
}

func (f *Font) FaceIsFixedWidth() int {
	return int(C.TTF_FontFaceIsFixedWidth(f.c))
}

func (f *Font) FaceFamilyName() string {
	return C.GoString(C.TTF_FontFaceFamilyName(f.c))
}

func (f *Font) FaceStyleName() string {
	return C.GoString(C.TTF_FontFaceStyleName(f.c))
}

func (f *Font) GlyphIsProvided(ch uint16) int {
	return int(C.TTF_GlyphIsProvided(f.c, C.Uint16(ch)))
}

func (f *Font) GlyphMetrics(ch uint16) (int, int, int, int, int, error) {
	minx := C.int(0)
	maxx := C.int(0)
	miny := C.int(0)
	maxy := C.int(0)
	advance := C.int(0)
	res := C.TTF_GlyphMetrics(f.c, C.Uint16(ch), &minx, &maxx, &miny, &maxy, &advance)
	var err error
	if res != 0 {
		err = getError()
	}
	return int(minx), int(maxx), int(miny), int(maxy), int(advance), err
}

func (f *Font) SizeText(text string) (int, int, error) {
	w := C.int(0)
	h := C.int(0)
	ctext := C.CString(text)
	res := C.TTF_SizeText(f.c, ctext, &w, &h)
	var err error
	if res != 0 {
		err = getError()
	}
	return int(w), int(h), err
}

func (f *Font) SizeUTF8(text string) (int, int, error) {
	w := C.int(0)
	h := C.int(0)
	ctext := C.CString(text)
	res := C.TTF_SizeUTF8(f.c, ctext, &w, &h)
	var err error
	if res != 0 {
		err = getError()
	}
	return int(w), int(h), err
}

func (f *Font) RenderTextSolid(text string, fg sdl.Color) (*sdl.Surface, error) {
	ctext := C.CString(text)
	ccolor := C.SDL_Color{C.Uint8(fg.R), C.Uint8(fg.G), C.Uint8(fg.B), C.Uint8(255)}
	s := C.TTF_RenderText_Solid(f.c, ctext, ccolor)
	if s == nil {
		return nil, getError()
	}
	return (*sdl.Surface)(unsafe.Pointer(s)), nil
}

func (f *Font) RenderUTF8Solid(text string, fg sdl.Color) (*sdl.Surface, error) {
	ctext := C.CString(text)
	ccolor := C.SDL_Color{C.Uint8(fg.R), C.Uint8(fg.G), C.Uint8(fg.B), C.Uint8(255)}
	s := C.TTF_RenderUTF8_Solid(f.c, ctext, ccolor)
	if s == nil {
		return nil, getError()
	}
	return (*sdl.Surface)(unsafe.Pointer(s)), nil
}

func (f *Font) Close() {
	C.TTF_CloseFont(f.c)
}

func VERSION() *sdl.Version {
	var v sdl.Version
	C.TTFVERSION(cVersion(&v))
	return &v
}

func Linked_Version() *sdl.Version {
	return (*sdl.Version)(unsafe.Pointer(C.TTF_Linked_Version()))
}

func ByteSwappedUnicode(swapped int) {
	C.TTF_ByteSwappedUNICODE(C.int(swapped))
}

func Init() error {
	if C.TTF_Init() != 0 {
		return getError()
	}
	return nil
}

func OpenFont(file string, ptsize int) (*Font, error) {
	cfile := C.CString(file)
	defer C.free(unsafe.Pointer(cfile))
	font := C.TTF_OpenFont(cfile, C.int(ptsize))
	if font == nil {
		return nil, getError()
	}
	return &Font{c: font}, nil
}

func OpenFontIndex(file string, ptsize int, index int64) (*Font, error) {
	cfile := C.CString(file)
	defer C.free(unsafe.Pointer(cfile))
	font := C.TTF_OpenFontIndex(cfile, C.int(ptsize), C.long(index))
	if font == nil {
		return nil, getError()
	}
	return &Font{c: font}, nil
}

func Quit() {
	C.TTF_Quit()
}

func WasInit() int {
	return int(C.TTF_WasInit())
}
