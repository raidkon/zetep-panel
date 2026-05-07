package transport

import (
	"os"
	"strings"
	"testing"

	"z-panel/internal/i18n"
)

func TestMain(m *testing.M) {
	i18n.Init()
	i18n.ApplyFromConfig("en")
	os.Exit(m.Run())
}

func TestParseSSHFromArgs_sshTools(t *testing.T) {
	t.Parallel()
	mode, h, rest, err := ParseSSHFromArgs([]string{"/bin/z-panel", "--ssh=h1", "ufw", "check", "tun0"})
	if err != nil || mode != RemoteLocalTools || h != "h1" || len(rest) != 4 || rest[1] != "ufw" {
		t.Fatalf("got mode=%v host=%q rest=%v err=%v", mode, h, rest, err)
	}
	mode, h, rest, err = ParseSSHFromArgs([]string{"/bin/z-panel", "--ssh", "user@h1", "version"})
	if err != nil || mode != RemoteLocalTools || h != "user@h1" || !strings.Contains(strings.Join(rest, " "), "version") {
		t.Fatalf("got mode=%v host=%q rest=%v err=%v", mode, h, rest, err)
	}
	mode, h, rest, err = ParseSSHFromArgs([]string{"/bin/z-panel", "--ssh", "h1", "ufw", "check", "lan"})
	if err != nil || mode != RemoteLocalTools || h != "h1" || len(rest) != 4 || rest[1] != "ufw" || rest[2] != "check" || rest[3] != "lan" {
		t.Fatalf("got mode=%v host=%q rest=%v err=%v", mode, h, rest, err)
	}
}

func TestParseSSHFromArgs_sshConnect(t *testing.T) {
	t.Parallel()
	mode, h, rest, err := ParseSSHFromArgs([]string{"/bin/z-panel", "--ssh-connect=p1", "ufw", "check"})
	if err != nil || mode != RemoteZPanelBinary || h != "p1" || len(rest) != 3 || rest[1] != "ufw" {
		t.Fatalf("got mode=%v host=%q rest=%v err=%v", mode, h, rest, err)
	}
	mode, h, rest, err = ParseSSHFromArgs([]string{"/bin/z-panel", "--ssh-connect", "user@p2", "daemon", "run"})
	if err != nil || mode != RemoteZPanelBinary || h != "user@p2" || rest[1] != "daemon" {
		t.Fatalf("got mode=%v host=%q rest=%v err=%v", mode, h, rest, err)
	}
}

func TestParseSSHFromArgs_conflict(t *testing.T) {
	t.Parallel()
	_, _, _, err := ParseSSHFromArgs([]string{"/bin/z-panel", "--ssh=h1", "--ssh=h2", "x"})
	if err == nil {
		t.Fatal("expected conflict error")
	}
	_, _, _, err = ParseSSHFromArgs([]string{"/bin/z-panel", "--ssh=h1", "--ssh-connect=h2", "x"})
	if err == nil {
		t.Fatal("expected conflict error")
	}
	_, _, _, err = ParseSSHFromArgs([]string{"/bin/z-panel", "--ssh-connect=h1", "--ssh=h2", "x"})
	if err == nil {
		t.Fatal("expected conflict error")
	}
	if err != nil && !strings.Contains(err.Error(), "ssh") {
		t.Fatal(err)
	}
}
