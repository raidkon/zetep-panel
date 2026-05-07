package installshell

import (
	"bytes"
	"os"
	"testing"

	"z-panel/internal/i18n"
)

func TestMain(m *testing.M) {
	i18n.Init()
	i18n.ApplyFromConfig("en")
	os.Exit(m.Run())
}

func TestHelp(t *testing.T) {
	var buf bytes.Buffer
	New().Help(&buf)
	if buf.Len() < 20 {
		t.Fatal(buf.String())
	}
}
