package daemon

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/signal"
	"syscall"

	"z-panel/internal/app"
	"z-panel/internal/i18n"
	"z-panel/internal/root"
	"z-panel/internal/settings"
	daemonpkg "z-panel/internal/daemon"
)

type Cmd struct{}

func New() *Cmd { return &Cmd{} }

func (c *Cmd) Name() string { return "daemon" }

func (c *Cmd) Run(args []string) error {
	if app.IsHelpRequest(args) {
		c.Help(os.Stdout)
		return nil
	}
	sub := "run"
	if len(args) >= 1 && !app.IsHelpRequest([]string{args[0]}) {
		sub = args[0]
	}
	switch sub {
	case "help", "-h", "--help":
		c.Help(os.Stdout)
		return nil
	case "run":
		if err := root.Require(); err != nil {
			return err
		}
		ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
		defer stop()
		return daemonpkg.RunForeground(ctx)
	default:
		return fmt.Errorf("%s", i18n.T("daemon.err_unknown_subcmd", sub))
	}
}

func (c *Cmd) Help(w io.Writer) {
	fmt.Fprintf(w, i18n.T("daemon.help"), settings.C.SocketPath)
}

func (c *Cmd) BashCompletionCase(w io.Writer) {
	fmt.Fprint(w, `	daemon)
		if [[ $ecword -eq 2 ]]; then
			mapfile -t COMPREPLY < <(compgen -W 'run help -h --help' -- "$cur")
		fi
		return
		;;
`)
}
