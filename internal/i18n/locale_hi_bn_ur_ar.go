package i18n

func hiStrings() map[string]string {
	return map[string]string{
		"root.unknown_command": "अज्ञात कमांड: %s\n\n",
		"root.help.tagline":    "z-panel — TUN के माध्यम से पॉलिसी रूटिंग (Xray हेतु wg-quick-शैली)।",
		"root.help.top": `शीर्ष स्तर:
  z-panel help | -h | --help     यह सारांश (सभी कमांड)
  z-panel version | -v | --version
  z-panel <command> [help | -h | --help]   एक कमांड की सहायता
  z-panel <command> …            बाकी तर्क उसी कमांड को

कमांड:
`,
		"root.help.cmdline": "  z-panel %s …\n", "root.help.ufw_note": "\nUFW टैग: '%s: …'\n\n",
		"root.help.section_rule": "─── %s ───\n", "root.need_root": "root अधिकार आवश्यक (sudo)",
		"install.help": `install [help] [<sshHost>]
  स्थानीय: बाइनरी %s पर कॉपी (chmod 755); root चाहिए।
  यदि %s नहीं — इंटरैक्टिव सेटअप।
  दूरस्थ: scp फिर SSH (-t): install; बिना config — config init।

`,
		"install.err.interrupted": "रद्द (Ctrl+C)", "install.err.interrupted_with": "रद्द (Ctrl+C): %w",
		"install.err.open_self": "स्वयं खोलने में: %w", "install.err.create_tmp": "%s बनाने में: %w",
		"install.err.copy": "कॉपी: %w", "install.err.rename": "%s पर नाम बदलने में: %w", "install.err.config": "कॉन्फ़िग: %w",
		"install.installed": "स्थापित: %s\n", "install.warn_completion": "चेतावनी bash completion: %v\n",
		"install.err.scp": "scp: %w", "install.err.ssh": "ssh: %w", "install.remote_done": "%s पर स्थापित: %s\n",
		"installshell.err.home": "होम डायरेक्टरी: %w", "installshell.err.mkdir": "mkdir: %w", "installshell.err.write": "लेखन %s: %w",
		"installshell.done": "bash completion स्थापित: %s\n",
		"installshell.hint_shell": "नया टर्मिनल या: source /usr/share/bash-completion/bash_completion",
		"installshell.hint_user": "(उपयोगकर्ता: ~/.bashrc जाँचें)",
		"installshell.help": `install-shell [help] [--user|-u]
  bash completion स्क्रिप्ट लिखता है। डिफ़ॉल्ट: %s (root)।
  bash 4+; bash-completion पैकेज।

`,
		"confcmd.err_unknown": "config: अज्ञात उप-कमांड %q (init या migrate)",
		"confcmd.help": `config [help] init [--force|-f]
  %s इंटरैक्टिव बनाएँ/अधिलेखित (--force)।
config [help] migrate
  अपग्रेड के बाद नई कुंजियाँ।

`,
		"version.help": `version
  संस्करण: %s (stderr की पहली पंक्ति भी)।
  समानार्थी: -v, --version

`,
		"settings.err.read": "पढ़ना %s: %w", "settings.err.parse": "पार्स %s: %w", "settings.err.mkdir": "mkdir %s: %w",
		"settings.err.write": "लेखन %s: %w", "settings.config_hdr": "# z-panel — कॉन्फ़िगरेशन\n\n",
		"settings.init_exists": "कॉन्फ़िग पहले से: %s (z-panel config init --force)\n",
		"settings.init_intro": "z-panel सेटअप — मान दर्ज करें या डिफ़ॉल्ट हेतु Enter।",
		"settings.saved": "\nसहेजा: %s\n",
		"settings.prompt.table": "रूटिंग टेबल / fwmark", "settings.prompt.state_dir": "state निर्देशिका (JSON)",
		"settings.prompt.systemd_network": "systemd-networkd निर्देशिका", "settings.prompt.lan_cidr": "UFW LAN CIDR",
		"settings.prompt.lan_dev": "UFW LAN इंटरफ़ेस", "settings.prompt.xray_addr": "Xray TUN डिफ़ॉल्ट पता",
		"settings.prompt.xray_peer": "Xray TUN डिफ़ॉल्ट पीयर", "settings.prompt.ufw_marker": "UFW टिप्पणी टैग",
		"settings.prompt.xray_mark": ".network मार्कर",
		"settings.prompt.language": "UI भाषा (" + LanguageListHint + ")",
		"settings.migrate_intro": "पुरानी z-panel कॉन्फ़िग। नए विकल्प सेट करें।",
		"settings.migrate_uptodate": "कॉन्फ़िग स्कीमा अप टू डेट।",
		"settings.migrate_no_file": "%s: कॉन्फ़िग फ़ाइल नहीं (z-panel config init)",
		"xrayredirect.want_up_down": "अपेक्षित: z-panel xray-redirect up|down …",
		"xrayredirect.want_down_iface": "अपेक्षित: z-panel xray-redirect down <interface>",
		"xrayredirect.bad_action": "xray-redirect: अज्ञात क्रिया %q",
		"xrayredirect.help": `xray-redirect [help] up [flags] <interface>
xray-redirect [help] down <interface>
  wg-quick जैसा पूरा टनल; cgroup v2। डिफ़ॉल्ट टेबल %s

`,
		"ufw.want_check": "अपेक्षित: z-panel ufw check …", "ufw.help": `ufw [help] check …
  UFW नियम %s। --lan-cidr (डि. %s), --lan-dev (%s)।

`,
		"ufw.err.lan_cidr_empty": "--lan-cidr खाली", "ufw.err.lan_cidr_need": "--lan-cidr के बाद मान",
		"ufw.err.lan_dev_empty": "--lan-dev खाली", "ufw.err.lan_dev_need": "--lan-dev के बाद मान",
		"ufw.err.unknown_flag": "अज्ञात फ़्लैग: %s", "ufw.err.too_many_iface": "अधिकतम एक इंटरफ़ेस, अतिरिक्त: %q",
		"ufw.ufw_status_failed": "ufw status verbose: %w\n%s", "ufw.section_rules": "=== UFW नियम %s ===",
		"ufw.no_lines": "(z-panel पंक्तियाँ नहीं)", "ufw.section_hints": "=== सुझाव ===",
		"ufw.hint_sysctl": `sudo ufw route allow in on %s out on %s from %s comment '%s: lan to tunnel'
`,
		"ufw.hint_return": "# यदि ufw उत्तर रोकता है:\n",
		"ufw.hint_return_cmd": "sudo ufw route allow in on %s out on %s comment '%s: return path'\n",
		"ufw.no_iface_hint1": "टनल इंटरफ़ेस निर्दिष्ट करें।\n", "ufw.no_iface_hint2": "टेम्पलेट: LAN %s dev %s\n",
		"ufw.section_full": "=== पूरा ufw status verbose ===",
		"xraytun.want_subcmd": "अपेक्षित: xray-tun up/down …", "xraytun.want_down": "अपेक्षित: xray-tun down <iface> ip",
		"xraytun.bad_action": "xray-tun: अज्ञात क्रिया %q",
		"xraytun.help": `xray-tun [help] up [flags] <iface> ip …
  systemd-networkd। डिफ़ॉल्ट %s और %s।

`,
		"xraytun.need_address_value": "--address के बाद मान", "xraytun.need_peer_value": "--peer के बाद मान",
		"xraytun.unknown_flag": "अज्ञात फ़्लैग: %s", "xraytun.want_up_usage": "up सिंटैक्स: help देखें",
		"xraytun.extra_args": "अतिरिक्त तर्क: %v", "xraytun.empty_address": "खाली पता",
		"xraytun.need_ipv4": "IPv4 अपेक्षित: %q", "xraytun.empty_iface": "खाली इंटरफ़ेस",
		"xraytun.bad_iface_char": "अमान्य वर्ण: %q", "xraytun.wrote": "लिखा %s\n",
		"xraytun.file_missing": "फ़ाइल नहीं: %s", "xraytun.not_managed": "अस्वीकृत: %s z-panel द्वारा प्रबंधित नहीं (%s)",
		"xraytun.removed": "हटाया %s\n", "xraytun.network_reload_fail": "नेटवर्क रीलोड विफल",
		"xraytun.ok_networkctl": "networkctl reload: ok", "xraytun.ok_systemd": "systemctl reload systemd-networkd: ok",
		"xraytun.file_managed_by": "z-panel xray-tun द्वारा बनाया गया।", "xraytun.file_remove_hint": "हटाएँ: z-panel xray-tun down %s ip",
		"redirect.need_value_after": "%q के बाद मान", "redirect.unknown_flag": "अज्ञात फ़्लैग: %s",
		"redirect.need_one_iface": "एक इंटरफ़ेस चाहिए: %v", "redirect.table_numeric": "--table संख्यात्मक: %w",
		"redirect.cgroup_missing": "cgroup सेट नहीं", "redirect.bypass_mark_fmt": " (यूनिट %s)",
		"redirect.mark_line": "z-panel: cgroup path=%q%s\n", "redirect.no_mark_line": "z-panel: --no-mark\n",
		"redirect.err.default_route": "डिफ़ॉल्ट रूट: %w", "redirect.err.rule_fwmark": "ip rule: %w",
		"redirect.err.rule_suppress": "suppress: %w", "redirect.err.route6": "IPv6 रूट: %w",
		"redirect.err.rule6_fw": "IPv6 नियम: %w", "redirect.err.rule6_sup": "IPv6 suppress: %w",
		"redirect.iface_not_found": "इंटरफ़ेस %q: %s", "redirect.down_no_state": "state नहीं %q",
		"redirect.down_bad_mode": "अमान्य state मोड %q", "redirect.down_done": "down: टेबल %s साफ\n",
		"redirect.ip6tables_missing": "ip6tables नहीं", "redirect.iptables_raw_fail": "iptables raw: %w",
		"redirect.ip6tables_raw_fail": "ip6tables raw: %w", "redirect.iptables_cgroup_fail": "iptables cgroup: %w",
		"redirect.fw_skip": "फ़ायरवॉल छोड़ा।", "redirect.nft_ok": "nft लागू", "redirect.ipt_ok": "iptables लागू",
		"redirect.ipt_cgroup_ok": "cgroup लागू", "redirect.ip6tables_cgroup_warn": "चेतावनी ip6tables: %v\n",
		"redirect.cg_systemctl": "ControlGroup %q: %w", "redirect.cg_empty": "खाली ControlGroup %q",
		"redirect.auto_unit": "ऑटो यूनिट: %s\n", "redirect.auto_fail": "यूनिट पता नहीं: %w",
		"state.state_file_err": "state फ़ाइल: %w", "state.up_line": "up: %s (state %s)\n",
		"state.summary_base": "mode=%s table=%s fwmark=%s dev %s",
		"state.summary_nomark": " no_bypass_mark=1", "state.summary_bypass": " bypass_cgroup=1",
		"bashcomp.line1": "# bash completion z-panel", "bashcomp.line2": "# z-panel install-shell",
		"bashcomp.line3": "# Bash 4+",
	}
}

