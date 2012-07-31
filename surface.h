#include <SDL_surface.h>
 
SDL_Surface *LoadBMP(const char *file);
int SaveBMP(SDL_Surface *s, const char *file);
int BlitSurface(SDL_Surface *src, const SDL_Rect *sr, SDL_Surface *dst, SDL_Rect *dr);
int BlitScaled(SDL_Surface *src, const SDL_Rect *sr, SDL_Surface *dst, SDL_Rect *dr);
