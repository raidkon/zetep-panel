package settings

import (
	"testing"
)

func TestApplyDefaults(t *testing.T) {
	ApplyDefaults()
	if C == nil || C.UfwMarker == "" {
		t.Fatal()
	}
}

func TestDaemonEnabled_variants(t *testing.T) {
	if (&Cfg{Daemon: 1}).DaemonEnabled() != true {
		t.Fatal()
	}
	if (&Cfg{Daemon: 0}).DaemonEnabled() != false {
		t.Fatal()
	}
	if (&Cfg{}).DaemonEnabled() != false {
		t.Fatal()
	}
}
