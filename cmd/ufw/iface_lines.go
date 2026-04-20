package ufw

import (
	"regexp"
	"strings"
)

// statusLinesReferencingIface returns ufw status lines where the interface appears
// as a UFW "on <iface>" device (e.g. "Anywhere on xray2tun ALLOW FWD …").
func statusLinesReferencingIface(lines []string, iface string) []string {
	if iface == "" {
		return nil
	}
	esc := regexp.QuoteMeta(iface)
	// ufw status: "… on xray2tun ALLOW …" or route-style "… in on lan0 out on xray2tun …"
	re := regexp.MustCompile(`(?:\s+on|\bout\s+on|\bin\s+on)\s+` + esc + `\s+`)
	var out []string
	for _, ln := range lines {
		s := strings.TrimRight(ln, "\r")
		if re.MatchString(s) {
			out = append(out, ln)
		}
	}
	return out
}

// fwdRuleCountForIface counts ufw status lines that both mention iface (on/in/out)
// and allow forwarding.
func fwdRuleCountForIface(lines []string, iface string) int {
	if iface == "" {
		return 0
	}
	esc := regexp.QuoteMeta(iface)
	re := regexp.MustCompile(`(?:\s+on|\bout\s+on|\bin\s+on)\s+` + esc + `\s+`)
	n := 0
	for _, ln := range lines {
		s := strings.TrimRight(ln, "\r")
		if !re.MatchString(s) {
			continue
		}
		if strings.Contains(strings.ToUpper(s), "ALLOW FWD") {
			n++
		}
	}
	return n
}

// statusAppearsToHaveReturnPath reports whether any status line looks like UFW
// permitting tunnel → LAN (return path). Forms seen in "ufw status verbose":
//   - ALLOW ROUTE … on <tun> … on <lan> (or swapped columns)
//   - ALLOW FWD with To = … on <lan> and From = … on <tun> (common for route add)
//   - literal "in on <tun>" / "out on <lan>" (rare)
func statusAppearsToHaveReturnPath(lines []string, tunnelIface, lanDev string) bool {
	if tunnelIface == "" || lanDev == "" {
		return false
	}
	tun := strings.ToLower(tunnelIface)
	lan := strings.ToLower(lanDev)
	inOn := "in on " + tun
	outOn := "out on " + lan
	reTun := regexp.MustCompile(`\s+on\s+` + regexp.QuoteMeta(tun) + `\s+`)
	reLan := regexp.MustCompile(`\s+on\s+` + regexp.QuoteMeta(lan) + `\s+`)

	for _, ln := range lines {
		low := strings.ToLower(strings.TrimRight(ln, "\r"))
		if strings.Contains(low, inOn) && strings.Contains(low, outOn) {
			return true
		}
		if strings.Contains(low, "allow route") && reTun.MatchString(low) && reLan.MatchString(low) {
			return true
		}
		// "Anywhere on lan0  ALLOW FWD  Anywhere on xray2tun" (return); not LAN→tun FWD
		// where To is tun and From is … on lan0.
		if idx := strings.Index(low, "allow fwd"); idx >= 0 {
			before := low[:idx]
			after := low[idx+len("allow fwd"):]
			if reLan.MatchString(before) && reTun.MatchString(after) {
				return true
			}
		}
	}
	return false
}
