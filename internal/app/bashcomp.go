package app

import (
	"fmt"
	"io"
	"strings"

	"z-panel/internal/config"
	"z-panel/internal/i18n"
)

// WriteBashCompletionScript writes the full bash completion script for z-panel.
func WriteBashCompletionScript(w io.Writer) {
	topWords := strings.Join(topLevelCompletionWords(), " ")
	fmt.Fprintf(w, `# %s
# %s
# %s
# z-panel completion (bundled with z-panel %s)

`, i18n.T("bashcomp.line1"), i18n.T("bashcomp.line2"), i18n.T("bashcomp.line3"), config.Version)

	// Bash body uses % in ${var//%d/...}; keep out of fmt.Fprintf to avoid accidental fmt verbs.
	io.WriteString(w, `# Host names from SSH client config (Host …), not known_hosts.
_z_panel_ssh_config_hosts() {
	local home f tok
	home="${HOME:-}"
	[[ -n "$home" ]] || return 0
	local -a cfgfiles
	if [[ -f "$home/.ssh/config" ]]; then
		cfgfiles+=("$home/.ssh/config")
		while read -r tok; do
			[[ -z "$tok" ]] && continue
			tok="${tok//\~/$home}"
			tok="${tok//%d/$home/.ssh}"
			tok="${tok//%h/$home}"
			shopt -s nullglob
			for f in $tok; do
				[[ -f "$f" ]] && cfgfiles+=("$f")
			done
			shopt -u nullglob
		done < <(awk '/^[[:space:]]*[Ii][Nn][Cc][Ll][Uu][Dd][Ee][[:space:]]/ {
			for (i = 2; i <= NF; i++) print $i
		}' "$home/.ssh/config" 2>/dev/null)
	fi
	if [[ -d "$home/.ssh/config.d" ]]; then
		for f in "$home"/.ssh/config.d/*.conf; do
			[[ -f "$f" ]] && cfgfiles+=("$f")
		done
	fi
	((${#cfgfiles[@]})) || return 0
	awk '
	/^Host[ \t]/ {
		for (i = 2; i <= NF; i++) {
			h = $i
			if (h == "*" || h ~ /^!/) continue
			print h
		}
	}
	' "${cfgfiles[@]}" 2>/dev/null | grep -vE '^\|[0-9]+\|' | sort -u
}

# Index of z-panel in COMP_WORDS (handles: z-panel, /usr/local/bin/z-panel, sudo z-panel …).
_z_panel_bin_idx() {
	local i w
	for ((i = 0; i < ${#COMP_WORDS[@]}; i++)); do
		w="${COMP_WORDS[i]}"
		case "$w" in
			z-panel | */z-panel)
				echo "$i"
				return 0
				;;
		esac
	done
	return 1
}

_z_panel_resolve_ssh() {
	local zi=$1
	_z_panel_ssh_host=""
	_z_panel_cmd_start=$((zi + 1))
	[[ ${#COMP_WORDS[@]} -gt $((zi + 1)) ]] || return 0
	local next="${COMP_WORDS[zi + 1]:-}"
	if [[ "$next" == --ssh-connect=* ]]; then
		_z_panel_ssh_host="${next#--ssh-connect=}"
		_z_panel_cmd_start=$((zi + 2))
	elif [[ "$next" == --ssh-connect ]]; then
		if [[ -n "${COMP_WORDS[zi + 2]:-}" ]]; then
			_z_panel_ssh_host="${COMP_WORDS[zi + 2]}"
			_z_panel_cmd_start=$((zi + 3))
		else
			_z_panel_cmd_start=$((zi + 2))
		fi
	elif [[ "$next" == --ssh=* ]]; then
		_z_panel_ssh_host="${next#--ssh=}"
		_z_panel_cmd_start=$((zi + 2))
	elif [[ "$next" == --ssh ]]; then
		if [[ -n "${COMP_WORDS[zi + 2]:-}" ]]; then
			_z_panel_ssh_host="${COMP_WORDS[zi + 2]}"
			_z_panel_cmd_start=$((zi + 3))
		else
			_z_panel_cmd_start=$((zi + 2))
		fi
	fi
}

_z_panel_interfaces() {
	if [[ -n "${_z_panel_ssh_host:-}" ]]; then
		ssh -o BatchMode=yes -o ConnectTimeout=2 "$_z_panel_ssh_host" 'ls -1 /sys/class/net 2>/dev/null' 2>/dev/null
	else
		local d
		for d in /sys/class/net/*; do
			[[ -e "$d" ]] || continue
			basename "$d"
		done
	fi
}

_z_panel_completion() {
	local cur cword ecword cmd zi
	compopt +o default +o bashdefault 2>/dev/null || true
	cur="${COMP_WORDS[COMP_CWORD]}"
	cword=$COMP_CWORD
	zi=$(_z_panel_bin_idx) || return 0
	_z_panel_resolve_ssh "$zi"
	ecword=$(( cword - _z_panel_cmd_start + 1 ))

	# Complete SSH host (not known_hosts): cursor on host word after --ssh or --ssh-connect.
	if [[ "${COMP_WORDS[zi + 1]:-}" == --ssh ]] && [[ $cword -eq $((zi + 2)) ]]; then
		mapfile -t COMPREPLY < <(compgen -W "$(_z_panel_ssh_config_hosts)" -- "$cur")
		((${#COMPREPLY[@]})) || compopt +o default +o bashdefault 2>/dev/null || true
		return
	fi
	if [[ "${COMP_WORDS[zi + 1]:-}" == --ssh-connect ]] && [[ $cword -eq $((zi + 2)) ]]; then
		mapfile -t COMPREPLY < <(compgen -W "$(_z_panel_ssh_config_hosts)" -- "$cur")
		((${#COMPREPLY[@]})) || compopt +o default +o bashdefault 2>/dev/null || true
		return
	fi

`)

	fmt.Fprintf(w, `	if [[ $cword -eq $((zi + 1)) ]]; then
		if [[ "$cur" == --ssh=* ]]; then
			mapfile -t COMPREPLY < <(compgen -W "$(_z_panel_ssh_config_hosts)" -P '--ssh=' -- "${cur#--ssh=}")
			((${#COMPREPLY[@]})) || compopt +o default +o bashdefault 2>/dev/null || true
			return
		fi
		if [[ "$cur" == --ssh-connect=* ]]; then
			mapfile -t COMPREPLY < <(compgen -W "$(_z_panel_ssh_config_hosts)" -P '--ssh-connect=' -- "${cur#--ssh-connect=}")
			((${#COMPREPLY[@]})) || compopt +o default +o bashdefault 2>/dev/null || true
			return
		fi
		mapfile -t COMPREPLY < <(compgen -W '%s' -- "$cur")
		return
	fi

	if [[ -n "$_z_panel_ssh_host" ]] && [[ $cword -eq $_z_panel_cmd_start ]]; then
		mapfile -t COMPREPLY < <(compgen -W '%s' -- "$cur")
		return
	fi

	cmd="${COMP_WORDS[_z_panel_cmd_start]}"

	case "$cmd" in
`, topWords, topWords)

	for _, c := range All() {
		c.BashCompletionCase(w)
	}

	fmt.Fprintf(w, `	*)
		return
		;;
	esac
}

complete -F _z_panel_completion z-panel
`)
}

func topLevelCompletionWords() []string {
	seen := make(map[string]struct{})
	var out []string
	add := func(s string) {
		s = strings.TrimSpace(s)
		if s == "" {
			return
		}
		if _, ok := seen[s]; ok {
			return
		}
		seen[s] = struct{}{}
		out = append(out, s)
	}
	for _, c := range All() {
		add(c.Name())
		if a, ok := c.(interface{ Aliases() []string }); ok {
			for _, al := range a.Aliases() {
				add(al)
			}
		}
	}
	add("help")
	add("-h")
	add("--help")
	add("--ssh")
	add("--ssh-connect")
	return out
}
