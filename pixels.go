package sdl

import (
	"reflect"
	"unsafe"
)

// #include <SDL.h>
import "C"

const (
	ALPHA_OPAQUE      = 255
	ALPHA_TRANSPARENT = 0
)

const (
	PIXELTYPE_UNKNOWN  = C.SDL_PIXELTYPE_UNKNOWN
	PIXELTYPE_INDEX1   = C.SDL_PIXELTYPE_INDEX1
	PIXELTYPE_INDEX4   = C.SDL_PIXELTYPE_INDEX4
	PIXELTYPE_INDEX8   = C.SDL_PIXELTYPE_INDEX8
	PIXELTYPE_PACKED8  = C.SDL_PIXELTYPE_PACKED8
	PIXELTYPE_PACKED16 = C.SDL_PIXELTYPE_PACKED16
	PIXELTYPE_PACKED32 = C.SDL_PIXELTYPE_PACKED32
	PIXELTYPE_ARRAYU8  = C.SDL_PIXELTYPE_ARRAYU8
	PIXELTYPE_ARRAYU16 = C.SDL_PIXELTYPE_ARRAYU16
	PIXELTYPE_ARRAYU32 = C.SDL_PIXELTYPE_ARRAYU32
	PIXELTYPE_ARRAYF16 = C.SDL_PIXELTYPE_ARRAYF16
	PIXELTYPE_ARRAYF32 = C.SDL_PIXELTYPE_ARRAYF32
)

const (
	BITMAPORDER_NONE = C.SDL_BITMAPORDER_NONE
	BITMAPORDER_4321 = C.SDL_BITMAPORDER_4321
	BITMAPORDER_1234 = C.SDL_BITMAPORDER_1234
)

const (
	PACKEDORDER_NONE = C.SDL_PACKEDORDER_NONE
	PACKEDORDER_XRGB = C.SDL_PACKEDORDER_XRGB
	PACKEDORDER_RGBX = C.SDL_PACKEDORDER_RGBX
	PACKEDORDER_ARGB = C.SDL_PACKEDORDER_ARGB
	PACKEDORDER_RGBA = C.SDL_PACKEDORDER_RGBA
	PACKEDORDER_XBGR = C.SDL_PACKEDORDER_XBGR
	PACKEDORDER_BGRX = C.SDL_PACKEDORDER_BGRX
	PACKEDORDER_ABGR = C.SDL_PACKEDORDER_ABGR
	PACKEDORDER_BGRA = C.SDL_PACKEDORDER_BGRA
)

const (
	ARRAYORDER_NONE = C.SDL_ARRAYORDER_NONE
	ARRAYORDER_RGB  = C.SDL_ARRAYORDER_RGB
	ARRAYORDER_RGBA = C.SDL_ARRAYORDER_RGBA
	ARRAYORDER_ARGB = C.SDL_ARRAYORDER_ARGB
	ARRAYORDER_BGR  = C.SDL_ARRAYORDER_BGR
	ARRAYORDER_BGRA = C.SDL_ARRAYORDER_BGRA
	ARRAYORDER_ABGR = C.SDL_ARRAYORDER_ABGR
)

const (
	PACKEDLAYOUT_NONE    = C.SDL_PACKEDLAYOUT_NONE
	PACKEDLAYOUT_332     = C.SDL_PACKEDLAYOUT_332
	PACKEDLAYOUT_4444    = C.SDL_PACKEDLAYOUT_4444
	PACKEDLAYOUT_1555    = C.SDL_PACKEDLAYOUT_1555
	PACKEDLAYOUT_5551    = C.SDL_PACKEDLAYOUT_5551
	PACKEDLAYOUT_565     = C.SDL_PACKEDLAYOUT_565
	PACKEDLAYOUT_8888    = C.SDL_PACKEDLAYOUT_8888
	PACKEDLAYOUT_2101010 = C.SDL_PACKEDLAYOUT_2101010
	PACKEDLAYOUT_1010102 = C.SDL_PACKEDLAYOUT_1010102
)

// TODO: Macros...

