#include <SDL_events.h>

void SetEventFilter(void *data);
void AddEventWatch(void *data);
void FilterEvents(void *data);

Uint8 GetEventState(Uint32 type);
