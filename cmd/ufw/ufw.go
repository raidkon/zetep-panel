package ufw

import (
	"fmt"
	"io"
	"os"
	"os/exec"
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
	if len(args) < 1 || args[0] != "check" {
		return fmt.Errorf("%s", i18n.T("ufw.want_check"))
	}
	iface, lanCIDR, lanDev, err := parseCheckArgs(args[1:])
	if err != nil {
		return err
	}
	return ufwCheck(iface, lanCIDR, lanDev)
}

func (c *Cmd) Help(w io.Writer) {
	fmt.Fprintf(w, i18n.T("ufw.help"), settings.C.UfwMarker, settings.C.DefaultLANCIDR, settings.C.DefaultLANDev)
}

func parseCheckArgs(args []string) (iface, lanCIDR, lanDev string, err error) {
	lanCIDR = settings.C.DefaultLANCIDR
	lanDev = settings.C.DefaultLANDev
	var pos []string
	for i := 0; i < len(args); i++ {
		a := args[i]
		switch {
		case strings.HasPrefix(a, "--lan-cidr="):
			lanCIDR = strings.TrimSpace(strings.TrimPrefix(a, "--lan-cidr="))
			if lanCIDR == "" {
				return "", "", "", fmt.Errorf("%s", i18n.T("ufw.err.lan_cidr_empty"))
			}
		case a == "--lan-cidr":
			if i+1 >= len(args) {
				return "", "", "", fmt.Errorf("%s", i18n.T("ufw.err.lan_cidr_need"))
			}
			i++
			lanCIDR = strings.TrimSpace(args[i])
			if lanCIDR == "" {
				return "", "", "", fmt.Errorf("%s", i18n.T("ufw.err.lan_cidr_empty"))
			}
		case strings.HasPrefix(a, "--lan-dev="):
			lanDev = strings.TrimSpace(strings.TrimPrefix(a, "--lan-dev="))
			if lanDev == "" {
				return "", "", "", fmt.Errorf("%s", i18n.T("ufw.err.lan_dev_empty"))
			}
		case a == "--lan-dev":
			if i+1 >= len(args) {
				return "", "", "", fmt.Errorf("%s", i18n.T("ufw.err.lan_dev_need"))
			}
			i++
			lanDev = strings.TrimSpace(args[i])
			if lanDev == "" {
				return "", "", "", fmt.Errorf("%s", i18n.T("ufw.err.lan_dev_empty"))
			}
		case strings.HasPrefix(a, "-"):
			return "", "", "", fmt.Errorf(i18n.T("ufw.err.unknown_flag"), a)
		default:
			pos = append(pos, a)
		}
	}
	if len(pos) > 1 {
		return "", "", "", fmt.Errorf(i18n.T("ufw.err.too_many_iface"), strings.Join(pos[1:], " "))
	}
	if len(pos) == 1 {
		iface = pos[0]
	}
	return iface, lanCIDR, lanDev, nil
}

func ufwCheck(iface, lanCIDR, lanDev string) error {
	out, err := exec.Command("ufw", "status", "verbose").CombinedOutput()
	if err != nil {
		return fmt.Errorf(i18n.T("ufw.ufw_status_failed"), err, out)
	}
	text := string(out)
	fmt.Printf("%s\n", i18n.T("ufw.section_rules", settings.C.UfwMarker))
	lines := strings.Split(text, "\n")
	found := false
	for _, ln := range lines {
		if strings.Contains(strings.ToLower(ln), settings.C.UfwMarker) {
			fmt.Println(ln)
			found = true
		}
	}
	if !found {
		fmt.Println(i18n.T("ufw.no_lines"))
	}
	fmt.Println()
	fmt.Println(i18n.T("ufw.section_hints"))
	if iface != "" {
		fmt.Printf(i18n.T("ufw.hint_sysctl"), lanDev, iface, lanCIDR, settings.C.UfwMarker)
		fmt.Println()
		fmt.Print(i18n.T("ufw.hint_return"))
		fmt.Printf(i18n.T("ufw.hint_return_cmd"), iface, lanDev, settings.C.UfwMarker)
	} else {
		fmt.Printf("%s", i18n.T("ufw.no_iface_hint1"))
		fmt.Printf(i18n.T("ufw.no_iface_hint2"), lanCIDR, lanDev)
	}
	fmt.Println()
	fmt.Println(i18n.T("ufw.section_full"))
	fmt.Print(text)
	return nil
}

func (c *Cmd) BashCompletionCase(w io.Writer) {
	fmt.Fprint(w, `	ufw)
		if [[ $cword -eq 2 ]]; then
			mapfile -t COMPREPLY < <(compgen -W 'check help -h --help' -- "$cur")
		elif [[ ${COMP_WORDS[2]} == check ]]; then
			if [[ $cur == -* ]]; then
				mapfile -t COMPREPLY < <(compgen -W '--lan-cidr --lan-dev' -- "$cur")
			else
				mapfile -t COMPREPLY < <(compgen -W "$(_z_panel_interfaces)" -- "$cur")
			fi
		fi
		return
		;;
`)
}
