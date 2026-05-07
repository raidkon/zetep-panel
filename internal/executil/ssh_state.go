package executil

import "sync"

var sshMu sync.RWMutex

var (
	sshHost        string
	sshControlPath string // OpenSSH ControlPath (-S); set while ControlMaster is active
	sshNoTTY       bool   // omit ssh -t (e.g. NOPASSWD sudo)
)

// SetRemoteSSHHost sets the SSH target for z-panel --ssh=… mode (local binary, remote commands).
// Pass empty to clear.
func SetRemoteSSHHost(h string) {
	sshMu.Lock()
	defer sshMu.Unlock()
	sshHost = h
}

// RemoteSSHHost returns the SSH target when running with --ssh, or empty locally / with --ssh-connect.
func RemoteSSHHost() string {
	sshMu.RLock()
	defer sshMu.RUnlock()
	return sshHost
}

// SetSSHNoTTY configures whether ssh uses a TTY (-t). When true, -t is omitted.
func SetSSHNoTTY(v bool) {
	sshMu.Lock()
	defer sshMu.Unlock()
	sshNoTTY = v
}

func sshWantTTY() bool {
	sshMu.RLock()
	defer sshMu.RUnlock()
	return !sshNoTTY
}

func setSSHControlPath(p string) {
	sshMu.Lock()
	defer sshMu.Unlock()
	sshControlPath = p
}

func sshControlPathForDial() string {
	sshMu.RLock()
	defer sshMu.RUnlock()
	return sshControlPath
}

// ResetSSHForTests clears in-process SSH routing state.
func ResetSSHForTests() {
	sshMu.Lock()
	defer sshMu.Unlock()
	sshHost = ""
	sshControlPath = ""
	sshNoTTY = false
}
