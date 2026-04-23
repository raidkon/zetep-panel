// Package config holds version and install paths (not from TOML).
package config

// Version is the user-visible program version.
// Default is the last tagged release; bump it when you cut a release.
// Local builds: use `make` or `go build -ldflags "-X z-panel/internal/config.Version=$(git describe --tags --always --dirty)"`
// so the binary reflects the exact commit (and dirty state).
var Version = "0.14.1"

const (
	InstallPath = "/usr/local/bin/z-panel"
	ConfigDir   = "/etc/z-panel"
	ConfigFile  = "/etc/z-panel/config.toml"

	// DefaultSocketPath is the Unix socket for the z-panel daemon (HTTP API + command execution).
	DefaultSocketPath = "/run/z-panel/z-panel.sock"
)
