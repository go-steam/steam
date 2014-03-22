// steam.go

package steam

// #cgo CFLAGS: -I$GOPATH/steamworks_sdk_128/public/steam
// #cgo LDFLAGS: -L$GOPATH/steamworks_sdk_128/redistributable_bin/linux64 -lsteam_api
// #define STEAM_API_NODLL
// #include "wrapper.h"
import "C"

import (
	"log"
)

// Shutdown calls STEAMAPI_Shutdown()
func Shutdown() {
	_, err := C.SteamShutdown()
	if err != nil {
		log.Fatal(err)
	}
}

// IsSteamRunning returns true if steam client is running.
func IsSteamRunning() bool {
	n, err := C.SteamIsSteamRunning()
	if err != nil {
		log.Fatal(err)
	}

	return n != 0
}

// Init calls STEAMAPI_InitSafe() and CSteamAPIContext.Init()
func Init() bool {
	n, err := C.SteamInit()
	if err != nil {
		log.Fatal(err)
	}

	return n != 0
}

//export GoLogFatal
/*
func GoLogFatal(s C.CString) {
	log.Fatal(C.GoString(s))
}
*/