const (
	PIXELFORMAT_UNKNOWN     = C.SDL_PIXELFORMAT_UNKNOWN
	PIXELFORMAT_INDEX1LSB   = C.SDL_PIXELFORMAT_INDEX1LSB
	PIXELFORMAT_INDEX1MSB   = C.SDL_PIXELFORMAT_INDEX1MSB
	PIXELFORMAT_INDEX4LSB   = C.SDL_PIXELFORMAT_INDEX4LSB
	PIXELFORMAT_INDEX4MSB   = C.SDL_PIXELFORMAT_INDEX4MSB
	PIXELFORMAT_INDEX8      = C.SDL_PIXELFORMAT_INDEX8
	PIXELFORMAT_RGB332      = C.SDL_PIXELFORMAT_RGB332
	PIXELFORMAT_RGB444      = C.SDL_PIXELFORMAT_RGB444
	PIXELFORMAT_RGB555      = C.SDL_PIXELFORMAT_RGB555
	PIXELFORMAT_BGR555      = C.SDL_PIXELFORMAT_BGR555
	PIXELFORMAT_ARGB4444    = C.SDL_PIXELFORMAT_ARGB4444
	PIXELFORMAT_RGBA4444    = C.SDL_PIXELFORMAT_RGBA4444
	PIXELFORMAT_ABGR4444    = C.SDL_PIXELFORMAT_ABGR4444
	PIXELFORMAT_BGRA4444    = C.SDL_PIXELFORMAT_BGRA4444
	PIXELFORMAT_ARGB1555    = C.SDL_PIXELFORMAT_ARGB1555
	PIXELFORMAT_RGBA5551    = C.SDL_PIXELFORMAT_RGBA5551
	PIXELFORMAT_ABGR1555    = C.SDL_PIXELFORMAT_ABGR1555
	PIXELFORMAT_BGRA5551    = C.SDL_PIXELFORMAT_BGRA5551
	PIXELFORMAT_RGB565      = C.SDL_PIXELFORMAT_RGB565
	PIXELFORMAT_BGR565      = C.SDL_PIXELFORMAT_BGR565
	PIXELFORMAT_RGB24       = C.SDL_PIXELFORMAT_RGB24
	PIXELFORMAT_BGR24       = C.SDL_PIXELFORMAT_BGR24
	PIXELFORMAT_RGB888      = C.SDL_PIXELFORMAT_RGB888
	PIXELFORMAT_RGBX8888    = C.SDL_PIXELFORMAT_RGBX8888
	PIXELFORMAT_BGR888      = C.SDL_PIXELFORMAT_BGR888
	PIXELFORMAT_BGRX8888    = C.SDL_PIXELFORMAT_BGRX8888
	PIXELFORMAT_ARGB8888    = C.SDL_PIXELFORMAT_ARGB8888
	PIXELFORMAT_RGBA8888    = C.SDL_PIXELFORMAT_RGBA8888
	PIXELFORMAT_ABGR8888    = C.SDL_PIXELFORMAT_ABGR8888
	PIXELFORMAT_BGRA8888    = C.SDL_PIXELFORMAT_BGRA8888
	PIXELFORMAT_ARGB2101010 = C.SDL_PIXELFORMAT_ARGB2101010

	PIXELFORMAT_IYUV = C.SDL_PIXELFORMAT_IYUV
	PIXELFORMAT_YUY2 = C.SDL_PIXELFORMAT_YUY2
	PIXELFORMAT_UYVY = C.SDL_PIXELFORMAT_UYVY
	PIXELFORMAT_YVYU = C.SDL_PIXELFORMAT_YVYU
)

type Color struct {
	R      uint8
	G      uint8
	B      uint
	unused uint8
}

func (c *Color) c() *C.SDL_Color {
	return (*C.SDL_Color)(unsafe.Pointer(c))
}

func (c Color) RGBA() (r, g, b, a uint32) {
	r = uint32(c.R) * 0xFFFF / 255
	g = uint32(c.G) * 0xFFFF / 255
	b = uint32(c.B) * 0xFFFF / 255
	a = 0xFFFF

	return
}

type Palette struct {
	ncolors  int32
	colors   uintptr
	Version  uint32
	Refcount int32
}

func (p *Palette) c() *C.SDL_Palette {
	return (*C.SDL_Palette)(unsafe.Pointer(p))
}

func (p *Palette) Colors() []Color {
	return *(*[]Color)(unsafe.Pointer(&reflect.SliceHeader{
		Data: p.colors,
		Len:  int(p.ncolors),
		Cap:  int(p.ncolors),
	}))
}

