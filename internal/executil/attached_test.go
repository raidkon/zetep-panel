package executil

import (
	"os/exec"
	"testing"

	"z-panel/internal/i18n"
)

func TestRunAttachedInterruptible_trueExits(t *testing.T) {
	i18n.Init()
	i18n.ApplyFromConfig("en")
	cmd := exec.Command("true")
	if err := RunAttachedInterruptible(cmd); err != nil {
		t.Fatal(err)
	}
}

func TestRunAttachedInterruptible_startFail(t *testing.T) {
	cmd := exec.Command("/nonexistent/binary/z-panel-test-xyz")
	err := RunAttachedInterruptible(cmd)
	if err == nil {
		t.Fatal("expected error")
	}
}
