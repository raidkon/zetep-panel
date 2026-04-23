# `z-panel xray-redirect`

Controls the **full-tunnel / redirect** stack: policy routing (`ip rule` / routing table), `suppress_prefixlength`, default route via the tunnel interface, sysctl `src_valid_mark`, nft anti-leak rules, and optional **cgroup v2** egress marking (iptables `cgroup` path), similar in spirit to **wg-quick**.

**`up` is idempotent for z-panel’s own rules:** re-running `up` removes the previous policy-routing set, the tunnel
table, and the nft/iptables lines tagged for this TUN, then applies a **single** copy (avoids duplicate `ip rule` /
mangle lines from repeated `up` without `down`).

## Subcommands

```bash
sudo z-panel xray-redirect up [flags] <interface>
sudo z-panel xray-redirect down <interface>
```

**`up` flags** (before the interface name) include:

- `--table=N` — routing table and fwmark (default from config)
- `--no-mark` — disable cgroup mark path
- `--ipv6` — IPv6 default `::/0` and IPv6 rules
- `--wan-lookup=auto|off|IP[/mask]` — add `ip -4 rule … from <WAN> lookup main` *before* the `not fwmark → table` rules
  so that replies for services on a public / uplink IPv4 (e.g. VLESS on WAN) are not steered into the tunnel by mistake.
  Default: `auto` (derive WAN from the lowest-metric `default` in `table main` that is *not* the TUN, then the first
  `scope global` address on that interface). Set `off` to omit; set an explicit IPv4 (or `/32` CIDR) if auto-detection
  is wrong.
- `--bypass-unit=auto|…` — systemd unit whose cgroup is used for bypass (default tries common units)
- `--bypass-cgroup=path` — explicit cgroup v2 path from the hierarchy root

**Examples:**

```bash
sudo z-panel xray-redirect up xray2tun
sudo z-panel xray-redirect up --bypass-unit=sing-box xray2tun
sudo z-panel xray-redirect up --table=51844 --ipv6 xray2tun
sudo z-panel xray-redirect down xray2tun
```

Requires root for `up` / `down`. Full flag list and defaults: `z-panel xray-redirect help`.
