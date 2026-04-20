package redirect

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"z-panel/internal/i18n"
	"z-panel/internal/root"
	"z-panel/internal/settings"
	"z-panel/internal/state"
)

// Up brings up policy routing + firewall (xray-redirect up).
func Up(iface string, opts UpOptions) error {
	if err := root.Require(); err != nil {
		return err
	}
	if err := ipLinkExists(iface); err != nil {
		return err
	}
	table := strings.TrimSpace(opts.Table)
	if table == "" {
		table = settings.C.Table
	}
	if _, err := strconv.ParseUint(table, 10, 32); err != nil {
		return fmt.Errorf(i18n.T("redirect.table_numeric"), err)
	}

	noMark := opts.NoMark
	bypassUnit := strings.TrimSpace(opts.BypassUnit)
	if bypassUnit == "" {
		bypassUnit = "auto"
	}
	mode := "wg"
	var cgroupPath, cgroupUnit string
	var err error
	if !noMark {
		cgroupPath, cgroupUnit, err = resolveBypassCgroupFromFlags(opts.BypassCgroup, bypassUnit)
		if err != nil {
			return err
		}
		if cgroupPath == "" {
			return fmt.Errorf("%s", i18n.T("redirect.cgroup_missing"))
		}
		suffix := ""
		if cgroupUnit != "" {
			suffix = fmt.Sprintf(i18n.T("redirect.bypass_mark_fmt"), cgroupUnit)
		}
		fmt.Fprintf(os.Stderr, i18n.T("redirect.mark_line"), cgroupPath, suffix)
	} else {
		fmt.Fprintf(os.Stderr, "%s", i18n.T("redirect.no_mark_line"))
	}

	if err := os.MkdirAll(settings.C.StateDir, 0o750); err != nil {
		return err
	}

	ipRouteFlushTableQuiet(table)
	if err := run("ip", "-4", "route", "add", "0.0.0.0/0", "dev", iface, "table", table); err != nil {
		return fmt.Errorf(i18n.T("redirect.err.default_route"), iface, table, err)
	}
	if err := run("ip", "-4", "rule", "add", "not", "fwmark", table, "table", table); err != nil {
		return fmt.Errorf(i18n.T("redirect.err.rule_fwmark"), err)
	}
	if err := run("ip", "-4", "rule", "add", "table", "main", "suppress_prefixlength", "0"); err != nil {
		return fmt.Errorf(i18n.T("redirect.err.rule_suppress"), err)
	}
	_ = run("sysctl", "-q", "net.ipv4.conf.all.src_valid_mark=1")

	ipv6 := opts.IPv6
	if ipv6 {
		if err := run("ip", "-6", "route", "add", "::/0", "dev", iface, "table", table); err != nil {
			return fmt.Errorf(i18n.T("redirect.err.route6"), err)
		}
		if err := run("ip", "-6", "rule", "add", "not", "fwmark", table, "table", table); err != nil {
			return fmt.Errorf(i18n.T("redirect.err.rule6_fw"), err)
		}
		if err := run("ip", "-6", "rule", "add", "table", "main", "suppress_prefixlength", "0"); err != nil {
			return fmt.Errorf(i18n.T("redirect.err.rule6_sup"), err)
		}
	}

	if err := addWGStyleFirewall(iface, table, ipv6, cgroupPath); err != nil {
		_ = downWGQuick(state.Partial(iface, table, "wg", ipv6), ipv6)
		return err
	}

	st := state.File{
		Mode:         mode,
		Interface:    iface,
		Table:        table,
		Fwmark:       table,
		WGIPv6:       ipv6,
		NoBypassMark: noMark,
		BypassCgroup: cgroupPath,
		BypassUnit:   cgroupUnit,
	}
	summary := i18n.T("state.summary_base", mode, table, table, iface)
	if noMark {
		summary += i18n.T("state.summary_nomark")
	} else {
		summary += i18n.T("state.summary_bypass")
	}
	return state.WriteAndPrint(st, summary)
}
