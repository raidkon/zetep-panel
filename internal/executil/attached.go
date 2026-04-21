// Package executil runs subprocesses attached to the terminal (stdin/stdout/stderr).
package executil

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"

	"z-panel/internal/i18n"
)

// RunAttachedInterruptible runs cmd attached to the terminal; on Ctrl+C/SIGTERM kills the child
// (remote sudo often swallows SIGINT).
func RunAttachedInterruptible(cmd *exec.Cmd) error {
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		return err
	}
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(sigCh)
	waitCh := make(chan error, 1)
	go func() { waitCh <- cmd.Wait() }()
	select {
	case err := <-waitCh:
		return err
	case <-sigCh:
		if cmd.Process != nil {
			_ = cmd.Process.Kill()
		}
		err := <-waitCh
		if err != nil {
			return fmt.Errorf(i18n.T("install.err.interrupted_with"), err)
		}
		return fmt.Errorf("%s", i18n.T("install.err.interrupted"))
	}
}
