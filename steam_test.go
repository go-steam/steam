// steam_test.go

package steam

import (
	"testing"
)

func TestIsSteamRunning(t *testing.T) {
	b := IsSteamRunning()
	if b != true {
		t.Errorf("Steam is not running!")
	}
}

func TestInit(t *testing.T) {
	b := Init()
	if b != true {
		t.Errorf("Init Failure")
	}
}
