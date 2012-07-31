#include <SDL_rwops.h>

long seek(SDL_RWops *ctx, long offset, int whence);
long tell(SDL_RWops *ctx);
size_t read(SDL_RWops *ctx, void *ptr, size_t s, size_t max);
size_t write(SDL_RWops *ctx, const void *ptr, size_t size, size_t num);
int close(SDL_RWops *ctx);

SDL_RWops *RWFromReadSeeker(void *r);
