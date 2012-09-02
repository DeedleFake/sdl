#include <SDL_log.h>

void Log(const char *str);
void LogVerbose(int cat, const char *str);
void LogDebug(int cat, const char *str);
void LogInfo(int category, const char *str);
void LogWarn(int category, const char *str);
void LogError(int category, const char *str);
void LogCritical(int category, const char *str);
void LogMessage(int cat, SDL_LogPriority pri, const char *str);

void LogSetOutputFunction(void *data);
