#include <SDL_audio.h>

Uint32 AUDIO_BITSIZE(Uint32 x);
Uint32 AUDIO_ISFLOAT(Uint32 x);
Uint32 AUDIO_ISBIGENDIAN(Uint32 x);
Uint32 AUDIO_ISSIGNED(Uint32 x);
Uint32 AUDIO_ISINT(Uint32 x);
Uint32 AUDIO_ISLITTLEENDIAN(Uint32 x);
Uint32 AUDIO_ISUNSIGNED(Uint32 x);

SDL_AudioSpec *LoadWAV(const char *file, SDL_AudioSpec *spec, Uint8 **buf, Uint32 *len);
