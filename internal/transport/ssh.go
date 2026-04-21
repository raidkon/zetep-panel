// Package transport implements optional remote execution (SSH) for z-panel.
package transport

import (
	"fmt"
	"os/exec"
	"strings"

	"z-panel/internal/config"
	"z-panel/internal/executil"
	"z-panel/internal/i18n"
)

// RemoteMode selects how z-panel talks to the remote host.
type RemoteMode int

const (
	// RemoteNone: normal local execution (no --ssh / --ssh-connect).
	RemoteNone RemoteMode = iota
	// RemoteLocalTools: --ssh — local z-panel binary; subcommands run system tools on the remote host via ssh+sudo (no remote z-panel install).
	RemoteLocalTools
	// RemoteZPanelBinary: --ssh-connect — run the remote installed z-panel (e.g. daemon on that host).
	RemoteZPanelBinary
)

// ParseSSHFromArgs extracts at most one of --ssh or --ssh-connect (with host) and returns the remaining
// arguments (including argv[0] as the program path). Mixing both flags is an error.
func ParseSSHFromArgs(args []string) (mode RemoteMode, host string, rest []string, err error) {
	if len(args) < 1 {
		return RemoteNone, "", nil, fmt.Errorf("%s", i18n.T("transport.ssh.err_argv"))
	}
	rest = append(rest, args[0])
	var modeSeen RemoteMode
	for i := 1; i < len(args); i++ {
		a := args[i]
		switch {
		case strings.HasPrefix(a, "--ssh="):
			if modeSeen != RemoteNone {
				return RemoteNone, "", nil, fmt.Errorf("%s", i18n.T("transport.ssh.err_conflict"))
			}
			modeSeen = RemoteLocalTools
			host = strings.TrimSpace(strings.TrimPrefix(a, "--ssh="))
			if host == "" {
				return RemoteNone, "", nil, fmt.Errorf("%s", i18n.T("transport.ssh.err_empty"))
			}
		case a == "--ssh":
			if modeSeen != RemoteNone {
				return RemoteNone, "", nil, fmt.Errorf("%s", i18n.T("transport.ssh.err_conflict"))
			}
			modeSeen = RemoteLocalTools
			if i+1 >= len(args) {
				return RemoteNone, "", nil, fmt.Errorf("%s", i18n.T("transport.ssh.err_missing"))
			}
			host = strings.TrimSpace(args[i+1])
			if host == "" {
				return RemoteNone, "", nil, fmt.Errorf("%s", i18n.T("transport.ssh.err_empty"))
			}
			i++
		case strings.HasPrefix(a, "--ssh-connect="):
			if modeSeen != RemoteNone {
				return RemoteNone, "", nil, fmt.Errorf("%s", i18n.T("transport.ssh.err_conflict"))
			}
			modeSeen = RemoteZPanelBinary
			host = strings.TrimSpace(strings.TrimPrefix(a, "--ssh-connect="))
			if host == "" {
				return RemoteNone, "", nil, fmt.Errorf("%s", i18n.T("transport.ssh.err_empty_connect"))
			}
		case a == "--ssh-connect":
			if modeSeen != RemoteNone {
				return RemoteNone, "", nil, fmt.Errorf("%s", i18n.T("transport.ssh.err_conflict"))
			}
			modeSeen = RemoteZPanelBinary
			if i+1 >= len(args) {
				return RemoteNone, "", nil, fmt.Errorf("%s", i18n.T("transport.ssh.err_missing_connect"))
			}
			host = strings.TrimSpace(args[i+1])
			if host == "" {
				return RemoteNone, "", nil, fmt.Errorf("%s", i18n.T("transport.ssh.err_empty_connect"))
			}
			i++
		default:
			rest = append(rest, a)
		}
	}
	return modeSeen, host, rest, nil
}

// RunZPanelOverSSH runs ssh -t host sudo <install-path> <rest of argv after program name>.
// Uses config.InstallPath so sudo finds the binary (sudo secure_path often omits /usr/local/bin).
// Used for --ssh-connect when the remote machine has z-panel installed.
func RunZPanelOverSSH(host string, argv []string) error {
	if len(argv) < 2 {
		return fmt.Errorf("%s", i18n.T("transport.ssh.err_no_cmd"))
	}
	inner := argv[1:]
	args := append([]string{"-t", host, "sudo", config.InstallPath}, inner...)
	cmd := exec.Command("ssh", args...)
	return executil.RunAttachedInterruptible(cmd)
}
