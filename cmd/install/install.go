package install

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"strconv"
	"syscall"

	"z-panel/cmd/installshell"
	"z-panel/internal/app"
	"z-panel/internal/config"
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
	if len(args) >= 1 {
		return remoteInstall(args[0])
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

// runAttachedInterruptible runs a command attached to the terminal; on Ctrl+C/SIGTERM kills the child
// (remote sudo often swallows SIGINT).
func runAttachedInterruptible(cmd *exec.Cmd) error {
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		return err
	}
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(sigCh)
	waitCh := make(chan error, 1)
	go func() { waitCh <- cmd.Wait() }()
	select {
	case err := <-waitCh:
		return err
	case <-sigCh:
		if cmd.Process != nil {
			_ = cmd.Process.Kill()
		}
		err := <-waitCh
		if err != nil {
			return fmt.Errorf(i18n.T("install.err.interrupted_with"), err)
		}
		return fmt.Errorf("%s", i18n.T("install.err.interrupted"))
	}
}

func localInstall() error {
	if err := root.Require(); err != nil {
		return err
	}
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
	// Target file mode (rename keeps mode of tmp)
	_ = os.Chmod(config.InstallPath, 0o755)
	fmt.Printf(i18n.T("install.installed"), config.InstallPath)
	if err := settings.InitInteractive(os.Stdin, os.Stdout, false); err != nil {
		return fmt.Errorf(i18n.T("install.err.config"), err)
	}
	if err := installshell.InstallSystem(); err != nil {
		fmt.Fprintf(os.Stderr, i18n.T("install.warn_completion"), err)
	}
	return nil
}

func remoteInstall(sshHost string) error {
	self, err := os.Executable()
	if err != nil {
		return err
	}
	self, err = filepath.EvalSymlinks(self)
	if err != nil {
		return err
	}
	remoteTmp := "/tmp/z-panel-install-" + strconv.Itoa(os.Getpid())
	scp := exec.Command("scp", "-C", self, sshHost+":"+remoteTmp)
	if err := runAttachedInterruptible(scp); err != nil {
		return fmt.Errorf(i18n.T("install.err.scp"), err)
	}
	// One session: install + interactive config init if config.toml missing (needs -t).
	remote := fmt.Sprintf(
		`sudo install -m 755 %s %s && rm -f %s && { if [ -f %s ]; then sudo z-panel config migrate; else sudo z-panel config init; fi; } && sudo z-panel install-shell`,
		remoteTmp, config.InstallPath, remoteTmp, config.ConfigFile,
	)
	ssh := exec.Command("ssh", "-t", sshHost, remote)
	if err := runAttachedInterruptible(ssh); err != nil {
		return fmt.Errorf(i18n.T("install.err.ssh"), err)
	}
	fmt.Printf(i18n.T("install.remote_done"), sshHost, config.InstallPath)
	return nil
}
