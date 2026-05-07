package i18n

func english() map[string]string {
	return map[string]string{
		"root.unknown_command": "unknown command: %s\n\n",

		"root.help.tagline": "z-panel — TUN policy routing (wg-quick-style for Xray).",
		"root.help.top": `Top level:
  z-panel help | -h | --help     this summary (all commands)
  z-panel version | -v | --version
  z-panel [--ssh=host | --ssh host] <command> …   local z-panel; run system tools on remote via ssh+sudo (no remote z-panel)
  z-panel [--ssh-connect=host | --ssh-connect host] <command> …   run remote installed z-panel (daemon/config on that host)
  z-panel <command> [help | -h | --help]   help for one command
  z-panel <command> …            all arguments after the command name go to that command

Commands:
`,
		"root.help.cmdline":      "  z-panel %s …\n",
		"root.help.ufw_note":     "\nUFW comment tag: '%s: …'\n\n",
		"root.help.section_rule": "─── %s ───\n",

		"root.need_root": "root privileges required (sudo)",

		"install.help": `install [help]
  Copies the binary to %s (chmod 755); requires root.
  If %s is missing — interactive prompts and saving settings.
  From your PC (upload this binary, then run install on the host; needs scp and ssh):
    z-panel --ssh=user@host install
  If z-panel is already on the server only (no local upload): z-panel --ssh-connect=user@host install

`,
		"install.err_remote_removed":  "install: unexpected argument %q (local: sudo z-panel install; remote from this machine: z-panel --ssh=host install with no extra args; on server only: z-panel --ssh-connect=host install)",
		"install.err.extra_with_ssh":  "install: extra argument %q with --ssh (use: z-panel --ssh=host install)",
		"install.err.need_scp":          "install: scp not found in PATH (needed for z-panel --ssh=… install)",
		"install.err.need_ssh":          "install: ssh not found in PATH (needed for z-panel --ssh=… install)",
		"install.err.scp":               "install: scp: %w",
		"install.err.ssh_run":         "install: ssh: %w",
		"install.ssh.uploading":       "install: copying this binary to %s:%s …\n",
		"install.err.interrupted":       "interrupted (Ctrl+C)",
		"install.err.interrupted_with":  "interrupted (Ctrl+C): %w",
		"install.err.open_self":         "open self: %w",
		"install.err.create_tmp":        "create %s: %w",
		"install.err.copy":              "copy: %w",
		"install.err.rename":            "rename to %s: %w",
		"install.err.config":            "config: %w",
		"install.installed":             "installed: %s\n",
		"install.new_version":           "New version: %s\n",
		"install.old_version":           "Previous version: %s\n",

		"installshell.err.home":   "home directory: %w",
		"installshell.err.mkdir":  "mkdir: %w",
		"installshell.err.write":  "write %s: %w",
		"installshell.done":       "bash completion installed: %s\n",
		"installshell.hint_shell": "Open a new shell or run: source /usr/share/bash-completion/bash_completion",
		"installshell.hint_user":  "(user install: ensure ~/.bashrc loads bash-completion)",
		"installshell.help": `install-shell [help] [--user|-u]
  Writes the bash completion script (generated from registered commands).
  Default: %s (requires root).
  --user: ~/.local/share/bash-completion/completions/z-panel (or $XDG_DATA_HOME/...).
  Requires bash 4+; for full support install the bash-completion package.

`,

		"confcmd.err_unknown": "config: unknown subcommand %q (expected init or migrate)",
		"confcmd.help": `config [help] init [--force|-f]
  Interactive create or overwrite %s (--force).
config [help] migrate
  Apply new config keys after upgrading z-panel (interactive prompts only for new schema versions).

`,

		"version.help": `version
  Version: %s (also printed as the first stderr line on every run).
  Top-level synonyms: -v, --version

`,

		"settings.err.read":   "read %s: %w",
		"settings.err.parse":  "parse %s: %w",
		"settings.err.mkdir":  "mkdir %s: %w",
		"settings.err.write":  "write %s: %w",
		"settings.config_hdr": "# z-panel — configuration\n\n",
		"settings.init_exists": "config already exists: %s (to overwrite: z-panel config init --force)\n",
		"settings.init_intro":  "z-panel setup — type a value or press Enter for the default.",
		"settings.saved":       "\nsaved: %s\n",
		"settings.prompt.table":           "Routing table / fwmark ID",
		"settings.prompt.systemd_network": "systemd-networkd unit directory",
		"settings.prompt.lan_cidr":        "UFW template: LAN CIDR",
		"settings.prompt.lan_dev":         "UFW template: LAN interface",
		"settings.prompt.xray_addr":       "Xray TUN: default address (ip mode)",
		"settings.prompt.xray_peer":       "Xray TUN: default peer",
		"settings.prompt.ufw_marker":      "UFW comment tag",
		"settings.prompt.xray_mark":       "Marker line in .network file",
		"settings.prompt.language":        "UI language (" + LanguageListHint + ")",
		"settings.migrate_intro":          "This config was written by an older z-panel. Please set values for new options.",
		"settings.migrate_uptodate":       "config schema is already up to date.",
		"settings.migrate_no_file":        "%s: config file not found (run z-panel config init)",

		"xrayredirect.want_up_down": "expected: z-panel xray-redirect up|down … (see z-panel xray-redirect help)",
		"xrayredirect.want_down_iface": "expected: z-panel xray-redirect down <interface>",
		"xrayredirect.bad_action":      "xray-redirect: unknown action %q (expected up or down)",
		"xrayredirect.help": `xray-redirect [help] up [flags] <interface>
xray-redirect [help] down <interface>
  Full tunnel like wg-quick: not fwmark → table, suppress_prefixlength, default dev <interface>,
  sysctl src_valid_mark, nft anti-leak, egress marking via cgroup v2 (iptables -m cgroup --path).
  up flags (before interface name):
    --bypass-unit=auto      default: try x-ui, sing-box, xray units
    --bypass-unit=x-ui      explicit systemd unit (may omit .service)
    --bypass-cgroup=path    explicit cgroup v2 path from root
    --table=N               table and fwmark (default %s)
    --no-mark               no cgroup mark
    --ipv6                  default ::/0 and IPv6 rules
    --wan-lookup=auto|off|IP[/mask]  "from <WAN> lookup main" for inbound on a public address (default auto)
  Examples:
    sudo z-panel xray-redirect up xray2tun
    sudo z-panel xray-redirect up --bypass-unit=sing-box xray2tun
    sudo z-panel xray-redirect up --table=51844 --ipv6 xray2tun
    sudo z-panel xray-redirect down xray2tun

`,

		"ufw.want_check":    "expected: z-panel ufw check [flags] [interface] (see z-panel ufw help)",
		"ufw.want_subcmd":   "expected: z-panel ufw check … (see z-panel ufw help)",
		"ufw.help": `ufw [help] check [--full] [--lan-cidr=CIDR] [--lan-dev=DEV] [interface]
  Checks UFW forwarding for the tunnel, MASQUERADE/SNAT (-o iface) in nat, and suggests fixes.
  First line: status (green / yellow / red). Without --full, only problems and fixes are printed in detail.
  Tag %s in comments. Templates: --lan-cidr (default %s), --lan-dev (%s).

ufw [help] masq-check [--lan-cidr=CIDR] <interface>
  Same as check <interface> (legacy alias).

`,
		"ufw.masq.verdict_ok":      "STATUS: OK — POSTROUTING MASQUERADE/SNAT for -o %s is present (%d matching rule(s)).",
		"ufw.masq.detail_heading":  "iptables-save -t nat — matching line(s):",
		"ufw.masq.verdict_missing": "STATUS: MISSING — no POSTROUTING MASQUERADE/SNAT rule with -o %s in table nat.\n",
		"ufw.masq.hint_add": `Suggested line (place with your other POSTROUTING MASQUERADE rules, e.g. second *nat block in /etc/ufw/before.rules), then: sudo ufw reload

-A POSTROUTING -s %s -o %s -j MASQUERADE
`,
		"ufw.masq.iptables_cmd": "iptables-save -t nat",
		"ufw.masq.want_iface":   "masq-check: exactly one interface name required (example: z-panel ufw masq-check xray2tun)",
		"ufw.err.lan_cidr_empty": "--lan-cidr: empty value",
		"ufw.err.lan_cidr_need":  "value required after --lan-cidr",
		"ufw.err.lan_dev_empty":  "--lan-dev: empty value",
		"ufw.err.lan_dev_need":   "value required after --lan-dev",
		"ufw.err.unknown_flag":   "unknown flag: %s",
		"ufw.err.too_many_iface": "at most one interface expected, extra: %q",
		"ufw.ufw_status_failed": "ufw status verbose: %w\n%s",
		"ufw.check.status_label":    "Status: ",
		"ufw.check.status_ok":       "all clear",
		"ufw.check.status_warn":     "possible issues",
		"ufw.check.status_bad":      "action required",
		"ufw.check.section_details": "=== What to address ===",
		"ufw.check.no_issues_full":  "(no issues — full report below)",
		"ufw.check.issue_no_iface":  "Tunnel interface not specified. Run: z-panel ufw check <interface> (and use --full for the full report).",
		"ufw.check.err_ipt_split":     "could not separate ufw and iptables-save output (unexpected remote shell output).",
		"ufw.check.issue_iptables":  "Could not read iptables nat table (iptables-save -t nat): %v",
		"ufw.check.issue_no_ufw":    "Interface %q does not appear in ufw status verbose — UFW is not referencing this tunnel.",
		"ufw.check.fix_no_ufw": `# Enable IP forwarding if needed: /etc/ufw/sysctl.conf → net.ipv4.ip_forward=1
Allow LAN → tunnel (adjust names/CIDR):
sudo ufw route allow in on %s out on %s from %s comment '%s: lan to tunnel'
`,
		"ufw.check.issue_no_fwd": "Interface %q is in ufw status but there is no ALLOW FWD rule for it — LAN traffic may not be forwarded into the tunnel.",
		"ufw.check.fix_no_fwd": `# /etc/ufw/sysctl.conf: net.ipv4.ip_forward=1
sudo ufw route allow in on %s out on %s from %s comment '%s: lan to tunnel'
`,
		"ufw.check.issue_no_masq": "No POSTROUTING MASQUERADE/SNAT rule with -o %q in table nat — outbound NAT from LAN via the tunnel will not work.",
		"ufw.check.issue_no_return": "No ufw status line looks like return forwarding (in on %s → out on %s). MASQUERADE only SNATs traffic leaving via the tunnel; it does not add a UFW forward/route rule for packets arriving on the tunnel toward LAN. If routed traffic is denied by default, replies to LAN clients may still be dropped.",
		"ufw.check.masq_none_in_full": "(no matching MASQUERADE/SNAT lines)",
		"ufw.section_rules":       "=== UFW rules mentioning %s ===",
		"ufw.no_lines":            "(no lines containing «z-panel»)",
		"ufw.section_iface_refs":   "=== UFW status lines referencing interface %s (any comment) ===",
		"ufw.no_iface_refs":       "(no ufw status lines reference %q — add route/forward rules if needed)",
		"ufw.section_hints":       "=== Suggestions (templates — verify for your setup) ===",
		"ufw.hint_sysctl": `# Routing and forwarding (kernel):
# /etc/ufw/sysctl.conf: net.ipv4.ip_forward=1
# LAN -> tunnel forward (example; adjust ufw zones if needed):
sudo ufw route allow in on %s out on %s from %s comment '%s: lan to tunnel'
`,
		"ufw.hint_return":     "# If ufw blocks replies:\n",
		"ufw.hint_return_cmd": "sudo ufw route allow in on %s out on %s comment '%s: return path'\n",
		"ufw.no_iface_hint1":  "Specify the tunnel interface for templates: z-panel ufw check [flags] <interface>\n",
		"ufw.no_iface_hint2":  "Current templates: LAN %s dev %s (change with --lan-cidr / --lan-dev)\n",
		"ufw.section_full":    "=== Full ufw status verbose (manual search) ===",

		"xraytun.want_subcmd": "expected: z-panel xray-tun up … or down … (see help)",
		"xraytun.want_down":   "expected: z-panel xray-tun down <interfaceName> ip",
		"xraytun.bad_action":  "xray-tun: unknown action %q (expected up or down)",
		"xraytun.help": `xray-tun [help] up [flags] <interfaceName> ip
xray-tun [help] up [flags] <interfaceName> <address[/mask]> [<peer[/mask]>]
xray-tun [help] down <interfaceName> ip
  systemd-networkd drop-in for the TUN address.
  “ip” mode: without --address/--peer defaults are %s and %s.
  Flags: --address=A (IPv4), --peer=P (optional).

`,
		"xraytun.need_address_value": "value required after --address",
		"xraytun.need_peer_value":    "value required after --peer",
		"xraytun.unknown_flag":       "unknown flag: %s",
		"xraytun.want_up_usage": "expected: z-panel xray-tun up [flags] <interface> ip\n" +
			"  or: z-panel xray-tun up [flags] <interface> <address[/mask]> [<peer[/mask]>]",
		"xraytun.extra_args":          "extra arguments: %v",
		"xraytun.empty_address":       "empty address",
		"xraytun.need_ipv4":           "IPv4 expected: %q",
		"xraytun.empty_iface":         "empty interface name",
		"xraytun.bad_iface_char":      "invalid character in interface name: %q",
		"xraytun.wrote":               "wrote %s\n",
		"xraytun.file_missing":        "file missing: %s",
		"xraytun.not_managed":         "refused: %s is not marked as z-panel-managed (missing %s) — delete manually if needed",
		"xraytun.removed":             "removed %s\n",
		"xraytun.network_reload_fail": "network reload: both networkctl reload and systemctl reload systemd-networkd failed",
		"xraytun.ok_networkctl":       "networkctl reload: ok",
		"xraytun.ok_systemd":          "systemctl reload systemd-networkd: ok",
		"xraytun.file_managed_by":     "Created by z-panel xray-tun. Manual edits not recommended.",
		"xraytun.file_remove_hint":    "Remove with: z-panel xray-tun down %s ip",

		"redirect.need_value_after": "value required after %q",
		"redirect.unknown_flag":     "unknown flag: %s",
		"redirect.need_one_iface":   "exactly one interface required, got: %v",
		"redirect.table_numeric":    "--table: routing table / fwmark must be numeric: %w",
		"redirect.cgroup_missing":   "cgroup for marking not set: use --bypass-cgroup=… or --bypass-unit=…, or --no-mark",
		"redirect.bypass_mark_fmt":  " (unit %s)",
		"redirect.mark_line":        "z-panel: bypass mark via cgroup v2: path=%q%s\n",
		"redirect.no_mark_line":     "z-panel: --no-mark — no cgroup mark; ip rules like wg-quick.\n",
		"redirect.err.default_route": "default 0.0.0.0/0 dev %s table %s: %w",
		"redirect.err.rule_fwmark":   "ip rule not fwmark: %w",
		"redirect.err.rule_suppress": "ip rule suppress_prefixlength: %w",
		"redirect.err.route6":        "ip -6 default route: %w",
		"redirect.err.rule6_fw":      "ip -6 rule not fwmark: %w",
		"redirect.err.rule6_sup":     "ip -6 rule suppress_prefixlength: %w",
		"redirect.iface_not_found":   "interface %q not found: %s",
		"redirect.down_no_state":     "no state for %q — run up first or remove ip rule/route manually (interface %q)",
		"redirect.down_bad_mode":     "unsupported state mode %q for %q — expected xray-redirect (wg/full); remove rules manually if needed",
		"redirect.down_done":         "down: flushed table %s, wg-style rules and firewall for %s\n",
		"redirect.ip6tables_missing": "ip6tables not found for IPv6 anti-leak",
		"redirect.iptables_raw_fail": "iptables raw: %w",
		"redirect.ip6tables_raw_fail": "ip6tables raw: %w",
		"redirect.iptables_cgroup_fail": "iptables cgroup (needs cgroup v2 and xt_cgroup, path from cgroup2 fs root): %w",
		"redirect.fw_skip":            "z-panel: wg-style firewall skipped (no TUN addresses for anti-leak and no cgroup mark).",
		"redirect.nft_ok":             "nft: anti-leak (preraw) applied",
		"redirect.ipt_ok":             "iptables: anti-leak (raw) applied",
		"redirect.ipt_cgroup_ok":      "iptables: OUTPUT mark via cgroup v2 (--path)",
		"redirect.ip6tables_cgroup_warn": "z-panel: warning: ip6tables cgroup: %v\n",
		"redirect.cg_systemctl":      "systemctl ControlGroup for %q: %w",
		"redirect.cg_empty":          "ControlGroup empty for %q — unit not active",
		"redirect.auto_unit":         "z-panel: auto-selected systemd unit for cgroup: %s\n",
		"redirect.auto_fail":         "could not auto-detect unit (tried %v): %w; set --bypass-unit=… or --bypass-cgroup=…",
		"redirect.wan_line":          "z-panel: ip rule pref %d: from %s lookup main (WAN bypass for services on a public / uplink address)\n",
		"redirect.wan_auto_err":        "z-panel: --wan-lookup=auto: failed to read IPv4 main table: %v\n",
		"redirect.wan_auto_skip":       "z-panel: --wan-lookup=auto: no public IPv4 on default in table main (tun %q) — skip WAN main rule; set --wan-lookup=IP/32 or fix routing\n",
		"redirect.err.wan_cidr":        "invalid --wan-lookup",
		"redirect.err.wan_rule":        "ip rule (from WAN, lookup main): %v",

		"state.state_file_err": "state file: %w",
		"state.up_line":        "up: %s (saved in %s)\n",
		"state.summary_base":   "mode=%s table=%s fwmark=%s (wg-quick style) default dev %s",
		"state.summary_nomark": " no_bypass_mark=1",
		"state.summary_bypass": " bypass_cgroup=1",

		"bashcomp.line1": "# bash completion for z-panel (generated from z-panel commands)",
		"bashcomp.line2": "# Install: z-panel install-shell; z-panel install runs this automatically (system-wide).",
		"bashcomp.line3": "# Requires bash 4+; for system-wide install use the bash-completion package.",

		"daemon.help": `daemon [help] [run]
  run — foreground daemon: HTTP API on Unix socket %s, periodic config reload (full state reconcile TODO).
  Requires root. Stop with SIGINT/SIGTERM.
  With daemon = 1 in config.toml, subcommands are sent to this process when reachable (except: z-panel daemon …).

`,
		"daemon.err_unknown_subcmd": "daemon: unknown subcommand %q (expected: run)",
		"daemon.fallback_warning":  "z-panel: daemon not reachable, running locally (start: z-panel daemon run)\n",

		"transport.ssh.err_argv":      "z-panel: internal argv error\n",
		"transport.ssh.err_duplicate": "z-panel: duplicate --ssh\n",
		"transport.ssh.err_conflict":  "z-panel: use only one of --ssh or --ssh-connect\n",
		"transport.ssh.err_empty":     "z-panel: empty --ssh host\n",
		"transport.ssh.err_missing":   "z-panel: missing value after --ssh\n",
		"transport.ssh.err_empty_connect": "z-panel: empty --ssh-connect host\n",
		"transport.ssh.err_missing_connect": "z-panel: missing value after --ssh-connect\n",
		"transport.ssh.err_no_cmd":    "z-panel: no subcommand after --ssh / --ssh-connect\n",
		"transport.remote_forbidden":    "%s: not available with --ssh (for install from this machine use: z-panel --ssh=host install; otherwise run on the server or use ufw / version; config/daemon/xray-* — on the host or --ssh-connect)\n",
	}
}

