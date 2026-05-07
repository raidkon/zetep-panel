package ufw

import (
	"os"
	"strings"
	"testing"

	"z-panel/internal/i18n"
	"z-panel/internal/settings"
)

func TestMain(m *testing.M) {
	i18n.Init()
	i18n.ApplyFromConfig("en")
	settings.C = &settings.Cfg{
		DefaultLANCIDR: "192.168.0.0/16",
		DefaultLANDev:  "lan0",
		UfwMarker:      "z-panel",
	}
	os.Exit(m.Run())
}

func TestParseCheckArgs(t *testing.T) {
	t.Parallel()
	settings.C = &settings.Cfg{
		DefaultLANCIDR: "192.168.0.0/16",
		DefaultLANDev:  "lan0",
	}
	iface, cidr, dev, full, err := parseCheckArgs(nil)
	if err != nil || iface != "" || cidr != "192.168.0.0/16" || dev != "lan0" || full {
		t.Fatalf("%v %q %q %q full=%v", err, iface, cidr, dev, full)
	}
	iface, cidr, dev, full, err = parseCheckArgs([]string{"--lan-cidr=10.0.0.0/8", "tun0"})
	if err != nil || iface != "tun0" || cidr != "10.0.0.0/8" || dev != "lan0" || full {
		t.Fatalf("%v %q %q %q full=%v", err, iface, cidr, dev, full)
	}
	_, _, _, _, err = parseCheckArgs([]string{"--lan-cidr"})
	if err == nil || !strings.Contains(err.Error(), "lan-cidr") {
		t.Fatalf("got %v", err)
	}
	_, _, _, _, err = parseCheckArgs([]string{"--lan-cidr="})
	if err == nil {
		t.Fatal()
	}
	_, _, _, _, err = parseCheckArgs([]string{"--oops"})
	if err == nil || !strings.Contains(strings.ToLower(err.Error()), "unknown") {
		t.Fatalf("got %v", err)
	}
	_, _, _, _, err = parseCheckArgs([]string{"a", "b"})
	if err == nil {
		t.Fatal()
	}
}
