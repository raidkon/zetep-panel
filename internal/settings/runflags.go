package settings

// ArgSkipDaemonForward is inserted by the daemon when re-invoking z-panel so the child does not
// delegate back to the Unix socket (avoids recursion). Not intended for interactive use.
const ArgSkipDaemonForward = "--z-panel-skip-daemon-forward"

var skipDaemonForward bool

// SetSkipDaemonForward forces local execution even when daemon=1 in config.toml.
func SetSkipDaemonForward(v bool) {
	skipDaemonForward = v
}

// SkipDaemonForward reports whether daemon delegation is disabled for this process.
func SkipDaemonForward() bool {
	return skipDaemonForward
}

// ResetInternalRunFlagsForTest clears in-process run flags (for tests).
func ResetInternalRunFlagsForTest() {
	skipDaemonForward = false
}

// StripInternalRunFlags removes process-internal flags from argv (after program name).
func StripInternalRunFlags(argv []string) []string {
	if len(argv) < 2 {
		return argv
	}
	out := make([]string, 0, len(argv))
	out = append(out, argv[0])
	for i := 1; i < len(argv); i++ {
		if argv[i] == ArgSkipDaemonForward {
			SetSkipDaemonForward(true)
			continue
		}
		out = append(out, argv[i])
	}
	return out
}
