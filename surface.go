package sdl

// #cgo pkg-config: sdl2
//
// #include <SDL.h>
//
// static SDL_Surface *LoadBMP(const char *file) { return SDL_LoadBMP(file); }
// static int SaveBMP(SDL_Surface *s, const char *file) { return SDL_SaveBMP(s, file); }
//
// static int BlitSurface(SDL_Surface *src, const SDL_Rect *sr, SDL_Surface *dst, SDL_Rect *dr) { return SDL_BlitSurface(src, sr, dst, dr); }
// static int BlitScaled(SDL_Surface *src, const SDL_Rect *sr, SDL_Surface *dst, SDL_Rect *dr) { return SDL_BlitScaled(src, sr, dst, dr); }
import "C"

const (
	SWSURFACE = 0
	PREALLOC  = 0x00000001
	RLEACCEL  = 0x00000002
	DONTFREE  = 0x00000004
)

func (s *Surface) MustLock() bool {
	return (s.Flags & RLEACCEL) != 0
}

type Surface struct {
	Flags  uint32
	Format *PixelFormat
	W, H   int32
	Pitch  int32
	Pixels uintptr

	Userdata uintptr

	Locked   int32
	LockData uintptr

	blitmap *BlitMap

	Refcount int32
}

func (s *Surface) c() *C.SDL_Surface {
	return (*C.SDL_Surface)(unsafe.Pointer(s))
}

func CreateRGBSurface(w, h, d int, rm, gm, bm, am uint32) (*Surface, error) {
	s := C.SDL_CreateRGBSurface(0,
		C.int(w),
		C.int(h),
		C.int(d),
		C.Uint32(rm),
		C.Uint32(gm),
		C.Uint32(bm),
		C.Uint32(am),
	)
	if s == nil {
		return nil, getError()
	}

	return (*Surface)(unsafe.Pointer(s)), nil
}

func CreateRGBSurfaceFrom(pix uintptr, w, h, d, p int, rm, gm, bm, am uint32) (*Surface, error) {
	s := C.SDL_CreateRGBSurfaceFrom(
		unsafe.Pointer(pix),
		C.int(w),
		C.int(h),
		C.int(d),
		C.int(p),
		C.Uint32(rm),
		C.Uint32(gm),
		C.Uint32(bm),
		C.Uint32(am),
	)
	if s == nil {
		return nil, getError()
	}

	return (*Surface)(unsafe.Pointer(s)), nil
}

func (s *Surface) Free() {
	C.SDL_FreeSurface(s.c())
}

func (s *Surface) SetPalette(pal *Palette) error {
	if C.SDL_SetSurfacePalette(s.c(), pal.c()) != 0 {
		return getError()
	}

	return nil
}

func (s *Surface) Lock() error {
	if C.SDL_LockSurface(s.c()) != 0 {
		return getError()
	}

	return nil
}

func (s *Surface) Unlock() {
	C.SDL_UnlockSurface(s.c())
}

func LoadBMP_RW(rw *RWops, free bool) (*Surface, error) {
	var cfree C.int
	if free {
		cfree = 1
	}

	s := C.SDL_LoadBMP_RW(rw.c(), cfree)
	if s == nil {
		return nil, getError()
	}

	return nil
}

func LoadBMP(file string) (*Surface, error) {
	cfile := C.CString(file)
	defer C.free(unsafe.Pointer(cfile))

	s := C.LoadBMP(cfile)
	if s == nil {
		return nil, getError()
	}

	return nil
}

func (s *Surface) SaveBMP_RW(rw *RWops, free bool) error {
	var cfree C.int
	if free {
		cfree = 1
	}

	if C.SDL_SaveBMP_RW(s.c(), rw.c(), cfree) != 0 {
		return getError()
	}

	return nil
}

func (s *Surface) SaveBMP(file string) error {
	cfile := C.CString(file)
	defer C.free(unsafe.Pointer(cfile))

	if C.SaveBMP(s.c(), cfile) != 0 {
		return getError()
	}

	return nil
}

func (s *Surface) SetSurfaceRLE(flag bool) error {
	var cflag C.int
	if flag {
		cflag = 1
	}

	if C.SDL_SetSurfaceRLE(s.c(), cflag) != 0 {
		return getError()
	}

	return nil
}

func (s *Surface) SetColorKey(flag bool, key uint32) error {
	var cflag C.int
	if flag {
		cflag = 1
	}

	if C.SDL_SetColorKey(s.c(), cflag, C.Uint32(key)) != 0 {
		return getError()
	}

	return nil
}

func (s *Surface) GetColorKey() (uint32, error) {
	var key C.Uint32
	if C.SDL_GetColorKey(s.c(), &key) != 0 {
		return 0, getError()
	}

	return key, nil
}

func (s *Surface) SetColorMod(r, g, b uint8) error {
	if C.SDL_SetSurfaceColorMod(s.c(), C.Uint8(r), C.Uint8(g), C.Uint8(b)) != 0 {
		return getError()
	}

	return nil
}