func bnStrings() map[string]string {
	return map[string]string{
		"root.unknown_command": "অজানা কমান্ড: %s\n\n",
		"root.help.tagline":    "z-panel — TUN-এর মাধ্যমে পলিসি রাউটিং (Xray-এর জন্য wg-quick-স্টাইল)।",
		"root.help.top": `শীর্ষ স্তর:
  z-panel help | -h | --help     এই সারাংশ
  z-panel version | -v | --version
  z-panel <command> …

কমান্ড:
`,
		"root.help.cmdline": "  z-panel %s …\n", "root.help.ufw_note": "\nUFW ট্যাগ: '%s: …'\n\n",
		"root.help.section_rule": "─── %s ───\n", "root.need_root": "root প্রয়োজন (sudo)",
		"install.help": `install [help] [<sshHost>]
  স্থানীয়: বাইনারি %s-এ কপি (chmod 755); root লাগবে।
  %s না থাকলে — ইন্টারঅ্যাক্টিভ সেটআপ।
  রিমোট: scp তারপর SSH (-t)।

`,
		"install.err.interrupted": "বাতিল (Ctrl+C)", "install.err.interrupted_with": "বাতিল (Ctrl+C): %w",
		"install.err.open_self": "নিজেকে খুলতে: %w", "install.err.create_tmp": "%s তৈরি: %w",
		"install.err.copy": "কপি: %w", "install.err.rename": "%s-এ নাম বদল: %w", "install.err.config": "কনফিগ: %w",
		"install.installed": "ইনস্টল: %s\n", "install.warn_completion": "সতর্কতা bash completion: %v\n",
		"install.err.scp": "scp: %w", "install.err.ssh": "ssh: %w", "install.remote_done": "%s-এ ইনস্টল: %s\n",
		"installshell.err.home": "হোম: %w", "installshell.err.mkdir": "mkdir: %w", "installshell.err.write": "লেখা %s: %w",
		"installshell.done": "bash completion ইনস্টল: %s\n",
		"installshell.hint_shell": "নতুন টার্মিনাল বা: source /usr/share/bash-completion/bash_completion",
		"installshell.hint_user": "(ব্যবহারকারী: ~/.bashrc)",
		"installshell.help": `install-shell [help] [--user|-u]
  bash completion লেখে। ডিফল্ট: %s (root)।

`,
		"confcmd.err_unknown": "config: অজানা উপ-কমান্ড %q",
		"confcmd.help": `config [help] init [--force|-f]
  %s ইন্টারঅ্যাক্টিভ তৈরি (--force)।
config [help] migrate
  আপগ্রেডের পর নতুন কী।

`,
		"version.help": `version
  সংস্করণ: %s
  -v, --version

`,
		"settings.err.read": "পড়া %s: %w", "settings.err.parse": "পার্স %s: %w", "settings.err.mkdir": "mkdir %s: %w",
		"settings.err.write": "লেখা %s: %w", "settings.config_hdr": "# z-panel — কনফিগারেশন\n\n",
		"settings.init_exists": "কনফিগ আগে থেকে: %s\n",
		"settings.init_intro": "z-panel সেটআপ — মান বা Enter।",
		"settings.saved": "\nসংরক্ষিত: %s\n",
		"settings.prompt.table": "রাউটিং টেবিল / fwmark", "settings.prompt.state_dir": "state ডিরেক্টরি",
		"settings.prompt.systemd_network": "systemd-networkd ডিরেক্টরি", "settings.prompt.lan_cidr": "UFW LAN CIDR",
		"settings.prompt.lan_dev": "UFW LAN ইন্টারফেস", "settings.prompt.xray_addr": "Xray TUN ডিফল্ট ঠিকানা",
		"settings.prompt.xray_peer": "Xray TUN পিয়ার", "settings.prompt.ufw_marker": "UFW ট্যাগ",
		"settings.prompt.xray_mark": ".network মার্কার",
		"settings.prompt.language": "UI ভাষা (" + LanguageListHint + ")",
		"settings.migrate_intro": "পুরনো z-panel কনফিগ। নতুন বিকল্প সেট করুন।",
		"settings.migrate_uptodate": "কনফিগ স্কিমা আপ টু ডেট।",
		"settings.migrate_no_file": "%s: কনফিগ ফাইল নেই",
		"xrayredirect.want_up_down": "প্রত্যাশিত: xray-redirect up|down …",
		"xrayredirect.want_down_iface": "প্রত্যাশিত: xray-redirect down <interface>",
		"xrayredirect.bad_action": "xray-redirect: অজানা অ্যাকশন %q",
		"xrayredirect.help": `xray-redirect [help] up [flags] <interface>
xray-redirect [help] down <interface>
  wg-quick-স্টাইল টানেল। ডিফল্ট টেবিল %s

`,
		"ufw.want_check": "প্রত্যাশিত: ufw check …", "ufw.help": `ufw [help] check …
  UFW নিয়ম %s। %s, %s।

`,
		"ufw.err.lan_cidr_empty": "--lan-cidr খালি", "ufw.err.lan_cidr_need": "--lan-cidr পরে মান",
		"ufw.err.lan_dev_empty": "--lan-dev খালি", "ufw.err.lan_dev_need": "--lan-dev পরে মান",
		"ufw.err.unknown_flag": "অজানা ফ্ল্যাগ: %s", "ufw.err.too_many_iface": "একটি ইন্টারফেস, অতিরিক্ত: %q",
		"ufw.ufw_status_failed": "ufw: %w\n%s", "ufw.section_rules": "=== UFW %s ===",
		"ufw.no_lines": "(কোনো z-panel লাইন নেই)", "ufw.section_hints": "=== পরামর্শ ===",
		"ufw.hint_sysctl": `sudo ufw route allow in on %s out on %s from %s comment '%s: lan to tunnel'
`,
		"ufw.hint_return": "# ufw উত্তর ব্লক করলে:\n",
		"ufw.hint_return_cmd": "sudo ufw route allow in on %s out on %s comment '%s: return path'\n",
		"ufw.no_iface_hint1": "টানেল ইন্টারফেস নির্দিষ্ট করুন।\n", "ufw.no_iface_hint2": "টেমপ্লেট: LAN %s dev %s\n",
		"ufw.section_full": "=== সম্পূর্ণ ufw status ===",
		"xraytun.want_subcmd": "প্রত্যাশিত: xray-tun …", "xraytun.want_down": "প্রত্যাশিত: xray-tun down <iface> ip",
		"xraytun.bad_action": "xray-tun: অজানা %q",
		"xraytun.help": `xray-tun [help] up …
  systemd-networkd। %s, %s।

`,
		"xraytun.need_address_value": "--address পরে মান", "xraytun.need_peer_value": "--peer পরে মান",
		"xraytun.unknown_flag": "অজানা ফ্ল্যাগ: %s", "xraytun.want_up_usage": "up সিনট্যাক্স: help",
		"xraytun.extra_args": "অতিরিক্ত: %v", "xraytun.empty_address": "খালি ঠিকানা",
		"xraytun.need_ipv4": "IPv4: %q", "xraytun.empty_iface": "খালি ইন্টারফেস",
		"xraytun.bad_iface_char": "অবৈধ অক্ষর: %q", "xraytun.wrote": "লেখা %s\n",
		"xraytun.file_missing": "ফাইল নেই: %s", "xraytun.not_managed": "প্রত্যাখ্যান: %s (%s)",
		"xraytun.removed": "মুছে ফেলা %s\n", "xraytun.network_reload_fail": "নেটওয়ার্ক রিলোড ব্যর্থ",
		"xraytun.ok_networkctl": "networkctl reload: ok", "xraytun.ok_systemd": "systemctl reload: ok",
		"xraytun.file_managed_by": "z-panel xray-tun তৈরি।", "xraytun.file_remove_hint": "সরান: z-panel xray-tun down %s ip",
		"redirect.need_value_after": "%q পরে মান", "redirect.unknown_flag": "অজানা ফ্ল্যাগ: %s",
		"redirect.need_one_iface": "একটি ইন্টারফেস: %v", "redirect.table_numeric": "--table সংখ্যা: %w",
		"redirect.cgroup_missing": "cgroup সেট নয়", "redirect.bypass_mark_fmt": " (ইউনিট %s)",
		"redirect.mark_line": "z-panel: cgroup path=%q%s\n", "redirect.no_mark_line": "z-panel: --no-mark\n",
		"redirect.err.default_route": "রুট: %w", "redirect.err.rule_fwmark": "ip rule: %w",
		"redirect.err.rule_suppress": "suppress: %w", "redirect.err.route6": "IPv6: %w",
		"redirect.err.rule6_fw": "IPv6 নিয়ম: %w", "redirect.err.rule6_sup": "IPv6 sup: %w",
		"redirect.iface_not_found": "ইন্টারফেস %q: %s", "redirect.down_no_state": "state নেই %q",
		"redirect.down_bad_mode": "অবৈধ মোড %q", "redirect.down_done": "down: টেবিল %s\n",
		"redirect.ip6tables_missing": "ip6tables নেই", "redirect.iptables_raw_fail": "iptables: %w",
		"redirect.ip6tables_raw_fail": "ip6tables: %w", "redirect.iptables_cgroup_fail": "cgroup: %w",
		"redirect.fw_skip": "ফায়ারওয়াল এড়ানো।", "redirect.nft_ok": "nft", "redirect.ipt_ok": "iptables",
		"redirect.ipt_cgroup_ok": "cgroup", "redirect.ip6tables_cgroup_warn": "সতর্কতা: %v\n",
		"redirect.cg_systemctl": "ControlGroup %q: %w", "redirect.cg_empty": "খালি %q",
		"redirect.auto_unit": "অটো ইউনিট: %s\n", "redirect.auto_fail": "সনাক্তকরণ ব্যর্থ: %w",
		"state.state_file_err": "state: %w", "state.up_line": "up: %s (%s)\n",
		"state.summary_base": "mode=%s table=%s fwmark=%s dev %s",
		"state.summary_nomark": " no_bypass_mark=1", "state.summary_bypass": " bypass_cgroup=1",
		"bashcomp.line1": "# bash completion z-panel", "bashcomp.line2": "# install-shell",
		"bashcomp.line3": "# Bash 4+",
	}
}

