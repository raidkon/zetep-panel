package redirect

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	"z-panel/internal/i18n"
)

var reIPAddrShow = regexp.MustCompile(`inet\s+([0-9.]+)/`)
var reIP6AddrShow = regexp.MustCompile(`inet6\s+([0-9a-fA-F:]+)/`)

func tunIfaceIPv4s(iface string) ([]string, error) {
	out, err := exec.Command("ip", "-o", "-4", "addr", "show", "dev", iface).Output()
	if err != nil {
		return nil, fmt.Errorf("ip addr show %s: %w", iface, err)
	}
	var ips []string
	for _, m := range reIPAddrShow.FindAllStringSubmatch(string(out), -1) {
		if len(m) > 1 {
			ips = append(ips, m[1])
		}
	}
	return ips, nil
}

func tunIfaceIPv6s(iface string) []string {
	out, err := exec.Command("ip", "-o", "-6", "addr", "show", "dev", iface).Output()
	if err != nil {
		return nil
	}
	var ips []string
	for _, m := range reIP6AddrShow.FindAllStringSubmatch(string(out), -1) {
		if len(m) > 1 && !strings.HasPrefix(m[1], "fe80:") {
			ips = append(ips, m[1])
		}
	}
	return ips
}

func addWGStyleFirewall(iface, table string, ipv6 bool, cgroupPath string) error {
	mark, err := strconv.Atoi(table)
	if err != nil {
		return err
	}
	v4addrs, err := tunIfaceIPv4s(iface)
	if err != nil {
		return err
	}
	v6list := tunIfaceIPv6s(iface)
	hasAntiLeak := len(v4addrs) > 0 || (ipv6 && len(v6list) > 0)
	hasMark := cgroupPath != ""
	if !hasAntiLeak && !hasMark {
		fmt.Fprintln(os.Stderr, i18n.T("redirect.fw_skip"))
		return nil
	}

	nftOK := false
	if _, err := exec.LookPath("nft"); err == nil {
		if err := nftApplyWGAntiLeak(iface, v4addrs, ipv6); err == nil {
			nftOK = true
			fmt.Println(i18n.T("redirect.nft_ok"))
		}
	}
	if !nftOK {
		if err := iptablesRawAntiLeak(iface, v4addrs, v6list, ipv6); err != nil {
			return err
		}
		fmt.Println(i18n.T("redirect.ipt_ok"))
	}

	if cgroupPath != "" {
		if err := iptablesMangleCgroupMark(cgroupPath, mark, iface); err != nil {
			return err
		}
	}
	return nil
}

