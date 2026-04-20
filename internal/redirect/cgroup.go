package redirect

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"z-panel/internal/i18n"
)

func resolveBypassCgroupFromFlags(explicitPath, bypassUnit string) (path string, unit string, err error) {
	explicitPath = strings.TrimSpace(explicitPath)
	if explicitPath != "" {
		return strings.TrimPrefix(explicitPath, "/"), "", nil
	}
	bu := strings.TrimSpace(strings.ToLower(bypassUnit))
	if bu == "" || bu == "auto" {
		return autoDetectBypassUnit()
	}
	u := normalizeSystemdUnit(bypassUnit)
	return cgroupFromUnit(u)
}

func cgroupFromUnit(u string) (path string, unit string, err error) {
	out, err := exec.Command("systemctl", "show", "-p", "ControlGroup", "--value", u).Output()
	if err != nil {
		return "", u, fmt.Errorf(i18n.T("redirect.cg_systemctl"), u, err)
	}
	cg := strings.TrimSpace(string(out))
	if cg == "" || strings.EqualFold(cg, "[not set]") {
		return "", u, fmt.Errorf(i18n.T("redirect.cg_empty"), u)
	}
	return strings.TrimPrefix(cg, "/"), u, nil
}

func autoDetectBypassUnit() (path string, unit string, err error) {
	candidates := []string{"x-ui", "sing-box", "xray"}
	var lastErr error
	for _, c := range candidates {
		u := normalizeSystemdUnit(c)
		p, un, err := cgroupFromUnit(u)
		if err == nil && p != "" {
			fmt.Fprintf(os.Stderr, i18n.T("redirect.auto_unit"), un)
			return p, un, nil
		}
		lastErr = err
	}
	return "", "", fmt.Errorf(i18n.T("redirect.auto_fail"), candidates, lastErr)
}

func normalizeSystemdUnit(s string) string {
	s = strings.TrimSpace(s)
	if s == "" || strings.Contains(s, ".") {
		return s
	}
	return s + ".service"
}
