package daemon

import "testing"

func TestForbiddenRemote(t *testing.T) {
	if !ForbiddenRemote("daemon") {
		t.Fatal("expected daemon forbidden")
	}
	if ForbiddenRemote("xray-redirect") || ForbiddenRemote("config") || ForbiddenRemote("version") {
		t.Fatal("unexpected forbidden")
	}
}
