package confcmd

import (
	"os"
	"testing"

	"z-panel/internal/i18n"
	"z-panel/internal/settings"
)

func TestMain(m *testing.M) {
	_ = os.Setenv("Z_PANEL_LANG", "en")
	i18n.Init()
	settings.C = &settings.Cfg{Language: "en"}
	os.Exit(m.Run())
}

func TestRun_emptyPrintsHelp(t *testing.T) {
	c := New()
	if err := c.Run([]string{}); err != nil {
		t.Fatal(err)
	}
}

func TestRun_helpFlag(t *testing.T) {
	c := New()
	if err := c.Run([]string{"help"}); err != nil {
		t.Fatal(err)
	}
}

func TestRun_unknownSubcommand(t *testing.T) {
	c := New()
	if err := c.Run([]string{"nope"}); err == nil {
		t.Fatal()
	}
}
