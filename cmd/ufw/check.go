package ufw

import (
	"fmt"
	"strings"

	"z-panel/internal/executil"
	"z-panel/internal/i18n"
	"z-panel/internal/settings"
)

type checkSeverity int

const (
	checkOK checkSeverity = iota
	checkWarn
	checkBad
)

// Printed between ufw and iptables-save inside one sudo sh -c (one SSH, one sudo prompt over --ssh).
const ufwCheckIptSplitMarker = "ZPANEL_UFWCHECK_IPT_SPLIT_LINE"

func runUnifiedCheck(iface, lanCIDR, lanDev string, full bool) error {
	var text string
	var masqLines []string
	var iptErr error

	if iface != "" {
		// One argv to sudo sh -c only — do not nest /bin/sh -c + quoting; ssh/sudo may drop -c and leave an interactive sudo sh (root #).
		script := fmt.Sprintf("/usr/sbin/ufw status verbose; echo; echo %s; /usr/sbin/iptables-save -t nat", ufwCheckIptSplitMarker)
		out, err := executil.RunTTYCombinedScript(script)
		if err != nil {
			return fmt.Errorf(i18n.T("ufw.ufw_status_failed"), err, out)
		}
		s := string(out)
		key := "\n" + ufwCheckIptSplitMarker + "\n"
		if i := strings.Index(s, key); i >= 0 {
			text = s[:i]
			ipt := strings.TrimSuffix(s[i+len(key):], "\n")
			masqLines = masqueradeLinesForIface(ipt, iface)
		} else {
			text = s
			iptErr = fmt.Errorf("%s", i18n.T("ufw.check.err_ipt_split"))
		}
	} else {
		out, err := executil.RunTTYCombined("ufw", "status", "verbose")
		if err != nil {
			return fmt.Errorf(i18n.T("ufw.ufw_status_failed"), err, out)
		}
		text = string(out)
	}

	lines := strings.Split(text, "\n")

	ifaceHits := statusLinesReferencingIface(lines, iface)
	nFwd := fwdRuleCountForIface(lines, iface)
	returnPath := statusAppearsToHaveReturnPath(lines, iface, lanDev)

	sev, issues := buildCheckIssues(iface, lanCIDR, lanDev, ifaceHits, nFwd, masqLines, iptErr, returnPath)

	printStatusLine(sev)
	fmt.Println()

	if full || sev != checkOK {
		if len(issues) > 0 {
			fmt.Println(i18n.T("ufw.check.section_details"))
			for _, s := range issues {
				fmt.Println(s)
				fmt.Println()
			}
		} else if full && sev == checkOK {
			fmt.Println(i18n.T("ufw.check.no_issues_full"))
			fmt.Println()
		}
	}

	if !full {
		return nil
	}

	markerFound := markerLinesPresent(lines, settings.C.UfwMarker)
	printCheckFullDetails(text, lines, iface, ifaceHits, masqLines, lanCIDR, lanDev, markerFound)
	return nil
}

func markerLinesPresent(lines []string, marker string) bool {
	m := strings.ToLower(marker)
	for _, ln := range lines {
		if strings.Contains(strings.ToLower(ln), m) {
			return true
		}
	}
	return false
}

func buildCheckIssues(iface, lanCIDR, lanDev string, ifaceHits []string, nFwd int, masqLines []string, iptErr error, returnPath bool) (checkSeverity, []string) {
	var issues []string
	var hasBad, hasWarn bool

	if iface == "" {
		hasWarn = true
		issues = append(issues, i18n.T("ufw.check.issue_no_iface"))
		return checkWarn, issues
	}

	if iptErr != nil {
		hasBad = true
		issues = append(issues, i18n.T("ufw.check.issue_iptables", iptErr))
	}

	if len(ifaceHits) == 0 {
		hasBad = true
		issues = append(issues, i18n.T("ufw.check.issue_no_ufw", iface)+"\n"+i18n.T("ufw.check.fix_no_ufw", lanDev, iface, lanCIDR, settings.C.UfwMarker))
	} else if nFwd == 0 {
		hasBad = true
		issues = append(issues, i18n.T("ufw.check.issue_no_fwd", iface)+"\n"+i18n.T("ufw.check.fix_no_fwd", lanDev, iface, lanCIDR, settings.C.UfwMarker))
	}

	if iptErr == nil && len(masqLines) == 0 {
		hasBad = true
		issues = append(issues, i18n.T("ufw.check.issue_no_masq", iface)+"\n"+i18n.T("ufw.masq.hint_add", lanCIDR, iface))
	}

	if !hasBad && !returnPath {
		hasWarn = true
		issues = append(issues, i18n.T("ufw.check.issue_no_return", iface, lanDev)+"\n"+i18n.T("ufw.hint_return")+i18n.T("ufw.hint_return_cmd", iface, lanDev, settings.C.UfwMarker))
	}

	var sev checkSeverity
	switch {
	case hasBad:
		sev = checkBad
	case hasWarn:
		sev = checkWarn
	default:
		sev = checkOK
	}
	return sev, issues
}

func printStatusLine(sev checkSeverity) {
	prefix := i18n.T("ufw.check.status_label")
	var phrase string
	switch sev {
	case checkOK:
		phrase = colorGreen(i18n.T("ufw.check.status_ok"))
	case checkWarn:
		phrase = colorYellow(i18n.T("ufw.check.status_warn"))
	case checkBad:
		phrase = colorRed(i18n.T("ufw.check.status_bad"))
	}
	fmt.Println(prefix + phrase)
}

func printCheckFullDetails(text string, lines []string, iface string, ifaceHits, masqLines []string, lanCIDR, lanDev string, markerFound bool) {
	fmt.Printf("%s\n", i18n.T("ufw.section_rules", settings.C.UfwMarker))
	if markerFound {
		for _, ln := range lines {
			if strings.Contains(strings.ToLower(ln), settings.C.UfwMarker) {
				fmt.Println(ln)
			}
		}
	} else {
		fmt.Println(i18n.T("ufw.no_lines"))
	}

	if iface != "" {
		fmt.Println()
		fmt.Printf("%s\n", i18n.T("ufw.section_iface_refs", iface))
		if len(ifaceHits) > 0 {
			for _, ln := range ifaceHits {
				fmt.Println(ln)
			}
		} else {
			fmt.Println(i18n.T("ufw.no_iface_refs", iface))
		}

		fmt.Println()
		fmt.Println(i18n.T("ufw.masq.detail_heading"))
		if len(masqLines) > 0 {
			for _, ln := range masqLines {
				fmt.Println(ln)
			}
		} else {
			fmt.Println(i18n.T("ufw.check.masq_none_in_full"))
		}
	}

	fmt.Println()
	fmt.Println(i18n.T("ufw.section_hints"))
	if iface != "" {
		fmt.Printf(i18n.T("ufw.hint_sysctl"), lanDev, iface, lanCIDR, settings.C.UfwMarker)
		fmt.Println()
		fmt.Print(i18n.T("ufw.hint_return"))
		fmt.Printf(i18n.T("ufw.hint_return_cmd"), iface, lanDev, settings.C.UfwMarker)
	} else {
		fmt.Printf("%s", i18n.T("ufw.no_iface_hint1"))
		fmt.Printf(i18n.T("ufw.no_iface_hint2"), lanCIDR, lanDev)
	}
	fmt.Println()
	fmt.Println(i18n.T("ufw.section_full"))
	fmt.Print(text)
}
