package state

import (
	"path/filepath"
	"strings"
	"testing"

	"z-panel/internal/settings"
)

func TestPath_sanitizesIface(t *testing.T) {
	t.Cleanup(func() { settings.C = nil })
	settings.C = &settings.Cfg{StateDir: "/var/lib/z-panel/state"}

	p := Path("tun0")
	if want := filepath.Join("/var/lib/z-panel/state", "tun0.json"); p != want {
		t.Fatalf("got %q want %q", p, want)
	}
	p = Path("../evil")
	if !strings.Contains(p, "_evil") || !strings.HasSuffix(p, ".json") {
		t.Fatalf("got %q", p)
	}
}

func TestPartial(t *testing.T) {
	f := Partial("eth0", "51843", "wg", true)
	if f.Interface != "eth0" || f.Table != "51843" || f.Mode != "wg" || !f.WGIPv6 {
		t.Fatalf("%+v", f)
	}
}
