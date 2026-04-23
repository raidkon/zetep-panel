package state

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"z-panel/internal/i18n"
	"z-panel/internal/settings"
)

// File is serialized state for a raised xray-redirect.
type File struct {
	Mode         string `json:"mode,omitempty"`
	Interface    string `json:"interface"`
	Table        string `json:"table"`
	Fwmark       string `json:"fwmark,omitempty"`
	WGIPv6       bool   `json:"wg_ipv6,omitempty"`
	NoBypassMark bool   `json:"no_bypass_mark,omitempty"`
	BypassCgroup string `json:"bypass_cgroup,omitempty"`
	BypassUnit   string `json:"bypass_unit,omitempty"`
	// optional policy: "from <WanCIDR> lookup main" for inbound services on a public / WAN IP
	WanRulePref int    `json:"wan_rule_pref,omitempty"`
	WanFromCIDR string `json:"wan_from_cidr,omitempty"`
}

func Path(iface string) string {
	safe := strings.Map(func(r rune) rune {
		if r == '/' || r == '.' {
			return '_'
		}
		return r
	}, iface)
	return filepath.Join(settings.C.StateDir, safe+".json")
}

func WriteAndPrint(st File, summary string) error {
	data, err := json.MarshalIndent(st, "", "  ")
	if err != nil {
		return err
	}
	if err := os.WriteFile(Path(st.Interface), data, 0o600); err != nil {
		return fmt.Errorf(i18n.T("state.state_file_err"), err)
	}
	fmt.Printf(i18n.T("state.up_line"), summary, Path(st.Interface))
	return nil
}

func Partial(iface, table, mode string, wgIPv6 bool) File {
	return File{Mode: mode, Interface: iface, Table: table, Fwmark: table, WGIPv6: wgIPv6}
}
