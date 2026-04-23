package client

import (
	"encoding/json"
	"errors"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"syscall"
	"testing"

	"z-panel/internal/daemon"
	"z-panel/internal/settings"
)

func TestExitError_Error(t *testing.T) {
	e := &ExitError{Code: 42}
	if e.Error() == "" {
		t.Fatal()
	}
}

func TestIsUnavailable(t *testing.T) {
	if !IsUnavailable(ErrDaemonUnavailable) {
		t.Fatal()
	}
	if IsUnavailable(errors.New("other")) {
		t.Fatal()
	}
}

func TestExitStatus(t *testing.T) {
	c, ok := ExitStatus(&ExitError{Code: 3})
	if !ok || c != 3 {
		t.Fatal()
	}
	c, ok = ExitStatus(errors.New("n"))
	if ok {
		t.Fatal()
	}
	_ = c
}

func TestForward_emptyArgs(t *testing.T) {
	err := Forward(nil)
	if err == nil {
		t.Fatal()
	}
}

func TestForward_daemonHTTP200(t *testing.T) {
	dir := t.TempDir()
	sock := filepath.Join(dir, "t.sock")
	ln, err := net.Listen("unix", sock)
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = ln.Close(); _ = os.Remove(sock) }()

	mux := http.NewServeMux()
	mux.HandleFunc("/v1/run", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(daemon.RunResponse{ExitCode: 0, Stdout: "ok\n"})
	})
	srv := &http.Server{Handler: mux}
	go func() { _ = srv.Serve(ln) }()
	defer func() { _ = srv.Close() }()

	t.Cleanup(func() { settings.C = nil })
	settings.ApplyDefaults()
	settings.C.SocketPath = sock

	if err := Forward([]string{"version"}); err != nil {
		t.Fatal(err)
	}
}

func TestForward_daemonHTTPErrorBody(t *testing.T) {
	dir := t.TempDir()
	sock := filepath.Join(dir, "e.sock")
	ln, err := net.Listen("unix", sock)
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = ln.Close(); _ = os.Remove(sock) }()
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/run", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "bad", http.StatusBadRequest)
	})
	srv := &http.Server{Handler: mux}
	go func() { _ = srv.Serve(ln) }()
	defer func() { _ = srv.Close() }()
	t.Cleanup(func() { settings.C = nil })
	settings.ApplyDefaults()
	settings.C.SocketPath = sock
	if err := Forward([]string{"x"}); err == nil {
		t.Fatal("expected error")
	}
}

func TestForward_decodeError(t *testing.T) {
	dir := t.TempDir()
	sock := filepath.Join(dir, "d.sock")
	ln, err := net.Listen("unix", sock)
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = ln.Close(); _ = os.Remove(sock) }()
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/run", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`not json`))
	})
	srv := &http.Server{Handler: mux}
	go func() { _ = srv.Serve(ln) }()
	defer func() { _ = srv.Close() }()
	t.Cleanup(func() { settings.C = nil })
	settings.ApplyDefaults()
	settings.C.SocketPath = sock
	if err := Forward([]string{"version"}); err == nil {
		t.Fatal()
	}
}

func TestForward_exitError(t *testing.T) {
	dir := t.TempDir()
	sock := filepath.Join(dir, "x.sock")
	ln, err := net.Listen("unix", sock)
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = ln.Close(); _ = os.Remove(sock) }()
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/run", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(daemon.RunResponse{ExitCode: 5})
	})
	srv := &http.Server{Handler: mux}
	go func() { _ = srv.Serve(ln) }()
	defer func() { _ = srv.Close() }()
	t.Cleanup(func() { settings.C = nil })
	settings.ApplyDefaults()
	settings.C.SocketPath = sock
	fwdErr := Forward([]string{"version"})
	var ee *ExitError
	if !errors.As(fwdErr, &ee) || ee.Code != 5 {
		t.Fatalf("err=%v", fwdErr)
	}
}

func TestIsUnixRefused_econnrefused(t *testing.T) {
	err := syscall.ECONNREFUSED
	if !isUnixRefused(err) {
		t.Fatal()
	}
}

func TestIsUnixRefused_message(t *testing.T) {
	if !isUnixRefused(errors.New("dial unix: connection refused")) {
		t.Fatal()
	}
}

func TestIsUnixRefused_nil(t *testing.T) {
	if isUnixRefused(nil) {
		t.Fatal()
	}
}

func TestIsUnixRefused_opError(t *testing.T) {
	op := &net.OpError{Err: syscall.ECONNREFUSED}
	if !isUnixRefused(op) {
		t.Fatal()
	}
}
