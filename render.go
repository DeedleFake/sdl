package sdl

import (
	"errors"
	"reflect"
	"unsafe"
)

// #include <SDL.h>
import "C"

type RendererFlags uint32

const (
	RENDERER_SOFTWARE      RendererFlags = C.SDL_RENDERER_SOFTWARE
	RENDERER_ACCELERATED   RendererFlags = C.SDL_RENDERER_ACCELERATED
	RENDERER_PRESENTVSYNC  RendererFlags = C.SDL_RENDERER_PRESENTVSYNC
	RENDERER_TARGETTEXTURE RendererFlags = C.SDL_RENDERER_TARGETTEXTURE
)

type RendererInfo struct {
	name              *C.char
	Flags             RendererFlags
	NumTextureFormats uint32
	TextureFormats    [16]uint32
	MaxTextureWidth   int32
	MaxTextureHeight  int32
}

func (r *RendererInfo) c() *C.SDL_RendererInfo {
	return (*C.SDL_RendererInfo)(unsafe.Pointer(r))
}

func (r *RendererInfo) Name() string {
	return C.GoString(r.name)
}

type TextureAccess int

const (
	TEXTUREACCESS_STATIC    TextureAccess = C.SDL_TEXTUREACCESS_STATIC
	TEXTUREACCESS_STREAMING TextureAccess = C.SDL_TEXTUREACCESS_STREAMING
	TEXTUREACCESS_TARGET    TextureAccess = C.SDL_TEXTUREACCESS_TARGET
)

type TextureModulate uint32

const (
	TEXTUREMODULATE_NONE  TextureModulate = C.SDL_TEXTUREMODULATE_NONE
	TEXTUREMODULATE_COLOR TextureModulate = C.SDL_TEXTUREMODULATE_COLOR
	TEXTUREMODULATE_ALPHA TextureModulate = C.SDL_TEXTUREMODULATE_ALPHA
)

type RendererFlip C.SDL_RendererFlip

func (r RendererFlip) c() C.SDL_RendererFlip {
	return C.SDL_RendererFlip(r)
}

const (
	FLIP_NONE       RendererFlip = C.SDL_FLIP_NONE
	FLIP_HORIZONTAL RendererFlip = C.SDL_FLIP_HORIZONTAL
	FLIP_VERTICAL   RendererFlip = C.SDL_FLIP_VERTICAL
)

type Renderer struct {
	c *C.SDL_Renderer
}

type Texture struct {
	c *C.SDL_Texture
}

func GetNumRenderDrivers() int {
	return int(C.SDL_GetNumRenderDrivers())
}

func GetRenderDriverInfo(i int) (*RendererInfo, error) {
	var r RendererInfo
	if C.SDL_GetRenderDriverInfo(C.int(i), r.c()) != 0 {
		return nil, getError()
	}

	return &r, nil
}

func CreateWindowAndRenderer(w, h int, flags WindowFlags) (*Window, *Renderer, error) {
	var win *C.SDL_Window
	var r *C.SDL_Renderer
	if C.SDL_CreateWindowAndRenderer(C.int(w), C.int(h), C.Uint32(flags), &win, &r) != 0 {
		return nil, nil, getError()
	}

	return &Window{win}, &Renderer{r}, nil
}

func (w *Window) CreateRenderer(i int, flags RendererFlags) (*Renderer, error) {
	r := C.SDL_CreateRenderer(w.c, C.int(i), C.Uint32(flags))
	if r == nil {
		return nil, getError()
	}

	return &Renderer{r}, nil
}

func (s *Surface) CreateSoftwareRenderer() (*Renderer, error) {
	r := C.SDL_CreateSoftwareRenderer(s.c())
	if r == nil {
		return nil, getError()
	}

	return &Renderer{r}, nil
}

func (w *Window) GetRenderer() (*Renderer, error) {
	r := C.SDL_GetRenderer(w.c)
	if r == nil {
		return nil, getError()
	}

	return &Renderer{r}, nil
}

