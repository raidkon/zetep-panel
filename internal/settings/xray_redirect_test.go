package settings

import (
	"errors"
	"fmt"
	"testing"

	"z-panel/internal/state"
)

func TestSanitizeIfaceKey(t *testing.T) {
	if sanitizeIfaceKey("tun0") != "tun0" {
		t.Fatal()
	}
	if sanitizeIfaceKey("../x") != "___x" {
		t.Fatalf("got %q", sanitizeIfaceKey("../x"))
	}
}

func TestErrNoXrayRedirectState_wrap(t *testing.T) {
	e := fmt.Errorf("iface tun0: %w", ErrNoXrayRedirectState)
	if !errors.Is(e, ErrNoXrayRedirectState) {
		t.Fatal()
	}
	_ = state.File{}
}
