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
	if EffectiveStoredSchema(3) != 3 {
		t.Fatal()
	}
}

func TestConfigPath(t *testing.T) {
	p := ConfigPath()
	if p == "" {
		t.Fatal()
	}
}

func TestDaemonEnabled(t *testing.T) {
	t.Parallel()
	var nilCfg *Cfg
	if nilCfg.DaemonEnabled() {
		t.Fatal()
	}
	c := &Cfg{Daemon: 0}
	if c.DaemonEnabled() {
		t.Fatal()
	}
	c.Daemon = 1
	if !c.DaemonEnabled() {
		t.Fatal()
	}
}
