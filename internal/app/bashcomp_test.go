package app

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"z-panel/internal/i18n"
	"z-panel/internal/settings"
)

func TestWriteBashCompletionScript_smoke(t *testing.T) {
	_ = os.Setenv("Z_PANEL_LANG", "en")
	i18n.Init()
	settings.C = &settings.Cfg{UfwMarker: "z-panel"}
	t.Cleanup(func() { registry = nil })
	Register(stubCmd{name: "demo"})
	var buf bytes.Buffer
	WriteBashCompletionScript(&buf)
	out := buf.String()
	if len(out) < 200 || !strings.Contains(out, "complete -F") {
		t.Fatalf("short or invalid script len=%d", len(out))
	}
}
