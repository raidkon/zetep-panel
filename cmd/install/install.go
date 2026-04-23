package install

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"z-panel/cmd/installshell"
	"z-panel/internal/app"
	"z-panel/internal/config"
	"z-panel/internal/executil"
	"z-panel/internal/i18n"
	"z-panel/internal/root"
	"z-panel/internal/settings"
)

type Cmd struct{}

func New() *Cmd { return &Cmd{} }

func (c *Cmd) Name() string { return "install" }

func (c *Cmd) Run(args []string) error {
	if app.IsHelpRequest(args) {
		c.Help(os.Stdout)
		return nil
	}
	if executil.RemoteSSHHost() != "" {
		if len(args) >= 1 {
			return fmt.Errorf("%s", i18n.T("install.err.extra_with_ssh", args[0]))
		}
		return installOverSSH()
	}
	if len(args) >= 1 {
		return fmt.Errorf("%s", i18n.T("install.err_remote_removed", args[0]))
	}
	return localInstall()
}

func (c *Cmd) Help(w io.Writer) {
	fmt.Fprintf(w, i18n.T("install.help"), config.InstallPath, config.ConfigFile)
}

func (c *Cmd) BashCompletionCase(w io.Writer) {
	fmt.Fprint(w, `	install)
		return
		;;
`)
}

func parseZPanelVersionOutput(output string) string {
	for _, line := range strings.Split(output, "\n") {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "z-panel ") {
			return strings.TrimSpace(strings.TrimPrefix(line, "z-panel "))
		}
	}
	return ""
}

// installedBinaryVersion runs the given z-panel binary and parses the first "z-panel x.y.z" line from combined output.
func installedBinaryVersion(bin string) string {
	st, err := os.Stat(bin)
	if err != nil || st.IsDir() {
		return ""
	}
	if st.Mode()&0111 == 0 {
		return ""
	}
	out, err := exec.Command(bin, "version").CombinedOutput()
	if err != nil {
		return ""
	}
	return parseZPanelVersionOutput(string(out))
}

func localInstall() error {
	if err := root.Require(); err != nil {
		return err
	}
	oldVer := installedBinaryVersion(config.InstallPath)
	self, err := os.Executable()
	if err != nil {
		return fmt.Errorf("os.Executable: %w", err)
	}
	self, err = filepath.EvalSymlinks(self)
	if err != nil {
		return fmt.Errorf("EvalSymlinks: %w", err)
	}
	in, err := os.Open(self)
	if err != nil {
		return fmt.Errorf(i18n.T("install.err.open_self"), err)
	}
	defer in.Close()
	tmp := config.InstallPath + ".new"
	out, err := os.OpenFile(tmp, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0o755)
	if err != nil {
		return fmt.Errorf(i18n.T("install.err.create_tmp"), tmp, err)
	}
	if _, err := io.Copy(out, in); err != nil {
		out.Close()
		os.Remove(tmp)
		return fmt.Errorf(i18n.T("install.err.copy"), err)
	}
	if err := out.Close(); err != nil {
		os.Remove(tmp)
		return err
	}
	if err := os.Rename(tmp, config.InstallPath); err != nil {
		os.Remove(tmp)
		return fmt.Errorf(i18n.T("install.err.rename"), config.InstallPath, err)
	}
	_ = os.Chmod(config.InstallPath, 0o755)
	fmt.Printf(i18n.T("install.installed"), config.InstallPath)
	fmt.Printf(i18n.T("install.new_version"), config.Version)
	if oldVer != "" {
		fmt.Printf(i18n.T("install.old_version"), oldVer)
	}
	if err := settings.InitInteractive(os.Stdin, os.Stdout, false); err != nil {
		return fmt.Errorf(i18n.T("install.err.config"), err)
	}
	return installshell.InstallSystem()
}

func installOverSSH() error {
	host := strings.TrimSpace(executil.RemoteSSHHost())
	if _, err := exec.LookPath("scp"); err != nil {
		return fmt.Errorf("%s", i18n.T("install.err.need_scp"))
	}
	if _, err := exec.LookPath("ssh"); err != nil {
		return fmt.Errorf("%s", i18n.T("install.err.need_ssh"))
	}
	self, err := os.Executable()
	if err != nil {
		return fmt.Errorf("os.Executable: %w", err)
	}
	self, err = filepath.EvalSymlinks(self)
	if err != nil {
		return fmt.Errorf("EvalSymlinks: %w", err)
	}
	b := make([]byte, 8)
	if _, err := rand.Read(b); err != nil {
		return err
	}
	remote := "/tmp/z-panel-stg-" + hex.EncodeToString(b)
	fmt.Fprintf(os.Stderr, i18n.T("install.ssh.uploading"), host, remote)
	scp := exec.Command("scp", self, host+":"+remote)
	if err := executil.RunAttachedInterruptible(scp); err != nil {
		return fmt.Errorf(i18n.T("install.err.scp"), err)
	}
	defer func() { _ = exec.Command("ssh", host, "rm", "-f", remote).Run() }()
	script := "chmod 755 " + remote + " && sudo " + remote + " install"
	cmd := exec.Command("ssh", "-t", host, script)
	if err := executil.RunAttachedInterruptible(cmd); err != nil {
		return fmt.Errorf(i18n.T("install.err.ssh_run"), err)
	}
	return nil
}
