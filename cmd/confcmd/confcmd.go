package confcmd

import (
	"fmt"
	"io"
	"os"

	"z-panel/internal/app"
	"z-panel/internal/config"
	"z-panel/internal/i18n"
	"z-panel/internal/settings"
)

type Cmd struct{}

func New() *Cmd { return &Cmd{} }

func (c *Cmd) Name() string { return "config" }

func (c *Cmd) Run(args []string) error {
	if len(args) == 0 || app.IsHelpRequest(args) {
		c.Help(os.Stdout)
		return nil
	}
	switch args[0] {
	case "init":
		force := false
		for _, a := range args[1:] {
			if a == "--force" || a == "-f" {
				force = true
			}
		}
		return settings.InitInteractive(os.Stdin, os.Stdout, force)
	case "migrate":
		return settings.MigrateInteractive(os.Stdin, os.Stdout)
	default:
		return fmt.Errorf(i18n.T("confcmd.err_unknown"), args[0])
	}
}

func (c *Cmd) Help(w io.Writer) {
	fmt.Fprintf(w, i18n.T("confcmd.help"), config.ConfigFile)
}

func (c *Cmd) BashCompletionCase(w io.Writer) {
	fmt.Fprint(w, `	config)
		if [[ $ecword -eq 2 ]]; then
			mapfile -t COMPREPLY < <(compgen -W 'init migrate help -h --help' -- "$cur")
		elif [[ ${COMP_WORDS[$((_z_panel_cmd_start+1))]} == init ]]; then
			mapfile -t COMPREPLY < <(compgen -W '--force -f' -- "$cur")
		fi
		return
		;;
`)
}
