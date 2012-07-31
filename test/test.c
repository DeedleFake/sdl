#include <SDL.h>

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

	SDL_Surface *bmp = SDL_LoadBMP("test.bmp");

	SDL_Surface *ws = SDL_GetWindowSurface(win);

	SDL_BlitSurface(ws, NULL, bmp, NULL);

	SDL_UpdateWindowSurface(win);

	SDL_Delay(3000);
}
