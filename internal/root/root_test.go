package root

import (
	"os"
	"testing"
)

func TestRequire_dependsOnUID(t *testing.T) {
	err := Require()
	if os.Geteuid() == 0 {
		if err != nil {
			t.Fatalf("as root: %v", err)
		}
	} else {
		if err == nil {
			t.Fatal("non-root: expected error")
		}
	}
}
