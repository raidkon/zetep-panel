package executil

import (
	"bytes"
	"io"
	"os"
	"os/exec"
	"strings"
)

// EnvSSHHost is set by main for: z-panel --ssh=host <subcommand> … (local binary; tools on remote via ssh+sudo).
// Not set for --ssh-connect (remote z-panel binary over ssh).
// Child commands use Command / CommandTTY to run tools on the remote host without a remote z-panel install.
const EnvSSHHost = "Z_PANEL_SSH_HOST"

// EnvSSHNoTTY, if set, omits ssh -t (NOPASSWD sudo, non-interactive).
const EnvSSHNoTTY = "Z_PANEL_SSH_NO_TTY"

// RemoteSSHHost returns the SSH target when running with --ssh, or empty for local or --ssh-connect execution.
func RemoteSSHHost() string {
	return os.Getenv(EnvSSHHost)
}

// shellSingleQuote wraps s in POSIX single quotes; embedded ' become '\''.
func shellSingleQuote(s string) string {
	if !strings.Contains(s, "'") {
		return "'" + s + "'"
	}
	var b strings.Builder
	b.WriteByte('\'')
	for _, r := range s {
		if r == '\'' {
			b.WriteString(`'"'"'`)
		} else {
			b.WriteRune(r)
		}
	}
	b.WriteByte('\'')
	return b.String()
}

// sshRemote builds: ssh [-S mux] [-t] host remoteWords...
// OpenSSH joins each argv after host with spaces into one exec line — do not pass a multiword -c script as separate argv.
func sshRemote(host string, withTTY bool, remoteWords ...string) *exec.Cmd {
	var args []string
	if mux := os.Getenv(EnvSSHMux); mux != "" {
		args = append(args, "-S", mux)
	}
	if withTTY && os.Getenv(EnvSSHNoTTY) == "" {
		args = append(args, "-t")
	}
	args = append(args, host)
	args = append(args, remoteWords...)
	return exec.Command("ssh", args...)
}

// sshRemoteQuoted runs: ssh … host "one remote shell command line" (one argv after host).
func sshRemoteQuoted(host string, withTTY bool, remoteLine string) *exec.Cmd {
	var args []string
	if mux := os.Getenv(EnvSSHMux); mux != "" {
		args = append(args, "-S", mux)
	}
	if withTTY && os.Getenv(EnvSSHNoTTY) == "" {
		args = append(args, "-t")
	}
	args = append(args, host, remoteLine)
	return exec.Command("ssh", args...)
}

// Command returns exec.Command(name, args...) locally, or ssh host sudo name args… on the remote host (no -T).
// Use when stdin will be piped (e.g. nft -f /dev/stdin) or for non-interactive sudo (NOPASSWD).
func Command(name string, arg ...string) *exec.Cmd {
	if h := RemoteSSHHost(); h != "" {
		words := append([]string{"sudo", name}, arg...)
		return sshRemote(h, false, words...)
	}
	return exec.Command(name, arg...)
}

// CommandTTY is like Command but uses ssh -t on the remote host so sudo can use a TTY for a password.
// Without -t, sudo often prints "A terminal is required to authenticate" and exits 1.
// If sudo is NOPASSWD for these commands, set Z_PANEL_SSH_NO_TTY=1 to omit -t.
func CommandTTY(name string, arg ...string) *exec.Cmd {
	if h := RemoteSSHHost(); h != "" {
		words := append([]string{"sudo", name}, arg...)
		cmd := sshRemote(h, true, words...)
		cmd.Stdin = os.Stdin
		return cmd
	}
	return exec.Command(name, arg...)
}

// RunTTYCombined runs CommandTTY and returns captured stdout+stderr. When Z_PANEL_SSH_HOST
// is set, output is also copied to the terminal: CombinedOutput() would hide sudo/ssh
// prompts and block the password on stdin, so we MultiWriter to os.Stdout/os.Stderr.
func RunTTYCombined(name string, arg ...string) ([]byte, error) {
	cmd := CommandTTY(name, arg...)
	if RemoteSSHHost() == "" {
		return cmd.CombinedOutput()
	}
	var buf bytes.Buffer
	cmd.Stdout = io.MultiWriter(&buf, os.Stdout)
	cmd.Stderr = io.MultiWriter(&buf, os.Stderr)
	cmd.Stdin = os.Stdin
	err := cmd.Run()
	return buf.Bytes(), err
}

// RunTTYCombinedScript runs sh -c script locally, or one ssh with sudo sh -c script remotely
// (one sudo for the whole script). With ControlMaster, reuse the same TCP connection.
func RunTTYCombinedScript(script string) ([]byte, error) {
	if RemoteSSHHost() == "" {
		return exec.Command("sh", "-c", script).CombinedOutput()
	}
	h := RemoteSSHHost()
	withTTY := os.Getenv(EnvSSHNoTTY) == ""
	// One ssh argv: sudo sh -c '…' — not ssh … sudo sh -c <unquoted>, or OpenSSH joins argv with spaces and breaks -c.
	remoteLine := "sudo sh -c " + shellSingleQuote(script)
	cmd := sshRemoteQuoted(h, withTTY, remoteLine)
	cmd.Stdin = os.Stdin
	var buf bytes.Buffer
	cmd.Stdout = io.MultiWriter(&buf, os.Stdout)
	cmd.Stderr = io.MultiWriter(&buf, os.Stderr)
	err := cmd.Run()
	return buf.Bytes(), err
}
