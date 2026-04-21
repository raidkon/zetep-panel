package xraytun

import (
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
	"strings"

	"z-panel/internal/app"
	"z-panel/internal/executil"
	"z-panel/internal/i18n"
	"z-panel/internal/root"
	"z-panel/internal/settings"
)

type Cmd struct{}

func New() *Cmd { return &Cmd{} }

func (c *Cmd) Name() string { return "xray-tun" }

func (c *Cmd) Run(args []string) error {
	if app.IsHelpRequest(args) {
		c.Help(os.Stdout)
		return nil
	}
	if len(args) < 1 {
		return fmt.Errorf("%s", i18n.T("xraytun.want_subcmd"))
	}
	switch args[0] {
	case "up":
		return handleUp(args[1:])
	case "down":
		if len(args) != 3 || args[2] != "ip" {
			return fmt.Errorf("%s", i18n.T("xraytun.want_down"))
		}
		return xrayTunDown(args[1])
	default:
		return fmt.Errorf(i18n.T("xraytun.bad_action"), args[0])
	}
}

func (c *Cmd) Help(w io.Writer) {
	fmt.Fprintf(w, i18n.T("xraytun.help"), settings.C.DefaultXrayAddr, settings.C.DefaultXrayPeer)
}

func (c *Cmd) BashCompletionCase(w io.Writer) {
	fmt.Fprint(w, `	xray-tun)
		if [[ $ecword -eq 2 ]]; then
			mapfile -t COMPREPLY < <(compgen -W 'up down help -h --help' -- "$cur")
		elif [[ ${COMP_WORDS[$((_z_panel_cmd_start+1))]} == up ]]; then
			if [[ $cur == -* ]]; then
				mapfile -t COMPREPLY < <(compgen -W '--address --peer' -- "$cur")
			else
				mapfile -t COMPREPLY < <(compgen -W "$(_z_panel_interfaces) ip" -- "$cur")
			fi
		elif [[ ${COMP_WORDS[$((_z_panel_cmd_start+1))]} == down ]]; then
			if [[ $ecword -eq 3 ]]; then
				mapfile -t COMPREPLY < <(compgen -W "$(_z_panel_interfaces)" -- "$cur")
			elif [[ $ecword -eq 4 ]]; then
				mapfile -t COMPREPLY < <(compgen -W 'ip' -- "$cur")
			fi
		fi
		return
		;;
`)
}

func handleUp(args []string) error {
	var flagAddr, flagPeer string
	var pos []string
	for i := 0; i < len(args); i++ {
		a := args[i]
		switch {
		case strings.HasPrefix(a, "--address="):
			flagAddr = strings.TrimSpace(strings.TrimPrefix(a, "--address="))
		case a == "--address":
			if i+1 >= len(args) {
				return fmt.Errorf("%s", i18n.T("xraytun.need_address_value"))
			}
			i++
			flagAddr = strings.TrimSpace(args[i])
		case strings.HasPrefix(a, "--peer="):
			flagPeer = strings.TrimSpace(strings.TrimPrefix(a, "--peer="))
		case a == "--peer":
			if i+1 >= len(args) {
				return fmt.Errorf("%s", i18n.T("xraytun.need_peer_value"))
			}
			i++
			flagPeer = strings.TrimSpace(args[i])
		case strings.HasPrefix(a, "-"):
			return fmt.Errorf(i18n.T("xraytun.unknown_flag"), a)
		default:
			pos = append(pos, a)
		}
	}
	if len(pos) < 2 {
		return fmt.Errorf("%s", i18n.T("xraytun.want_up_usage"))
	}
	iface := pos[0]
	if pos[1] == "ip" {
		addr := flagAddr
		if addr == "" {
			addr = settings.C.DefaultXrayAddr
		}
		peer := flagPeer
		if peer == "" {
			peer = settings.C.DefaultXrayPeer
		}
		return xrayTunWrite(iface, addr, peer)
	}
	addr := pos[1]
	peerPos := ""
	if len(pos) >= 3 {
		peerPos = pos[2]
	}
	if len(pos) > 3 {
		return fmt.Errorf(i18n.T("xraytun.extra_args"), pos[3:])
	}
	peer := peerPos
	if peer == "" {
		peer = flagPeer
	}
	return xrayTunWrite(iface, addr, peer)
}

