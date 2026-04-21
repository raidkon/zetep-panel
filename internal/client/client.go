// Package client forwards delegated commands to the z-panel daemon over a Unix socket.
package client

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"strings"
	"syscall"
	"time"

	"z-panel/internal/daemon"
	"z-panel/internal/settings"
)

// ErrDaemonUnavailable means the Unix socket is not accepting connections (daemon stopped).
var ErrDaemonUnavailable = errors.New("z-panel daemon unavailable")

// ExitError carries a subprocess exit code from a delegated run.
type ExitError struct {
	Code int
}

func (e *ExitError) Error() string {
	return fmt.Sprintf("exit status %d", e.Code)
}

// IsUnavailable reports whether err indicates the daemon is not running.
func IsUnavailable(err error) bool {
	return errors.Is(err, ErrDaemonUnavailable)
}

// ExitStatus returns (code, true) if err is *ExitError.
func ExitStatus(err error) (int, bool) {
	var ee *ExitError
	if errors.As(err, &ee) {
		return ee.Code, true
	}
	return 0, false
}

// Forward sends the subcommand argv (without the program name) to the daemon and copies
// stdout/stderr to the process streams.
func Forward(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("client: empty args")
	}
	if settings.C == nil {
		if err := settings.Load(); err != nil {
			return err
		}
	}
	socketPath := settings.C.SocketPath

	reqBody, err := json.Marshal(daemon.RunRequest{Args: args})
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	tr := &http.Transport{
		DialContext: func(ctx context.Context, _, _ string) (net.Conn, error) {
			var d net.Dialer
			return d.DialContext(ctx, "unix", socketPath)
		},
	}
	hc := &http.Client{Transport: tr}
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, "http://localhost/v1/run", bytes.NewReader(reqBody))
	if err != nil {
		return err
	}
	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := hc.Do(httpReq)
	if err != nil {
		if isUnixRefused(err) {
			return ErrDaemonUnavailable
		}
		return err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(io.LimitReader(resp.Body, 64<<20))
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("daemon: HTTP %d: %s", resp.StatusCode, bytes.TrimSpace(body))
	}
	var out daemon.RunResponse
	if err := json.Unmarshal(body, &out); err != nil {
		return fmt.Errorf("daemon: decode response: %w", err)
	}
	_, _ = os.Stdout.WriteString(out.Stdout)
	_, _ = os.Stderr.WriteString(out.Stderr)
	if out.ExitCode != 0 {
		return &ExitError{Code: out.ExitCode}
	}
	return nil
}

func isUnixRefused(err error) bool {
	if err == nil {
		return false
	}
	if errors.Is(err, syscall.ECONNREFUSED) {
		return true
	}
	return strings.Contains(strings.ToLower(err.Error()), "connection refused")
}