func (r *Renderer) GetInfo() (*RendererInfo, error) {
	var ri RendererInfo
	if C.SDL_GetRendererInfo(r.c, ri.c()) != 0 {
		return nil, getError()
	}

	return &ri, nil
}

func (r *Renderer) CreateTexture(format uint32, access TextureAccess, w, h int) (*Texture, error) {
	t := C.SDL_CreateTexture(r.c,
		C.Uint32(format),
		C.int(access),
		C.int(w),
		C.int(h),
	)
	if t == nil {
		return nil, getError()
	}

	return &Texture{t}, nil
}

func (r *Renderer) CreateTextureFromSurface(s *Surface) (*Texture, error) {
	t := C.SDL_CreateTextureFromSurface(r.c, s.c())
	if t == nil {
		return nil, getError()
	}

	return &Texture{t}, nil
}

func (t *Texture) Query() (uint32, TextureAccess, int, int, error) {
	var format C.Uint32
	var access C.int
	var w, h C.int
	if C.SDL_QueryTexture(t.c, &format, &access, &w, &h) != 0 {
		return 0, 0, 0, 0, getError()
	}

	return uint32(format), TextureAccess(access), int(w), int(h), nil
}

func (t *Texture) SetColorMod(r, g, b uint8) error {
	if C.SDL_SetTextureColorMod(t.c, C.Uint8(r), C.Uint8(g), C.Uint8(b)) != 0 {
		return getError()
	}

	return nil
}

func (t *Texture) GetColorMode() (r, g, b uint8, err error) {
	var cr, cg, cb C.Uint8
	if C.SDL_GetTextureColorMod(t.c, &cr, &cg, &cb) != 0 {
		return 0, 0, 0, getError()
	}

	return uint8(cr), uint8(cg), uint8(cb), nil
}

func (t *Texture) SetAlphaMod(a uint8) error {
	if C.SDL_SetTextureAlphaMod(t.c, C.Uint8(a)) != 0 {
		return getError()
	}

	return nil
}

func (t *Texture) GetAlphaMod() (uint8, error) {
	var a C.Uint8
	if C.SDL_GetTextureAlphaMod(t.c, &a) != 0 {
		return 0, getError()
	}

	return uint8(a), nil
}

func (t *Texture) SetBlendMode(mode BlendMode) error {
	if C.SDL_SetTextureBlendMode(t.c, mode.c()) != 0 {
		return getError()
	}

	return nil
}

func (t *Texture) GetBlendMode() (BlendMode, error) {
	var mode C.SDL_BlendMode
	if C.SDL_GetTextureBlendMode(t.c, &mode) != 0 {
		return 0, getError()
	}

	return BlendMode(mode), nil
}

func (t *Texture) Update8(r *Rect, pix []uint8, pitch int) error {
	if r == nil {
		_, _, w, h, err := t.Query()
		if err != nil {
			return err
		}

		r = &Rect{
			W: int32(w),
			H: int32(h),
		}
	}

	if len(pix) < int(r.W)*int(r.H) {
		return errors.New("len(pix) < r.W * r.H")
	}

	if C.SDL_UpdateTexture(t.c, r.c(), unsafe.Pointer(&pix[0]), C.int(pitch)) != 0 {
		return getError()
	}

	return nil
}

func (t *Texture) Update16(r *Rect, pix []uint16, pitch int) error {
	if r == nil {
		_, _, w, h, err := t.Query()
		if err != nil {
			return err
		}

		r = &Rect{
			W: int32(w),
			H: int32(h),
		}
	}

	if len(pix) < int(r.W)*int(r.H) {
		return errors.New("len(pix) < r.W * r.H")
	}

	if C.SDL_UpdateTexture(t.c, r.c(), unsafe.Pointer(&pix[0]), C.int(pitch)) != 0 {
		return getError()
	}

	return nil
}

