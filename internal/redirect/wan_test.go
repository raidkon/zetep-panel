package redirect

import "testing"

func Test_wanMainBypassPrefV4(t *testing.T) {
	// 51843 % 50 = 43 → 100+43
	if got := wanMainBypassPrefV4("51843"); got != 143 {
		t.Fatalf("51843: got %d", got)
	}
	if got := wanMainBypassPrefV4("70000"); got != 100 { // 70000 % 50 == 0
		t.Fatalf("70000: got %d", got)
	}
}

func Test_policyPrefV4(t *testing.T) {
	// 51843 % 500 = 343
	if s := policyPrefSuppressV4("51843"); s != 32000+343 {
		t.Fatalf("suppress 51843: got %d", s)
	}
	if n := policyPrefNotFwmarkV4("51843"); n != 32000+344 {
		t.Fatalf("notfw 51843: got %d", n)
	}
}

func Test_defaultIPv4WanDevs(t *testing.T) {
	main := "default dev ppp0 scope link\n" +
		"default via 192.168.8.1 dev wan1 proto dhcp src 192.168.8.6 metric 2000\n"
	d := defaultIPv4WanDevs("xray2tun", main)
	if len(d) < 1 || d[0] != "ppp0" {
		t.Fatalf("got %v", d)
	}
	// TUN is skipped: only the other default
	main2 := "default dev xray2tun scope link\n" +
		"default dev ppp0 scope link\n"
	d2 := defaultIPv4WanDevs("xray2tun", main2)
	if len(d2) != 1 || d2[0] != "ppp0" {
		t.Fatalf("got %v", d2)
	}
}

func Test_normalizeWanCIDRInput(t *testing.T) {
	c, err := normalizeWanCIDRInput("192.0.2.1")
	if err != nil || c != "192.0.2.1/32" {
		t.Fatalf("got %q %v", c, err)
	}
	c2, err := normalizeWanCIDRInput("192.0.2.0/24")
	if err != nil || c2 != "192.0.2.0/24" {
		t.Fatalf("got %q %v", c2, err)
	}
	_, err = normalizeWanCIDRInput("nope")
	if err == nil {
		t.Fatal("expected error")
	}
}
