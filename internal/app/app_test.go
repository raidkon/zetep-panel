package app

import (
	"bytes"
	"io"
	"os"
	"testing"

	"z-panel/internal/i18n"
	"z-panel/internal/settings"
)

type stubCmd struct {
	name string
}

func (s stubCmd) Name() string { return s.name }
func (s stubCmd) Run([]string) error { return nil }
func (s stubCmd) Help(w io.Writer) { _, _ = io.WriteString(w, s.name+" help\n") }
func (s stubCmd) BashCompletionCase(w io.Writer) {}

func TestIsHelpRequest(t *testing.T) {
	t.Parallel()
	if !IsHelpRequest([]string{"help"}) || !IsHelpRequest([]string{"-h"}) || !IsHelpRequest([]string{"--help"}) {
		t.Fatal()
	}
	if IsHelpRequest([]string{}) || IsHelpRequest([]string{"up"}) {
		t.Fatal()
	}
}

func TestFindCommand(t *testing.T) {
	t.Cleanup(func() { registry = nil })
	Register(stubCmd{name: "alpha"})
	Register(stubCmd{name: "beta"})
	if FindCommand("alpha") == nil || FindCommand("beta") == nil {
		t.Fatal()
	}
	if FindCommand("gamma") != nil {
		t.Fatal()
	}
}

func TestTopLevelCompletionWords_unique(t *testing.T) {
	t.Cleanup(func() { registry = nil })
	Register(stubCmd{name: "one"})
	Register(stubCmd{name: "two"})
	w := topLevelCompletionWords()
	seen := map[string]bool{}
	for _, s := range w {
		if seen[s] {
			t.Fatalf("dup %q", s)
		}
		seen[s] = true
	}
}

func TestPrintRootHelp_smoke(t *testing.T) {
	i18n.Init()
	_ = os.Setenv("Z_PANEL_LANG", "en")
	i18n.Init()
	settings.C = &settings.Cfg{UfwMarker: "z-panel"}
	t.Cleanup(func() { registry = nil })
	Register(stubCmd{name: "stub"})
	var buf bytes.Buffer
	PrintRootHelp(&buf)
	if buf.Len() < 50 {
		t.Fatalf("short output: %q", buf.String())
	}
}
