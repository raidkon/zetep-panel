package redirect

import (
	"fmt"
	"strings"

	"z-panel/internal/i18n"
)

// UpOptions are xray-redirect up arguments (after "up").
type UpOptions struct {
	Table        string
	BypassUnit   string
	BypassCgroup string
	NoMark       bool
	IPv6         bool
	// WanLookup: "" (default) = same as "auto" — "from <WAN> lookup main" to fix inbound on public IP with cgroup mark.
	// "off" disables. Otherwise explicit IPv4 or CIDR (e.g. 203.0.113.1/32).
	WanLookup string
}

// ParseUpArgs parses flags and one positional interface name.
func ParseUpArgs(args []string) (iface string, opts UpOptions, err error) {
	opts.BypassUnit = "auto"
	var pos []string
	for i := 0; i < len(args); i++ {
		a := args[i]
		switch {
		case a == "--no-mark":
			opts.NoMark = true
		case a == "--ipv6":
			opts.IPv6 = true
		case strings.HasPrefix(a, "--table="):
			opts.Table = strings.TrimSpace(strings.TrimPrefix(a, "--table="))
		case a == "--table":
			if i+1 >= len(args) {
				return "", opts, fmt.Errorf(i18n.T("redirect.need_value_after"), a)
			}
			i++
			opts.Table = strings.TrimSpace(args[i])
		case strings.HasPrefix(a, "--bypass-cgroup="):
			opts.BypassCgroup = strings.TrimSpace(strings.TrimPrefix(a, "--bypass-cgroup="))
		case a == "--bypass-cgroup":
			if i+1 >= len(args) {
				return "", opts, fmt.Errorf(i18n.T("redirect.need_value_after"), a)
			}
			i++
			opts.BypassCgroup = strings.TrimSpace(args[i])
		case strings.HasPrefix(a, "--bypass-unit="):
			opts.BypassUnit = strings.TrimPrefix(a, "--bypass-unit=")
		case a == "--bypass-unit":
			if i+1 >= len(args) {
				return "", opts, fmt.Errorf(i18n.T("redirect.need_value_after"), a)
			}
			i++
			opts.BypassUnit = args[i]
		case strings.HasPrefix(a, "--wan-lookup="):
			opts.WanLookup = strings.TrimSpace(strings.TrimPrefix(a, "--wan-lookup="))
		case a == "--wan-lookup":
			if i+1 >= len(args) {
				return "", opts, fmt.Errorf(i18n.T("redirect.need_value_after"), a)
			}
			i++
			opts.WanLookup = strings.TrimSpace(args[i])
		case strings.HasPrefix(a, "-"):
			return "", opts, fmt.Errorf(i18n.T("redirect.unknown_flag"), a)
		default:
			pos = append(pos, a)
		}
	}
	if len(pos) != 1 {
		return "", opts, fmt.Errorf(i18n.T("redirect.need_one_iface"), pos)
	}
	return pos[0], opts, nil
}
