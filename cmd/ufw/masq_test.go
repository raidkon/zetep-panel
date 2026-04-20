package ufw

import (
	"testing"
)

func TestPostroutingNatLines(t *testing.T) {
	t.Parallel()
	sample := `# comment
*nat
:PREROUTING ACCEPT [0:0]
-A PREROUTING -j TEST
-A POSTROUTING -s 192.168.0.0/22 -o wan1 -j MASQUERADE
COMMIT
*nat
-A POSTROUTING -s 192.168.0.0/22 -o xray2tun -j MASQUERADE
-A POSTROUTING -j MASQUERADE
COMMIT
`
	lines := postroutingNatLines(sample)
	if len(lines) != 3 {
		t.Fatalf("got %d lines: %v", len(lines), lines)
	}
}

func TestMasqueradeLinesForIface(t *testing.T) {
	t.Parallel()
	sample := `*nat
-A POSTROUTING -s 192.168.0.0/22 -o wan1 -j MASQUERADE
-A POSTROUTING -s 192.168.0.0/22 -o wan10 -j MASQUERADE
-A POSTROUTING -s 10.0.0.0/8 -o wan1 -j SNAT --to-source 1.1.1.1
COMMIT
`
	h := masqueradeLinesForIface(sample, "wan1")
	if len(h) != 2 {
		t.Fatalf("got %v", h)
	}
	if len(masqueradeLinesForIface(sample, "wan10")) != 1 {
		t.Fatal()
	}
	if len(masqueradeLinesForIface(sample, "ppp0")) != 0 {
		t.Fatal()
	}
}
