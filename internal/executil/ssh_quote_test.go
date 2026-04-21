package executil

import "testing"

func TestShellSingleQuote(t *testing.T) {
	if got := shellSingleQuote("a b"); got != "'a b'" {
		t.Fatalf("got %q", got)
	}
	if got := shellSingleQuote("it's"); got != `'it'"'"'s'` {
		t.Fatalf("got %q", got)
	}
}
