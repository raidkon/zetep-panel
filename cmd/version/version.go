package version

import (
	"fmt"
	"io"
	"os"

	"z-panel/internal/app"
	"z-panel/internal/config"
	"z-panel/internal/i18n"
)

type Cmd struct{}

func New() *Cmd { return &Cmd{} }

func (c *Cmd) Name() string { return "version" }

func (c *Cmd) Aliases() []string { return []string{"-v", "--version"} }

func (c *Cmd) Run(args []string) error {
	if app.IsHelpRequest(args) {
		c.Help(os.Stdout)
		return nil
	}
	return nil
}

func (c *Cmd) Help(w io.Writer) {
	fmt.Fprintf(w, i18n.T("version.help"), config.Version)
}

func (c *Cmd) BashCompletionCase(w io.Writer) {
	fmt.Fprint(w, `	-v|--version|version)
		return
		;;
`)
}