func russian() map[string]string {
	return map[string]string{
		"root.unknown_command": "неизвестная команда: %s\n\n",

		"root.help.tagline": "z-panel — маршрутизация через TUN (policy routing, стиль wg-quick для Xray).",
		"root.help.top": `Верхний уровень:
  z-panel help | -h | --help     эта справка (сводка по всем командам)
  z-panel version | -v | --version
  z-panel [--ssh=хост | --ssh хост] <команда> …   локальный z-panel; утилиты на удалённом хосте через ssh+sudo (без z-panel там)
  z-panel [--ssh-connect=хост | --ssh-connect хост] <команда> …   запуск установленного на сервере z-panel (демон, конфиг там)
  z-panel <command> [help | -h | --help]   справка только по команде
  z-panel <command> …            все аргументы после имени команды — в пакет команды

Команды:
`,
		"root.help.cmdline":      "  z-panel %s …\n",
		"root.help.ufw_note":     "\nПометка в UFW: comment '%s: …'\n\n",
		"root.help.section_rule": "─── %s ───\n",

		"root.need_root": "нужны права root (sudo)",

		"install.help": `install [help]
  Копирует бинарник в %s (chmod 755), нужен root.
  Если %s ещё нет — интерактивный опрос и сохранение настроек.
  С этой машины (залить этот бинарник и выполнить install на хосте; нужны scp и ssh):
    z-panel --ssh=user@host install
  Если z-panel только на сервере (без заливки с ПК): z-panel --ssh-connect=user@host install

`,
		"install.err_remote_removed":  "install: лишний аргумент %q (локально: sudo z-panel install; с этой машины: z-panel --ssh=хост install без лишних аргументов; только на сервере: z-panel --ssh-connect=хост install)",
		"install.err.extra_with_ssh":  "install: лишний аргумент %q с --ssh (нужно: z-panel --ssh=хост install)",
		"install.err.need_scp":          "install: в PATH нет scp (нужен для z-panel --ssh=… install)",
		"install.err.need_ssh":          "install: в PATH нет ssh (нужен для z-panel --ssh=… install)",
		"install.err.scp":               "install: scp: %w",
		"install.err.ssh_run":         "install: ssh: %w",
		"install.ssh.uploading":       "install: копирование этого бинарника на %s:%s …\n",
		"install.err.interrupted":      "прервано (Ctrl+C)",
		"install.err.interrupted_with": "прервано (Ctrl+C): %w",
		"install.err.open_self":        "открыть себя: %w",
		"install.err.create_tmp":       "создать %s: %w",
		"install.err.copy":             "копирование: %w",
		"install.err.rename":           "переименовать в %s: %w",
		"install.err.config":           "конфиг: %w",
		"install.installed":            "установлено: %s\n",
		"install.new_version":          "Новая версия: %s\n",
		"install.old_version":          "Старая версия: %s\n",

		"installshell.err.home":   "домашний каталог: %w",
		"installshell.err.mkdir":  "mkdir: %w",
		"installshell.err.write":  "запись %s: %w",
		"installshell.done":       "установлено автодополнение bash: %s\n",
		"installshell.hint_shell": "Откройте новый терминал или выполните: source /usr/share/bash-completion/bash_completion",
		"installshell.hint_user":  "(для пользователя: убедитесь, что ~/.bashrc подключает bash-completion)",
		"installshell.help": `install-shell [help] [--user|-u]
  Записывает скрипт автодополнения bash (сгенерированный из зарегистрированных команд).
  По умолчанию: %s (нужен root).
  --user: ~/.local/share/bash-completion/completions/z-panel (или $XDG_DATA_HOME/...).
  Нужен bash 4+; для полной поддержки установите пакет bash-completion.

`,

		"confcmd.err_unknown": "config: неизвестная подкоманда %q (ожидалось init или migrate)",
		"confcmd.help": `config [help] init [--force|-f]
  Интерактивное создание или перезапись %s (--force).
config [help] migrate
  Применить новые ключи конфига после обновления z-panel (интерактивно только для новых версий схемы).

`,

		"version.help": `version
  Версия: %s (также печатается первой строкой на stderr при любом запуске).
  Синонимы верхнего уровня: -v, --version

`,

		"settings.err.read":   "чтение %s: %w",
		"settings.err.parse":  "разбор %s: %w",
		"settings.err.mkdir":  "mkdir %s: %w",
		"settings.err.write":  "запись %s: %w",
		"settings.config_hdr": "# z-panel — конфигурация\n\n",
		"settings.init_exists": "конфиг уже есть: %s (для перезаписи: z-panel config init --force)\n",
		"settings.init_intro":  "Настройка z-panel — введите значение или Enter для значения по умолчанию.",
		"settings.saved":       "\nсохранено: %s\n",
		"settings.prompt.table":           "Таблица маршрутизации и fwmark",
		"settings.prompt.systemd_network": "Каталог unit-файлов systemd-networkd",
		"settings.prompt.lan_cidr":        "UFW шаблон: LAN CIDR",
		"settings.prompt.lan_dev":         "UFW шаблон: интерфейс LAN",
		"settings.prompt.xray_addr":       "Xray TUN: адрес по умолчанию (режим ip)",
		"settings.prompt.xray_peer":       "Xray TUN: peer по умолчанию",
		"settings.prompt.ufw_marker":      "Метка в комментариях UFW",
		"settings.prompt.xray_mark":     "Маркер в .network (строка в файле)",
		"settings.prompt.language":      "Язык интерфейса (" + LanguageListHint + ")",
		"settings.migrate_intro":        "Конфиг создан старой версией z-panel. Задайте значения для новых параметров.",
		"settings.migrate_uptodate":     "схема конфига уже актуальна.",
		"settings.migrate_no_file":      "%s: файл конфига не найден (выполните z-panel config init)",

		"xrayredirect.want_up_down":    "ожидалось: z-panel xray-redirect up|down … (см. z-panel xray-redirect help)",
		"xrayredirect.want_down_iface": "ожидалось: z-panel xray-redirect down <interface>",
		"xrayredirect.bad_action":      "xray-redirect: неизвестное действие %q (ожидалось up или down)",
		"xrayredirect.help": `xray-redirect [help] up [флаги] <interface>
xray-redirect [help] down <interface>
  Полный туннель как wg-quick: not fwmark → table, suppress_prefixlength, default dev <interface>,
  sysctl src_valid_mark, nft anti-leak, пометка исходящего трафика по cgroup v2 (iptables -m cgroup --path).
  Флаги up (до имени интерфейса):
    --bypass-unit=auto      по умолчанию: перебор юнитов x-ui, sing-box, xray
    --bypass-unit=x-ui      явный systemd-юнит (можно без .service)
    --bypass-cgroup=путь    явный путь cgroup v2 от корня
    --table=N               таблица и fwmark (по умолчанию %s)
    --no-mark               без пометки cgroup
    --ipv6                  default ::/0 и правила IPv6
    --wan-lookup=auto|off|IP[/маска]  «from <WAN> lookup main» для входа на публичный адрес (по умолчанию auto)
  Примеры:
    sudo z-panel xray-redirect up xray2tun
    sudo z-panel xray-redirect up --bypass-unit=sing-box xray2tun
    sudo z-panel xray-redirect up --table=51844 --ipv6 xray2tun
    sudo z-panel xray-redirect down xray2tun

`,

		"ufw.want_check":  "ожидалось: z-panel ufw check [флаги] [interface] (см. z-panel ufw help)",
		"ufw.want_subcmd": "ожидалось: z-panel ufw check … (см. z-panel ufw help)",
		"ufw.help": `ufw [help] check [--full] [--lan-cidr=CIDR] [--lan-dev=DEV] [interface]
  Проверка forward UFW для туннеля, MASQUERADE/SNAT (-o интерфейс) в nat и подсказки по исправлению.
  Первая строка: статус (зелёный / жёлтый / красный). Без --full подробности только по проблемам и что сделать.
  Пометка %s в комментариях. Шаблоны: --lan-cidr (по умолчанию %s), --lan-dev (%s).

ufw [help] masq-check [--lan-cidr=CIDR] <интерфейс>
  То же, что check <интерфейс> (старый алиас).

`,
		"ufw.masq.verdict_ok":      "СТАТУС: ЕСТЬ — в nat для исхода через %s уже задано MASQUERADE/SNAT в POSTROUTING (совпадений: %d).",
		"ufw.masq.detail_heading":  "Соответствующие строки (iptables-save -t nat):",
		"ufw.masq.verdict_missing": "СТАТУС: НЕТ — в таблице nat не найдено правило POSTROUTING с MASQUERADE/SNAT и -o %s.\n",
		"ufw.masq.hint_add": `Добавьте строку (рядом с остальными POSTROUTING MASQUERADE, часто второй блок *nat в /etc/ufw/before.rules), затем: sudo ufw reload

-A POSTROUTING -s %s -o %s -j MASQUERADE
`,
		"ufw.masq.iptables_cmd": "iptables-save -t nat",
		"ufw.masq.want_iface":   "masq-check: нужен ровно один интерфейс (пример: z-panel ufw masq-check xray2tun)",
		"ufw.err.lan_cidr_empty": "--lan-cidr: пустое значение",
		"ufw.err.lan_cidr_need":  "нужно значение после --lan-cidr",
		"ufw.err.lan_dev_empty":  "--lan-dev: пустое значение",
		"ufw.err.lan_dev_need":   "нужно значение после --lan-dev",
		"ufw.err.unknown_flag":   "неизвестный флаг: %s",
		"ufw.err.too_many_iface": "ожидается не больше одного интерфейса, лишнее: %q",
		"ufw.ufw_status_failed":  "ufw status verbose: %w\n%s",
		"ufw.check.status_label":    "Статус: ",
		"ufw.check.status_ok":       "все отлично",
		"ufw.check.status_warn":     "возможны проблемы",
		"ufw.check.status_bad":      "необходимы разрешения",
		"ufw.check.section_details": "=== Что сделать ===",
		"ufw.check.no_issues_full":  "(замечаний нет — ниже полный отчёт)",
		"ufw.check.issue_no_iface":  "Не указан интерфейс туннеля. Запустите: z-panel ufw check <интерфейс> (и --full для полного вывода).",
		"ufw.check.err_ipt_split":     "не удалось отделить вывод ufw и iptables-save (неожиданный вывод на удалённой стороне).",
		"ufw.check.issue_iptables":  "Не удалось прочитать таблицу nat (iptables-save -t nat): %v",
		"ufw.check.issue_no_ufw":    "Интерфейса %q нет в выводе «ufw status verbose» — UFW не ссылается на этот туннель.",
		"ufw.check.fix_no_ufw": `# При необходимости: /etc/ufw/sysctl.conf → net.ipv4.ip_forward=1
Разрешить LAN → туннель (подставьте свои имена/CIDR):
sudo ufw route allow in on %s out on %s from %s comment '%s: lan to tunnel'
`,
		"ufw.check.issue_no_fwd": "Интерфейс %q есть в статусе ufw, но нет правила ALLOW FWD — трафик LAN в туннель может не пересылаться.",
		"ufw.check.fix_no_fwd": `# /etc/ufw/sysctl.conf: net.ipv4.ip_forward=1
sudo ufw route allow in on %s out on %s from %s comment '%s: lan to tunnel'
`,
		"ufw.check.issue_no_masq": "В таблице nat нет POSTROUTING MASQUERADE/SNAT с -o %q — исходящий NAT с LAN в туннель не задан.",
		"ufw.check.issue_no_return": "В статусе ufw не видно пересылки для обратного пути (in on %s → out on %s). MASQUERADE в nat только делает SNAT для исходящего трафика в туннель; отдельное правило forward/route для пакетов, входящих с туннеля к LAN, этим не заменяется — при deny (routed) ответы к клиентам всё ещё могут резаться.",
		"ufw.check.masq_none_in_full": "(нет подходящих строк MASQUERADE/SNAT)",
		"ufw.section_rules":      "=== Правила UFW с пометкой %s ===",
		"ufw.no_lines":           "(нет строк, содержащих «z-panel»)",
		"ufw.section_iface_refs": "=== Строки ufw status, где есть интерфейс %s (любой комментарий) ===",
		"ufw.no_iface_refs":      "(в ufw status нет строк с «on %s» — при необходимости добавьте route/forward)",
		"ufw.section_hints":      "=== Рекомендации (шаблоны, проверьте под свою схему) ===",
		"ufw.hint_sysctl": `# Маршрутизация и forward (ядро):
# /etc/ufw/sysctl.conf: net.ipv4.ip_forward=1
# Пересылка LAN -> туннель (пример; подставьте зоны ufw при необходимости):
sudo ufw route allow in on %s out on %s from %s comment '%s: lan to tunnel'
`,
		"ufw.hint_return":     "# Если ufw режет ответы:\n",
		"ufw.hint_return_cmd": "sudo ufw route allow in on %s out on %s comment '%s: return path'\n",
		"ufw.no_iface_hint1":  "Укажите интерфейс туннеля для шаблонов: z-panel ufw check [флаги] <interface>\n",
		"ufw.no_iface_hint2":  "Текущие шаблоны: LAN %s dev %s (смена: --lan-cidr / --lan-dev)\n",
		"ufw.section_full":    "=== Полный вывод ufw status verbose (для ручного поиска) ===",

		"xraytun.want_subcmd": "ожидалось: z-panel xray-tun up … или down … (см. help)",
		"xraytun.want_down":   "ожидалось: z-panel xray-tun down <interfaceName> ip",
		"xraytun.bad_action":  "xray-tun: неизвестное действие %q (ожидалось up или down)",
		"xraytun.help": `xray-tun [help] up [флаги] <interfaceName> ip
xray-tun [help] up [флаги] <interfaceName> <address[/mask]> [<peer[/mask]>]
xray-tun [help] down <interfaceName> ip
  systemd-networkd drop-in для адреса TUN.
  Режим «ip»: без --address/--peer подставляются %s и %s.
  Флаги: --address=A (IPv4), --peer=P (опционально).

`,
		"xraytun.need_address_value": "нужно значение после --address",
		"xraytun.need_peer_value":    "нужно значение после --peer",
		"xraytun.unknown_flag":       "неизвестный флаг: %s",
		"xraytun.want_up_usage": "ожидалось: z-panel xray-tun up [флаги] <interface> ip\n" +
			"  или: z-panel xray-tun up [флаги] <interface> <адрес[/маска]> [<peer[/маска]>]",
		"xraytun.extra_args":          "лишние аргументы: %v",
		"xraytun.empty_address":       "пустой адрес",
		"xraytun.need_ipv4":           "ожидается IPv4: %q",
		"xraytun.empty_iface":         "пустое имя интерфейса",
		"xraytun.bad_iface_char":      "недопустимый символ в имени интерфейса: %q",
		"xraytun.wrote":               "записан %s\n",
		"xraytun.file_missing":        "файл отсутствует: %s",
		"xraytun.not_managed":         "отказ: %s не помечен как созданный z-panel (нет %s) — удалите вручную при необходимости",
		"xraytun.removed":             "удалён %s\n",
		"xraytun.network_reload_fail": "перезагрузка конфигурации сети: networkctl reload и systemctl reload systemd-networkd не удались",
		"xraytun.ok_networkctl":       "networkctl reload: ok",
		"xraytun.ok_systemd":          "systemctl reload systemd-networkd: ok",
		"xraytun.file_managed_by":     "Создано z-panel xray-tun. Редактировать вручную не рекомендуется.",
		"xraytun.file_remove_hint":    "Снять: z-panel xray-tun down %s ip",

		"redirect.need_value_after": "нужно значение после %q",
		"redirect.unknown_flag":     "неизвестный флаг: %s",
		"redirect.need_one_iface":   "нужен ровно один интерфейс, получено: %v",
		"redirect.table_numeric":    "--table: нужен числовой идентификатор таблицы/fwmark: %w",
		"redirect.cgroup_missing":   "cgroup для пометки не задан: укажите --bypass-cgroup=… или --bypass-unit=…, либо --no-mark",
		"redirect.bypass_mark_fmt":  " (юнит %s)",
		"redirect.mark_line":        "z-panel: пометка обхода по cgroup v2: path=%q%s\n",
		"redirect.no_mark_line":     "z-panel: --no-mark — без пометки cgroup; ip rule как у wg-quick.\n",
		"redirect.err.default_route": "default 0.0.0.0/0 dev %s table %s: %w",
		"redirect.err.rule_fwmark":   "ip rule not fwmark: %w",
		"redirect.err.rule_suppress": "ip rule suppress_prefixlength: %w",
		"redirect.err.route6":        "маршрут ip -6 default: %w",
		"redirect.err.rule6_fw":      "правило ip -6 not fwmark: %w",
		"redirect.err.rule6_sup":     "правило ip -6 suppress_prefixlength: %w",
		"redirect.iface_not_found":   "интерфейс %q не найден: %s",
		"redirect.down_no_state":     "нет state для %q — сначала up или удалите вручную ip rule/route (интерфейс %q)",
		"redirect.down_bad_mode":     "неподдерживаемый режим state %q для %q — ожидался xray-redirect (wg/full); при необходимости снимите правила вручную",
		"redirect.down_done":         "down: очищены table %s, wg-стиль правила и firewall для %s\n",
		"redirect.ip6tables_missing": "ip6tables не найден для anti-leak IPv6",
		"redirect.iptables_raw_fail": "iptables raw: %w",
		"redirect.ip6tables_raw_fail": "ip6tables raw: %w",
		"redirect.iptables_cgroup_fail": "iptables cgroup (нужны cgroup v2 и xt_cgroup, путь от корня fs cgroup2): %w",
		"redirect.fw_skip":            "z-panel: firewall wg-стиль пропущен (нет адресов на TUN для anti-leak и нет пометки cgroup).",
		"redirect.nft_ok":             "nft: anti-leak (preraw) применён",
		"redirect.ipt_ok":             "iptables: anti-leak (raw) применён",
		"redirect.ipt_cgroup_ok":      "iptables: пометка OUTPUT по cgroup v2 (--path)",
		"redirect.ip6tables_cgroup_warn": "z-panel: предупреждение: ip6tables cgroup: %v\n",
		"redirect.cg_systemctl":      "systemctl ControlGroup для %q: %w",
		"redirect.cg_empty":          "ControlGroup пуст для %q — юнит не активен",
		"redirect.auto_unit":         "z-panel: авто-выбран systemd-юнит для cgroup: %s\n",
		"redirect.auto_fail":         "не удалось авто-определить юнит (пробовали %v): %w; укажите --bypass-unit=… или --bypass-cgroup=…",
		"redirect.wan_line":          "z-panel: ip rule pref %d: from %s lookup main (обход: входящие на публичный/uplink-адрес)\n",
		"redirect.wan_auto_err":        "z-panel: --wan-lookup=auto: не удалось прочитать main (IPv4): %v\n",
		"redirect.wan_auto_skip":       "z-panel: --wan-lookup=auto: нет подходящего IPv4 на default в main (tun %q) — пропуск правила; задайте --wan-lookup=IP/32 или маршрут\n",
		"redirect.err.wan_cidr":        "неверный --wan-lookup",
		"redirect.err.wan_rule":        "ip rule (from WAN, lookup main): %v",

		"state.state_file_err": "файл состояния: %w",
		"state.up_line":        "up: %s (сохранено в %s)\n",
		"state.summary_base":   "mode=%s table=%s fwmark=%s (wg-quick-стиль) default dev %s",
		"state.summary_nomark": " no_bypass_mark=1",
		"state.summary_bypass": " bypass_cgroup=1",

		"daemon.help": `daemon [help] [run]
  run — демон на переднем плане: HTTP API на Unix-сокете %s, периодическая перезагрузка конфига (полная сверка состояния — позже).
  Нужен root. Остановка: SIGINT/SIGTERM.
  При daemon = 1 в config.toml подкоманды уходят в этот процесс, если он доступен (исключение: z-panel daemon …).

`,
		"daemon.err_unknown_subcmd": "daemon: неизвестная подкоманда %q (ожидалось: run)",
		"daemon.fallback_warning":  "z-panel: демон недоступен, выполняю локально (запуск: z-panel daemon run)\n",

		"transport.ssh.err_argv":      "z-panel: внутренняя ошибка argv\n",
		"transport.ssh.err_duplicate": "z-panel: повторный флаг --ssh\n",
		"transport.ssh.err_conflict":  "z-panel: укажите только один из флагов --ssh или --ssh-connect\n",
		"transport.ssh.err_empty":     "z-panel: пустой хост для --ssh\n",
		"transport.ssh.err_missing":   "z-panel: нет значения после --ssh\n",
		"transport.ssh.err_empty_connect": "z-panel: пустой хост для --ssh-connect\n",
		"transport.ssh.err_missing_connect": "z-panel: нет значения после --ssh-connect\n",
		"transport.ssh.err_no_cmd":    "z-panel: нет подкоманды после --ssh / --ssh-connect\n",
		"transport.remote_forbidden":  "%s: недоступно с --ssh (установка с этой машины: z-panel --ssh=хост install; иначе — на сервере или ufw / version; config/daemon/xray-* — на хосте или --ssh-connect)\n",

		"bashcomp.line1": "# bash completion для z-panel (генерируется из команд z-panel)",
		"bashcomp.line2": "# Установка: z-panel install-shell; z-panel install выполняет это автоматически (системно).",
		"bashcomp.line3": "# Нужен bash 4+; для системной установки рекомендуется пакет bash-completion.",
	}
}
