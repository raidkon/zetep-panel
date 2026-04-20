package i18n

func zhStrings() map[string]string {
	return map[string]string{
		"root.unknown_command": "未知命令：%s\n\n",
		"root.help.tagline":    "z-panel — 经 TUN 的策略路由（适用于 Xray 的 wg-quick 风格）。",
		"root.help.top": `顶层用法：
  z-panel help | -h | --help     本帮助（所有命令摘要）
  z-panel version | -v | --version
  z-panel <command> [help | -h | --help]   单个命令的帮助
  z-panel <command> …            命令名后的参数交给该命令

命令：
`,
		"root.help.cmdline":      "  z-panel %s …\n",
		"root.help.ufw_note":     "\nUFW 注释标记：'%s: …'\n\n",
		"root.help.section_rule": "─── %s ───\n",
		"root.need_root":         "需要 root 权限（sudo）",

		"install.help": `install [help] [<sshHost>]
  本地：复制二进制到 %s（chmod 755），需 root。
  若尚无 %s — 交互式询问并保存设置。
  远程：scp，然后一次 SSH（-t）：install；若无配置则 config init（交互）。

`,
		"install.err.interrupted":      "已中断（Ctrl+C）",
		"install.err.interrupted_with": "已中断（Ctrl+C）：%w",
		"install.err.open_self":        "打开自身：%w",
		"install.err.create_tmp":       "创建 %s：%w",
		"install.err.copy":             "复制：%w",
		"install.err.rename":           "重命名为 %s：%w",
		"install.err.config":           "配置：%w",
		"install.installed":            "已安装：%s\n",
		"install.warn_completion":      "警告：bash 补全：%v\n",
		"install.err.scp":              "scp：%w",
		"install.err.ssh":              "ssh install/init：%w",
		"install.remote_done":          "已在 %s 上安装：%s\n",

		"installshell.err.home":   "主目录：%w",
		"installshell.err.mkdir":  "mkdir：%w",
		"installshell.err.write":  "写入 %s：%w",
		"installshell.done":       "已安装 bash 补全：%s\n",
		"installshell.hint_shell": "请打开新终端或执行：source /usr/share/bash-completion/bash_completion",
		"installshell.hint_user":  "（用户安装：请确认 ~/.bashrc 加载 bash-completion）",
		"installshell.help": `install-shell [help] [--user|-u]
  写入 bash 补全脚本（由已注册命令生成）。
  默认：%s（需 root）。
  --user：~/.local/share/bash-completion/completions/z-panel（或 $XDG_DATA_HOME/...）。
  需要 bash 4+；完整功能请安装 bash-completion 包。

`,

		"confcmd.err_unknown": "config：未知子命令 %q（应为 init 或 migrate）",
		"confcmd.help": `config [help] init [--force|-f]
  交互式创建或覆盖 %s（--force）。
config [help] migrate
  升级 z-panel 后应用新配置项（仅对新 schema 版本交互询问）。

`,

		"version.help": `version
  版本：%s（每次运行也会作为 stderr 首行打印）。
  顶层同义：-v、--version

`,

		"settings.err.read":               "读取 %s：%w",
		"settings.err.parse":              "解析 %s：%w",
		"settings.err.mkdir":              "mkdir %s：%w",
		"settings.err.write":              "写入 %s：%w",
		"settings.config_hdr":             "# z-panel — 配置\n\n",
		"settings.init_exists":            "配置已存在：%s（覆盖请用：z-panel config init --force）\n",
		"settings.init_intro":             "z-panel 设置 — 输入值或按 Enter 使用默认值。",
		"settings.saved":                  "\n已保存：%s\n",
		"settings.prompt.table":           "路由表 / fwmark ID",
		"settings.prompt.state_dir":       "state 目录（JSON），通常在配置旁",
		"settings.prompt.systemd_network": "systemd-networkd 单元目录",
		"settings.prompt.lan_cidr":        "UFW 模板：LAN CIDR",
		"settings.prompt.lan_dev":         "UFW 模板：LAN 网卡",
		"settings.prompt.xray_addr":       "Xray TUN：默认地址（ip 模式）",
		"settings.prompt.xray_peer":       "Xray TUN：默认对端",
		"settings.prompt.ufw_marker":      "UFW 注释标记",
		"settings.prompt.xray_mark":       ".network 文件中的标记行",
		"settings.prompt.language":        "界面语言（" + LanguageListHint + "）",
		"settings.migrate_intro":          "此配置由较旧版本的 z-panel 写入。请为新选项设置值。",
		"settings.migrate_uptodate":       "配置 schema 已是最新。",
		"settings.migrate_no_file":        "%s：未找到配置文件（请运行 z-panel config init）",

		"xrayredirect.want_up_down":    "应为：z-panel xray-redirect up|down …（见 z-panel xray-redirect help）",
		"xrayredirect.want_down_iface": "应为：z-panel xray-redirect down <interface>",
		"xrayredirect.bad_action":      "xray-redirect：未知操作 %q（应为 up 或 down）",
		"xrayredirect.help": `xray-redirect [help] up [flags] <interface>
xray-redirect [help] down <interface>
  完整隧道类似 wg-quick：not fwmark → table、suppress_prefixlength、default dev <interface>，
  sysctl src_valid_mark、nft 防泄漏、经 cgroup v2 的出口标记（iptables -m cgroup --path）。
  up 的标志（在网卡名之前）：
    --bypass-unit=auto      默认：尝试 x-ui、sing-box、xray 单元
    --bypass-unit=x-ui      显式 systemd 单元（可省略 .service）
    --bypass-cgroup=path    自根起的 cgroup v2 路径
    --table=N               表与 fwmark（默认 %s）
    --no-mark               无 cgroup 标记
    --ipv6                  default ::/0 与 IPv6 规则
  示例：
    sudo z-panel xray-redirect up xray2tun
    sudo z-panel xray-redirect up --bypass-unit=sing-box xray2tun
    sudo z-panel xray-redirect up --table=51844 --ipv6 xray2tun
    sudo z-panel xray-redirect down xray2tun

`,

		"ufw.want_check": "应为：z-panel ufw check [flags] [interface]（见 z-panel ufw help）",
		"ufw.help": `ufw [help] check [--lan-cidr=CIDR] [--lan-dev=DEV] [interface]
  带标记 %s 的 UFW 规则。模板：--lan-cidr（默认 %s）、--lan-dev（%s）。

`,
		"ufw.err.lan_cidr_empty": "--lan-cidr：空值",
		"ufw.err.lan_cidr_need":  "--lan-cidr 后需要值",
		"ufw.err.lan_dev_empty":  "--lan-dev：空值",
		"ufw.err.lan_dev_need":   "--lan-dev 后需要值",
		"ufw.err.unknown_flag":   "未知标志：%s",
		"ufw.err.too_many_iface": "最多一个接口，多余：%q",
		"ufw.ufw_status_failed":  "ufw status verbose：%w\n%s",
		"ufw.section_rules":      "=== 提及 %s 的 UFW 规则 ===",
		"ufw.no_lines":           "（无包含 «z-panel» 的行）",
		"ufw.section_hints":      "=== 建议（模板 — 请按环境核实）===",
		"ufw.hint_sysctl": `# 路由与转发（内核）：
# /etc/ufw/sysctl.conf: net.ipv4.ip_forward=1
# LAN -> 隧道转发（示例；按需调整 ufw 区域）：
sudo ufw route allow in on %s out on %s from %s comment '%s: lan to tunnel'
`,
		"ufw.hint_return":     "# 若 ufw 拦截回程：\n",
		"ufw.hint_return_cmd": "sudo ufw route allow in on %s out on %s comment '%s: return path'\n",
		"ufw.no_iface_hint1":  "请指定隧道接口以生成模板：z-panel ufw check [flags] <interface>\n",
		"ufw.no_iface_hint2":  "当前模板：LAN %s dev %s（用 --lan-cidr / --lan-dev 修改）\n",
		"ufw.section_full":    "=== 完整 ufw status verbose（手动查找）===",

		"xraytun.want_subcmd": "应为：z-panel xray-tun up … 或 down …（见 help）",
		"xraytun.want_down":   "应为：z-panel xray-tun down <interfaceName> ip",
		"xraytun.bad_action":  "xray-tun：未知操作 %q（应为 up 或 down）",
		"xraytun.help": `xray-tun [help] up [flags] <interfaceName> ip
xray-tun [help] up [flags] <interfaceName> <address[/mask]> [<peer[/mask]>]
xray-tun [help] down <interfaceName> ip
  systemd-networkd 的 TUN 地址 drop-in。
  「ip」模式：无 --address/--peer 时默认为 %s 与 %s。
  标志：--address=A（IPv4）、--peer=P（可选）。

`,
		"xraytun.need_address_value": "--address 后需要值",
		"xraytun.need_peer_value":    "--peer 后需要值",
		"xraytun.unknown_flag":       "未知标志：%s",
		"xraytun.want_up_usage": "应为：z-panel xray-tun up [flags] <interface> ip\n" +
			"  或：z-panel xray-tun up [flags] <interface> <address[/mask]> [<peer[/mask]>]",
		"xraytun.extra_args":          "多余参数：%v",
		"xraytun.empty_address":       "空地址",
		"xraytun.need_ipv4":           "需要 IPv4：%q",
		"xraytun.empty_iface":         "空接口名",
		"xraytun.bad_iface_char":      "接口名含非法字符：%q",
		"xraytun.wrote":               "已写入 %s\n",
		"xraytun.file_missing":        "文件不存在：%s",
		"xraytun.not_managed":         "拒绝：%s 未标记为 z-panel 管理（缺少 %s）— 必要时请手动删除",
		"xraytun.removed":             "已删除 %s\n",
		"xraytun.network_reload_fail": "网络重载：networkctl reload 与 systemctl reload systemd-networkd 均失败",
		"xraytun.ok_networkctl":       "networkctl reload：ok",
		"xraytun.ok_systemd":          "systemctl reload systemd-networkd：ok",
		"xraytun.file_managed_by":     "由 z-panel xray-tun 创建。不建议手动编辑。",
		"xraytun.file_remove_hint":    "移除：z-panel xray-tun down %s ip",

		"redirect.need_value_after": "%q 后需要值",
		"redirect.unknown_flag":     "未知标志：%s",
		"redirect.need_one_iface":   "需要恰好一个接口，得到：%v",
		"redirect.table_numeric":    "--table：路由表/fwmark 须为数字：%w",
		"redirect.cgroup_missing":   "未设置用于标记的 cgroup：请使用 --bypass-cgroup=… 或 --bypass-unit=…，或 --no-mark",
		"redirect.bypass_mark_fmt":  "（单元 %s）",
		"redirect.mark_line":        "z-panel：经 cgroup v2 绕过标记：path=%q%s\n",
		"redirect.no_mark_line":     "z-panel：--no-mark — 无 cgroup 标记；ip 规则同 wg-quick。\n",
		"redirect.err.default_route": "default 0.0.0.0/0 dev %s table %s：%w",
		"redirect.err.rule_fwmark":   "ip rule not fwmark：%w",
		"redirect.err.rule_suppress": "ip rule suppress_prefixlength：%w",
		"redirect.err.route6":        "ip -6 默认路由：%w",
		"redirect.err.rule6_fw":      "ip -6 rule not fwmark：%w",
		"redirect.err.rule6_sup":     "ip -6 rule suppress_prefixlength：%w",
		"redirect.iface_not_found":   "接口 %q 不存在：%s",
		"redirect.down_no_state":     "无 %q 的 state — 请先 up 或手动删除 ip rule/route（接口 %q）",
		"redirect.down_bad_mode":     "不支持的 state 模式 %q（%q）— 预期 xray-redirect（wg/full）；必要时手动删除规则",
		"redirect.down_done":         "down：已清理 table %s、wg 风格规则与 %s 的防火墙\n",
		"redirect.ip6tables_missing": "未找到 ip6tables，无法 IPv6 anti-leak",
		"redirect.iptables_raw_fail": "iptables raw：%w",
		"redirect.ip6tables_raw_fail": "ip6tables raw：%w",
		"redirect.iptables_cgroup_fail": "iptables cgroup（需要 cgroup v2 与 xt_cgroup，路径自 cgroup2 根）：%w",
		"redirect.fw_skip":            "z-panel：已跳过 wg 风格防火墙（TUN 无地址做 anti-leak 且无 cgroup 标记）。",
		"redirect.nft_ok":             "nft：anti-leak（preraw）已应用",
		"redirect.ipt_ok":             "iptables：anti-leak（raw）已应用",
		"redirect.ipt_cgroup_ok":      "iptables：经 cgroup v2（--path）标记 OUTPUT",
		"redirect.ip6tables_cgroup_warn": "z-panel：警告：ip6tables cgroup：%v\n",
		"redirect.cg_systemctl":      "systemctl ControlGroup（%q）：%w",
		"redirect.cg_empty":          "ControlGroup 对 %q 为空 — 单元未激活",
		"redirect.auto_unit":         "z-panel：自动选择 cgroup 的 systemd 单元：%s\n",
		"redirect.auto_fail":         "无法自动检测单元（已试 %v）：%w；请设置 --bypass-unit=… 或 --bypass-cgroup=…",

		"state.state_file_err": "state 文件：%w",
		"state.up_line":        "up：%s（state %s）\n",
		"state.summary_base":   "mode=%s table=%s fwmark=%s（wg-quick 风格）default dev %s",
		"state.summary_nomark": " no_bypass_mark=1",
		"state.summary_bypass": " bypass_cgroup=1",

		"bashcomp.line1": "# z-panel 的 bash 补全（由 z-panel 命令生成）",
		"bashcomp.line2": "# 安装：z-panel install-shell；z-panel install 会自动执行（系统级）。",
		"bashcomp.line3": "# 需要 bash 4+；系统级安装建议使用 bash-completion 包。",
	}
}
