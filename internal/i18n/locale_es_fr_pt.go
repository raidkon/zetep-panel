package i18n

func esStrings() map[string]string {
	return map[string]string{
		"root.unknown_command": "comando desconocido: %s\n\n",
		"root.help.tagline":    "z-panel — enrutamiento por políticas vía TUN (estilo wg-quick para Xray).",
		"root.help.top": `Nivel superior:
  z-panel help | -h | --help     este resumen (todos los comandos)
  z-panel version | -v | --version
  z-panel <command> [help | -h | --help]   ayuda de un comando
  z-panel <command> …            el resto va al comando

Comandos:
`,
		"root.help.cmdline": "  z-panel %s …\n", "root.help.ufw_note": "\nEtiqueta UFW: '%s: …'\n\n",
		"root.help.section_rule": "─── %s ───\n", "root.need_root": "se requieren privilegios root (sudo)",
		"install.help": `install [help] [<sshHost>]
  Local: copia el binario a %s (chmod 755); requiere root.
  Si falta %s — asistente interactivo y guardado.
  Remoto: scp y una sesión SSH (-t): install; si no hay config — config init (interactivo).

`,
		"install.err.interrupted": "interrumpido (Ctrl+C)", "install.err.interrupted_with": "interrumpido (Ctrl+C): %w",
		"install.err.open_self": "abrir propio binario: %w", "install.err.create_tmp": "crear %s: %w",
		"install.err.copy": "copia: %w", "install.err.rename": "renombrar a %s: %w", "install.err.config": "config: %w",
		"install.installed": "instalado: %s\n", "install.new_version": "Nueva versión: %s\n", "install.old_version": "Versión anterior: %s\n",
		"install.err.scp": "scp: %w", "install.err.ssh": "ssh install/init: %w", "install.remote_done": "instalado en %s: %s\n",
		"installshell.err.home": "directorio home: %w", "installshell.err.mkdir": "mkdir: %w", "installshell.err.write": "escribir %s: %w",
		"installshell.done": "bash completion instalado: %s\n",
		"installshell.hint_shell": "Abra un terminal nuevo o ejecute: source /usr/share/bash-completion/bash_completion",
		"installshell.hint_user": "(usuario: asegúrese de que ~/.bashrc carga bash-completion)",
		"installshell.help": `install-shell [help] [--user|-u]
  Escribe el script de bash completion (generado desde los comandos).
  Por defecto: %s (requiere root).
  --user: ~/.local/share/bash-completion/completions/z-panel (o $XDG_DATA_HOME/...).
  Requiere bash 4+; para soporte completo instale bash-completion.

`,
		"confcmd.err_unknown": "config: subcomando desconocido %q (se esperaba init o migrate)",
		"confcmd.help": `config [help] init [--force|-f]
  Crear o sobrescribir %s de forma interactiva (--force).
config [help] migrate
  Aplicar claves nuevas tras actualizar z-panel (solo versiones de esquema nuevas).

`,
		"version.help": `version
  Versión: %s (también como primera línea en stderr).
  Sinónimos: -v, --version

`,
		"settings.err.read": "leer %s: %w", "settings.err.parse": "analizar %s: %w", "settings.err.mkdir": "mkdir %s: %w",
		"settings.err.write": "escribir %s: %w", "settings.config_hdr": "# z-panel — configuración\n\n",
		"settings.init_exists": "la config ya existe: %s (sobrescribir: z-panel config init --force)\n",
		"settings.init_intro": "Configuración z-panel — escriba un valor o Enter para el predeterminado.",
		"settings.saved": "\nguardado: %s\n",
		"settings.prompt.table": "Tabla de enrutamiento / fwmark",
		"settings.prompt.systemd_network": "Directorio systemd-networkd", "settings.prompt.lan_cidr": "Plantilla UFW: LAN CIDR",
		"settings.prompt.lan_dev": "Plantilla UFW: interfaz LAN", "settings.prompt.xray_addr": "Xray TUN: dirección por defecto",
		"settings.prompt.xray_peer": "Xray TUN: peer por defecto", "settings.prompt.ufw_marker": "Etiqueta en comentarios UFW",
		"settings.prompt.xray_mark": "Marcador en .network",
		"settings.prompt.language": "Idioma de la UI (" + LanguageListHint + ")",
		"settings.migrate_intro": "Esta config es de una versión antigua. Establezca los valores nuevos.",
		"settings.migrate_uptodate": "el esquema de config ya está actualizado.",
		"settings.migrate_no_file": "%s: no hay archivo de config (ejecute z-panel config init)",
		"xrayredirect.want_up_down": "se esperaba: z-panel xray-redirect up|down …",
		"xrayredirect.want_down_iface": "se esperaba: z-panel xray-redirect down <interfaz>",
		"xrayredirect.bad_action": "xray-redirect: acción desconocida %q",
		"xrayredirect.help": `xray-redirect [help] up [flags] <interfaz>
xray-redirect [help] down <interfaz>
  Túnel completo estilo wg-quick; marca cgroup v2.
  Flags up: --bypass-unit, --bypass-cgroup, --table (def. %s), --no-mark, --ipv6
  Ejemplos: sudo z-panel xray-redirect up xray2tun

`,
		"ufw.want_check": "se esperaba: z-panel ufw check …", "ufw.want_subcmd": "se esperaba: ufw check … o ufw masq-check <iface>",
		"ufw.help": `ufw [help] check [--lan-cidr=CIDR] [--lan-dev=DEV] [interfaz]
  Reglas UFW con etiqueta %s. Plantillas: --lan-cidr (def. %s), --lan-dev (%s).

ufw [help] masq-check [--lan-cidr=CIDR] <interfaz>
  Busca MASQUERADE/SNAT POSTROUTING con -o <interfaz> (iptables-save). Si falta, sugiere línea en before.rules. CIDR LAN por defecto %s.

`,
		"ufw.masq.verdict_ok": "ESTADO: OK — hay MASQUERADE/SNAT en POSTROUTING para -o %s (%d regla(s)).", "ufw.masq.detail_heading": "Líneas (iptables-save -t nat):", "ufw.masq.verdict_missing": "ESTADO: FALTA — no hay POSTROUTING MASQUERADE/SNAT con -o %s en nat.\n",
		"ufw.masq.hint_add": "Línea sugerida (junto a otros MASQUERADE), luego: sudo ufw reload\n\n-A POSTROUTING -s %s -o %s -j MASQUERADE\n",
		"ufw.masq.iptables_cmd": "iptables-save -t nat", "ufw.masq.want_iface": "masq-check: una interfaz (ej. z-panel ufw masq-check xray2tun)",
		"ufw.err.lan_cidr_empty": "--lan-cidr: vacío", "ufw.err.lan_cidr_need": "valor tras --lan-cidr",
		"ufw.err.lan_dev_empty": "--lan-dev: vacío", "ufw.err.lan_dev_need": "valor tras --lan-dev",
		"ufw.err.unknown_flag": "flag desconocido: %s", "ufw.err.too_many_iface": "máx. una interfaz, extra: %q",
		"ufw.ufw_status_failed": "ufw status verbose: %w\n%s", "ufw.section_rules": "=== Reglas UFW con %s ===",
		"ufw.no_lines": "(sin líneas «z-panel»)", "ufw.section_iface_refs": "=== Líneas ufw status con interfaz %s (cualquier comentario) ===",
		"ufw.no_iface_refs": "(ninguna línea «on %s» en ufw status — añada route/forward si hace falta)", "ufw.section_hints": "=== Sugerencias (plantillas) ===",
		"ufw.hint_sysctl": `# Enrutamiento:
# /etc/ufw/sysctl.conf: net.ipv4.ip_forward=1
sudo ufw route allow in on %s out on %s from %s comment '%s: lan to tunnel'
`,
		"ufw.hint_return": "# Si ufw bloquea respuestas:\n",
		"ufw.hint_return_cmd": "sudo ufw route allow in on %s out on %s comment '%s: return path'\n",
		"ufw.no_iface_hint1": "Indique interfaz del túnel: z-panel ufw check … <interfaz>\n",
		"ufw.no_iface_hint2": "Plantillas actuales: LAN %s dev %s\n", "ufw.section_full": "=== ufw status verbose completo ===",
		"xraytun.want_subcmd": "se esperaba: z-panel xray-tun up/down …", "xraytun.want_down": "se esperaba: z-panel xray-tun down <iface> ip",
		"xraytun.bad_action": "xray-tun: acción desconocida %q",
		"xraytun.help": `xray-tun [help] up [flags] <iface> ip | … <addr> [<peer>]
  systemd-networkd drop-in. Modo ip: defectos %s y %s.

`,
		"xraytun.need_address_value": "valor tras --address", "xraytun.need_peer_value": "valor tras --peer",
		"xraytun.unknown_flag": "flag desconocido: %s",
		"xraytun.want_up_usage": "se esperaba: z-panel xray-tun up [flags] <iface> ip\n  o con dirección explícita",
		"xraytun.extra_args": "argumentos extra: %v", "xraytun.empty_address": "dirección vacía",
		"xraytun.need_ipv4": "se espera IPv4: %q", "xraytun.empty_iface": "nombre de interfaz vacío",
		"xraytun.bad_iface_char": "carácter inválido en interfaz: %q", "xraytun.wrote": "escrito %s\n",
		"xraytun.file_missing": "falta archivo: %s", "xraytun.not_managed": "rechazado: %s no gestionado por z-panel (%s)",
		"xraytun.removed": "eliminado %s\n", "xraytun.network_reload_fail": "fallo al recargar red",
		"xraytun.ok_networkctl": "networkctl reload: ok", "xraytun.ok_systemd": "systemctl reload systemd-networkd: ok",
		"xraytun.file_managed_by": "Creado por z-panel xray-tun.", "xraytun.file_remove_hint": "Quitar: z-panel xray-tun down %s ip",
		"redirect.need_value_after": "valor tras %q", "redirect.unknown_flag": "flag desconocido: %s",
		"redirect.need_one_iface": "se necesita una interfaz, recibido: %v",
		"redirect.table_numeric": "--table debe ser numérico: %w", "redirect.cgroup_missing": "cgroup no definido; use flags o --no-mark",
		"redirect.bypass_mark_fmt": " (unidad %s)", "redirect.mark_line": "z-panel: marca cgroup v2: path=%q%s\n",
		"redirect.no_mark_line": "z-panel: --no-mark\n", "redirect.err.default_route": "ruta default: %w",
		"redirect.err.rule_fwmark": "ip rule: %w", "redirect.err.rule_suppress": "suppress_prefixlength: %w",
		"redirect.err.route6": "ruta IPv6: %w", "redirect.err.rule6_fw": "regla IPv6: %w", "redirect.err.rule6_sup": "regla IPv6 sup: %w",
		"redirect.iface_not_found": "interfaz %q no encontrada: %s",
		"redirect.down_no_state": "sin state para %q — ejecute up antes (interfaz %q)",
		"redirect.down_bad_mode": "modo state no soportado %q para %q",
		"redirect.down_done": "down: limpiado table %s para %s\n", "redirect.ip6tables_missing": "ip6tables no encontrado",
		"redirect.iptables_raw_fail": "iptables raw: %w", "redirect.ip6tables_raw_fail": "ip6tables raw: %w",
		"redirect.iptables_cgroup_fail": "iptables cgroup: %w", "redirect.fw_skip": "firewall wg omitido.",
		"redirect.nft_ok": "nft anti-leak aplicado", "redirect.ipt_ok": "iptables anti-leak aplicado",
		"redirect.ipt_cgroup_ok": "iptables cgroup aplicado", "redirect.ip6tables_cgroup_warn": "aviso ip6tables cgroup: %v\n",
		"redirect.cg_systemctl": "ControlGroup %q: %w", "redirect.cg_empty": "ControlGroup vacío %q",
		"redirect.auto_unit": "unidad systemd auto: %s\n", "redirect.auto_fail": "no se detectó unidad: %w",
		"state.state_file_err": "archivo state: %w", "state.up_line": "up: %s (config %s)\n",
		"state.summary_base": "mode=%s table=%s fwmark=%s default dev %s",
		"state.summary_nomark": " no_bypass_mark=1", "state.summary_bypass": " bypass_cgroup=1",
		"bashcomp.line1": "# bash completion z-panel", "bashcomp.line2": "# Instalar: z-panel install-shell",
		"bashcomp.line3": "# Requiere bash 4+",
	}
}

