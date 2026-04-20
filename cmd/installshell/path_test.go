package installshell

import (
	"path/filepath"
	"testing"

	"z-panel/internal/i18n"
)

func TestInstallUserPath_xdg(t *testing.T) {
	t.Setenv("HOME", "/home/testuser")
	t.Setenv("XDG_DATA_HOME", "/xdg/data")
	i18n.Init()
	p, err := installUserPath()
	if err != nil {
		t.Fatal(err)
	}
	want := filepath.Join("/xdg/data", "bash-completion", "completions", "z-panel")
	if p != want {
		t.Fatalf("got %q want %q", p, want)
	}
}

func TestInstallUserPath_default(t *testing.T) {
	t.Setenv("HOME", "/home/u2")
	t.Setenv("XDG_DATA_HOME", "")
	i18n.Init()
	p, err := installUserPath()
	if err != nil {
		t.Fatal(err)
	}
	want := filepath.Join("/home/u2", ".local", "share", "bash-completion", "completions", "z-panel")
	if p != want {
		t.Fatalf("got %q want %q", p, want)
	}
}
