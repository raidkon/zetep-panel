// Package i18n provides localized user-visible strings.
//
// Language selection for Init() (call once from main before settings.Load):
//   - Standard locale: LANGUAGE (colon-separated), LC_ALL, LC_MESSAGES, LANG
//
// After loading config.toml, call ApplyFromConfig(settings.C.Language):
//   - auto — same rules as Init() (system locale only)
//   - en | zh | hi | es | fr | ar | bn | pt | ru | ur — fixed UI language from the config file
//
// z-panel does not read Z_PANEL_* environment variables for UI or routing; use config.toml
// (see settings.config_hdr and keys no_banner, ssh_no_tty, ssh_no_multiplex, language, daemon).
//
// Default is English (en). Unknown locales fall back to English.
// C/POSIX is treated as English.
package i18n

import (
	"fmt"
	"os"
	"strings"
)

// Lang is a supported UI language.
type Lang string

const (
	EN Lang = "en"
	RU Lang = "ru"
)

var (
	catalog map[Lang]map[string]string
	current Lang
)

func init() {
	catalog = map[Lang]map[string]string{
		EN: english(),
		RU: russian(),
	}
}

// Init selects the active language from the system locale only (call once from main before Load).
func Init() {
	current = detect()
}

// ApplyFromConfig applies language from config.toml. Call after settings.Load.
func ApplyFromConfig(lang string) {
	s := strings.ToLower(strings.TrimSpace(lang))
	if s == "" || s == "auto" {
		current = detect()
		return
	}
	if l, ok := CanonicalLanguage(s); ok {
		current = l
		return
	}
	current = EN
}

// Language returns the active language.
func Language() Lang {
	return current
}

func detect() Lang {
	if v := strings.TrimSpace(os.Getenv("LANGUAGE")); v != "" {
		for _, p := range strings.Split(v, ":") {
			p = strings.TrimSpace(p)
			if p == "" {
				continue
			}
			if l := normLang(p); l != "" {
				return l
			}
		}
	}
	for _, v := range []string{os.Getenv("LC_ALL"), os.Getenv("LC_MESSAGES"), os.Getenv("LANG")} {
		if l := normLang(v); l != "" {
			return l
		}
	}
	return EN
}

func normLang(s string) Lang {
	if strings.TrimSpace(s) == "" {
		return ""
	}
	raw := strings.TrimSpace(s)
	base := strings.ToLower(strings.Split(strings.Split(raw, "@")[0], ".")[0])
	base = strings.ReplaceAll(base, "-", "_")
	if l := langFromLocaleTag(base); l != "" {
		return l
	}
	return EN
}

// T returns the localized string for key. Optional args are passed to fmt.Sprintf.
func T(key string, args ...any) string {
	tab := catalog[current]
	s := tab[key]
	if s == "" {
		s = catalog[EN][key]
	}
	if s == "" {
		return key
	}
	if len(args) > 0 {
		return fmt.Sprintf(s, args...)
	}
	return s
}
