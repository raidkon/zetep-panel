// z-panel: TUN policy routing (wg-quick-style default route).
package main

import (
	"fmt"
	"os"

	"z-panel/cmd/confcmd"
	"z-panel/cmd/install"
	"z-panel/cmd/installshell"
	"z-panel/cmd/ufw"
	"z-panel/cmd/version"
	"z-panel/cmd/xrayredirect"
	"z-panel/cmd/xraytun"
	"z-panel/internal/app"
	"z-panel/internal/config"
	"z-panel/internal/i18n"
	"z-panel/internal/settings"
)

func init() {
	app.Register(version.New())
	app.Register(install.New())
	app.Register(installshell.New())
	app.Register(confcmd.New())
	app.Register(xrayredirect.New())
	app.Register(ufw.New())
	app.Register(xraytun.New())
}

func main() {
	i18n.Init()
	fmt.Fprintf(os.Stderr, "z-panel %s\n", config.Version)
	if err := settings.Load(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	i18n.ApplyFromConfig(settings.C.Language)

	if len(os.Args) < 2 {
		app.PrintRootHelp(os.Stdout)
		os.Exit(0)
	}

	arg1 := os.Args[1]
	if arg1 == "help" || arg1 == "-h" || arg1 == "--help" {
		app.PrintRootHelp(os.Stdout)
		os.Exit(0)
	}

	if arg1 == "-v" || arg1 == "--version" {
		arg1 = "version"
	}

	cmd := app.FindCommand(arg1)
	if cmd == nil {
		fmt.Fprint(os.Stderr, i18n.T("root.unknown_command", os.Args[1]))
		app.PrintRootHelp(os.Stdout)
		os.Exit(1)
	}

	if err := cmd.Run(os.Args[2:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
