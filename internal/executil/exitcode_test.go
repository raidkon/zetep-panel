package executil

import (
	"fmt"
	"os/exec"
	"testing"
)

func TestExitCode_exitError(t *testing.T) {
	cmd := exec.Command("false")
	err := cmd.Run()
	code, ok := ExitCode(err)
	if !ok || code != 1 {
		t.Fatalf("got code=%d ok=%v err=%v", code, ok, err)
	}
}

func TestExitCode_wrappedExitError(t *testing.T) {
	cmd := exec.Command("false")
	inner := cmd.Run()
	err := fmt.Errorf("wrap: %w", inner)
	code, ok := ExitCode(err)
	if !ok || code != 1 {
		t.Fatalf("code=%d ok=%v", code, ok)
	}
}

func TestExitCode_notExit(t *testing.T) {
	_, ok := ExitCode(fmt.Errorf("plain"))
	if ok {
		t.Fatal()
	}
}