func nftApplyWGAntiLeak(iface string, v4addrs []string, ipv6 bool) error {
	tab := "z-panel-" + iface
	_ = exec.Command("nft", "delete", "table", "ip", tab).Run()
	var nftcmd strings.Builder
	fmt.Fprintf(&nftcmd, "add table ip %s\n", tab)
	fmt.Fprintf(&nftcmd, "add chain ip %s preraw { type filter hook prerouting priority -300; }\n", tab)
	for _, a := range v4addrs {
		fmt.Fprintf(&nftcmd, "add rule ip %s preraw iifname != \"%s\" ip daddr %s fib saddr type != local drop\n", tab, iface, a)
	}
	if ipv6 {
		tab6 := "z-panel6-" + iface
		v6addrs := tunIfaceIPv6s(iface)
		_ = exec.Command("nft", "delete", "table", "ip6", tab6).Run()
		fmt.Fprintf(&nftcmd, "add table ip6 %s\n", tab6)
		fmt.Fprintf(&nftcmd, "add chain ip6 %s preraw { type filter hook prerouting priority -300; }\n", tab6)
		for _, a := range v6addrs {
			fmt.Fprintf(&nftcmd, "add rule ip6 %s preraw iifname != \"%s\" ip6 daddr %s fib saddr type != local drop\n", tab6, iface, a)
		}
	}
	cmd := exec.Command("nft", "-f", "/dev/stdin")
	cmd.Stdin = strings.NewReader(nftcmd.String())
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func iptablesRawAntiLeak(iface string, v4addrs, v6list []string, ipv6 bool) error {
	marker := fmt.Sprintf("z-panel(8) rule for %s", iface)
	for _, a := range v4addrs {
		if err := run("iptables", "-t", "raw", "-I", "PREROUTING", "!", "-i", iface, "-d", a, "-m", "addrtype", "!", "--src-type", "LOCAL", "-j", "DROP", "-m", "comment", "--comment", marker); err != nil {
			return fmt.Errorf(i18n.T("redirect.iptables_raw_fail"), err)
		}
	}
	if ipv6 {
		if _, err := exec.LookPath("ip6tables"); err != nil {
			return fmt.Errorf("%s", i18n.T("redirect.ip6tables_missing"))
		}
		for _, a := range v6list {
			if err := run("ip6tables", "-t", "raw", "-I", "PREROUTING", "!", "-i", iface, "-d", a, "-m", "addrtype", "!", "--src-type", "LOCAL", "-j", "DROP", "-m", "comment", "--comment", marker); err != nil {
				return fmt.Errorf(i18n.T("redirect.ip6tables_raw_fail"), err)
			}
		}
	}
	return nil
}

func iptablesMangleCgroupMark(cgroupPath string, mark int, iface string) error {
	markS := strconv.Itoa(mark)
	markComment := fmt.Sprintf("z-panel(8) mark for %s", iface)
	if err := run("iptables", "-t", "mangle", "-I", "OUTPUT", "-m", "cgroup", "--path", cgroupPath, "-j", "MARK", "--set-mark", markS, "-m", "comment", "--comment", markComment); err != nil {
		return fmt.Errorf(i18n.T("redirect.iptables_cgroup_fail"), err)
	}
	if _, err := exec.LookPath("ip6tables"); err == nil {
		if err := run("ip6tables", "-t", "mangle", "-I", "OUTPUT", "-m", "cgroup", "--path", cgroupPath, "-j", "MARK", "--set-mark", markS, "-m", "comment", "--comment", markComment); err != nil {
			fmt.Fprintf(os.Stderr, i18n.T("redirect.ip6tables_cgroup_warn"), err)
		}
	}
	fmt.Println(i18n.T("redirect.ipt_cgroup_ok"))
	return nil
}

func removeWGStyleFirewall(iface string) {
	_ = exec.Command("nft", "delete", "table", "ip", "z-panel-"+iface).Run()
	_ = exec.Command("nft", "delete", "table", "ip6", "z-panel6-"+iface).Run()
	removeIptablesZPanelComments(iface)
}

func removeIptablesZPanelComments(iface string) {
	marker := fmt.Sprintf("z-panel(8) rule for %s", iface)
	markM := fmt.Sprintf("z-panel(8) mark for %s", iface)
	for _, ipt := range []string{"iptables", "ip6tables"} {
		if _, err := exec.LookPath(ipt); err != nil {
			continue
		}
		restore := buildIptablesDeleteRestore(ipt, marker, markM)
		if restore == "" {
			continue
		}
		cmd := exec.Command(ipt+"-restore", "-n")
		cmd.Stdin = strings.NewReader(restore)
		_ = cmd.Run()
	}
}

func buildIptablesDeleteRestore(iptables, marker, markM string) string {
	out, err := exec.Command(iptables + "-save").Output()
	if err != nil {
		return ""
	}
	curTable := ""
	var rawD, mangleD []string
	for _, line := range strings.Split(string(out), "\n") {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "*") {
			curTable = strings.TrimPrefix(line, "*")
			continue
		}
		if !strings.HasPrefix(line, "-A ") {
			continue
		}
		if !strings.Contains(line, marker) && !strings.Contains(line, markM) {
			continue
		}
		d := strings.Replace(line, "-A", "-D", 1)
		switch curTable {
		case "raw":
			rawD = append(rawD, d)
		case "mangle":
			mangleD = append(mangleD, d)
		}
	}
	var b strings.Builder
	if len(rawD) > 0 {
		b.WriteString("*raw\n")
		for _, x := range rawD {
			b.WriteString(x + "\n")
		}
		b.WriteString("COMMIT\n")
	}
	if len(mangleD) > 0 {
		b.WriteString("*mangle\n")
		for _, x := range mangleD {
			b.WriteString(x + "\n")
		}
		b.WriteString("COMMIT\n")
	}
	return b.String()
}
