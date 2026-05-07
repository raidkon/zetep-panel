package executil

import "testing"

func TestTryStartSSHMultiplex_noMux(t *testing.T) {
	t.Cleanup(ResetSSHForTests)
	stop, err := TryStartSSHMultiplex("user@host", true)
	if err != nil || stop != nil {
		t.Fatalf("stop=%p err=%v", stop, err)
	}
}

func TestTryStartSSHMultiplex_emptyHost(t *testing.T) {
	t.Cleanup(ResetSSHForTests)
	stop, err := TryStartSSHMultiplex("", false)
	if stop != nil || err != nil {
		t.Fatal()
	}
}
