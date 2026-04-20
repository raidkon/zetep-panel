package installshell

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"z-panel/internal/app"
	"z-panel/internal/i18n"
	"z-panel/internal/root"
)

const systemCompletionPath = "/usr/share/bash-completion/completions/z-panel"

type Cmd struct{}

func New() *Cmd { return &Cmd{} }

func (c *Cmd) Name() string { return "install-shell" }

// InstallSystem writes the completion script to the system path (requires root).
func InstallSystem() error {
	if err := root.Require(); err != nil {
		return err
	}
	return writeCompletion(systemCompletionPath)
}

func installUserPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf(i18n.T("installshell.err.home"), err)
	}
	if xdg := os.Getenv("XDG_DATA_HOME"); xdg != "" {
		return filepath.Join(xdg, "bash-completion", "completions", "z-panel"), nil
	}
	return filepath.Join(home, ".local", "share", "bash-completion", "completions", "z-panel"), nil
}

func writeCompletion(dest string) error {
	if err := os.MkdirAll(filepath.Dir(dest), 0o755); err != nil {
		return fmt.Errorf(i18n.T("installshell.err.mkdir"), err)
	}
	var buf bytes.Buffer
	app.WriteBashCompletionScript(&buf)
	tmp := dest + ".tmp"
	if err := os.WriteFile(tmp, buf.Bytes(), 0o644); err != nil {
		return err
	}
	if err := os.Rename(tmp, dest); err != nil {
		os.Remove(tmp)
		return fmt.Errorf(i18n.T("installshell.err.write"), dest, err)
	}
	return nil
}

func (c *Cmd) Run(args []string) error {
	if app.IsHelpRequest(args) {
		c.Help(os.Stdout)
		return nil
	}
	userOnly := false
	for _, a := range args {
		switch a {
		case "--user", "-u":
			userOnly = true
		}
	}

	var dest string
	var err error
	if userOnly {
		dest, err = installUserPath()
		if err != nil {
			return err
		}
	} else {
		if err := root.Require(); err != nil {
			return err
		}
		dest = systemCompletionPath
	}

	if err := writeCompletion(dest); err != nil {
		return err
	}

	fmt.Printf(i18n.T("installshell.done"), dest)
	fmt.Println(i18n.T("installshell.hint_shell"))
	if userOnly {
		fmt.Println(i18n.T("installshell.hint_user"))
	}
	return nil
}

func (c *Cmd) Help(w io.Writer) {
	fmt.Fprintf(w, i18n.T("installshell.help"), systemCompletionPath)
}

func (c *Cmd) BashCompletionCase(w io.Writer) {
	fmt.Fprint(w, `	install-shell)
		mapfile -t COMPREPLY < <(compgen -W '--user -u help -h --help' -- "$cur")
		return
		;;
`)
}