func urStrings() map[string]string {
	return map[string]string{
		"root.unknown_command": "نامعلوم کمانڈ: %s\n\n",
		"root.help.tagline":    "z-panel — TUN کے ذریعے پالیسی روٹنگ (Xray کے لیے wg-quick-انداز)۔",
		"root.help.top": `اوپری سطح:
  z-panel help | -h | --help     یہ خلاصہ
  z-panel version | -v | --version
  z-panel <command> …

کمانڈز:
`,
		"root.help.cmdline": "  z-panel %s …\n", "root.help.ufw_note": "\nUFW ٹیگ: '%s: …'\n\n",
		"root.help.section_rule": "─── %s ───\n", "root.need_root": "root درکار (sudo)",
		"install.help": `install [help] [<sshHost>]
  مقامی: بائنری %s پر کاپی (chmod 755)؛ root۔
  اگر %s نہیں — انٹرایکٹو سیٹ اپ۔
  ریموٹ: scp پھر SSH (-t)۔

`,
		"install.err.interrupted": "منقطع (Ctrl+C)", "install.err.interrupted_with": "منقطع (Ctrl+C): %w",
		"install.err.open_self": "کھولنے میں: %w", "install.err.create_tmp": "%s بنانا: %w",
		"install.err.copy": "کاپی: %w", "install.err.rename": "%s نام: %w", "install.err.config": "کنفیگ: %w",
		"install.installed": "انسٹال: %s\n", "install.warn_completion": "انتباہ bash completion: %v\n",
		"install.err.scp": "scp: %w", "install.err.ssh": "ssh: %w", "install.remote_done": "%s پر: %s\n",
		"installshell.err.home": "ہوم: %w", "installshell.err.mkdir": "mkdir: %w", "installshell.err.write": "لکھنا %s: %w",
		"installshell.done": "bash completion: %s\n",
		"installshell.hint_shell": "نیا ٹرمینل یا: source /usr/share/bash-completion/bash_completion",
		"installshell.hint_user": "(صارف: ~/.bashrc)",
		"installshell.help": `install-shell [help] [--user|-u]
  bash completion لکھتا ہے۔ ڈیفالٹ: %s (root)۔

`,
		"confcmd.err_unknown": "config: نامعلوم ذیلی کمانڈ %q",
		"confcmd.help": `config [help] init [--force|-f]
  %s انٹرایکٹو (--force)۔
config [help] migrate
  اپ گریڈ کے بعد نئی کلیدیں۔

`,
		"version.help": `version
  ورژن: %s
  -v, --version

`,
		"settings.err.read": "پڑھنا %s: %w", "settings.err.parse": "پارس %s: %w", "settings.err.mkdir": "mkdir %s: %w",
		"settings.err.write": "لکھنا %s: %w", "settings.config_hdr": "# z-panel — ترتیب\n\n",
		"settings.init_exists": "کنفیگ پہلے سے: %s\n",
		"settings.init_intro": "z-panel سیٹ اپ — قدر یا Enter۔",
		"settings.saved": "\nمحفوظ: %s\n",
		"settings.prompt.table": "روٹنگ ٹیبل / fwmark", "settings.prompt.state_dir": "state فولڈر",
		"settings.prompt.systemd_network": "systemd-networkd", "settings.prompt.lan_cidr": "UFW LAN CIDR",
		"settings.prompt.lan_dev": "UFW LAN انٹرفیس", "settings.prompt.xray_addr": "Xray TUN ڈیفالٹ",
		"settings.prompt.xray_peer": "Xray TUN peer", "settings.prompt.ufw_marker": "UFW ٹیگ",
		"settings.prompt.xray_mark": ".network مارکر",
		"settings.prompt.language": "UI زبان (" + LanguageListHint + ")",
		"settings.migrate_intro": "پرانی z-panel کنفیگ۔ نئے اختیارات۔",
		"settings.migrate_uptodate": "سکیم تازہ۔",
		"settings.migrate_no_file": "%s: کنفیگ فائل نہیں",
		"xrayredirect.want_up_down": "متوقع: xray-redirect up|down …",
		"xrayredirect.want_down_iface": "متوقع: xray-redirect down <interface>",
		"xrayredirect.bad_action": "xray-redirect: نامعلوم %q",
		"xrayredirect.help": `xray-redirect [help] up [flags] <interface>
xray-redirect [help] down <interface>
  wg-quick جیسا ٹنل۔ ڈیفالٹ ٹیبل %s

`,
		"ufw.want_check": "متوقع: ufw check …", "ufw.help": `ufw [help] check …
  UFW %s۔ %s, %s۔

`,
		"ufw.err.lan_cidr_empty": "--lan-cidr خالی", "ufw.err.lan_cidr_need": "--lan-cidr بعد قدر",
		"ufw.err.lan_dev_empty": "--lan-dev خالی", "ufw.err.lan_dev_need": "--lan-dev بعد قدر",
		"ufw.err.unknown_flag": "نامعلوم فلیگ: %s", "ufw.err.too_many_iface": "ایک انٹرفیس، اضافی: %q",
		"ufw.ufw_status_failed": "ufw: %w\n%s", "ufw.section_rules": "=== UFW %s ===",
		"ufw.no_lines": "(کوئی z-panel لائن نہیں)", "ufw.section_hints": "=== تجاویز ===",
		"ufw.hint_sysctl": `sudo ufw route allow in on %s out on %s from %s comment '%s: lan to tunnel'
`,
		"ufw.hint_return": "# اگر ufw جواب روکے:\n",
		"ufw.hint_return_cmd": "sudo ufw route allow in on %s out on %s comment '%s: return path'\n",
		"ufw.no_iface_hint1": "ٹنل انٹرفیس بتائیں۔\n", "ufw.no_iface_hint2": "سانچے: LAN %s dev %s\n",
		"ufw.section_full": "=== مکمل ufw status ===",
		"xraytun.want_subcmd": "متوقع: xray-tun …", "xraytun.want_down": "متوقع: xray-tun down <iface> ip",
		"xraytun.bad_action": "xray-tun: نامعلوم %q",
		"xraytun.help": `xray-tun [help] up …
  systemd-networkd۔ %s, %s۔

`,
		"xraytun.need_address_value": "--address بعد", "xraytun.need_peer_value": "--peer بعد",
		"xraytun.unknown_flag": "نامعلوم فلیگ: %s", "xraytun.want_up_usage": "up: help",
		"xraytun.extra_args": "اضافی: %v", "xraytun.empty_address": "خالی پتہ",
		"xraytun.need_ipv4": "IPv4: %q", "xraytun.empty_iface": "خالی انٹرفیس",
		"xraytun.bad_iface_char": "غلط حرف: %q", "xraytun.wrote": "لکھا %s\n",
		"xraytun.file_missing": "فائل نہیں: %s", "xraytun.not_managed": "مسترد: %s (%s)",
		"xraytun.removed": "حذف %s\n", "xraytun.network_reload_fail": "ری لوڈ ناکام",
		"xraytun.ok_networkctl": "networkctl reload: ok", "xraytun.ok_systemd": "systemctl reload: ok",
		"xraytun.file_managed_by": "z-panel xray-tun۔", "xraytun.file_remove_hint": "ہٹائیں: z-panel xray-tun down %s ip",
		"redirect.need_value_after": "%q بعد", "redirect.unknown_flag": "نامعلوم فلیگ: %s",
		"redirect.need_one_iface": "ایک انٹرفیس: %v", "redirect.table_numeric": "--table عدد: %w",
		"redirect.cgroup_missing": "cgroup سیٹ نہیں", "redirect.bypass_mark_fmt": " (یونٹ %s)",
		"redirect.mark_line": "z-panel: cgroup path=%q%s\n", "redirect.no_mark_line": "z-panel: --no-mark\n",
		"redirect.err.default_route": "روٹ: %w", "redirect.err.rule_fwmark": "ip rule: %w",
		"redirect.err.rule_suppress": "suppress: %w", "redirect.err.route6": "IPv6: %w",
		"redirect.err.rule6_fw": "IPv6: %w", "redirect.err.rule6_sup": "IPv6 sup: %w",
		"redirect.iface_not_found": "انٹرفیس %q: %s", "redirect.down_no_state": "state نہیں %q",
		"redirect.down_bad_mode": "غلط موڈ %q", "redirect.down_done": "down: ٹیبل %s\n",
		"redirect.ip6tables_missing": "ip6tables نہیں", "redirect.iptables_raw_fail": "iptables: %w",
		"redirect.ip6tables_raw_fail": "ip6tables: %w", "redirect.iptables_cgroup_fail": "cgroup: %w",
		"redirect.fw_skip": "فائر وال چھوڑا۔", "redirect.nft_ok": "nft", "redirect.ipt_ok": "iptables",
		"redirect.ipt_cgroup_ok": "cgroup", "redirect.ip6tables_cgroup_warn": "انتباہ: %v\n",
		"redirect.cg_systemctl": "ControlGroup %q: %w", "redirect.cg_empty": "خالی %q",
		"redirect.auto_unit": "آٹو یونٹ: %s\n", "redirect.auto_fail": "ناکام: %w",
		"state.state_file_err": "state: %w", "state.up_line": "up: %s (%s)\n",
		"state.summary_base": "mode=%s table=%s fwmark=%s dev %s",
		"state.summary_nomark": " no_bypass_mark=1", "state.summary_bypass": " bypass_cgroup=1",
		"bashcomp.line1": "# bash completion z-panel", "bashcomp.line2": "# install-shell",
		"bashcomp.line3": "# Bash 4+",
	}
}

