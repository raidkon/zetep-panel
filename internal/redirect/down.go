package redirect

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"z-panel/internal/i18n"
	"z-panel/internal/root"
	"z-panel/internal/state"
)

// Down tears down configuration using the interface state file.
func Down(iface string) error {
	if err := root.Require(); err != nil {
		return err
	}
	data, err := os.ReadFile(state.Path(iface))
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf(i18n.T("redirect.down_no_state"), iface, iface)
		}
		return err
	}
	var st state.File
	if err := json.Unmarshal(data, &st); err != nil {
		return err
	}
	if st.Mode == "wg" || st.Mode == "full" {
		return downWG(st)
	}
	return fmt.Errorf(i18n.T("redirect.down_bad_mode"), st.Mode, iface)
}

func downWG(st state.File) error {
	removeWGStyleFirewall(st.Interface)
	table := st.Table
	for strings.Contains(ipRuleShow("-4"), "lookup "+table) {
		runIPQuiet("ip", "-4", "rule", "delete", "table", table)
	}
	for strings.Contains(ipRuleShow("-4"), "suppress_prefixlength 0") {
		runIPQuiet("ip", "-4", "rule", "delete", "table", "main", "suppress_prefixlength", "0")
	}
	if st.WGIPv6 {
		for strings.Contains(ipRuleShow("-6"), "lookup "+table) {
			runIPQuiet("ip", "-6", "rule", "delete", "table", table)
		}
		for strings.Contains(ipRuleShow("-6"), "suppress_prefixlength 0") {
			runIPQuiet("ip", "-6", "rule", "delete", "table", "main", "suppress_prefixlength", "0")
		}
	}
	ipRouteFlushTableQuiet(table)
	_ = os.Remove(state.Path(st.Interface))
	fmt.Printf(i18n.T("redirect.down_done"), table, st.Interface)
	return nil
}

func downWGQuick(st state.File, ipv6 bool) error {
	removeWGStyleFirewall(st.Interface)
	table := st.Table
	for strings.Contains(ipRuleShow("-4"), "lookup "+table) {
		runIPQuiet("ip", "-4", "rule", "delete", "table", table)
	}
	for strings.Contains(ipRuleShow("-4"), "suppress_prefixlength 0") {
		runIPQuiet("ip", "-4", "rule", "delete", "table", "main", "suppress_prefixlength", "0")
	}
	if ipv6 {
		for strings.Contains(ipRuleShow("-6"), "lookup "+table) {
			runIPQuiet("ip", "-6", "rule", "delete", "table", table)
		}
		for strings.Contains(ipRuleShow("-6"), "suppress_prefixlength 0") {
			runIPQuiet("ip", "-6", "rule", "delete", "table", "main", "suppress_prefixlength", "0")
		}
	}
	ipRouteFlushTableQuiet(table)
	_ = os.Remove(state.Path(st.Interface))
	return nil
}
