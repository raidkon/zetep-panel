package daemon

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"

	"z-panel/internal/config"
	"z-panel/internal/settings"
)

// Serve listens on the configured Unix socket and runs until ctx is cancelled.
func Serve(ctx context.Context, socketPath string) error {
	if err := os.MkdirAll(filepath.Dir(socketPath), 0o755); err != nil {
		return fmt.Errorf("mkdir socket dir: %w", err)
	}
	_ = os.Remove(socketPath)
	ln, err := net.Listen("unix", socketPath)
	if err != nil {
		return err
	}
	if err := os.Chmod(socketPath, 0o600); err != nil {
		_ = ln.Close()
		return fmt.Errorf("chmod socket: %w", err)
	}

	var runMu sync.Mutex
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/health", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, `{"ok":true}`+"\n")
	})
	mux.HandleFunc("/v1/run", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(w, `{"error":"Content-Type must be application/json"}`, http.StatusBadRequest)
			return
		}
		var req RunRequest
		if err := json.NewDecoder(io.LimitReader(r.Body, 1<<20)).Decode(&req); err != nil {
			http.Error(w, fmt.Sprintf(`{"error":%q}`, err.Error()), http.StatusBadRequest)
			return
		}
		if len(req.Args) < 1 {
			http.Error(w, `{"error":"args required"}`, http.StatusBadRequest)
			return
		}
		if ForbiddenRemote(req.Args[0]) {
			http.Error(w, `{"error":"command forbidden for remote run"}`, http.StatusBadRequest)
			return
		}

		runMu.Lock()
		defer runMu.Unlock()

		exe, err := os.Executable()
		if err != nil {
			http.Error(w, fmt.Sprintf(`{"error":%q}`, err.Error()), http.StatusInternalServerError)
			return
		}
		cmd := exec.Command(exe, append([]string{settings.ArgSkipDaemonForward}, req.Args...)...)
		cmd.Env = os.Environ()
		var stdout, stderr bytes.Buffer
		cmd.Stdout = &stdout
		cmd.Stderr = &stderr
		runErr := cmd.Run()
		exit := 0
		if runErr != nil {
			var ee *exec.ExitError
			if errors.As(runErr, &ee) {
				exit = ee.ExitCode()
			} else {
				http.Error(w, fmt.Sprintf(`{"error":%q}`, runErr.Error()), http.StatusInternalServerError)
				return
			}
		}
		resp := RunResponse{ExitCode: exit, Stdout: stdout.String(), Stderr: stderr.String()}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(resp)
	})

	srv := &http.Server{Handler: mux}
	errCh := make(chan error, 1)
	go func() {
		errCh <- srv.Serve(ln)
	}()

	go func() {
		<-ctx.Done()
		_ = srv.Shutdown(context.Background())
	}()

	if err := <-errCh; err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

// RunForeground loads settings, writes the PID file, starts the watch loop, and serves the socket until ctx ends.
func RunForeground(ctx context.Context) error {
	if settings.C == nil {
		if err := settings.Load(); err != nil {
			return err
		}
	}
	sock := settings.C.SocketPath
	pidPath := strings.TrimSpace(settings.C.PidPath)
	if pidPath == "" {
		pidPath = config.DefaultPidPath
	}
	if err := os.MkdirAll(filepath.Dir(pidPath), 0o755); err != nil {
		return fmt.Errorf("mkdir pid dir: %w", err)
	}
	if err := os.WriteFile(pidPath, []byte(fmt.Sprintf("%d\n", os.Getpid())), 0o644); err != nil {
		return fmt.Errorf("write pid file: %w", err)
	}
	defer func() { _ = os.Remove(pidPath) }()

	log.SetPrefix("z-panel: ")
	log.Printf("daemon listening on %s (pid file %s)", sock, pidPath)
	wctx, cancel := context.WithCancel(ctx)
	defer cancel()
	go Loop(wctx)
	return Serve(ctx, sock)
}
