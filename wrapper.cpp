// wrapper.cpp

#define STEAM_API_NODLL
#define VERSION_SAFE_STEAM_API_INTERFACES
#include "steam_api.h"

extern "C" {
	#include "wrapper.h"
	#include "_cgo_export.h"
}

CSteamAPIContext myContext;

int SteamInit() {
	if (SteamAPI_InitSafe())
	{		
		myContext.Init();
		return 1;
	}
	else
		GoLogFatal("SteamAPI_InitSafe failed.");
	return 0;
}

int SteamIsSteamRunning() {
	if (SteamAPI_IsSteamRunning())
		return 1;
	else
		return 0;
}

void SteamShutdown() {
	SteamAPI_Shutdown();
}
