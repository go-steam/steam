// steam.go
// +build linux,amd64 windows,amd64

package steam

/*
#cgo linux amd64 LDFLAGS: -L$GOPATH/steamworks_sdk_128/redistributable_bin/linux64 -lsteam_api
#cgo linux 386 LDFLAGS: -L$GOPATH/steamworks_sdk_128/redistributable_bin/linux32
#cgo darwin LDFLAGS: -L$GOPATH/steamworks_sdk_128/redistributable_bin/osx32
#cgo windows amd64 LDFLAGS: -L$GOPATH/steamworks_sdk_128/redistributable_bin/win64
#cgo windows 386 LDFLAGS: -L$GOPATH/steamworks_sdk_128/redistributable_bin
#cgo windows CFLAGS: -DIS_WINDOWS=1
#cgo !windows CFLAGS: -DIS_WINDOWS=0 -DSTEAM_API_NODLL
#include "wrapper.h"
*/
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
func GoLogFatal(s *C.char) {
	log.Fatalln(C.GoString(s))
}
