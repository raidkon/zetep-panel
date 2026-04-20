package app

import (
	"fmt"
	"io"
	"strings"

	"z-panel/internal/i18n"
)

// WriteBashCompletionScript writes the full bash completion script for z-panel.
func WriteBashCompletionScript(w io.Writer) {
	fmt.Fprintf(w, `# %s
# %s
# %s

_z_panel_interfaces() {
	local d
	for d in /sys/class/net/*; do
		[[ -e "$d" ]] || continue
		basename "$d"
	done
}

_z_panel_completion() {
	local cur cword
	cur="${COMP_WORDS[COMP_CWORD]}"
	cword=$COMP_CWORD

	if [[ $cword -eq 1 ]]; then
		mapfile -t COMPREPLY < <(compgen -W '%s' -- "$cur")
		return
	fi

	local cmd="${COMP_WORDS[1]}"

	case "$cmd" in
`, i18n.T("bashcomp.line1"), i18n.T("bashcomp.line2"), i18n.T("bashcomp.line3"), strings.Join(topLevelCompletionWords(), " "))

	for _, c := range All() {
		c.BashCompletionCase(w)
	}

	fmt.Fprintf(w, `	esac
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
	return out
}
