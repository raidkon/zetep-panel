package redirect

import (
	"fmt"
	"net"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"z-panel/internal/executil"
	"z-panel/internal/i18n"
)

// wanMainBypassPrefV4 is the ip rule priority for "from <WAN> lookup main".
// It must be LOWER than suppressNotFwmarkPrefs (earlier in rule list) than not-fwmark → table.
func wanMainBypassPrefV4(table string) int {
	n, _ := strconv.Atoi(table)
	if n < 0 {
		n = 0
	}
	// 100–149: always before z-panel suppress/not-fwmark (32xxx) and typical 3276x rules.
	return 100 + (n % 50)
}

// policyPrefSuppressV4 / policyPrefNotFwmarkV4 are explicit prefs for the two wg-style rules so the
// kernel does not auto-fill 3184x between our WAN rule and them (which would break bypass order).
func policyPrefSuppressV4(table string) int {
	n, _ := strconv.Atoi(table)
	if n < 0 {
		n = 0
	}
	return 32000 + (n % 500)
}

func policyPrefNotFwmarkV4(table string) int {
	return policyPrefSuppressV4(table) + 1
}

// legacyWanRulePrefV4 was used in z-panel 0.14.0 before ordering fix; remove on teardown.
func legacyWanRulePrefV4(table string) int {
	n, _ := strconv.Atoi(table)
	if n < 0 {
		n = 0
	}
	return 30000 + (n % 2000)
}

var reMainDefault = regexp.MustCompile(`\bdefault\b.*\bdev\s+(\S+)`)
var reMainMetric = regexp.MustCompile(`\bmetric\s+(\d+)\b`)

type mainDefaultRoute struct {
	iface  string
	metric int
}

// defaultIPv4WanDevs returns default gateway device names from `table main` (IPv4)
// in ascending metric order, excluding the tunnel interface. Used to pick WAN for policy bypass.
func defaultIPv4WanDevs(tunIface, mainRouteList string) []string {
	var out []mainDefaultRoute
	for _, line := range strings.Split(mainRouteList, "\n") {
		line = strings.TrimSpace(line)
		if !strings.HasPrefix(line, "default") {
			continue
		}
		m := reMainDefault.FindStringSubmatch(line)
		if m == nil {
			continue
		}
		iface := m[1]
		if iface == tunIface {
			continue
		}
		metric := 0
		if m2 := reMainMetric.FindStringSubmatch(line); m2 != nil {
			if v, err := strconv.Atoi(m2[1]); err == nil {
				metric = v
			}
		}
		out = append(out, mainDefaultRoute{iface: iface, metric: metric})
	}
	if len(out) == 0 {
		return nil
	}
	sort.SliceStable(out, func(i, j int) bool {
		if out[i].metric != out[j].metric {
			return out[i].metric < out[j].metric
		}
		return out[i].iface < out[j].iface
	})
	seen := make(map[string]struct{})
	var ifaces []string
	for _, e := range out {
		if _, ok := seen[e.iface]; ok {
			continue
		}
		seen[e.iface] = struct{}{}
		ifaces = append(ifaces, e.iface)
	}
	return ifaces
}