func (t *Texture) Update32(r *Rect, pix []uint32, pitch int) error {
	if r == nil {
		_, _, w, h, err := t.Query()
		if err != nil {
			return err
		}

		r = &Rect{
			W: int32(w),
			H: int32(h),
		}
	}

	if len(pix) < int(r.W)*int(r.H) {
		return errors.New("len(pix) < r.W * r.H")
	}

	if C.SDL_UpdateTexture(t.c, r.c(), unsafe.Pointer(&pix[0]), C.int(pitch)) != 0 {
		return getError()
	}

	return nil
}

func (t *Texture) Update64(r *Rect, pix []uint64, pitch int) error {
	if r == nil {
		_, _, w, h, err := t.Query()
		if err != nil {
			return err
		}

		r = &Rect{
			W: int32(w),
			H: int32(h),
		}
	}

	if len(pix) < int(r.W)*int(r.H) {
		return errors.New("len(pix) < r.W * r.H")
	}

	if C.SDL_UpdateTexture(t.c, r.c(), unsafe.Pointer(&pix[0]), C.int(pitch)) != 0 {
		return getError()
	}

	return nil
}

func (t *Texture) Lock8(r *Rect) (pix []uint8, pitch int, err error) {
	if r == nil {
		_, _, w, h, err := t.Query()
		if err != nil {
			return nil, 0, err
		}

		r = &Rect{
			W: int32(w),
			H: int32(h),
		}
	}

	var cpix unsafe.Pointer
	var cpitch C.int
	if C.SDL_LockTexture(t.c, r.c(), &cpix, &cpitch) != 0 {
		return nil, 0, getError()
	}

	pix = *(*[]uint8)(unsafe.Pointer(&reflect.SliceHeader{
		Data: uintptr(cpix),
		Len:  int(r.W) * int(r.H),
		Cap:  int(r.W) * int(r.H),
	}))

	return pix, int(cpitch), nil
}

func (t *Texture) Lock16(r *Rect) (pix []uint16, pitch int, err error) {
	if r == nil {
		_, _, w, h, err := t.Query()
		if err != nil {
			return nil, 0, err
		}

		r = &Rect{
			W: int32(w),
			H: int32(h),
		}
	}

	var cpix unsafe.Pointer
	var cpitch C.int
	if C.SDL_LockTexture(t.c, r.c(), &cpix, &cpitch) != 0 {
		return nil, 0, getError()
	}

	pix = *(*[]uint16)(unsafe.Pointer(&reflect.SliceHeader{
		Data: uintptr(cpix),
		Len:  int(r.W) * int(r.H),
		Cap:  int(r.W) * int(r.H),
	}))

	return pix, int(cpitch), nil
}

func (t *Texture) Lock32(r *Rect) (pix []uint32, pitch int, err error) {
	if r == nil {
		_, _, w, h, err := t.Query()
		if err != nil {
			return nil, 0, err
		}

		r = &Rect{
			W: int32(w),
			H: int32(h),
		}
	}

	var cpix unsafe.Pointer
	var cpitch C.int
	if C.SDL_LockTexture(t.c, r.c(), &cpix, &cpitch) != 0 {
		return nil, 0, getError()
	}

	pix = *(*[]uint32)(unsafe.Pointer(&reflect.SliceHeader{
		Data: uintptr(cpix),
		Len:  int(r.W) * int(r.H),
		Cap:  int(r.W) * int(r.H),
	}))

	return pix, int(cpitch), nil
}

func (t *Texture) Lock64(r *Rect) (pix []uint64, pitch int, err error) {
	if r == nil {
		_, _, w, h, err := t.Query()
		if err != nil {
			return nil, 0, err
		}

		r = &Rect{
			W: int32(w),
			H: int32(h),
		}
	}

	var cpix unsafe.Pointer
	var cpitch C.int
	if C.SDL_LockTexture(t.c, r.c(), &cpix, &cpitch) != 0 {
		return nil, 0, getError()
	}

	pix = *(*[]uint64)(unsafe.Pointer(&reflect.SliceHeader{
		Data: uintptr(cpix),
		Len:  int(r.W) * int(r.H),
		Cap:  int(r.W) * int(r.H),
	}))

	return pix, int(cpitch), nil
}

