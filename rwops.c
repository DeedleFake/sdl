#include "rwops.h"

#include "_cgo_export.h"

long sdl2_seek(SDL_RWops *ctx, long offset, int whence)
{
	return SDL_RWseek(ctx, offset, whence);
}

long sdl2_tell(SDL_RWops *ctx)
{
	return SDL_RWtell(ctx);
}

size_t sdl2_read(SDL_RWops *ctx, void *ptr, size_t s, size_t max)
{
	return SDL_RWread(ctx, ptr, s, max);
}

size_t sdl2_write(SDL_RWops *ctx, const void *ptr, size_t size, size_t num)
{
	return SDL_RWwrite(ctx, ptr, size, num);
}

int sdl2_close(SDL_RWops *ctx)
{
	return SDL_RWclose(ctx);
}

const char *sdl_cstring(GoString str)
{
	char *n = malloc(str.n);
	memcpy(n, str.p, str.n);
	return n;
}

long seekRWReadSeeker(SDL_RWops *ctx, long off, int wh)
{
	struct seekReadSeeker_return r = seekReadSeeker(ctx->hidden.unknown.data1, off, wh);
	if (r.r0 < 0)
	{
		SDL_SetError(sdl_cstring(r.r1));
		return -1;
	}

	return r.r0;
}

size_t readRWReadSeeker(SDL_RWops *ctx, void *ptr, size_t size, size_t max)
{
	struct readReadSeeker_return r = readReadSeeker(ctx->hidden.unknown.data1, ptr, size, max);
	if (r.r0 == 0)
	{
		return 0;
	}
	else if (r.r0 < 0)
	{
		SDL_SetError(sdl_cstring(r.r1));
		return r.r0;
	}

	return r.r0;
}

size_t writeRWReadSeeker(SDL_RWops *ctx, const void *data, size_t size, size_t num)
{
	SDL_SetError("Can't write to RWops created from io.ReadSeeker");
	return -1;
}

int closeRWReadSeeker(SDL_RWops *ctx)
{
	if (ctx)
	{
		SDL_FreeRW(ctx);
	}

	return 0;
}

SDL_RWops *RWFromReadSeeker(void *r)
{
	SDL_RWops *rw = SDL_AllocRW();
	rw->seek = seekRWReadSeeker;
	rw->read = readRWReadSeeker;
	rw->write = writeRWReadSeeker;
	rw->close = closeRWReadSeeker;
	rw->hidden.unknown.data1 = r;

	return rw;
}