func frStrings() map[string]string {
	return map[string]string{
		"root.unknown_command": "commande inconnue : %s\n\n",
		"root.help.tagline":    "z-panel — routage par politique via TUN (style wg-quick pour Xray).",
		"root.help.top": `Niveau racine :
  z-panel help | -h | --help     ce résumé
  z-panel version | -v | --version
  z-panel <commande> [help | -h | --help]   aide d'une commande
  z-panel <commande> …            arguments passés à la commande

Commandes :
`,
		"root.help.cmdline": "  z-panel %s …\n", "root.help.ufw_note": "\nÉtiquette UFW : '%s: …'\n\n",
		"root.help.section_rule": "─── %s ───\n", "root.need_root": "droits root requis (sudo)",
		"install.help": `install [help] [<sshHost>]
  Local : copie le binaire vers %s (chmod 755) ; root requis.
  Si %s manque — assistant interactif.
  Distant : scp puis SSH (-t) : install ; sinon config init.

`,
		"install.err.interrupted": "interrompu (Ctrl+C)", "install.err.interrupted_with": "interrompu (Ctrl+C) : %w",
		"install.err.open_self": "ouverture du binaire : %w", "install.err.create_tmp": "créer %s : %w",
		"install.err.copy": "copie : %w", "install.err.rename": "renommer vers %s : %w", "install.err.config": "config : %w",
		"install.installed": "installé : %s\n", "install.new_version": "Nouvelle version : %s\n", "install.old_version": "Version précédente : %s\n",
		"install.err.scp": "scp : %w", "install.err.ssh": "ssh : %w", "install.remote_done": "installé sur %s : %s\n",
		"installshell.err.home": "répertoire personnel : %w", "installshell.err.mkdir": "mkdir : %w", "installshell.err.write": "écriture %s : %w",
		"installshell.done": "bash completion installé : %s\n",
		"installshell.hint_shell": "Nouveau terminal ou : source /usr/share/bash-completion/bash_completion",
		"installshell.hint_user": "(utilisateur : vérifiez ~/.bashrc)",
		"installshell.help": `install-shell [help] [--user|-u]
  Écrit le script bash completion. Défaut : %s (root).
  --user : ~/.local/share/bash-completion/completions/z-panel
  Bash 4+ ; paquet bash-completion recommandé.

`,
		"confcmd.err_unknown": "config : sous-commande inconnue %q (init ou migrate)",
		"confcmd.help": `config [help] init [--force|-f]
  Création interactive ou écrasement de %s (--force).
config [help] migrate
  Nouvelles clés après mise à jour de z-panel.

`,
		"version.help": `version
  Version : %s (aussi 1re ligne stderr).
  Synonymes : -v, --version

`,
		"settings.err.read": "lecture %s : %w", "settings.err.parse": "analyse %s : %w", "settings.err.mkdir": "mkdir %s : %w",
		"settings.err.write": "écriture %s : %w", "settings.config_hdr": "# z-panel — configuration\n\n",
		"settings.init_exists": "config existe : %s (écraser : z-panel config init --force)\n",
		"settings.init_intro": "Configuration z-panel — valeur ou Entrée pour défaut.",
		"settings.saved": "\nenregistré : %s\n",
		"settings.prompt.table": "Table / fwmark",
		"settings.prompt.systemd_network": "Répertoire systemd-networkd", "settings.prompt.lan_cidr": "UFW modèle LAN CIDR",
		"settings.prompt.lan_dev": "UFW interface LAN", "settings.prompt.xray_addr": "Xray TUN adresse défaut",
		"settings.prompt.xray_peer": "Xray TUN peer défaut", "settings.prompt.ufw_marker": "Étiquette UFW",
		"settings.prompt.xray_mark": "Marqueur .network",
		"settings.prompt.language": "Langue UI (" + LanguageListHint + ")",
		"settings.migrate_intro": "Config d'une ancienne version. Renseignez les nouvelles options.",
		"settings.migrate_uptodate": "schéma config à jour.",
		"settings.migrate_no_file": "%s : fichier config absent (z-panel config init)",
		"xrayredirect.want_up_down": "attendu : z-panel xray-redirect up|down …",
		"xrayredirect.want_down_iface": "attendu : z-panel xray-redirect down <interface>",
		"xrayredirect.bad_action": "xray-redirect : action inconnue %q",
		"xrayredirect.help": `xray-redirect [help] up [flags] <interface>
xray-redirect [help] down <interface>
  Tunnel complet style wg-quick ; cgroup v2.
  Flags : --bypass-unit, --table (déf. %s), --no-mark, --ipv6

`,
		"ufw.want_check": "attendu : z-panel ufw check …", "ufw.want_subcmd": "attendu : ufw check … ou ufw masq-check <iface>",
		"ufw.help": `ufw [help] check [--lan-cidr=] [--lan-dev=] [interface]
  Règles UFW %s. Modèles : --lan-cidr (déf. %s), --lan-dev (%s).

ufw [help] masq-check [--lan-cidr=] <interface>
  Cherche MASQUERADE/SNAT POSTROUTING -o <interface>. Sinon ligne pour before.rules. CIDR LAN déf. %s.

`,
		"ufw.masq.verdict_ok": "ÉTAT : OK — MASQUERADE/SNAT POSTROUTING pour -o %s présent (%d règle(s)).", "ufw.masq.detail_heading": "Lignes (iptables-save -t nat) :", "ufw.masq.verdict_missing": "ÉTAT : MANQUANT — aucune règle POSTROUTING MASQUERADE/SNAT avec -o %s.\n",
		"ufw.masq.hint_add": "Ligne (avec les autres MASQUERADE), puis : sudo ufw reload\n\n-A POSTROUTING -s %s -o %s -j MASQUERADE\n",
		"ufw.masq.iptables_cmd": "iptables-save -t nat", "ufw.masq.want_iface": "masq-check : une interface (ex. z-panel ufw masq-check xray2tun)",
		"ufw.err.lan_cidr_empty": "--lan-cidr vide", "ufw.err.lan_cidr_need": "valeur après --lan-cidr",
		"ufw.err.lan_dev_empty": "--lan-dev vide", "ufw.err.lan_dev_need": "valeur après --lan-dev",
		"ufw.err.unknown_flag": "option inconnue : %s", "ufw.err.too_many_iface": "une seule interface, extra : %q",
		"ufw.ufw_status_failed": "ufw status verbose : %w\n%s", "ufw.section_rules": "=== Règles UFW %s ===",
		"ufw.no_lines": "(aucune ligne z-panel)", "ufw.section_iface_refs": "=== Lignes ufw status mentionnant l’interface %s ===",
		"ufw.no_iface_refs": "(aucune ligne «on %s» — ajoutez route/forward si besoin)", "ufw.section_hints": "=== Suggestions ===",
		"ufw.hint_sysctl": `# Routage :
sudo ufw route allow in on %s out on %s from %s comment '%s: lan to tunnel'
`,
		"ufw.hint_return": "# Si ufw bloque :\n", "ufw.hint_return_cmd": "sudo ufw route allow in on %s out on %s comment '%s: return path'\n",
		"ufw.no_iface_hint1": "Indiquez l'interface tunnel.\n", "ufw.no_iface_hint2": "Modèles : LAN %s dev %s\n",
		"ufw.section_full": "=== ufw status verbose ===",
		"xraytun.want_subcmd": "attendu : xray-tun up/down …", "xraytun.want_down": "attendu : xray-tun down <iface> ip",
		"xraytun.bad_action": "xray-tun : action inconnue %q",
		"xraytun.help": `xray-tun [help] up [flags] <iface> ip …
  systemd-networkd. Mode ip : défauts %s et %s.

`,
		"xraytun.need_address_value": "valeur après --address", "xraytun.need_peer_value": "valeur après --peer",
		"xraytun.unknown_flag": "option inconnue : %s", "xraytun.want_up_usage": "syntaxe up : voir help",
		"xraytun.extra_args": "arguments en trop : %v", "xraytun.empty_address": "adresse vide",
		"xraytun.need_ipv4": "IPv4 attendu : %q", "xraytun.empty_iface": "interface vide",
		"xraytun.bad_iface_char": "caractère invalide : %q", "xraytun.wrote": "écrit %s\n",
		"xraytun.file_missing": "fichier absent : %s", "xraytun.not_managed": "refus : %s non géré (%s)",
		"xraytun.removed": "supprimé %s\n", "xraytun.network_reload_fail": "échec rechargement réseau",
		"xraytun.ok_networkctl": "networkctl reload : ok", "xraytun.ok_systemd": "systemctl reload systemd-networkd : ok",
		"xraytun.file_managed_by": "Créé par z-panel xray-tun.", "xraytun.file_remove_hint": "Retirer : z-panel xray-tun down %s ip",
		"redirect.need_value_after": "valeur après %q", "redirect.unknown_flag": "option inconnue : %s",
		"redirect.need_one_iface": "une interface requise : %v", "redirect.table_numeric": "--table numérique : %w",
		"redirect.cgroup_missing": "cgroup non défini", "redirect.bypass_mark_fmt": " (unité %s)",
		"redirect.mark_line": "z-panel : cgroup path=%q%s\n", "redirect.no_mark_line": "z-panel : --no-mark\n",
		"redirect.err.default_route": "route : %w", "redirect.err.rule_fwmark": "ip rule : %w",
		"redirect.err.rule_suppress": "suppress : %w", "redirect.err.route6": "route IPv6 : %w",
		"redirect.err.rule6_fw": "règle IPv6 : %w", "redirect.err.rule6_sup": "règle sup IPv6 : %w",
		"redirect.iface_not_found": "interface %q : %s", "redirect.down_no_state": "pas de state %q",
		"redirect.down_bad_mode": "mode state %q inconnu", "redirect.down_done": "down : table %s nettoyée\n",
		"redirect.ip6tables_missing": "ip6tables absent", "redirect.iptables_raw_fail": "iptables raw : %w",
		"redirect.ip6tables_raw_fail": "ip6tables raw : %w", "redirect.iptables_cgroup_fail": "iptables cgroup : %w",
		"redirect.fw_skip": "pare-feu wg ignoré.", "redirect.nft_ok": "nft appliqué", "redirect.ipt_ok": "iptables appliqué",
		"redirect.ipt_cgroup_ok": "cgroup appliqué", "redirect.ip6tables_cgroup_warn": "avertissement ip6tables : %v\n",
		"redirect.cg_systemctl": "ControlGroup %q : %w", "redirect.cg_empty": "ControlGroup vide %q",
		"redirect.auto_unit": "unité auto : %s\n", "redirect.auto_fail": "détection unité : %w",
		"state.state_file_err": "fichier state : %w", "state.up_line": "up : %s (config %s)\n",
		"state.summary_base": "mode=%s table=%s fwmark=%s dev %s",
		"state.summary_nomark": " no_bypass_mark=1", "state.summary_bypass": " bypass_cgroup=1",
		"bashcomp.line1": "# bash completion z-panel", "bashcomp.line2": "# z-panel install-shell",
		"bashcomp.line3": "# Bash 4+",
	}
}

