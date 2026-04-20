package i18n

import (
	"os"
	"testing"
)

func TestCanonicalLanguage(t *testing.T) {
	t.Parallel()
	tab := []struct {
		in   string
		want Lang
		ok   bool
	}{
		{"en", EN, true},
		{"EN", EN, true},
		{"english", EN, true},
		{"ru", RU, true},
		{"zh", ZH, true},
		{"es", ES, true},
		{"fr", FR, true},
		{"pt", PT, true},
		{"  hi  ", HI, true},
		{"nope", "", false},
	}
	for _, tc := range tab {
		l, ok := CanonicalLanguage(tc.in)
		if ok != tc.ok || l != tc.want {
			t.Fatalf("%q: got (%v,%v) want (%v,%v)", tc.in, l, ok, tc.want, tc.ok)
		}
	}
}

func TestLangFromLocaleTag(t *testing.T) {
	t.Parallel()
	if langFromLocaleTag("ru_RU.UTF-8") != RU {
		t.Fatal()
	}
	if langFromLocaleTag("zh_CN") != ZH {
		t.Fatal()
	}
	if langFromLocaleTag("c") != EN {
		t.Fatal()
	}
	if langFromLocaleTag("posix") != EN {
		t.Fatal()
	}
	if langFromLocaleTag("zz_ZZ") != "" {
		t.Fatalf("unknown should be empty, got %v", langFromLocaleTag("zz_ZZ"))
	}
}

func TestApplyFromConfig_and_T(t *testing.T) {
	t.Cleanup(func() {
		_ = os.Unsetenv("Z_PANEL_LANG")
		Init()
	})
	_ = os.Unsetenv("Z_PANEL_LANG")
	Init()
	ApplyFromConfig("ru")
	if Language() != RU {
		t.Fatal()
	}
	if s := T("root.need_root"); s == "" || s == "root.need_root" {
		t.Fatalf("missing translation: %q", s)
	}
	ApplyFromConfig("auto")
	// unknown fixed lang → EN
	ApplyFromConfig("not-a-lang")
	if Language() != EN {
		t.Fatalf("got %v", Language())
	}
}

func TestT_missingKey(t *testing.T) {
	Init()
	ApplyFromConfig("en")
	if got := T("this.key.does.not.exist"); got != "this.key.does.not.exist" {
		t.Fatalf("got %q", got)
	}
}