func (t *Texture) Unlock() {
	C.SDL_UnlockTexture(t.c)
}

func (r *Renderer) RenderTargetSupported() bool {
	return C.SDL_RenderTargetSupported(r.c) == C.SDL_TRUE
}

func (r *Renderer) SetRenderTarget(t *Texture) error {
	if C.SDL_SetRenderTarget(r.c, t.c) != 0 {
		return getError()
	}

	return nil
}

func (r *Renderer) SetViewport(rect *Rect) error {
	if C.SDL_RenderSetViewport(r.c, rect.c()) != 0 {
		return getError()
	}

	return nil
}

func (r *Renderer) GetViewport() *Rect {
	var rect Rect
	C.SDL_RenderGetViewport(r.c, rect.c())

	return &rect
}

func (r *Renderer) SetDrawColor(red, g, b, a uint8) error {
	if C.SDL_SetRenderDrawColor(r.c, C.Uint8(red), C.Uint8(g), C.Uint8(b), C.Uint8(a)) != 0 {
		return getError()
	}

	return nil
}

func (r *Renderer) GetDrawColor() (red, g, b, a uint8, err error) {
	var cr, cg, cb, ca C.Uint8
	if C.SDL_GetRenderDrawColor(r.c, &cr, &cg, &cb, &ca) != 0 {
		return 0, 0, 0, 0, getError()
	}

	return uint8(cr), uint8(cg), uint8(cb), uint8(ca), nil
}

func (r *Renderer) SetDrawBlendMode(mode BlendMode) error {
	if C.SDL_SetRenderDrawBlendMode(r.c, mode.c()) != 0 {
		return getError()
	}

	return nil
}

func (r *Renderer) GetDrawBlendMode() (BlendMode, error) {
	var mode C.SDL_BlendMode
	if C.SDL_GetRenderDrawBlendMode(r.c, &mode) != 0 {
		return 0, getError()
	}

	return BlendMode(mode), nil
}

func (r *Renderer) Clear() error {
	if C.SDL_RenderClear(r.c) != 0 {
		return getError()
	}

	return nil
}

func (r *Renderer) DrawPoint(x, y int) error {
	if C.SDL_RenderDrawPoint(r.c, C.int(x), C.int(y)) != 0 {
		return getError()
	}

	return nil
}

func (r *Renderer) DrawPoints(points []Point) error {
	if C.SDL_RenderDrawPoints(r.c, (*C.SDL_Point)(unsafe.Pointer(&points[0])), C.int(len(points))) != 0 {
		return getError()
	}

	return nil
}

func (r *Renderer) DrawLine(x1, y1, x2, y2 int) error {
	if C.SDL_RenderDrawLine(r.c, C.int(x1), C.int(y1), C.int(x2), C.int(y2)) != 0 {
		return getError()
	}

	return nil
}

func (r *Renderer) DrawRect(rect *Rect) error {
	if C.SDL_RenderDrawRect(r.c, rect.c()) != 0 {
		return getError()
	}

	return nil
}

func (r *Renderer) DrawRects(rects []Rect) error {
	if C.SDL_RenderDrawRects(r.c, (*C.SDL_Rect)(unsafe.Pointer(&rects[0])), C.int(len(rects))) != 0 {
		return getError()
	}

	return nil
}

func (r *Renderer) FillRect(rect *Rect) error {
	if C.SDL_RenderFillRect(r.c, rect.c()) != 0 {
		return getError()
	}

	return nil
}

func (r *Renderer) FillRects(rects []Rect) error {
	if C.SDL_RenderFillRects(r.c, (*C.SDL_Rect)(unsafe.Pointer(&rects[0])), C.int(len(rects))) != 0 {
		return getError()
	}

	return nil
}

