package redirect

import (
	"os"
	"path/filepath"
	"testing"

	"z-panel/internal/executil"
)

func TestIpLinkExists_fakeIP(t *testing.T) {
	t.Cleanup(func() { _ = os.Unsetenv(executil.EnvSSHHost) })
	_ = os.Unsetenv(executil.EnvSSHHost)
	dir := t.TempDir()
	ip := filepath.Join(dir, "ip")
	script := "#!/bin/sh\nif [ \"$1\" = link ]; then echo ok; fi\nexit 0\n"
	if err := os.WriteFile(ip, []byte(script), 0o755); err != nil {
		t.Fatal(err)
	}
	t.Setenv("PATH", dir)
	if err := ipLinkExists("eth0"); err != nil {
		t.Fatal(err)
	}
}

func TestIpLinkExists_missingIface(t *testing.T) {
	t.Cleanup(func() { _ = os.Unsetenv(executil.EnvSSHHost) })
	dir := t.TempDir()
	ip := filepath.Join(dir, "ip")
	if err := os.WriteFile(ip, []byte("#!/bin/sh\nexit 1\n"), 0o755); err != nil {
		t.Fatal(err)
	}
	t.Setenv("PATH", dir)
	if err := ipLinkExists("nope"); err == nil {
		t.Fatal()
	}
}

func TestRun_smoke(t *testing.T) {
	t.Cleanup(func() { _ = os.Unsetenv(executil.EnvSSHHost) })
	if err := run("true"); err != nil {
		t.Fatal(err)
	}
}

func TestIPRuleShow(t *testing.T) {
	s := ipRuleShow("ipv4")
	_ = s
}
