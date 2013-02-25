#include <SDL_audio.h>

Uint32 AUDIO_BITSIZE(Uint32 x)
{
	return SDL_AUDIO_BITSIZE(x);
}

Uint32 AUDIO_ISFLOAT(Uint32 x)
{
	return SDL_AUDIO_ISFLOAT(x);
}

Uint32 AUDIO_ISBIGENDIAN(Uint32 x)
{
	return SDL_AUDIO_ISBIGENDIAN(x);
}

Uint32 AUDIO_ISSIGNED(Uint32 x)
{
	return SDL_AUDIO_ISSIGNED(x);
}

Uint32 AUDIO_ISINT(Uint32 x)
{
	return SDL_AUDIO_ISINT(x);
}

Uint32 AUDIO_ISLITTLEENDIAN(Uint32 x)
{
	return SDL_AUDIO_ISLITTLEENDIAN(x);
}

Uint32 AUDIO_ISUNSIGNED(Uint32 x)
{
	return SDL_AUDIO_ISUNSIGNED(x);
}

SDL_AudioSpec *LoadWAV(const char *file, SDL_AudioSpec *spec, Uint8 **buf, Uint32 *len)
{
	return SDL_LoadWAV(file, spec, buf, len);
}
