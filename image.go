package sdl

// #include <SDL_image.h>
import "C"
import "unsafe"

func Load(name string) (*Surface, error) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	s := C.IMG_Load(cname)
	if s == nil {
		return nil, getError()
	}
	return (*Surface)(unsafe.Pointer(s)), nil
}

func LoadRW(rw *RWops, freesrc bool) (*Surface, error) {
	cfreesrc := C.int(0)
	if freesrc {
		cfreesrc = 1
	}

	s := C.IMG_Load_RW(rw.c, cfreesrc)
	if s == nil {
		return nil, getError()
	}
	return (*Surface)(unsafe.Pointer(s)), nil
}


func (r *Renderer) LoadTexture(name string) (*Texture, error) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	s := C.IMG_LoadTexture(r.c, cname)
	if s == nil {
		return nil, getError()
	}
	return &Texture{c: s}, nil
}

func (r *Renderer) LoadTextureRW(rw *RWops, freesrc bool) (*Texture, error) {
	cfreesrc := C.int(0)
	if freesrc {
		cfreesrc = 1
	}

	s := C.IMG_LoadTexture_RW(r.c, rw.c, cfreesrc)
	if s == nil {
		return nil, getError()
	}
	return &Texture{c: s}, nil
}