type PixelFormat struct {
	Format        uint32
	Palette       *Palette
	BitsPerPixel  uint8
	BytesPerPixel uint8
	padding       [2]uint8
	Rmask         uint32
	Gmask         uint32
	Bmask         uint32
	Amask         uint32
	Rloss         uint8
	Gloss         uint8
	Bloss         uint8
	Aloss         uint8
	Rshift        uint8
	Gshift        uint8
	Bshift        uint8
	Ashift        uint8
	Refcount      int32
	Next          *PixelFormat
}

func (p *PixelFormat) c() *C.SDL_PixelFormat {
	return (*C.SDL_PixelFormat)(unsafe.Pointer(p))
}

func GetPixelFormatName(f uint32) string {
	return C.GoString(C.SDL_GetPixelFormatName(C.Uint32(f)))
}

func PixelFormatEnumToMasks(f uint32) (bpp int, rm, gm, bm, am uint32, err error) {
	var cbpp C.int
	var crm, cgm, cbm, cam C.Uint32
	if C.SDL_PixelFormatEnumToMasks(C.Uint32(f), &cbpp, &crm, &cgm, &cbm, &cam) == C.SDL_FALSE {
		return 0, 0, 0, 0, 0, getError()
	}

	return int(cbpp), uint32(crm), uint32(cgm), uint32(cbm), uint32(cam), nil
}

func MasksToPixelFormatEnum(bpp int, rm, gm, bm, am uint32) (uint32, error) {
	f := C.SDL_MasksToPixelFormatEnum(C.int(bpp), C.Uint32(rm), C.Uint32(gm), C.Uint32(bm), C.Uint32(am))
	if f == PIXELFORMAT_UNKNOWN {
		return 0, getError()
	}

	return uint32(f), nil
}

func AllocFormat(f uint32) (*PixelFormat, error) {
	p := C.SDL_AllocFormat(C.Uint32(f))
	if p == nil {
		return nil, getError()
	}

	return (*PixelFormat)(unsafe.Pointer(p)), nil
}

func (p *PixelFormat) Free() {
	C.SDL_FreeFormat(p.c())
}

func AllocPalette(ncolors int) (*Palette, error) {
	p := C.SDL_AllocPalette(C.int(ncolors))
	if p == nil {
		return nil, getError()
	}

	return (*Palette)(unsafe.Pointer(p)), nil
}

func (p *PixelFormat) SetPalette(pal *Palette) error {
	if C.SDL_SetPixelFormatPalette(p.c(), pal.c()) != 0 {
		return getError()
	}

	return nil
}

func (p *Palette) SetColors(c []Color, firstcolor int) error {
	if C.SDL_SetPaletteColors(p.c(), c[0].c(), C.int(firstcolor), C.int(len(c))) != 0 {
		return getError()
	}

	return nil
}

func (p *Palette) Free() {
	C.SDL_FreePalette(p.c())
}

func (p *PixelFormat) MapRGB(r, g, b uint8) uint32 {
	return uint32(C.SDL_MapRGB(p.c(), C.Uint8(r), C.Uint8(g), C.Uint8(b)))
}

func (p *PixelFormat) MapRGBA(r, g, b, a uint8) uint32 {
	return uint32(C.SDL_MapRGBA(
		p.c(),
		C.Uint8(r),
		C.Uint8(g),
		C.Uint8(b),
		C.Uint8(a),
	))
}

func (p *PixelFormat) GetRGB(pix uint32) (r, g, b uint8) {
	var cr, cg, cb C.Uint8
	C.SDL_GetRGB(C.Uint32(pix), p.c(), &cr, &cg, &cb)

	return uint8(cr), uint8(cg), uint8(cb)
}

func (p *PixelFormat) GetRGBA(pix uint32) (r, g, b, a uint8) {
	var cr, cg, cb, ca C.Uint8
	C.SDL_GetRGBA(C.Uint32(pix), p.c(), &cr, &cg, &cb, &ca)

	return uint8(cr), uint8(cg), uint8(cb), uint8(ca)
}

func CalculateGammaRamp(gamma float32, ramp []uint16) {
	if len(ramp) < 256 {
		panic("len(ramp) must be > 256")
	}

	C.SDL_CalculateGammaRamp(C.float(gamma), (*C.Uint16)(&ramp[0]))
}
