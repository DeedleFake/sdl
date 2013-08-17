#include <SDL_rwops.h>

long sdl2_seek(SDL_RWops *ctx, long offset, int whence);
long sdl2_tell(SDL_RWops *ctx);
size_t sdl2_read(SDL_RWops *ctx, void *ptr, size_t s, size_t max);
size_t sdl2_write(SDL_RWops *ctx, const void *ptr, size_t size, size_t num);
int sdl2_close(SDL_RWops *ctx);

SDL_RWops *RWFromReadSeeker(void *r);
