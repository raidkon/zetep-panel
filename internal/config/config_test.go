package config

import "testing"

func TestConstants_nonEmpty(t *testing.T) {
	if Version == "" || InstallPath == "" || ConfigDir == "" || ConfigFile == "" {
		t.Fatalf("Version=%q InstallPath=%q ConfigDir=%q ConfigFile=%q", Version, InstallPath, ConfigDir, ConfigFile)
	}
	if DefaultSocketPath == "" || DefaultPidPath == "" {
		t.Fatalf("DefaultSocketPath=%q DefaultPidPath=%q", DefaultSocketPath, DefaultPidPath)
	}
}