func ptStrings() map[string]string {
	return map[string]string{
		"root.unknown_command": "comando desconhecido: %s\n\n",
		"root.help.tagline":    "z-panel — roteamento por política via TUN (estilo wg-quick para Xray).",
		"root.help.top": `Nível superior:
  z-panel help | -h | --help     este resumo
  z-panel version | -v | --version
  z-panel <comando> [help | -h | --help]   ajuda de um comando
  z-panel <comando> …            argumentos ao comando

Comandos:
`,
		"root.help.cmdline": "  z-panel %s …\n", "root.help.ufw_note": "\nEtiqueta UFW: '%s: …'\n\n",
		"root.help.section_rule": "─── %s ───\n", "root.need_root": "privilégios root necessários (sudo)",
		"install.help": `install [help] [<sshHost>]
  Local: copia o binário para %s (chmod 755); precisa root.
  Se faltar %s — assistente interativo.
  Remoto: scp e SSH (-t): install; senão config init.

`,
		"install.err.interrupted": "interrompido (Ctrl+C)", "install.err.interrupted_with": "interrompido (Ctrl+C): %w",
		"install.err.open_self": "abrir binário: %w", "install.err.create_tmp": "criar %s: %w",
		"install.err.copy": "cópia: %w", "install.err.rename": "renomear para %s: %w", "install.err.config": "config: %w",
		"install.installed": "instalado: %s\n", "install.new_version": "Nova versão: %s\n", "install.old_version": "Versão anterior: %s\n",
		"install.err.scp": "scp: %w", "install.err.ssh": "ssh: %w", "install.remote_done": "instalado em %s: %s\n",
		"installshell.err.home": "diretório home: %w", "installshell.err.mkdir": "mkdir: %w", "installshell.err.write": "escrever %s: %w",
		"installshell.done": "bash completion instalado: %s\n",
		"installshell.hint_shell": "Novo terminal ou: source /usr/share/bash-completion/bash_completion",
		"installshell.hint_user": "(usuário: verifique ~/.bashrc)",
		"installshell.help": `install-shell [help] [--user|-u]
  Escreve bash completion. Padrão: %s (root).
  Bash 4+; instale bash-completion.

`,
		"confcmd.err_unknown": "config: subcomando desconhecido %q (init ou migrate)",
		"confcmd.help": `config [help] init [--force|-f]
  Criar ou sobrescrever %s (--force).
config [help] migrate
  Novas chaves após atualizar z-panel.

`,
		"version.help": `version
  Versão: %s (também na 1ª linha do stderr).
  Sinônimos: -v, --version

`,
		"settings.err.read": "ler %s: %w", "settings.err.parse": "analisar %s: %w", "settings.err.mkdir": "mkdir %s: %w",
		"settings.err.write": "escrever %s: %w", "settings.config_hdr": "# z-panel — configuração\n\n",
		"settings.init_exists": "config já existe: %s (z-panel config init --force)\n",
		"settings.init_intro": "Configuração z-panel — valor ou Enter para padrão.",
		"settings.saved": "\nsalvo: %s\n",
		"settings.prompt.table": "Tabela / fwmark",
		"settings.prompt.systemd_network": "Diretório systemd-networkd", "settings.prompt.lan_cidr": "Modelo UFW LAN CIDR",
		"settings.prompt.lan_dev": "Interface LAN UFW", "settings.prompt.xray_addr": "Xray TUN endereço padrão",
		"settings.prompt.xray_peer": "Xray TUN peer padrão", "settings.prompt.ufw_marker": "Etiqueta UFW",
		"settings.prompt.xray_mark": "Marcador .network",
		"settings.prompt.language": "Idioma da UI (" + LanguageListHint + ")",
		"settings.migrate_intro": "Config de versão antiga. Defina novas opções.",
		"settings.migrate_uptodate": "esquema de config atualizado.",
		"settings.migrate_no_file": "%s: arquivo de config ausente (z-panel config init)",
		"xrayredirect.want_up_down": "esperado: z-panel xray-redirect up|down …",
		"xrayredirect.want_down_iface": "esperado: z-panel xray-redirect down <interface>",
		"xrayredirect.bad_action": "xray-redirect: ação desconhecida %q",
		"xrayredirect.help": `xray-redirect [help] up [flags] <interface>
xray-redirect [help] down <interface>
  Túnel estilo wg-quick ; cgroup v2. --table padrão %s

`,
		"ufw.want_check": "esperado: z-panel ufw check …", "ufw.want_subcmd": "esperado: ufw check … ou ufw masq-check <iface>",
		"ufw.help": `ufw [help] check …
  Regras UFW %s. --lan-cidr (padrão %s), --lan-dev (%s).

ufw [help] masq-check [--lan-cidr=CIDR] <interface>
  Procura MASQUERADE/SNAT -o <interface> em nat. Se faltar, linha para before.rules. CIDR LAN padrão %s.

`,
		"ufw.masq.verdict_ok": "ESTADO: OK — MASQUERADE/SNAT POSTROUTING para -o %s presente (%d regra(s)).", "ufw.masq.detail_heading": "Linhas (iptables-save -t nat):", "ufw.masq.verdict_missing": "ESTADO: FALTANDO — sem regra POSTROUTING MASQUERADE/SNAT com -o %s.\n",
		"ufw.masq.hint_add": "Linha sugerida; depois: sudo ufw reload\n\n-A POSTROUTING -s %s -o %s -j MASQUERADE\n",
		"ufw.masq.iptables_cmd": "iptables-save -t nat", "ufw.masq.want_iface": "masq-check: uma interface (ex. z-panel ufw masq-check xray2tun)",
		"ufw.err.lan_cidr_empty": "--lan-cidr vazio", "ufw.err.lan_cidr_need": "valor após --lan-cidr",
		"ufw.err.lan_dev_empty": "--lan-dev vazio", "ufw.err.lan_dev_need": "valor após --lan-dev",
		"ufw.err.unknown_flag": "flag desconhecido: %s", "ufw.err.too_many_iface": "no máx. uma interface, extra: %q",
		"ufw.ufw_status_failed": "ufw status verbose: %w\n%s", "ufw.section_rules": "=== Regras UFW %s ===",
		"ufw.no_lines": "(sem linhas z-panel)", "ufw.section_iface_refs": "=== Linhas ufw status com interface %s ===",
		"ufw.no_iface_refs": "(nenhuma linha «on %s» — adicione route/forward se precisar)", "ufw.section_hints": "=== Sugestões ===",
		"ufw.hint_sysctl": `sudo ufw route allow in on %s out on %s from %s comment '%s: lan to tunnel'
`,
		"ufw.hint_return": "# Se ufw bloquear:\n", "ufw.hint_return_cmd": "sudo ufw route allow in on %s out on %s comment '%s: return path'\n",
		"ufw.no_iface_hint1": "Especifique a interface do túnel.\n", "ufw.no_iface_hint2": "Modelos: LAN %s dev %s\n",
		"ufw.section_full": "=== ufw status verbose completo ===",
		"xraytun.want_subcmd": "esperado: xray-tun up/down …", "xraytun.want_down": "esperado: xray-tun down <iface> ip",
		"xraytun.bad_action": "xray-tun: ação desconhecida %q",
		"xraytun.help": `xray-tun [help] up [flags] <iface> ip …
  systemd-networkd. Padrões %s e %s.

`,
		"xraytun.need_address_value": "valor após --address", "xraytun.need_peer_value": "valor após --peer",
		"xraytun.unknown_flag": "flag desconhecido: %s", "xraytun.want_up_usage": "veja help para sintaxe up",
		"xraytun.extra_args": "argumentos extras: %v", "xraytun.empty_address": "endereço vazio",
		"xraytun.need_ipv4": "IPv4 esperado: %q", "xraytun.empty_iface": "interface vazia",
		"xraytun.bad_iface_char": "caractere inválido: %q", "xraytun.wrote": "gravado %s\n",
		"xraytun.file_missing": "arquivo ausente: %s", "xraytun.not_managed": "recusado: %s não gerenciado (%s)",
		"xraytun.removed": "removido %s\n", "xraytun.network_reload_fail": "falha ao recarregar rede",
		"xraytun.ok_networkctl": "networkctl reload: ok", "xraytun.ok_systemd": "systemctl reload systemd-networkd: ok",
		"xraytun.file_managed_by": "Criado por z-panel xray-tun.", "xraytun.file_remove_hint": "Remover: z-panel xray-tun down %s ip",
		"redirect.need_value_after": "valor após %q", "redirect.unknown_flag": "flag desconhecido: %s",
		"redirect.need_one_iface": "uma interface necessária: %v", "redirect.table_numeric": "--table numérico: %w",
		"redirect.cgroup_missing": "cgroup não definido", "redirect.bypass_mark_fmt": " (unidade %s)",
		"redirect.mark_line": "z-panel: cgroup path=%q%s\n", "redirect.no_mark_line": "z-panel: --no-mark\n",
		"redirect.err.default_route": "rota: %w", "redirect.err.rule_fwmark": "ip rule: %w",
		"redirect.err.rule_suppress": "suppress: %w", "redirect.err.route6": "rota IPv6: %w",
		"redirect.err.rule6_fw": "regra IPv6: %w", "redirect.err.rule6_sup": "regra sup IPv6: %w",
		"redirect.iface_not_found": "interface %q: %s", "redirect.down_no_state": "sem state %q",
		"redirect.down_bad_mode": "modo state %q inválido", "redirect.down_done": "down: tabela %s limpa\n",
		"redirect.ip6tables_missing": "ip6tables ausente", "redirect.iptables_raw_fail": "iptables raw: %w",
		"redirect.ip6tables_raw_fail": "ip6tables raw: %w", "redirect.iptables_cgroup_fail": "iptables cgroup: %w",
		"redirect.fw_skip": "firewall wg ignorado.", "redirect.nft_ok": "nft aplicado", "redirect.ipt_ok": "iptables aplicado",
		"redirect.ipt_cgroup_ok": "cgroup aplicado", "redirect.ip6tables_cgroup_warn": "aviso ip6tables: %v\n",
		"redirect.cg_systemctl": "ControlGroup %q: %w", "redirect.cg_empty": "ControlGroup vazio %q",
		"redirect.auto_unit": "unidade auto: %s\n", "redirect.auto_fail": "falha ao detectar unidade: %w",
		"state.state_file_err": "arquivo state: %w", "state.up_line": "up: %s (config %s)\n",
		"state.summary_base": "mode=%s table=%s fwmark=%s dev %s",
		"state.summary_nomark": " no_bypass_mark=1", "state.summary_bypass": " bypass_cgroup=1",
		"bashcomp.line1": "# bash completion z-panel", "bashcomp.line2": "# z-panel install-shell",
		"bashcomp.line3": "# Bash 4+",
	}
}
