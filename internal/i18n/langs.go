package i18n

import "strings"

// Top 10 languages by number of speakers (broadly): en, zh, hi, es, fr, ar, bn, pt, ru, ur.
const (
	ZH Lang = "zh" // Chinese (Simplified UI copy; locale zh_*)
	HI Lang = "hi" // Hindi
	ES Lang = "es" // Spanish
	FR Lang = "fr" // French
	AR Lang = "ar" // Arabic
	BN Lang = "bn" // Bengali
	PT Lang = "pt" // Portuguese
	UR Lang = "ur" // Urdu
)

// AllUILangs lists fixed UI languages (excluding auto).
var AllUILangs = []Lang{EN, ZH, HI, ES, FR, AR, BN, PT, RU, UR}

// LanguageListHint is a compact hint for prompts and docs.
const LanguageListHint = "auto | en | zh | hi | es | fr | ar | bn | pt | ru | ur"

// CanonicalLanguage maps user/config input to a Lang. Second return is false if unknown.
func CanonicalLanguage(s string) (Lang, bool) {
	s = strings.ToLower(strings.TrimSpace(s))
	switch s {
	case "en", "english":
		return EN, true
	case "zh", "chinese", "cmn", "hans":
		return ZH, true
	case "hi", "hindi":
		return HI, true
	case "es", "spanish", "espanol", "español":
		return ES, true
	case "fr", "french", "francais", "français":
		return FR, true
	case "ar", "arabic":
		return AR, true
	case "bn", "bengali", "bangla":
		return BN, true
	case "pt", "portuguese":
		return PT, true
	case "ru", "russian":
		return RU, true
	case "ur", "urdu":
		return UR, true
	default:
		return "", false
	}
}

func langFromLocaleTag(tag string) Lang {
	tag = strings.ToLower(strings.TrimSpace(tag))
	tag = strings.ReplaceAll(tag, "-", "_")
	if tag == "" {
		return ""
	}
	if l, ok := CanonicalLanguage(tag); ok {
		return l
	}
	if i := strings.IndexByte(tag, '_'); i > 0 {
		prefix := tag[:i]
		if l, ok := CanonicalLanguage(prefix); ok {
			return l
		}
	}
	if len(tag) >= 2 {
		if l, ok := CanonicalLanguage(tag[:2]); ok {
			return l
		}
	}
	switch tag {
	case "c", "posix":
		return EN
	}
	return ""
}
