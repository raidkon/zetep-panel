package transport

import (
	"os"
	"path/filepath"
	"testing"

	"z-panel/internal/executil"
	"z-panel/internal/i18n"
)

func TestRunZPanelOverSSH_fakeSSH(t *testing.T) {
	dir := t.TempDir()
	ssh := filepath.Join(dir, "ssh")
	if err := os.WriteFile(ssh, []byte("#!/bin/sh\nexit 0\n"), 0o755); err != nil {
		t.Fatal(err)
	}
	t.Setenv("PATH", dir)
	i18n.Init()
	i18n.ApplyFromConfig("en")
	err := RunZPanelOverSSH("u@h", []string{"/usr/local/bin/z-panel", "version"})
	if err != nil {
		t.Fatal(err)
	}
}

func TestRunZPanelOverSSH_errNoCmd(t *testing.T) {
	i18n.Init()
	i18n.ApplyFromConfig("en")
	err := RunZPanelOverSSH("h", []string{"prog"})
	if err == nil {
		t.Fatal()
	}
}

func TestRunZPanelOverSSH_usesRunAttached(t *testing.T) {
	dir := t.TempDir()
	ssh := filepath.Join(dir, "ssh")
	if err := os.WriteFile(ssh, []byte("#!/bin/sh\nexit 7\n"), 0o755); err != nil {
		t.Fatal(err)
	}
	t.Setenv("PATH", dir)
	i18n.Init()
	i18n.ApplyFromConfig("en")
	err := RunZPanelOverSSH("h", []string{"/bin/z", "x"})
	code, ok := executil.ExitCode(err)
	if !ok || code != 7 {
		t.Fatalf("err=%v code=%d ok=%v", err, code, ok)
	}
}
