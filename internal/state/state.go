package state

// File is serialized state for a raised xray-redirect (stored in config.toml under [xray_redirect.<iface>]).
type File struct {
	Mode         string `json:"mode,omitempty" toml:"mode,omitempty"`
	Interface    string `json:"interface" toml:"interface"`
	Table        string `json:"table" toml:"table"`
	Fwmark       string `json:"fwmark,omitempty" toml:"fwmark,omitempty"`
	WGIPv6       bool   `json:"wg_ipv6,omitempty" toml:"wg_ipv6,omitempty"`
	NoBypassMark bool   `json:"no_bypass_mark,omitempty" toml:"no_bypass_mark,omitempty"`
	BypassCgroup string `json:"bypass_cgroup,omitempty" toml:"bypass_cgroup,omitempty"`
	BypassUnit   string `json:"bypass_unit,omitempty" toml:"bypass_unit,omitempty"`
	WanRulePref  int    `json:"wan_rule_pref,omitempty" toml:"wan_rule_pref,omitempty"`
	WanFromCIDR  string `json:"wan_from_cidr,omitempty" toml:"wan_from_cidr,omitempty"`
}

func Partial(iface, table, mode string, wgIPv6 bool) File {
	return File{Mode: mode, Interface: iface, Table: table, Fwmark: table, WGIPv6: wgIPv6}
}
