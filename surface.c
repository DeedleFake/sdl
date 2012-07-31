#include "surface.h"

SDL_Surface *LoadBMP(const char *file)
{
	return SDL_LoadBMP(file);
}

int SaveBMP(SDL_Surface *s, const char *file)
{
	return SDL_SaveBMP(s, file);
}

int BlitSurface(SDL_Surface *src, const SDL_Rect *sr, SDL_Surface *dst, SDL_Rect *dr)
{
	return SDL_BlitSurface(src, sr, dst, dr);
}

int BlitScaled(SDL_Surface *src, const SDL_Rect *sr, SDL_Surface *dst, SDL_Rect *dr)
{
	return SDL_BlitScaled(src, sr, dst, dr);
}