func (s *Surface) GetColorMod() (r, g, b uint8, err error) {
	var cr, cg, cb C.Uint8
	if C.SDL_GetSurfaceColorMod(s.c, &cr, &cg, &cb) != 0 {
		return 0, 0, 0, getError()
	}

	return uint8(cr), uint8(cg), uint8(cb), nil
}

func (s *Surface) SetAlphaMod(a uint8) error {
	if C.SDL_SetSurfaceAlphaMod(s.c(), C.Uint8(a)) != 0 {
		return getError()
	}

	return nil
}

func (s *Surface) GetAlphaMod() (uint8, error) {
	var a C.Uint8
	if C.SDL_GetSurfaceAlphaMod(s.c(), &a) != 0 {
		return 0, getError()
	}

	return uint8(a), nil
}

func (s *Surface) SetBlendMode(m BlendMode) error {
	if C.SDL_SetSurfaceBlendMode(s.c(), m.c()) != 0 {
		return getError()
	}

	return nil
}

func (s *Surface) GetBlendMode() (BlendMode, error) {
	var m C.SDL_BlendMode
	if C.SDL_GetSurfaceBlendMode(s.c(), &m) != 0 {
		return 0, getError()
	}

	return BlendMode(m), nil
}

func (s *Surface) SetClipRect(r *Rect) bool {
	return C.SDL_SetClipRect(s.c(), r.c()) == C.SDL_TRUE
}

func (s *Surface) GetClipRect() *Rect {
	var rect Rect
	C.SDL_GetClipRect(s.c(), rect.c())
}

func (s *Surface) Convert(f *PixelFormat, flags uint32) (*Surface, error) {
	cs := C.SDL_ConvertSurface(s.c(), f.c(), C.Uint32(flags))
	if cs == nil {
		return nil, getError()
	}

	return (*Surface)(unsafe.Pointer(cs)), nil
}

func (s *Surface) ConvertFormat(pf uint32, flags uint32) (*Surface, error) {
	cs := C.SDL_ConvertSurfaceFormat(s.c(), C.Uint32(pf), C.Uint32(flags))
	if cs == nil {
		return nil, getError()
	}

	return (*Surface)(unsafe.Pointer(cs)), nil
}

func ConvertPixels(w, h int, sf uint32, src uintptr, sp int, df uint32, dst uintptr, dp int) error {
	en := C.SDL_ConvertPixels(
		C.int(w),
		C.int(h),
		C.Uint32(sf),
		unsafe.Pointer(src),
		C.int(sp),
		C.Uint32(df),
		unsafe.Pointer(dst),
		C.int(dp),
	)
	if en != 0 {
		return getError()
	}

	return nil
}

func (s *Surface) FillRect(r *Rect, c uint32) error {
	if C.SDL_FillRect(s.c(), r.c(), C.Uint32(c)) != 0 {
		return getError()
	}

	return nil
}

func (s *Surface) FillRects(r []Rect, c uint32) error {
	if C.SDL_FillRects(s.c(), (*C.SDL_Rect)(unsafe.Pointer(&r[0])), C.int(len(r)), C.Uint32(c)) != 0 {
		return getError()
	}

	return nil
}

func (s *Surface) Blit(sr *Rect, dst *Surface, dr *Rect) error {
	if C.BlitSurface(s.c(), sr.c(), dst.c(), dr.c()) != 0 {
		return getError()
	}

	return nil
}

func (s *Surface) UpperBlit(sr *Rect, dst *Surface, dr *Rect) error {
	if C.SDL_UpperBlit(s.c(), sr.c(), dst.c(), dr.c()) != 0 {
		return getError()
	}

	return nil
}

func (s *Surface) LowerBlit(sr *Rect, dst *Surface, dr *Rect) error {
	if C.SDL_LowerBlit(s.c(), sr.c(), dst.c(), dr.c()) != 0 {
		return getError()
	}

	return nil
}

func (s *Surface) SoftStretch(sr *Rect, dst *Surface, dr *Rect) error {
	if C.SDL_SoftStretch(s.c(), sr.c(), dst.c(), dr.c()) != 0 {
		return getError()
	}

	return nil
}

func (s *Surface) BlitScaled(sr *Rect, dst *Surface, dr *Rect) error {
	if C.BlitScaled(s.c(), sr.c(), dst.c(), dr.c()) != 0 {
		return getError()
	}

	return nil
}

func (s *Surface) UpperBlitScaled(sr *Rect, dst *Surface, dr *Rect) error {
	if C.SDL_UpperBlitScaled(s.c(), sr.c(), dst.c(), dr.c()) != 0 {
		return getError()
	}

	return nil
}

func (s *Surface) LowerBlitScaled(sr *Rect, dst *Surface, dr *Rect) error {
	if C.SDL_LowerBlitScaled(s.c(), sr.c(), dst.c(), dr.c()) != 0 {
		return getError()
	}

	return nil
}
