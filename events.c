#include "events.h"

#include "_cgo_export.h"

void SetEventFilter(void *data)
{
	SDL_SetEventFilter(eventFilter, data);
}

void AddEventWatch(void *data)
{
	SDL_AddEventWatch(eventFilter, data);
}

void FilterEvents(void *data)
{
	SDL_FilterEvents(eventFilter, data);
}

Uint8 GetEventState(Uint32 type)
{
	return SDL_GetEventState(type);
}
