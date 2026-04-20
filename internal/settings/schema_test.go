package settings

import "testing"

func TestEffectiveStoredSchema(t *testing.T) {
	t.Parallel()
	if EffectiveStoredSchema(0) != 1 {
		t.Fatal()
	}
	if EffectiveStoredSchema(-1) != 1 {
		t.Fatal()
	}
	if EffectiveStoredSchema(1) != 1 {
		t.Fatal()
	}
	if EffectiveStoredSchema(2) != 2 {
		t.Fatal()
	}
}

func TestConfigPath(t *testing.T) {
	p := ConfigPath()
	if p == "" {
		t.Fatal()
	}
}
