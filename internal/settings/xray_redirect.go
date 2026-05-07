package settings

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"z-panel/internal/config"
	"z-panel/internal/root"
	"z-panel/internal/state"
)

// ErrNoXrayRedirectState means neither config nor legacy JSON has a snapshot for the interface.
var ErrNoXrayRedirectState = errors.New("no xray-redirect state")

func sanitizeIfaceKey(iface string) string {
	return strings.Map(func(r rune) rune {
		if r == '/' || r == '.' {
			return '_'
		}
		return r
	}, iface)
}

// mergeLegacyStateFiles imports /etc/z-panel/state/<iface>.json into XrayRedirect when the map is empty.
func mergeLegacyStateFiles(c *Cfg) {
	if len(c.XrayRedirect) > 0 {
		return
	}
	legacyDir := filepath.Join(config.ConfigDir, "state")
	entries, err := os.ReadDir(legacyDir)
	if err != nil {
		return
	}
	m := make(map[string]state.File)
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".json") {
			continue
		}
		path := filepath.Join(legacyDir, e.Name())
		b, err := os.ReadFile(path)
		if err != nil {
			continue
		}
		var st state.File
		if err := json.Unmarshal(b, &st); err != nil || strings.TrimSpace(st.Interface) == "" {
			continue
		}
		m[sanitizeIfaceKey(st.Interface)] = st
	}
	if len(m) == 0 {
		return
	}
	c.XrayRedirect = m
}

// PersistXrayRedirect saves xray-redirect snapshot into config.toml.
func PersistXrayRedirect(st state.File) error {
	if err := root.Require(); err != nil {
		return err
	}
	if C == nil {
		if err := Load(); err != nil {
			return err
		}
	}
	if C.XrayRedirect == nil {
		C.XrayRedirect = make(map[string]state.File)
	}
	key := sanitizeIfaceKey(st.Interface)
	C.XrayRedirect[key] = st
	if err := Write(*C); err != nil {
		return err
	}
	_ = os.Remove(legacyStateJSONPath(st.Interface))
	return nil
}

// LoadXrayRedirect returns persisted state for xray-redirect down, or legacy JSON file.
func LoadXrayRedirect(iface string) (state.File, error) {
	if C == nil {
		if err := Load(); err != nil {
			return state.File{}, err
		}
	}
	key := sanitizeIfaceKey(iface)
	if st, ok := C.XrayRedirect[key]; ok {
		return st, nil
	}
	// Legacy per-file JSON (pre single-config design).
	legacy := legacyStateJSONPath(iface)
	b, err := os.ReadFile(legacy)
	if err != nil {
		if os.IsNotExist(err) {
			return state.File{}, fmt.Errorf("%s: %w", iface, ErrNoXrayRedirectState)
		}
		return state.File{}, err
	}
	var st state.File
	if err := json.Unmarshal(b, &st); err != nil {
		return state.File{}, err
	}
	return st, nil
}

// RemoveXrayRedirectEntry deletes snapshot from config (after successful down).
func RemoveXrayRedirectEntry(iface string) error {
	if err := root.Require(); err != nil {
		return err
	}
	if C == nil {
		if err := Load(); err != nil {
			return err
		}
	}
	key := sanitizeIfaceKey(iface)
	if C.XrayRedirect != nil {
		delete(C.XrayRedirect, key)
	}
	if err := Write(*C); err != nil {
		return err
	}
	_ = os.Remove(legacyStateJSONPath(iface))
	return nil
}

func legacyStateJSONPath(iface string) string {
	return filepath.Join(config.ConfigDir, "state", sanitizeIfaceKey(iface)+".json")
}
