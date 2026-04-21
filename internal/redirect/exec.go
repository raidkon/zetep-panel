package redirect

import (
	"fmt"
	"io"
	"os"
	"strings"

	"z-panel/internal/executil"
	"z-panel/internal/i18n"
)

func ipLinkExists(iface string) error {
	out, err := executil.RunTTYCombined("ip", "link", "show", "dev", iface)
	if err != nil {
		return fmt.Errorf(i18n.T("redirect.iface_not_found"), iface, strings.TrimSpace(string(out)))
	}
	return nil
}

func run(name string, args ...string) error {
	cmd := executil.CommandTTY(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func ipRouteFlushTableQuiet(table string) {
	cmd := executil.Command("ip", "route", "flush", "table", table)
	cmd.Stdout = io.Discard
	cmd.Stderr = io.Discard
	_ = cmd.Run()
}

func runIPQuiet(name string, args ...string) {
	cmd := executil.Command(name, args...)
	cmd.Stdout = io.Discard
	cmd.Stderr = io.Discard
	_ = cmd.Run()
}

func ipRuleShow(proto string) string {
	out, _ := executil.Command("ip", proto, "rule", "show").Output()
	return string(out)
}
