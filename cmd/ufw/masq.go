package ufw

import (
	"fmt"
	"regexp"
	"strings"

	"z-panel/internal/i18n"
	"z-panel/internal/settings"
)

// postroutingNatLines returns non-empty trimmed lines inside *nat … COMMIT blocks.
func postroutingNatLines(iptablesSave string) []string {
	inNat := false
	var out []string
	for _, raw := range strings.Split(iptablesSave, "\n") {
		line := strings.TrimSpace(raw)
		if line == "*nat" {
			inNat = true
			continue
		}
		if line == "COMMIT" {
			inNat = false
			continue
		}
		if !inNat || line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		if !strings.HasPrefix(line, "-A ") {
			continue
		}
		if !strings.Contains(line, "POSTROUTING") {
			continue
		}
		if strings.Contains(line, "MASQUERADE") || strings.Contains(line, "-j SNAT") {
			out = append(out, line)
		}
	}
	return out
}

func lineMatchesOutIface(line, iface string) bool {
	if iface == "" {
		return false
	}
	re := regexp.MustCompile(`-o ` + regexp.QuoteMeta(iface) + `(\s|$)`)
	return re.MatchString(line)
}

func masqueradeLinesForIface(iptablesSave, iface string) []string {
	var hit []string
	for _, ln := range postroutingNatLines(iptablesSave) {
		if lineMatchesOutIface(ln, iface) {
			hit = append(hit, ln)
		}
	}
	return hit
}

func parseMasqCheckArgs(args []string) (iface, lanCIDR string, err error) {
	lanCIDR = settings.C.DefaultLANCIDR
	var pos []string
	for i := 0; i < len(args); i++ {
		a := args[i]
		switch {
		case strings.HasPrefix(a, "--lan-cidr="):
			lanCIDR = strings.TrimSpace(strings.TrimPrefix(a, "--lan-cidr="))
			if lanCIDR == "" {
				return "", "", fmt.Errorf("%s", i18n.T("ufw.err.lan_cidr_empty"))
			}
		case a == "--lan-cidr":
			if i+1 >= len(args) {
				return "", "", fmt.Errorf("%s", i18n.T("ufw.err.lan_cidr_need"))
			}
			i++
			lanCIDR = strings.TrimSpace(args[i])
			if lanCIDR == "" {
				return "", "", fmt.Errorf("%s", i18n.T("ufw.err.lan_cidr_empty"))
			}
		case strings.HasPrefix(a, "-"):
			return "", "", fmt.Errorf(i18n.T("ufw.err.unknown_flag"), a)
		default:
			pos = append(pos, a)
		}
	}
	if len(pos) != 1 {
		return "", "", fmt.Errorf("%s", i18n.T("ufw.masq.want_iface"))
	}
	return pos[0], lanCIDR, nil
}
