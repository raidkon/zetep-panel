package state

import "testing"

func TestPartial(t *testing.T) {
	f := Partial("eth0", "51843", "wg", true)
	if f.Interface != "eth0" || f.Table != "51843" || f.Mode != "wg" || !f.WGIPv6 {
		t.Fatalf("%+v", f)
	}
}
