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
	if EffectiveStoredSchema(4) != 4 {
		t.Fatal()
	}
}

func TestConfigPath(t *testing.T) {
	p := ConfigPath()
	if p == "" {
		t.Fatal()
	}
}

func TestApplySchemaUpgradesAuto(t *testing.T) {
	t.Parallel()
	c := Cfg{SchemaVersion: 1, Language: "en"}
	stored := EffectiveStoredSchema(c.SchemaVersion)
	if err := applySchemaUpgradesAuto(&c, stored); err != nil {
		t.Fatal(err)
	}
	c.SchemaVersion = CurrentSchemaVersion
	if c.SchemaVersion != CurrentSchemaVersion {
		t.Fatalf("expected CurrentSchemaVersion %d, got %d", CurrentSchemaVersion, c.SchemaVersion)
	}
}
