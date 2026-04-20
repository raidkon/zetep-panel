package ufw

import (
	"reflect"
	"testing"
)

func TestStatusLinesReferencingIface(t *testing.T) {
	lines := []string{
		"Anywhere on xray2tun       ALLOW FWD   192.168.0.0/22 on lan0     # comment",
		"53/udp on lan0             ALLOW IN    192.168.0.0/22",
		"",
	}
	got := statusLinesReferencingIface(lines, "xray2tun")
	want := []string{lines[0]}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("xray2tun: got %#v want %#v", got, want)
	}
	if h := statusLinesReferencingIface(lines, "lan0"); len(h) != 2 {
		t.Fatalf("lan0: want 2 lines, got %d %#v", len(h), h)
	}
	if h := statusLinesReferencingIface(lines, "ppp9"); len(h) != 0 {
		t.Fatalf("ppp9: want 0 lines, got %#v", h)
	}
}

func TestFwdRuleCountForIface(t *testing.T) {
	lines := []string{
		"Anywhere on xray2tun       ALLOW FWD   192.168.0.0/22 on lan0",
		"53/udp on xray2tun             ALLOW IN    Anywhere",
	}
	if n := fwdRuleCountForIface(lines, "xray2tun"); n != 1 {
		t.Fatalf("fwd count: got %d want 1", n)
	}
}

func TestStatusAppearsToHaveReturnPath(t *testing.T) {
	lines := []string{
		"foo in on xray2tun bar out on lan0 baz",
	}
	if !statusAppearsToHaveReturnPath(lines, "xray2tun", "lan0") {
		t.Fatal("expected true")
	}
	if statusAppearsToHaveReturnPath([]string{"on xray2tun ALLOW FWD"}, "xray2tun", "lan0") {
		t.Fatal("expected false without in/out pair")
	}
	routeVerbose := "Anywhere on lan0             ALLOW ROUTE Anywhere on xray2tun             # z-panel: return path"
	if !statusAppearsToHaveReturnPath([]string{routeVerbose}, "xray2tun", "lan0") {
		t.Fatal("expected true for ufw status ALLOW ROUTE line")
	}
	routeVerbose2 := "Anywhere on xray2tun       ALLOW ROUTE Anywhere on lan0       "
	if !statusAppearsToHaveReturnPath([]string{routeVerbose2}, "xray2tun", "lan0") {
		t.Fatal("expected true when tunnel appears first in ROUTE row")
	}
	fwdReturn := "Anywhere on lan0           ALLOW FWD   Anywhere on xray2tun       # z-panel: return path"
	if !statusAppearsToHaveReturnPath([]string{fwdReturn}, "xray2tun", "lan0") {
		t.Fatal("expected true for ALLOW FWD return-path row (lan in To, tun in From)")
	}
	fwdLan2Tun := "Anywhere on xray2tun       ALLOW FWD   192.168.0.0/22 on lan0     # LAN to tunnel"
	if statusAppearsToHaveReturnPath([]string{fwdLan2Tun}, "xray2tun", "lan0") {
		t.Fatal("LAN→tun FWD must not count as return path")
	}
}
