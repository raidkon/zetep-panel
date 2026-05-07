package version

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"z-panel/internal/config"
	"z-panel/internal/i18n"
)

func TestMain(m *testing.M) {
	i18n.Init()
	i18n.ApplyFromConfig("en")
	os.Exit(m.Run())
}

func TestHelp_containsVersion(t *testing.T) {
	var buf bytes.Buffer
	New().Help(&buf)
	if !strings.Contains(buf.String(), config.Version) {
		t.Fatalf("%q", buf.String())
	}
}

func TestRun_help(t *testing.T) {
	if err := New().Run([]string{"help"}); err != nil {
		t.Fatal(err)
	}
}

func TestAliases(t *testing.T) {
	a := New().Aliases()
	if len(a) < 2 {
		t.Fatal(a)
	}
}
