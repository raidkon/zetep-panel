// Package app defines the subcommand contract and root help.
package app

import (
	"fmt"
	"io"
	"strings"

	"z-panel/internal/i18n"
	"z-panel/internal/settings"
)

// Command is one top-level subcommand (everything after the name is passed to Run).
type Command interface {
	Name() string
	Run(args []string) error
	Help(w io.Writer)
	// BashCompletionCase writes case branch(es) for _z_panel_completion (tab-indented, through ;;).
	BashCompletionCase(w io.Writer)
}

// IsHelpRequest: first token after the command is help / -h / --help.
func IsHelpRequest(args []string) bool {
	if len(args) == 0 {
		return false
	}
	switch strings.ToLower(strings.TrimSpace(args[0])) {
	case "help", "-h", "--help":
		return true
	default:
		return false
	}
}

// PrintRootHelp prints the header and Help for each registered command.
func PrintRootHelp(w io.Writer) {
	cmds := All()
	fmt.Fprintln(w, i18n.T("root.help.tagline"))
	fmt.Fprint(w, i18n.T("root.help.top"))
	for _, c := range cmds {
		fmt.Fprintf(w, i18n.T("root.help.cmdline"), c.Name())
	}
	fmt.Fprintf(w, i18n.T("root.help.ufw_note"), settings.C.UfwMarker)
	for _, c := range cmds {
		fmt.Fprintf(w, i18n.T("root.help.section_rule"), c.Name())
		c.Help(w)
		fmt.Fprintln(w)
	}
}
