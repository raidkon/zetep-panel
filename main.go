// z-panel: TUN policy routing (wg-quick-style default route).
package main

import (
	"fmt"
	"os"

	"z-panel/cmd/confcmd"
	daemoncmd "z-panel/cmd/daemon"
	"z-panel/cmd/install"
	"z-panel/cmd/installshell"
	"z-panel/cmd/ufw"
	"z-panel/cmd/version"
	"z-panel/cmd/xrayredirect"
	"z-panel/cmd/xraytun"
	"z-panel/internal/app"
	"z-panel/internal/client"
	"z-panel/internal/config"
	"z-panel/internal/daemon"
	"z-panel/internal/executil"
	"z-panel/internal/i18n"
	"z-panel/internal/settings"
	"z-panel/internal/transport"
)

func init() {
	app.Register(version.New())
	app.Register(install.New())
	app.Register(installshell.New())
	app.Register(confcmd.New())
	app.Register(daemoncmd.New())
	app.Register(xrayredirect.New())
	app.Register(ufw.New())
	app.Register(xraytun.New())
}

func main() {
	i18n.Init()
	os.Exit(runMain(os.Args))
}

// runMain implements the root command flow; returns an exit code (process should not call os.Exit).
func runMain(args []string) int {
	remoteMode, sshHost, rest, err := transport.ParseSSHFromArgs(args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 2
	}
	switch remoteMode {
	case transport.RemoteZPanelBinary:
		if os.Getenv("Z_PANEL_NO_BANNER") == "" {
			fmt.Fprintf(os.Stderr, "z-panel %s\n", config.Version)
		}
		settings.ApplyDefaults()
		i18n.ApplyFromConfig(settings.C.Language)
		if len(rest) < 2 {
			app.PrintRootHelp(os.Stdout)
			return 0
		}
		arg1 := rest[1]
		if arg1 == "help" || arg1 == "-h" || arg1 == "--help" {
			app.PrintRootHelp(os.Stdout)
			return 0
		}
		if err := transport.RunZPanelOverSSH(sshHost, rest); err != nil {
			fmt.Fprintln(os.Stderr, err)
			if code, ok := executil.ExitCode(err); ok {
				return code
			}
			return 1
		}
		return 0
	case transport.RemoteLocalTools:
		if os.Getenv("Z_PANEL_NO_BANNER") == "" {
			fmt.Fprintf(os.Stderr, "z-panel %s\n", config.Version)
		}
		_ = os.Setenv(executil.EnvSSHHost, sshHost)
		defer func() { _ = os.Unsetenv(executil.EnvSSHHost) }()
		if muxStop, _ := executil.TryStartSSHMultiplex(sshHost); muxStop != nil {
			defer muxStop()
		}
		settings.ApplyDefaults()
		i18n.ApplyFromConfig(settings.C.Language)
		if len(rest) < 2 {
			app.PrintRootHelp(os.Stdout)
			return 0
		}
		arg1 := rest[1]
		if arg1 == "help" || arg1 == "-h" || arg1 == "--help" {
			app.PrintRootHelp(os.Stdout)
			return 0
		}
		if arg1 == "-v" || arg1 == "--version" {
			arg1 = "version"
		}
		switch arg1 {
		case "install-shell", "config", "daemon", "xray-tun", "xray-redirect":
			fmt.Fprint(os.Stderr, i18n.T("transport.remote_forbidden", arg1))
			return 1
		}
		cmd := app.FindCommand(arg1)
		if cmd == nil {
			fmt.Fprint(os.Stderr, i18n.T("root.unknown_command", arg1))
			app.PrintRootHelp(os.Stdout)
			return 1
		}
		if err := cmd.Run(rest[2:]); err != nil {
			fmt.Fprintln(os.Stderr, err)
			if code, ok := executil.ExitCode(err); ok {
				return code
			}
			return 1
		}
		return 0
	}

	if os.Getenv("Z_PANEL_NO_BANNER") == "" {
		fmt.Fprintf(os.Stderr, "z-panel %s\n", config.Version)
	}
	if err := settings.Load(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	i18n.ApplyFromConfig(settings.C.Language)

	skipDaemon := os.Getenv("Z_PANEL_SKIP_DAEMON") != ""
	if !skipDaemon && len(args) >= 2 && settings.C.DaemonEnabled() && !daemon.ForbiddenRemote(args[1]) {
		err := client.Forward(args[1:])
		if err == nil {
			return 0
		}
		if client.IsUnavailable(err) {
			fmt.Fprintf(os.Stderr, "%s", i18n.T("daemon.fallback_warning"))
			_ = os.Setenv("Z_PANEL_SKIP_DAEMON", "1")
		} else {
			fmt.Fprintln(os.Stderr, err)
			if code, ok := client.ExitStatus(err); ok {
				return code
			}
			return 1
		}
	}

	if len(args) < 2 {
		app.PrintRootHelp(os.Stdout)
		return 0
	}

	arg1 := args[1]
	if arg1 == "help" || arg1 == "-h" || arg1 == "--help" {
		app.PrintRootHelp(os.Stdout)
		return 0
	}

	if arg1 == "-v" || arg1 == "--version" {
		arg1 = "version"
	}

	cmd := app.FindCommand(arg1)
	if cmd == nil {
		fmt.Fprint(os.Stderr, i18n.T("root.unknown_command", args[1]))
		app.PrintRootHelp(os.Stdout)
		return 1
	}

	if err := cmd.Run(args[2:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	return 0
}
