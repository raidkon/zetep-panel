package i18n

import "testing"

func TestMergeMaps(t *testing.T) {
	t.Parallel()
	base := map[string]string{"a": "1", "b": "2"}
	over := map[string]string{"b": "x", "c": "3"}
	got := mergeMaps(base, over)
	if got["a"] != "1" || got["b"] != "x" || got["c"] != "3" {
		t.Fatalf("%v", got)
	}
	// base unchanged
	if base["b"] != "2" {
		t.Fatal("mutated base")
	}
}

func TestMergeMaps_emptyOver(t *testing.T) {
	t.Parallel()
	base := map[string]string{"k": "v"}
	got := mergeMaps(base, nil)
	if len(got) != 1 || got["k"] != "v" {
		t.Fatalf("%v", got)
	}
}
