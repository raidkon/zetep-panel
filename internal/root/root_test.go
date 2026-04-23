package root

import (
	"os"
	"testing"

	"z-panel/internal/executil"
)

func TestRequire_dependsOnUID(t *testing.T) {
	t.Cleanup(func() { _ = os.Unsetenv(executil.EnvSSHHost) })
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
	t.Cleanup(func() { _ = os.Unsetenv(executil.EnvSSHHost) })
	t.Setenv(executil.EnvSSHHost, "h")
	if err := Require(); err != nil {
		t.Fatal(err)
	}
}
