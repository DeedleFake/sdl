#include "version.h"

void VERSION(SDL_version *v)
{
	SDL_VERSION(v);
}

Uint16 VERSIONNUM(Uint8 x, Uint8 y, Uint8 z)
{
	return SDL_VERSIONNUM(x, y, z);
}

int VERSION_ATLEAST(Uint8 x, Uint8 y, Uint8 z)
{
	return SDL_VERSION_ATLEAST(x, y, z);
}
