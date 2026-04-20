package redirect

import (
	"os"
	"strings"
	"testing"

	"z-panel/internal/i18n"
)

func TestMain(m *testing.M) {
	os.Setenv("Z_PANEL_LANG", "en")
	i18n.Init()
	os.Exit(m.Run())
}

func TestParseUpArgs(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name    string
		args    []string
		wantIF  string
		want    UpOptions
		wantErr string
	}{
		{
			name:   "iface only",
			args:   []string{"tun0"},
			wantIF: "tun0",
			want:   UpOptions{BypassUnit: "auto"},
		},
		{
			name:   "flags",
			args:   []string{"--no-mark", "--ipv6", "--table", "99", "xray2tun"},
			wantIF: "xray2tun",
			want: UpOptions{
				BypassUnit: "auto",
				NoMark:     true,
				IPv6:       true,
				Table:      "99",
			},
		},
		{
			name:   "equals forms",
			args:   []string{"--table=42", "--bypass-cgroup=/sys/fs/cgroup/foo", "--bypass-unit=myapp", "eth9"},
			wantIF: "eth9",
			want: UpOptions{
				BypassUnit:   "myapp",
				BypassCgroup: "/sys/fs/cgroup/foo",
				Table:        "42",
			},
		},
		{
			name:    "unknown flag",
			args:    []string{"--nope", "tun0"},
			wantErr: "unknown flag",
		},
		{
			name:    "table without value",
			args:    []string{"--table"},
			wantErr: "value required",
		},
		{
			name:    "no iface",
			args:    []string{"--ipv6"},
			wantErr: "exactly one interface",
		},
		{
			name:    "two ifaces",
			args:    []string{"a", "b"},
			wantErr: "exactly one interface",
		},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			iface, opts, err := ParseUpArgs(tc.args)
			if tc.wantErr != "" {
				if err == nil {
					t.Fatalf("want error containing %q", tc.wantErr)
				}
				if !strings.Contains(strings.ToLower(err.Error()), strings.ToLower(tc.wantErr)) {
					t.Fatalf("err=%v want substring %q", err, tc.wantErr)
				}
				return
			}
			if err != nil {
				t.Fatal(err)
			}
			if iface != tc.wantIF {
				t.Fatalf("iface=%q want %q", iface, tc.wantIF)
			}
			if opts != tc.want {
				t.Fatalf("opts=%+v want %+v", opts, tc.want)
			}
		})
	}
}
