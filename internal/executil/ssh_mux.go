package executil

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

// TryStartSSHMultiplex starts an SSH ControlMaster (-fN) so all later ssh calls in this process can use
// ssh -S socket … on the same host (one TCP session, less handshake). If OpenSSH is absent or the
// master fails to start, returns (nil, nil) — callers should proceed without multiplex.
// When disableMux is true (e.g. ssh_no_multiplex in config), multiplex is not started.
func TryStartSSHMultiplex(host string, disableMux bool) (stop func(), _ error) {
	if host == "" || disableMux {
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
	setSSHControlPath(sock)
	return func() {
		setSSHControlPath("")
		c := exec.Command("ssh", "-o", "ControlPath="+sock, "-O", "exit", host)
		c.Stderr = io.Discard
		_ = c.Run()
		_ = os.RemoveAll(dir)
	}, nil
}
