package redirect

import (
	"strconv"
	"strings"
)

// removeXrayRedirectPolicyRouting removes all policy rules and the default route in `table` created by
// xray-redirect, including the "from <WAN> lookup main" rule. Safe to call repeatedly.
func removeXrayRedirectPolicyRouting(table string, ipv6 bool) {
	deleteWanMainLookupV4ByTable(table)
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
}

func deleteWanMainLookupV4ByTable(table string) {
	runIPQuiet("ip", "-4", "rule", "del", "pref", strconv.Itoa(wanMainBypassPrefV4(table)))
	// z-panel 0.14.0 first build: WAN pref collided with auto-assigned not-fwmark order
	runIPQuiet("ip", "-4", "rule", "del", "pref", strconv.Itoa(legacyWanRulePrefV4(table)))
}
