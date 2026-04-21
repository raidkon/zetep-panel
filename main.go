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
	remoteMode, sshHost, rest, err := transport.ParseSSHFromArgs(os.Args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
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
			os.Exit(0)
		}
		arg1 := rest[1]
		if arg1 == "help" || arg1 == "-h" || arg1 == "--help" {
			app.PrintRootHelp(os.Stdout)
			os.Exit(0)
		}
		if err := transport.RunZPanelOverSSH(sshHost, rest); err != nil {
			fmt.Fprintln(os.Stderr, err)
			if code, ok := executil.ExitCode(err); ok {
				os.Exit(code)
			}
			os.Exit(1)
		}
		return
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
			os.Exit(0)
		}
		arg1 := rest[1]
		if arg1 == "help" || arg1 == "-h" || arg1 == "--help" {
			app.PrintRootHelp(os.Stdout)
			os.Exit(0)
		}
		if arg1 == "-v" || arg1 == "--version" {
			arg1 = "version"
		}
		switch arg1 {
		case "install", "install-shell", "config", "daemon", "xray-tun", "xray-redirect":
			fmt.Fprint(os.Stderr, i18n.T("transport.remote_forbidden", arg1))
			os.Exit(1)
		}
		cmd := app.FindCommand(arg1)
		if cmd == nil {
			fmt.Fprint(os.Stderr, i18n.T("root.unknown_command", arg1))
			app.PrintRootHelp(os.Stdout)
			os.Exit(1)
		}
		if err := cmd.Run(rest[2:]); err != nil {
			fmt.Fprintln(os.Stderr, err)
			if code, ok := executil.ExitCode(err); ok {
				os.Exit(code)
			}
			os.Exit(1)
		}
		return
	}

	if os.Getenv("Z_PANEL_NO_BANNER") == "" {
		fmt.Fprintf(os.Stderr, "z-panel %s\n", config.Version)
	}
	if err := settings.Load(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	i18n.ApplyFromConfig(settings.C.Language)

	skipDaemon := os.Getenv("Z_PANEL_SKIP_DAEMON") != ""
	if !skipDaemon && len(os.Args) >= 2 && settings.C.DaemonEnabled() && !daemon.ForbiddenRemote(os.Args[1]) {
		err := client.Forward(os.Args[1:])
		if err == nil {
			return
		}
		if client.IsUnavailable(err) {
			fmt.Fprintf(os.Stderr, "%s", i18n.T("daemon.fallback_warning"))
			_ = os.Setenv("Z_PANEL_SKIP_DAEMON", "1")
		} else {
			fmt.Fprintln(os.Stderr, err)
			if code, ok := client.ExitStatus(err); ok {
				os.Exit(code)
			}
			os.Exit(1)
		}
	}

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