func arStrings() map[string]string {
	return map[string]string{
		"root.unknown_command": "أمر غير معروف: %s\n\n",
		"root.help.tagline":    "z-panel — توجيه بالسياسات عبر TUN (نمط wg-quick لـ Xray).",
		"root.help.top": `المستوى الأعلى:
  z-panel help | -h | --help     هذا الملخص
  z-panel version | -v | --version
  z-panel <command> …

الأوامر:
`,
		"root.help.cmdline": "  z-panel %s …\n", "root.help.ufw_note": "\nوسم UFW: '%s: …'\n\n",
		"root.help.section_rule": "─── %s ───\n", "root.need_root": "صلاحيات الجذر مطلوبة (sudo)",
		"install.help": `install [help] [<sshHost>]
  محلي: نسخ الثنائي إلى %s (chmod 755)؛ يحتاج root.
  إن لم يوجد %s — إعداد تفاعلي.
  بعيد: scp ثم SSH (-t)۔

`,
		"install.err.interrupted": "أُلغي (Ctrl+C)", "install.err.interrupted_with": "أُلغي (Ctrl+C): %w",
		"install.err.open_self": "فتح الملف: %w", "install.err.create_tmp": "إنشاء %s: %w",
		"install.err.copy": "نسخ: %w", "install.err.rename": "إعادة تسمية إلى %s: %w", "install.err.config": "إعداد: %w",
		"install.installed": "تم التثبيت: %s\n", "install.warn_completion": "تحذير bash completion: %v\n",
		"install.err.scp": "scp: %w", "install.err.ssh": "ssh: %w", "install.remote_done": "على %s: %s\n",
		"installshell.err.home": "المجلد الرئيسي: %w", "installshell.err.mkdir": "mkdir: %w", "installshell.err.write": "كتابة %s: %w",
		"installshell.done": "bash completion: %s\n",
		"installshell.hint_shell": "طرفية جديدة أو: source /usr/share/bash-completion/bash_completion",
		"installshell.hint_user": "(مستخدم: ~/.bashrc)",
		"installshell.help": `install-shell [help] [--user|-u]
  يكتب bash completion. الافتراضي: %s (root)۔

`,
		"confcmd.err_unknown": "config: أمر فرعي غير معروف %q",
		"confcmd.help": `config [help] init [--force|-f]
  إنشاء %s تفاعلياً (--force)۔
config [help] migrate
  مفاتيح جديدة بعد الترقية۔

`,
		"version.help": `version
  الإصدار: %s
  -v, --version

`,
		"settings.err.read": "قراءة %s: %w", "settings.err.parse": "تحليل %s: %w", "settings.err.mkdir": "mkdir %s: %w",
		"settings.err.write": "كتابة %s: %w", "settings.config_hdr": "# z-panel — إعداد\n\n",
		"settings.init_exists": "الإعداد موجود: %s\n",
		"settings.init_intro": "إعداد z-panel — أدخل قيمة أو Enter للافتراضي.",
		"settings.saved": "\nحُفظ: %s\n",
		"settings.prompt.table": "جدول التوجيه / fwmark", "settings.prompt.state_dir": "مجلد state",
		"settings.prompt.systemd_network": "systemd-networkd", "settings.prompt.lan_cidr": "UFW LAN CIDR",
		"settings.prompt.lan_dev": "UFW واجهة LAN", "settings.prompt.xray_addr": "Xray TUN افتراضي",
		"settings.prompt.xray_peer": "Xray TUN peer", "settings.prompt.ufw_marker": "وسم UFW",
		"settings.prompt.xray_mark": "علامة .network",
		"settings.prompt.language": "لغة الواجهة (" + LanguageListHint + ")",
		"settings.migrate_intro": "إعداد من إصدار قديم. عيّن الخيارات الجديدة.",
		"settings.migrate_uptodate": "المخطط محدّث.",
		"settings.migrate_no_file": "%s: لا ملف إعداد",
		"xrayredirect.want_up_down": "متوقع: xray-redirect up|down …",
		"xrayredirect.want_down_iface": "متوقع: xray-redirect down <interface>",
		"xrayredirect.bad_action": "xray-redirect: إجراء غير معروف %q",
		"xrayredirect.help": `xray-redirect [help] up [flags] <interface>
xray-redirect [help] down <interface>
  نفق مثل wg-quick. الجدول الافتراضي %s

`,
		"ufw.want_check": "متوقع: ufw check …", "ufw.help": `ufw [help] check …
  قواعد UFW %s۔ %s، %s۔

`,
		"ufw.err.lan_cidr_empty": "--lan-cidr فارغ", "ufw.err.lan_cidr_need": "قيمة بعد --lan-cidr",
		"ufw.err.lan_dev_empty": "--lan-dev فارغ", "ufw.err.lan_dev_need": "قيمة بعد --lan-dev",
		"ufw.err.unknown_flag": "خيار غير معروف: %s", "ufw.err.too_many_iface": "واجهة واحدة كحد أقصى، زائد: %q",
		"ufw.ufw_status_failed": "ufw: %w\n%s", "ufw.section_rules": "=== قواعد UFW %s ===",
		"ufw.no_lines": "(لا أسطر z-panel)", "ufw.section_hints": "=== اقتراحات ===",
		"ufw.hint_sysctl": `sudo ufw route allow in on %s out on %s from %s comment '%s: lan to tunnel'
`,
		"ufw.hint_return": "# إذا حظر ufw الردود:\n",
		"ufw.hint_return_cmd": "sudo ufw route allow in on %s out on %s comment '%s: return path'\n",
		"ufw.no_iface_hint1": "حدد واجهة النفق.\n", "ufw.no_iface_hint2": "قوالب: LAN %s dev %s\n",
		"ufw.section_full": "=== ufw status كامل ===",
		"xraytun.want_subcmd": "متوقع: xray-tun …", "xraytun.want_down": "متوقع: xray-tun down <iface> ip",
		"xraytun.bad_action": "xray-tun: غير معروف %q",
		"xraytun.help": `xray-tun [help] up …
  systemd-networkd۔ %s، %s۔

`,
		"xraytun.need_address_value": "قيمة بعد --address", "xraytun.need_peer_value": "قيمة بعد --peer",
		"xraytun.unknown_flag": "خيار غير معروف: %s", "xraytun.want_up_usage": "انظر help",
		"xraytun.extra_args": "وسائط زائدة: %v", "xraytun.empty_address": "عنوان فارغ",
		"xraytun.need_ipv4": "IPv4 متوقع: %q", "xraytun.empty_iface": "واجهة فارغة",
		"xraytun.bad_iface_char": "حرف غير صالح: %q", "xraytun.wrote": "كُتب %s\n",
		"xraytun.file_missing": "الملف مفقود: %s", "xraytun.not_managed": "مرفوض: %s (%s)",
		"xraytun.removed": "حُذف %s\n", "xraytun.network_reload_fail": "فشل إعادة تحميل الشبكة",
		"xraytun.ok_networkctl": "networkctl reload: ok", "xraytun.ok_systemd": "systemctl reload: ok",
		"xraytun.file_managed_by": "أنشأه z-panel xray-tun.", "xraytun.file_remove_hint": "إزالة: z-panel xray-tun down %s ip",
		"redirect.need_value_after": "قيمة بعد %q", "redirect.unknown_flag": "خيار غير معروف: %s",
		"redirect.need_one_iface": "واجهة واحدة: %v", "redirect.table_numeric": "--table رقم: %w",
		"redirect.cgroup_missing": "cgroup غير مضبوط", "redirect.bypass_mark_fmt": " (وحدة %s)",
		"redirect.mark_line": "z-panel: cgroup path=%q%s\n", "redirect.no_mark_line": "z-panel: --no-mark\n",
		"redirect.err.default_route": "مسار: %w", "redirect.err.rule_fwmark": "ip rule: %w",
		"redirect.err.rule_suppress": "suppress: %w", "redirect.err.route6": "IPv6: %w",
		"redirect.err.rule6_fw": "IPv6: %w", "redirect.err.rule6_sup": "IPv6 sup: %w",
		"redirect.iface_not_found": "واجهة %q: %s", "redirect.down_no_state": "لا state لـ %q",
		"redirect.down_bad_mode": "وضع غير مدعوم %q", "redirect.down_done": "down: جدول %s\n",
		"redirect.ip6tables_missing": "لا ip6tables", "redirect.iptables_raw_fail": "iptables: %w",
		"redirect.ip6tables_raw_fail": "ip6tables: %w", "redirect.iptables_cgroup_fail": "cgroup: %w",
		"redirect.fw_skip": "تم تخطي الجدار.", "redirect.nft_ok": "nft", "redirect.ipt_ok": "iptables",
		"redirect.ipt_cgroup_ok": "cgroup", "redirect.ip6tables_cgroup_warn": "تحذير: %v\n",
		"redirect.cg_systemctl": "ControlGroup %q: %w", "redirect.cg_empty": "فارغ %q",
		"redirect.auto_unit": "وحدة تلقائية: %s\n", "redirect.auto_fail": "فشل الكشف: %w",
		"state.state_file_err": "state: %w", "state.up_line": "up: %s (%s)\n",
		"state.summary_base": "mode=%s table=%s fwmark=%s dev %s",
		"state.summary_nomark": " no_bypass_mark=1", "state.summary_bypass": " bypass_cgroup=1",
		"bashcomp.line1": "# bash completion z-panel", "bashcomp.line2": "# install-shell",
		"bashcomp.line3": "# Bash 4+",
	}
}
