package ufw

import (
	"errors"
	"testing"
)

func TestMarkerLinesPresent(t *testing.T) {
	if !markerLinesPresent([]string{"x z-panel Y"}, "Z-Panel") {
		t.Fatal()
	}
	if markerLinesPresent([]string{"nothing"}, "z-panel") {
		t.Fatal()
	}
}

func TestBuildCheckIssues_noIface(t *testing.T) {
	sev, issues := buildCheckIssues("", "192.168.0.0/16", "lan0", nil, 0, nil, nil, true)
	if sev != checkWarn || len(issues) != 1 {
		t.Fatalf("%v %v", sev, issues)
	}
}

func TestBuildCheckIssues_iptErr(t *testing.T) {
	sev, issues := buildCheckIssues("tun0", "192.168.0.0/16", "lan0", []string{"line"}, 1, []string{"m"}, errors.New("boom"), true)
	if sev != checkBad {
		t.Fatal(sev)
	}
	var found bool
	for _, s := range issues {
		if len(s) > 0 {
			found = true
			break
		}
	}
	if !found {
		t.Fatal(issues)
	}
}

func TestBuildCheckIssues_okPath(t *testing.T) {
	sev, _ := buildCheckIssues("eth0", "10.0.0.0/8", "lan0", []string{"eth0 on"}, 3, []string{"MASQUERADE"}, nil, true)
	if sev != checkOK {
		t.Fatalf("got %v want OK", sev)
	}
}

func TestBuildCheckIssues_noReturnWarn(t *testing.T) {
	sev, issues := buildCheckIssues("eth0", "10.0.0.0/8", "lan0", []string{"eth0"}, 1, []string{"-A POSTROUTING -j MASQUERADE"}, nil, false)
	if sev != checkWarn || len(issues) == 0 {
		t.Fatalf("%v %d", sev, len(issues))
	}
}
