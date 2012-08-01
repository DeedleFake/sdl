// +build ignore

#include <SDL.h>

SDL_Texture *LoadTexture(SDL_Renderer *ren, const char *file)
{
	SDL_Surface *bmp = SDL_LoadBMP("test.bmp");

	SDL_Texture *tex = SDL_CreateTextureFromSurface(ren, bmp);

	SDL_FreeSurface(bmp);

	return tex;
}

int main(int argc, char *argv[])
{
	SDL_Init(SDL_INIT_EVERYTHING);

	SDL_Window *win = SDL_CreateWindow(
			"Test",
			SDL_WINDOWPOS_UNDEFINED,
			SDL_WINDOWPOS_UNDEFINED,
			640,
			480,
			SDL_WINDOW_SHOWN
	);

	SDL_Renderer *ren = SDL_CreateRenderer(win, -1, SDL_RENDERER_ACCELERATED);

	SDL_Texture *bmp = LoadTexture(ren, "test.bmp");

	SDL_RenderClear(ren);
	SDL_RenderCopy(ren, bmp, NULL, NULL);
	SDL_RenderPresent(ren);

	SDL_Delay(3000);

	return 0;
}
