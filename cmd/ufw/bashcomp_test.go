package ufw

import (
	"strings"
	"testing"
)

func TestBashCompletionCase(t *testing.T) {
	var b strings.Builder
	New().BashCompletionCase(&b)
	if !strings.Contains(b.String(), "ufw") {
		t.Fatal()
	}
}
