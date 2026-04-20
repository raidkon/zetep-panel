package xrayredirect

import (
	"os"
	"testing"

	"z-panel/internal/i18n"
	"z-panel/internal/settings"
)

func TestMain(m *testing.M) {
	_ = os.Setenv("Z_PANEL_LANG", "en")
	i18n.Init()
	settings.C = &settings.Cfg{Table: "51843"}
	os.Exit(m.Run())
}

func TestRun_help(t *testing.T) {
	c := New()
	if err := c.Run([]string{"help"}); err != nil {
		t.Fatal(err)
	}
}

func TestRun_missingAction(t *testing.T) {
	c := New()
	if err := c.Run([]string{}); err == nil {
		t.Fatal()
	}
}

func TestRun_badAction(t *testing.T) {
	c := New()
	if err := c.Run([]string{"pause"}); err == nil {
		t.Fatal()
	}
}

func TestRun_downWrongArity(t *testing.T) {
	c := New()
	if err := c.Run([]string{"down"}); err == nil {
		t.Fatal()
	}
	if err := c.Run([]string{"down", "a", "b"}); err == nil {
		t.Fatal()
	}
}