func xrayTunNetworkFile(iface string) string {
	return filepath.Join(settings.C.SystemdNetworkDir, iface+".network")
}

func normalizeIPCIDR(s string) (string, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return "", fmt.Errorf("%s", i18n.T("xraytun.empty_address"))
	}
	if strings.Contains(s, "/") {
		ip, ipnet, err := net.ParseCIDR(s)
		if err != nil {
			return "", err
		}
		if ip.To4() == nil {
			return "", fmt.Errorf(i18n.T("xraytun.need_ipv4"), s)
		}
		ones, _ := ipnet.Mask.Size()
		return fmt.Sprintf("%s/%d", ip.To4().String(), ones), nil
	}
	ip := net.ParseIP(s)
	if ip == nil || ip.To4() == nil {
		return "", fmt.Errorf(i18n.T("xraytun.need_ipv4"), s)
	}
	return ip.To4().String() + "/32", nil
}

func xrayTunWrite(iface string, addrStr, peerStr string) error {
	if err := root.Require(); err != nil {
		return err
	}
	if err := sanitizeIfaceName(iface); err != nil {
		return err
	}
	addr, err := normalizeIPCIDR(addrStr)
	if err != nil {
		return fmt.Errorf("address: %w", err)
	}
	peer := ""
	if strings.TrimSpace(peerStr) != "" {
		peer, err = normalizeIPCIDR(peerStr)
		if err != nil {
			return fmt.Errorf("peer: %w", err)
		}
	}
	netBlock := fmt.Sprintf("Address=%s\n", addr)
	if peer != "" {
		netBlock += fmt.Sprintf("Peer=%s\n", peer)
	}
	body := fmt.Sprintf(`%s
# %s
# %s

[Match]
Name=%s

[Network]
%sKeepConfiguration=yes

[Link]
ActivationPolicy=manual
RequiredForOnline=no
`, settings.C.XrayTunManagedMark, i18n.T("xraytun.file_managed_by"), i18n.T("xraytun.file_remove_hint", iface), iface, netBlock)

	path := xrayTunNetworkFile(iface)
	if err := os.MkdirAll(settings.C.SystemdNetworkDir, 0o755); err != nil {
		return fmt.Errorf(i18n.T("settings.err.mkdir"), settings.C.SystemdNetworkDir, err)
	}
	tmp := path + ".tmp"
	if err := os.WriteFile(tmp, []byte(body), 0o644); err != nil {
		return err
	}
	if err := os.Rename(tmp, path); err != nil {
		os.Remove(tmp)
		return fmt.Errorf(i18n.T("settings.err.write"), path, err)
	}
	fmt.Printf(i18n.T("xraytun.wrote"), path)
	return networkdReload()
}

func xrayTunDown(iface string) error {
	if err := root.Require(); err != nil {
		return err
	}
	if err := sanitizeIfaceName(iface); err != nil {
		return err
	}
	path := xrayTunNetworkFile(iface)
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf(i18n.T("xraytun.file_missing"), path)
		}
		return err
	}
	if !strings.Contains(string(data), settings.C.XrayTunManagedMark) {
		return fmt.Errorf(i18n.T("xraytun.not_managed"), path, settings.C.XrayTunManagedMark)
	}
	if err := os.Remove(path); err != nil {
		return err
	}
	fmt.Printf(i18n.T("xraytun.removed"), path)
	return networkdReload()
}

func sanitizeIfaceName(iface string) error {
	if iface == "" {
		return fmt.Errorf("%s", i18n.T("xraytun.empty_iface"))
	}
	for _, r := range iface {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') || r == '-' || r == '_' || r == '.' || r == '@' {
			continue
		}
		return fmt.Errorf(i18n.T("xraytun.bad_iface_char"), iface)
	}
	return nil
}

func networkdReload() error {
	if executil.CommandTTY("networkctl", "reload").Run() == nil {
		fmt.Println(i18n.T("xraytun.ok_networkctl"))
		return nil
	}
	cmd := executil.CommandTTY("systemctl", "reload", "systemd-networkd")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("%s", i18n.T("xraytun.network_reload_fail"))
	}
	fmt.Println(i18n.T("xraytun.ok_systemd"))
	return nil
}
