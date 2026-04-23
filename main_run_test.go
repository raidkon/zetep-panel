package main

import (
	"os"
	"path/filepath"
	"testing"

	"z-panel/internal/settings"
)

func TestRunMain_version(t *testing.T) {
	t.Setenv("Z_PANEL_SKIP_DAEMON", "1")
	t.Setenv("Z_PANEL_NO_BANNER", "1")
	settings.C = nil
	code := runMain([]string{"z-panel", "version"})
	if code != 0 {
		t.Fatalf("code=%d", code)
	}
}

func TestRunMain_help(t *testing.T) {
	t.Setenv("Z_PANEL_SKIP_DAEMON", "1")
	t.Setenv("Z_PANEL_NO_BANNER", "1")
	settings.C = nil
	code := runMain([]string{"z-panel", "help"})
	if code != 0 {
		t.Fatalf("code=%d", code)
	}
}

func TestRunMain_badArg(t *testing.T) {
	t.Setenv("Z_PANEL_SKIP_DAEMON", "1")
	t.Setenv("Z_PANEL_NO_BANNER", "1")
	settings.C = nil
	code := runMain([]string{"z-panel", "not-a-command-xyz"})
	if code != 1 {
		t.Fatalf("code=%d", code)
	}
}

func TestRunMain_sshConnect_version(t *testing.T) {
	dir := t.TempDir()
	ssh := filepath.Join(dir, "ssh")
	if err := os.WriteFile(ssh, []byte("#!/bin/sh\nexit 0\n"), 0o755); err != nil {
		t.Fatal(err)
	}
	t.Setenv("PATH", dir)
	t.Setenv("Z_PANEL_SKIP_DAEMON", "1")
	t.Setenv("Z_PANEL_NO_BANNER", "1")
	settings.C = nil
	code := runMain([]string{"z-panel", "--ssh-connect=u@h", "version"})
	if code != 0 {
		t.Fatalf("code=%d", code)
	}
}

func TestRunMain_parseError(t *testing.T) {
	t.Setenv("Z_PANEL_SKIP_DAEMON", "1")
	t.Setenv("Z_PANEL_NO_BANNER", "1")
	settings.C = nil
	code := runMain([]string{"z-panel", "--ssh=h1", "--ssh=h2", "x"})
	if code != 2 {
		t.Fatalf("code=%d", code)
	}
}

func TestRunMain_onlyProg(t *testing.T) {
	t.Setenv("Z_PANEL_SKIP_DAEMON", "1")
	t.Setenv("Z_PANEL_NO_BANNER", "1")
	settings.C = nil
	code := runMain([]string{"z-panel"})
	if code != 0 {
		t.Fatalf("code=%d", code)
	}
}

func TestRunMain_sshConnect_helpOnly(t *testing.T) {
	t.Setenv("Z_PANEL_SKIP_DAEMON", "1")
	t.Setenv("Z_PANEL_NO_BANNER", "1")
	settings.C = nil
	code := runMain([]string{"z-panel", "--ssh-connect=user@host"})
	if code != 0 {
		t.Fatalf("code=%d", code)
	}
}

func TestRunMain_sshLocal_install_usesScpSsh(t *testing.T) {
	dir := t.TempDir()
	ssh := filepath.Join(dir, "ssh")
	if err := os.WriteFile(ssh, []byte("#!/bin/sh\nexit 0\n"), 0o755); err != nil {
		t.Fatal(err)
	}
	scp := filepath.Join(dir, "scp")
	if err := os.WriteFile(scp, []byte("#!/bin/sh\nexit 0\n"), 0o755); err != nil {
		t.Fatal(err)
	}
	t.Setenv("PATH", dir+string(os.PathListSeparator)+os.Getenv("PATH"))
	t.Setenv("Z_PANEL_SKIP_DAEMON", "1")
	t.Setenv("Z_PANEL_NO_BANNER", "1")
	settings.C = nil
	code := runMain([]string{"z-panel", "--ssh", "h1", "install"})
	if code != 0 {
		t.Fatalf("code=%d", code)
	}
}

func TestRunMain_sshLocal_config_forbidden(t *testing.T) {
	t.Setenv("Z_PANEL_SKIP_DAEMON", "1")
	t.Setenv("Z_PANEL_NO_BANNER", "1")
	settings.C = nil
	code := runMain([]string{"z-panel", "--ssh", "h1", "config"})
	if code != 1 {
		t.Fatalf("code=%d", code)
	}
}
