// Package daemon hosts the long-running process: Unix socket API and state watch.
package daemon

// ForbiddenRemote reports whether the subcommand must not be run via the daemon socket
// (recursive control: starting a nested daemon from the daemon).
func ForbiddenRemote(cmdName string) bool {
	return cmdName == "daemon"
}
