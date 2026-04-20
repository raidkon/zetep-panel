package xraytun

import (
	"os"
	"strings"
	"testing"

	"z-panel/internal/i18n"
)

func TestMain(m *testing.M) {
	_ = os.Setenv("Z_PANEL_LANG", "en")
	i18n.Init()
	os.Exit(m.Run())
}

func TestNormalizeIPCIDR(t *testing.T) {
	t.Parallel()
	got, err := normalizeIPCIDR("10.1.2.3/24")
	if err != nil || got != "10.1.2.3/24" {
		t.Fatalf("%v %q", err, got)
	}
	got, err = normalizeIPCIDR("10.1.2.3")
	if err != nil || got != "10.1.2.3/32" {
		t.Fatalf("%v %q", err, got)
	}
	_, err = normalizeIPCIDR("")
	if err == nil {
		t.Fatal()
	}
	_, err = normalizeIPCIDR("2001:db8::1/64")
	if err == nil || !strings.Contains(err.Error(), "IPv4") {
		t.Fatalf("got %v", err)
	}
}

func TestSanitizeIfaceName(t *testing.T) {
	t.Parallel()
	if err := sanitizeIfaceName("tun0"); err != nil {
		t.Fatal(err)
	}
	if err := sanitizeIfaceName(""); err == nil {
		t.Fatal()
	}
	if err := sanitizeIfaceName("bad name"); err == nil {
		t.Fatal()
	}
}
