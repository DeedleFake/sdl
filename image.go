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

func (r *Renderer) LoadTexture(name string) (*Texture, error) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	s := C.IMG_LoadTexture(r.c, cname)
	if s == nil {
		return nil, getError()
	}
	return &Texture{c: s}, nil
}
