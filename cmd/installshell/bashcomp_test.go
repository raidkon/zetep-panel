package installshell

import (
	"strings"
	"testing"
)

func TestBashCompletionCase(t *testing.T) {
	var b strings.Builder
	New().BashCompletionCase(&b)
	if b.Len() < 20 {
		t.Fatal()
	}
}