func (r *Renderer) Copy(t *Texture, sr *Rect, dr *Rect) error {
	if C.SDL_RenderCopy(r.c, t.c, sr.c(), dr.c()) != 0 {
		return getError()
	}

	return nil
}

func (r *Renderer) CopyEx(t *Texture, sr *Rect, dr *Rect, a float64, c *Point, flip RendererFlip) error {
	en := C.SDL_RenderCopyEx(r.c,
		t.c,
		sr.c(),
		dr.c(),
		C.double(a),
		c.c(),
		flip.c(),
	)
	if en != 0 {
		return getError()
	}

	return nil
}

// TODO: Find a way to do SDL_RenderReadPixels().

//func (r *Renderer) ReadPixels8(rect *Rect, format uint32, pix []uint8, pitch int) error {
//	if rect == nil {
//		_, _, w, h, err := t.Query()
//		if err != nil {
//			return 0, err
//		}
//
//		rect = &Rect{
//			W: int32(w),
//			H: int32(h),
//		}
//	}
//
//	if len(pix) < int(rect.W)*int(rect.H) {
//		return errors.New("len(pix) < rect.W * rect.H")
//	}
//
//	en := C.SDL_RenderReadPixels(r.c,
//		C.Uint32(format),
//		unsafe.Pointer(&pix[0]),
//		C.int(pitch),
//	)
//	if en != 0 {
//		return getError()
//	}
//
//	return nil
//}
//
//func (r *Renderer) ReadPixels16(rect *Rect, format uint32, pix []uint16, pitch int) error {
//	if rect == nil {
//		_, _, w, h, err := t.Query()
//		if err != nil {
//			return 0, err
//		}
//
//		rect = &Rect{
//			W: int32(w),
//			H: int32(h),
//		}
//	}
//
//	if len(pix) < int(rect.W)*int(rect.H) {
//		return errors.New("len(pix) < rect.W * rect.H")
//	}
//
//	en := C.SDL_RenderReadPixels(r.c,
//		C.Uint32(format),
//		unsafe.Pointer(&pix[0]),
//		C.int(pitch),
//	)
//	if en != 0 {
//		return getError()
//	}
//
//	return nil
//}
//
//func (r *Renderer) ReadPixels32(rect *Rect, format uint32, pix []uint32, pitch int) error {
//	if rect == nil {
//		_, _, w, h, err := t.Query()
//		if err != nil {
//			return 0, err
//		}
//
//		rect = &Rect{
//			W: int32(w),
//			H: int32(h),
//		}
//	}
//
//	if len(pix) < int(rect.W)*int(rect.H) {
//		return errors.New("len(pix) < rect.W * rect.H")
//	}
//
//	en := C.SDL_RenderReadPixels(r.c,
//		C.Uint32(format),
//		unsafe.Pointer(&pix[0]),
//		C.int(pitch),
//	)
//	if en != 0 {
//		return getError()
//	}
//
//	return nil
//}
//
//func (r *Renderer) ReadPixels64(rect *Rect, format uint32, pix []uint64, pitch int) error {
//	if rect == nil {
//		_, _, w, h, err := t.Query()
//		if err != nil {
//			return err
//		}
//
//		rect = &Rect{
//			W: int32(w),
//			H: int32(h),
//		}
//	}
//
//	if len(pix) < int(rect.W)*int(rect.H) {
//		return errors.New("len(pix) < rect.W * rect.H")
//	}
//
//	en := C.SDL_RenderReadPixels(r.c,
//		C.Uint32(format),
//		unsafe.Pointer(&pix[0]),
//		C.int(pitch),
//	)
//	if en != 0 {
//		return getError()
//	}
//
//	return nil
//}

func (r *Renderer) Present() {
	C.SDL_RenderPresent(r.c)
}

func (t *Texture) Destroy() {
	C.SDL_DestroyTexture(t.c)
}

func (r *Renderer) Destroy() {
	C.SDL_DestroyRenderer(r.c)
}
