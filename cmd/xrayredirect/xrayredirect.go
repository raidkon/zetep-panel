package xrayredirect

import (
	"fmt"
	"io"
	"os"

	"z-panel/internal/app"
	"z-panel/internal/i18n"
	"z-panel/internal/redirect"
	"z-panel/internal/settings"
)

type Cmd struct{}

func New() *Cmd { return &Cmd{} }

func (c *Cmd) Name() string { return "xray-redirect" }

func (c *Cmd) Run(args []string) error {
	if app.IsHelpRequest(args) {
		c.Help(os.Stdout)
		return nil
	}
	if len(args) < 1 {
		return fmt.Errorf("%s", i18n.T("xrayredirect.want_up_down"))
	}
	switch args[0] {
	case "up":
		iface, opts, err := redirect.ParseUpArgs(args[1:])
		if err != nil {
			return err
		}
		return redirect.Up(iface, opts)
	case "down":
		if len(args) != 2 {
			return fmt.Errorf("%s", i18n.T("xrayredirect.want_down_iface"))
		}
		return redirect.Down(args[1])
	default:
		return fmt.Errorf(i18n.T("xrayredirect.bad_action"), args[0])
	}
}

func (c *Cmd) Help(w io.Writer) {
	fmt.Fprintf(w, i18n.T("xrayredirect.help"), settings.C.Table)
}

func (c *Cmd) BashCompletionCase(w io.Writer) {
	fmt.Fprint(w, `	xray-redirect)
		if [[ $ecword -eq 2 ]]; then
			mapfile -t COMPREPLY < <(compgen -W 'up down help -h --help' -- "$cur")
		elif [[ ${COMP_WORDS[$((_z_panel_cmd_start+1))]} == up ]]; then
			if [[ $cur == -* ]]; then
				mapfile -t COMPREPLY < <(compgen -W '--no-mark --ipv6 --table --bypass-cgroup --bypass-unit' -- "$cur")
			else
				mapfile -t COMPREPLY < <(compgen -W "$(_z_panel_interfaces)" -- "$cur")
			fi
		elif [[ ${COMP_WORDS[$((_z_panel_cmd_start+1))]} == down ]] && [[ $ecword -eq 3 ]]; then
			mapfile -t COMPREPLY < <(compgen -W "$(_z_panel_interfaces)" -- "$cur")
		fi
		return
		;;
`)
}
