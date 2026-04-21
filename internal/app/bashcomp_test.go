package app

import (
	"bytes"
	"os"
	"os/exec"
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
	if _, err := exec.LookPath("bash"); err != nil {
		t.Skip("bash not in PATH")
	}
	cmd := exec.Command("bash", "-n", "/dev/stdin")
	cmd.Stdin = &buf
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("bash -n: %v\n%s", err, out)
	}
}
