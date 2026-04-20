package i18n

func english() map[string]string {
	return map[string]string{
		"root.unknown_command": "unknown command: %s\n\n",

		"root.help.tagline": "z-panel — TUN policy routing (wg-quick-style for Xray).",
		"root.help.top": `Top level:
  z-panel help | -h | --help     this summary (all commands)
  z-panel version | -v | --version
  z-panel <command> [help | -h | --help]   help for one command
  z-panel <command> …            all arguments after the command name go to that command

Commands:
`,
		"root.help.cmdline":      "  z-panel %s …\n",
		"root.help.ufw_note":     "\nUFW comment tag: '%s: …'\n\n",
		"root.help.section_rule": "─── %s ───\n",

		"root.need_root": "root privileges required (sudo)",

		"install.help": `install [help] [<sshHost>]
  Local: copies the binary to %s (chmod 755); requires root.
  If %s is missing — interactive prompts and saving settings.
  Remote: scp, then one SSH session (-t): install and, if no config — config init (interactive).

`,
		"install.err.interrupted":       "interrupted (Ctrl+C)",
		"install.err.interrupted_with":  "interrupted (Ctrl+C): %w",
		"install.err.open_self":         "open self: %w",
		"install.err.create_tmp":        "create %s: %w",
		"install.err.copy":              "copy: %w",
		"install.err.rename":            "rename to %s: %w",
		"install.err.config":            "config: %w",
		"install.installed":             "installed: %s\n",
		"install.warn_completion":       "warning: bash completion: %v\n",
		"install.err.scp":             "scp: %w",
		"install.err.ssh":             "ssh install/init: %w",
		"install.remote_done":         "installed on %s: %s\n",

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
		"settings.prompt.state_dir":       "State directory (JSON), usually next to config",
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
  Examples:
    sudo z-panel xray-redirect up xray2tun
    sudo z-panel xray-redirect up --bypass-unit=sing-box xray2tun
    sudo z-panel xray-redirect up --table=51844 --ipv6 xray2tun
    sudo z-panel xray-redirect down xray2tun

`,

		"ufw.want_check": "expected: z-panel ufw check [flags] [interface] (see z-panel ufw help)",
		"ufw.help": `ufw [help] check [--lan-cidr=CIDR] [--lan-dev=DEV] [interface]
  UFW rules with tag %s. Templates: --lan-cidr (default %s), --lan-dev (%s).

`,
		"ufw.err.lan_cidr_empty": "--lan-cidr: empty value",
		"ufw.err.lan_cidr_need":  "value required after --lan-cidr",
		"ufw.err.lan_dev_empty":  "--lan-dev: empty value",
		"ufw.err.lan_dev_need":   "value required after --lan-dev",
		"ufw.err.unknown_flag":   "unknown flag: %s",
		"ufw.err.too_many_iface": "at most one interface expected, extra: %q",
		"ufw.ufw_status_failed": "ufw status verbose: %w\n%s",
		"ufw.section_rules":       "=== UFW rules mentioning %s ===",
		"ufw.no_lines":            "(no lines containing «z-panel»)",
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

		"state.state_file_err": "state file: %w",
		"state.up_line":        "up: %s (state %s)\n",
		"state.summary_base":   "mode=%s table=%s fwmark=%s (wg-quick style) default dev %s",
		"state.summary_nomark": " no_bypass_mark=1",
		"state.summary_bypass": " bypass_cgroup=1",

		"bashcomp.line1": "# bash completion for z-panel (generated from z-panel commands)",
		"bashcomp.line2": "# Install: z-panel install-shell; z-panel install runs this automatically (system-wide).",
		"bashcomp.line3": "# Requires bash 4+; for system-wide install use the bash-completion package.",
	}
}

func russian() map[string]string {
	return map[string]string{
		"root.unknown_command": "неизвестная команда: %s\n\n",

		"root.help.tagline": "z-panel — маршрутизация через TUN (policy routing, стиль wg-quick для Xray).",
		"root.help.top": `Верхний уровень:
  z-panel help | -h | --help     эта справка (сводка по всем командам)
  z-panel version | -v | --version
  z-panel <command> [help | -h | --help]   справка только по команде
  z-panel <command> …            все аргументы после имени команды — в пакет команды

Команды:
`,
		"root.help.cmdline":      "  z-panel %s …\n",
		"root.help.ufw_note":     "\nПометка в UFW: comment '%s: …'\n\n",
		"root.help.section_rule": "─── %s ───\n",

		"root.need_root": "нужны права root (sudo)",

		"install.help": `install [help] [<sshHost>]
  Локально: копирует бинарник в %s (chmod 755), нужен root.
  Если %s ещё нет — интерактивный опрос и сохранение настроек.
  Удалённо: scp, затем одна SSH-сессия (-t): install и при отсутствии конфига — config init (интерактивно).

`,
		"install.err.interrupted":      "прервано (Ctrl+C)",
		"install.err.interrupted_with": "прервано (Ctrl+C): %w",
		"install.err.open_self":        "открыть себя: %w",
		"install.err.create_tmp":       "создать %s: %w",
		"install.err.copy":             "копирование: %w",
		"install.err.rename":           "переименовать в %s: %w",
		"install.err.config":           "конфиг: %w",
		"install.installed":            "установлено: %s\n",
		"install.warn_completion":      "предупреждение: bash completion: %v\n",
		"install.err.scp":            "scp: %w",
		"install.err.ssh":            "ssh install/init: %w",
		"install.remote_done":        "установлено на %s: %s\n",

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
		"settings.prompt.state_dir":       "Каталог state (JSON), обычно рядом с конфигом",
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
  Примеры:
    sudo z-panel xray-redirect up xray2tun
    sudo z-panel xray-redirect up --bypass-unit=sing-box xray2tun
    sudo z-panel xray-redirect up --table=51844 --ipv6 xray2tun
    sudo z-panel xray-redirect down xray2tun

`,

		"ufw.want_check": "ожидалось: z-panel ufw check [флаги] [interface] (см. z-panel ufw help)",
		"ufw.help": `ufw [help] check [--lan-cidr=CIDR] [--lan-dev=DEV] [interface]
  Правила UFW с пометкой %s. Шаблоны: --lan-cidr (по умолчанию %s), --lan-dev (%s).

`,
		"ufw.err.lan_cidr_empty": "--lan-cidr: пустое значение",
		"ufw.err.lan_cidr_need":  "нужно значение после --lan-cidr",
		"ufw.err.lan_dev_empty":  "--lan-dev: пустое значение",
		"ufw.err.lan_dev_need":   "нужно значение после --lan-dev",
		"ufw.err.unknown_flag":   "неизвестный флаг: %s",
		"ufw.err.too_many_iface": "ожидается не больше одного интерфейса, лишнее: %q",
		"ufw.ufw_status_failed":  "ufw status verbose: %w\n%s",
		"ufw.section_rules":      "=== Правила UFW с пометкой %s ===",
		"ufw.no_lines":           "(нет строк, содержащих «z-panel»)",
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

		"state.state_file_err": "файл состояния: %w",
		"state.up_line":        "up: %s (state %s)\n",
		"state.summary_base":   "mode=%s table=%s fwmark=%s (wg-quick-стиль) default dev %s",
		"state.summary_nomark": " no_bypass_mark=1",
		"state.summary_bypass": " bypass_cgroup=1",

		"bashcomp.line1": "# bash completion для z-panel (генерируется из команд z-panel)",
		"bashcomp.line2": "# Установка: z-panel install-shell; z-panel install выполняет это автоматически (системно).",
		"bashcomp.line3": "# Нужен bash 4+; для системной установки рекомендуется пакет bash-completion.",
	}
}
