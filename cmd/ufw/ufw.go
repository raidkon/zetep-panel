package ufw

import (
	"fmt"
	"io"
	"os"
	"strings"

	"z-panel/internal/app"
	"z-panel/internal/i18n"
	"z-panel/internal/settings"
)

type Cmd struct{}

func New() *Cmd { return &Cmd{} }

func (c *Cmd) Name() string { return "ufw" }

func (c *Cmd) Run(args []string) error {
	if app.IsHelpRequest(args) {
		c.Help(os.Stdout)
		return nil
	}
	if len(args) < 1 {
		return fmt.Errorf("%s", i18n.T("ufw.want_subcmd"))
	}
	switch args[0] {
	case "check":
		iface, lanCIDR, lanDev, full, err := parseCheckArgs(args[1:])
		if err != nil {
			return err
		}
		return runUnifiedCheck(iface, lanCIDR, lanDev, full)
	case "masq-check":
		iface, lanCIDR, err := parseMasqCheckArgs(args[1:])
		if err != nil {
			return err
		}
		return runUnifiedCheck(iface, lanCIDR, settings.C.DefaultLANDev, false)
	default:
		return fmt.Errorf("%s", i18n.T("ufw.want_subcmd"))
	}
}

func (c *Cmd) Help(w io.Writer) {
	fmt.Fprintf(w, i18n.T("ufw.help"), settings.C.UfwMarker, settings.C.DefaultLANCIDR, settings.C.DefaultLANDev, settings.C.DefaultLANCIDR)
}

func parseCheckArgs(args []string) (iface, lanCIDR, lanDev string, full bool, err error) {
	lanCIDR = settings.C.DefaultLANCIDR
	lanDev = settings.C.DefaultLANDev
	var pos []string
	for i := 0; i < len(args); i++ {
		a := args[i]
		switch {
		case a == "--full":
			full = true
		case strings.HasPrefix(a, "--lan-cidr="):
			lanCIDR = strings.TrimSpace(strings.TrimPrefix(a, "--lan-cidr="))
			if lanCIDR == "" {
				return "", "", "", false, fmt.Errorf("%s", i18n.T("ufw.err.lan_cidr_empty"))
			}
		case a == "--lan-cidr":
			if i+1 >= len(args) {
				return "", "", "", false, fmt.Errorf("%s", i18n.T("ufw.err.lan_cidr_need"))
			}
			i++
			lanCIDR = strings.TrimSpace(args[i])
			if lanCIDR == "" {
				return "", "", "", false, fmt.Errorf("%s", i18n.T("ufw.err.lan_cidr_empty"))
			}
		case strings.HasPrefix(a, "--lan-dev="):
			lanDev = strings.TrimSpace(strings.TrimPrefix(a, "--lan-dev="))
			if lanDev == "" {
				return "", "", "", false, fmt.Errorf("%s", i18n.T("ufw.err.lan_dev_empty"))
			}
		case a == "--lan-dev":
			if i+1 >= len(args) {
				return "", "", "", false, fmt.Errorf("%s", i18n.T("ufw.err.lan_dev_need"))
			}
			i++
			lanDev = strings.TrimSpace(args[i])
			if lanDev == "" {
				return "", "", "", false, fmt.Errorf("%s", i18n.T("ufw.err.lan_dev_empty"))
			}
		case strings.HasPrefix(a, "-"):
			return "", "", "", false, fmt.Errorf(i18n.T("ufw.err.unknown_flag"), a)
		default:
			pos = append(pos, a)
		}
	}
	if len(pos) > 1 {
		return "", "", "", false, fmt.Errorf(i18n.T("ufw.err.too_many_iface"), strings.Join(pos[1:], " "))
	}
	if len(pos) == 1 {
		iface = pos[0]
	}
	return iface, lanCIDR, lanDev, full, nil
}

func (c *Cmd) BashCompletionCase(w io.Writer) {
	fmt.Fprint(w, `	ufw)
		if [[ $cword -eq 2 ]]; then
			mapfile -t COMPREPLY < <(compgen -W 'check masq-check help -h --help' -- "$cur")
		elif [[ ${COMP_WORDS[2]} == check ]]; then
			if [[ $cur == -* ]]; then
				mapfile -t COMPREPLY < <(compgen -W '--full --lan-cidr --lan-dev' -- "$cur")
			else
				mapfile -t COMPREPLY < <(compgen -W "$(_z_panel_interfaces)" -- "$cur")
			fi
		elif [[ ${COMP_WORDS[2]} == masq-check ]]; then
			if [[ $cur == -* ]]; then
				mapfile -t COMPREPLY < <(compgen -W '--lan-cidr' -- "$cur")
			else
				mapfile -t COMPREPLY < <(compgen -W "$(_z_panel_interfaces)" -- "$cur")
			fi
		fi
		return
		;;
`)
}
