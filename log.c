#include "log.h"

#include "_cgo_export.h"

void Log(const char *str)
{
	SDL_Log(str);
}

void LogVerbose(int cat, const char *str)
{
	SDL_LogVerbose(cat, str);
}

void LogDebug(int cat, const char *str)
{
	SDL_LogDebug(cat, str);
}

void LogInfo(int category, const char *str)
{
	SDL_LogInfo(category, str);
}

void LogWarn(int category, const char *str)
{
	SDL_LogWarn(category, str);
}

void LogError(int category, const char *str)
{
	SDL_LogError(category, str);
}

void LogCritical(int category, const char *str)
{
	SDL_LogCritical(category, str);
}

void LogMessage(int cat, SDL_LogPriority pri, const char *str)
{
	SDL_LogMessage(cat, pri, str);
}

void LogSetOutputFunction(void *data)
{
	SDL_LogSetOutputFunction((SDL_LogOutputFunction)logOutputFunction, data);
}
