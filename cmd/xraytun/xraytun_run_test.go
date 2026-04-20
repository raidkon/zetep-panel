package xraytun

import (
	"testing"

	"z-panel/internal/settings"
)

func TestRun_help(t *testing.T) {
	settings.C = &settings.Cfg{
		DefaultXrayAddr: "10.0.0.1/30",
		DefaultXrayPeer: "10.0.0.2/30",
	}
	c := New()
	if err := c.Run([]string{"help"}); err != nil {
		t.Fatal(err)
	}
}

func TestRun_missingSub(t *testing.T) {
	c := New()
	if err := c.Run([]string{}); err == nil {
		t.Fatal()
	}
}

func TestRun_badSub(t *testing.T) {
	c := New()
	if err := c.Run([]string{"pause"}); err == nil {
		t.Fatal()
	}
}
