package executil

import (
	"os"
	"testing"
)

func TestTryStartSSHMultiplex_noMux(t *testing.T) {
	t.Setenv(EnvSSHNoMux, "1")
	stop, err := TryStartSSHMultiplex("user@host")
	if err != nil || stop != nil {
		t.Fatalf("stop=%p err=%v", stop, err)
	}
}

func TestTryStartSSHMultiplex_emptyHost(t *testing.T) {
	t.Cleanup(func() { _ = os.Unsetenv(EnvSSHNoMux) })
	stop, err := TryStartSSHMultiplex("")
	if stop != nil || err != nil {
		t.Fatal()
	}
}
