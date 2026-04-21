package main

import (
	"os"
	"testing"

	"z-panel/internal/app"
	"z-panel/internal/i18n"
	"z-panel/internal/settings"
)

func TestMain(m *testing.M) {
	_ = os.Setenv("Z_PANEL_LANG", "en")
	i18n.Init()
	if err := settings.Load(); err != nil {
		// /etc/z-panel/config.toml may be unreadable in CI — still set C on success path only
		_ = err
	}
	if settings.C == nil {
		settings.C = &settings.Cfg{UfwMarker: "z-panel", Language: "en"}
	}
	os.Exit(m.Run())
}

func TestFindCommand_allRegistered(t *testing.T) {
	for _, name := range []string{
		"version", "-v", "--version",
		"install", "install-shell", "config", "daemon",
		"xray-redirect", "ufw", "xray-tun",
	} {
		if app.FindCommand(name) == nil {
			t.Fatalf("missing %q", name)
		}
	}
}

func TestFindCommand_unknown(t *testing.T) {
	if app.FindCommand("not-a-real-command") != nil {
		t.Fatal()
	}
}
