package executil

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestShellSingleQuote_edgeCases(t *testing.T) {
	if shellSingleQuote("noquote") != "'noquote'" {
		t.Fatal(shellSingleQuote("noquote"))
	}
	q := shellSingleQuote("a'b")
	if !strings.Contains(q, "'") {
		t.Fatal(q)
	}
}

func TestCommand_local(t *testing.T) {
	t.Cleanup(func() { _ = os.Unsetenv(EnvSSHHost) })
	_ = os.Unsetenv(EnvSSHHost)
	c := Command("echo", "x")
	if filepath.Base(c.Args[0]) != "echo" {
		t.Fatalf("args %v", c.Args)
	}
}

func TestCommand_remote(t *testing.T) {
	t.Cleanup(func() { _ = os.Unsetenv(EnvSSHHost) })
	t.Setenv(EnvSSHHost, "user@host")
	c := Command("true", "a", "b")
	args := strings.Join(c.Args, " ")
	if !strings.Contains(args, "ssh") || !strings.Contains(args, "user@host") || !strings.Contains(args, "sudo") {
		t.Fatalf("args=%v", c.Args)
	}
}

func TestCommand_remote_mux(t *testing.T) {
	t.Cleanup(func() {
		_ = os.Unsetenv(EnvSSHHost)
		_ = os.Unsetenv(EnvSSHMux)
	})
	t.Setenv(EnvSSHHost, "h1")
	t.Setenv(EnvSSHMux, "/tmp/zpanel-mux-test.sock")
	c := Command("ls")
	found := false
	for _, a := range c.Args {
		if a == "-S" {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("args=%v", c.Args)
	}
}

func TestCommandTTY_remote_ttyFlag(t *testing.T) {
	t.Cleanup(func() { _ = os.Unsetenv(EnvSSHHost) })
	t.Setenv(EnvSSHHost, "h")
	c := CommandTTY("ls")
	if c.Args[0] != "ssh" {
		t.Fatalf("%v", c.Args)
	}
	hasT := false
	for _, a := range c.Args {
		if a == "-t" {
			hasT = true
			break
		}
	}
	if !hasT {
		t.Fatalf("expected -t in %v", c.Args)
	}
}

func TestCommandTTY_remote_noTTY(t *testing.T) {
	t.Cleanup(func() {
		_ = os.Unsetenv(EnvSSHHost)
		_ = os.Unsetenv(EnvSSHNoTTY)
	})
	t.Setenv(EnvSSHHost, "h")
	t.Setenv(EnvSSHNoTTY, "1")
	c := CommandTTY("ls")
	for _, a := range c.Args {
		if a == "-t" {
			t.Fatal("unexpected -t", c.Args)
		}
	}
}

func TestRunTTYCombined_local(t *testing.T) {
	t.Cleanup(func() { _ = os.Unsetenv(EnvSSHHost) })
	out, err := RunTTYCombined("echo", "ok")
	if err != nil || string(out) != "ok\n" {
		t.Fatalf("%q %v", out, err)
	}
}

func TestRunTTYCombinedScript_local(t *testing.T) {
	t.Cleanup(func() { _ = os.Unsetenv(EnvSSHHost) })
	out, err := RunTTYCombinedScript("echo hello; echo more")
	if err != nil || !strings.Contains(string(out), "hello") {
		t.Fatalf("%q %v", out, err)
	}
}

func TestRunTTYCombinedScript_remote_quotedLine(t *testing.T) {
	t.Cleanup(func() {
		_ = os.Unsetenv(EnvSSHHost)
		_ = os.Unsetenv(EnvSSHNoTTY)
	})
	dir := t.TempDir()
	sshPath := filepath.Join(dir, "ssh")
	content := "#!/bin/sh\nexit 0\n"
	if err := os.WriteFile(sshPath, []byte(content), 0o755); err != nil {
		t.Fatal(err)
	}
	t.Setenv("PATH", dir)
	t.Setenv(EnvSSHHost, "u@remote")
	out, err := RunTTYCombinedScript("true")
	if err != nil {
		t.Fatal(err)
	}
	_ = out
}
