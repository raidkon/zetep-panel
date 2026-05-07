package root

import (
	"os"
	"testing"

	"z-panel/internal/executil"
)

func TestRequire_dependsOnUID(t *testing.T) {
	t.Cleanup(executil.ResetSSHForTests)
	err := Require()
	if os.Geteuid() == 0 {
		if err != nil {
			t.Fatalf("as root: %v", err)
		}
	} else {
		if err == nil {
			t.Fatal("non-root: expected error")
		}
	}
}

func TestRequire_skipsRootCheckUnderSSH(t *testing.T) {
	t.Cleanup(executil.ResetSSHForTests)
	executil.SetRemoteSSHHost("h")
	if err := Require(); err != nil {
		t.Fatal(err)
	}
}
