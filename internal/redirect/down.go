package redirect

import (
	"encoding/json"
	"fmt"
	"os"

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
	removeXrayRedirectPolicyRouting(st.Table, st.WGIPv6)
	_ = os.Remove(state.Path(st.Interface))
	fmt.Printf(i18n.T("redirect.down_done"), st.Table, st.Interface)
	return nil
}

func downWGQuick(st state.File, ipv6 bool) error {
	removeWGStyleFirewall(st.Interface)
	removeXrayRedirectPolicyRouting(st.Table, ipv6)
	_ = os.Remove(state.Path(st.Interface))
	return nil
}
