package install

import (
	"bytes"
	"os"
	"testing"

	"z-panel/internal/i18n"
)

func TestMain(m *testing.M) {
	_ = os.Setenv("Z_PANEL_LANG", "en")
	i18n.Init()
	os.Exit(m.Run())
}

func TestHelp_containsPaths(t *testing.T) {
	var buf bytes.Buffer
	New().Help(&buf)
	s := buf.String()
	if len(s) < 30 {
		t.Fatal(s)
	}
}
