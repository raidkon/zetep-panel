package redirect

import (
	"errors"
	"fmt"

	"z-panel/internal/i18n"
	"z-panel/internal/root"
	"z-panel/internal/settings"
	"z-panel/internal/state"
)

// Down tears down configuration using persisted state in config.toml (or legacy JSON).
func Down(iface string) error {
	if err := root.Require(); err != nil {
		return err
	}
	st, err := settings.LoadXrayRedirect(iface)
	if err != nil {
		if errors.Is(err, settings.ErrNoXrayRedirectState) {
			return fmt.Errorf(i18n.T("redirect.down_no_state"), iface, iface)
		}
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
	if err := settings.RemoveXrayRedirectEntry(st.Interface); err != nil {
		return err
	}
	fmt.Printf(i18n.T("redirect.down_done"), st.Table, st.Interface)
	return nil
}

func downWGQuick(st state.File, ipv6 bool) error {
	removeWGStyleFirewall(st.Interface)
	removeXrayRedirectPolicyRouting(st.Table, ipv6)
	_ = settings.RemoveXrayRedirectEntry(st.Interface)
	return nil
}
