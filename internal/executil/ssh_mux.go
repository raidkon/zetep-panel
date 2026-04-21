package executil

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

// EnvSSHMux is the control socket path (OpenSSH -S) for an active ControlMaster; set while the multiplex is alive.
const EnvSSHMux = "Z_PANEL_SSH_MUX"

// EnvSSHNoMux, if non-empty, skips starting ControlMaster (one new TCP connection per ssh invocation).
const EnvSSHNoMux = "Z_PANEL_SSH_NO_MUX"

// TryStartSSHMultiplex starts an SSH ControlMaster (-fN) so all later ssh calls in this process can use
// ssh -S socket … on the same host (one TCP session, less handshake). If OpenSSH is absent or the
// master fails to start, returns (nil, nil) — callers should proceed without multiplex.
func TryStartSSHMultiplex(host string) (stop func(), _ error) {
	if host == "" || os.Getenv(EnvSSHNoMux) != "" {
		return nil, nil
	}
	dir, err := os.MkdirTemp("", "z-panel-ssh-mux-*")
	if err != nil {
		return nil, err
	}
	sock := filepath.Join(dir, "control.sock")
	cmd := exec.Command("ssh",
		"-fN",
		"-o", "ControlMaster=yes",
		"-o", "ControlPersist=300",
		"-o", "ControlPath="+sock,
		host,
	)
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		_ = os.RemoveAll(dir)
		fmt.Fprintf(os.Stderr, "z-panel: ssh multiplex unavailable: %v\n", err)
		return nil, nil
	}
	deadline := time.Now().Add(3 * time.Second)
	for time.Now().Before(deadline) {
		if _, err := os.Stat(sock); err == nil {
			break
		}
		time.Sleep(10 * time.Millisecond)
	}
	if _, err := os.Stat(sock); err != nil {
		_ = os.RemoveAll(dir)
		fmt.Fprintf(os.Stderr, "z-panel: ssh multiplex: control socket did not appear\n")
		return nil, nil
	}
	_ = os.Setenv(EnvSSHMux, sock)
	return func() {
		_ = os.Unsetenv(EnvSSHMux)
		c := exec.Command("ssh", "-o", "ControlPath="+sock, "-O", "exit", host)
		c.Stderr = io.Discard
		_ = c.Run()
		_ = os.RemoveAll(dir)
	}, nil
}