// firstIPv4OnDev returns a /32 CIDR of the first suitable global unicast address on the device.
// PPP "peer" lines without slash are handled.
func firstIPv4CIDR32OnDev(dev string) (string, error) {
	data, err := executil.RunTTYCombined("ip", "-4", "-o", "addr", "show", "dev", dev)
	if err != nil {
		return "", fmt.Errorf("ip addr show dev %s: %w", dev, err)
	}
	// e.g. "2: ppp0    inet 1.2.3.4 peer 5.6.7.8/32 scope global ppp0"
	// e.g. "2: eth0    inet 192.168.1.1/24 brd ..."
	for _, line := range strings.Split(string(data), "\n") {
		line = strings.TrimSpace(line)
		if !strings.Contains(line, "inet ") {
			continue
		}
		// "scope link" (secondary addresses) and "scope host" are skipped
		if strings.Contains(line, "scope") && !strings.Contains(line, "scope global") {
			continue
		}
		idx := strings.Index(line, " inet ")
		if idx == -1 {
			continue
		}
		rest := line[idx+5:]
		fields := strings.Fields(rest)
		if len(fields) == 0 {
			continue
		}
		raw := fields[0]
		var host string
		if strings.Contains(raw, "/") {
			if ip, _, err := net.ParseCIDR(raw); err == nil && !ip.IsUnspecified() {
				if ip4 := ip.To4(); ip4 != nil {
					if !isSkippableLocal(ip4) {
						host = ip4.String()
					}
				}
			}
		} else {
			if ip4 := net.ParseIP(raw).To4(); ip4 != nil && !isSkippableLocal(ip4) {
				host = ip4.String()
			}
		}
		if host == "" {
			continue
		}
		return host + "/32", nil
	}
	return "", nil
}

func isSkippableLocal(ip4 net.IP) bool {
	if ip4 == nil {
		return true
	}
	if ip4.IsLoopback() {
		return true
	}
	if v := ip4[0]; v == 169 && ip4[1] == 254 { // link-local
		return true
	}
	// 100.64/10 — CGNAT, still often a valid "WAN" in double-NAT; do not skip
	// 10/8, 172.16/12, 192.168/16: prefer actual public; if only private, still useful for "from" rule
	_ = ip4
	return false
}

// wanCIDRForAutoMode picks primary WAN IPv4/32 from `ip -4 route show table main` (not via ip rule / route get).
func wanCIDRForAutoMode(tunIface string) (string, error) {
	out, err := executil.RunTTYCombined("ip", "-4", "route", "show", "table", "main")
	if err != nil {
		return "", err
	}
	devs := defaultIPv4WanDevs(tunIface, string(out))
	for _, dev := range devs {
		cidr, err := firstIPv4CIDR32OnDev(dev)
		if err != nil {
			return "", err
		}
		if cidr != "" {
			return cidr, nil
		}
	}
	return "", nil
}

// ResolveWanMainLookupCIDR returns the IPv4 CIDR for --wan-lookup, or "" if the rule should be omitted.
func ResolveWanMainLookupCIDR(tunIface, wanMode string) (cidr string, err error) {
	mode := strings.TrimSpace(wanMode)
	if mode == "" {
		mode = "auto"
	}
	m := strings.ToLower(mode)
	switch m {
	case "off", "no", "false", "0":
		return "", nil
	case "auto", "on", "yes", "true":
		c, werr := wanCIDRForAutoMode(tunIface)
		if werr != nil {
			fmt.Fprint(os.Stderr, i18n.T("redirect.wan_auto_err", werr))
		}
		if c == "" {
			fmt.Fprint(os.Stderr, i18n.T("redirect.wan_auto_skip", tunIface))
		}
		return c, nil
	default:
		return normalizeWanCIDRInput(mode)
	}
}

// normalizeWanCIDRInput turns "1.2.3.4" into "1.2.3.4/32"; validates CIDR.
func normalizeWanCIDRInput(s string) (string, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return "", fmt.Errorf("empty CIDR")
	}
	if !strings.Contains(s, "/") {
		if ip := net.ParseIP(s); ip != nil {
			if v4 := ip.To4(); v4 != nil {
				return v4.String() + "/32", nil
			}
		}
		return "", fmt.Errorf("not an IPv4 address: %q", s)
	}
	ip, n, err := net.ParseCIDR(s)
	if err != nil {
		return "", err
	}
	if ip4 := ip.To4(); ip4 == nil {
		return "", fmt.Errorf("IPv4 CIDR expected: %q", s)
	}
	ones, _ := n.Mask.Size()
	if ones > 32 {
		return "", fmt.Errorf("prefix too long: %q", s)
	}
	// use canonical
	return fmt.Sprintf("%s/%d", ip.To4().String(), ones), nil
}
