package redirect

import (
	"testing"
)

func TestNormalizeSystemdUnit(t *testing.T) {
	t.Parallel()
	if got := normalizeSystemdUnit("xray"); got != "xray.service" {
		t.Fatalf("got %q", got)
	}
	if got := normalizeSystemdUnit("xray.service"); got != "xray.service" {
		t.Fatalf("got %q", got)
	}
	if got := normalizeSystemdUnit("  sing-box.service "); got != "sing-box.service" {
		t.Fatalf("got %q", got)
	}
}

func TestResolveBypassCgroupFromFlags_explicit(t *testing.T) {
	path, unit, err := resolveBypassCgroupFromFlags("/slice/foo", "auto")
	if err != nil {
		t.Fatal(err)
	}
	if unit != "" {
		t.Fatalf("unit=%q want empty", unit)
	}
	if path != "slice/foo" {
		t.Fatalf("path=%q", path)
	}
}

func TestResolveBypassCgroupFromFlags_explicitNoLeadingSlash(t *testing.T) {
	path, _, err := resolveBypassCgroupFromFlags("user.slice/app", "anything")
	if err != nil {
		t.Fatal(err)
	}
	if path != "user.slice/app" {
		t.Fatalf("path=%q", path)
	}
}
