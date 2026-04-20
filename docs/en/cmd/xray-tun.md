# `z-panel xray-tun`

Writes **systemd-networkd** drop-in `.network` files so a TUN interface gets a stable IPv4 address (and optional peer route), then reloads **networkd**.

## Syntax

```bash
sudo z-panel xray-tun up [flags] <interfaceName> ip
sudo z-panel xray-tun up [flags] <interfaceName> <address[/mask]> [<peer[/mask]>]
sudo z-panel xray-tun down <interfaceName> ip
```

- **`ip` mode** — after `up … <iface> ip`, address/peer come from config defaults unless overridden with **`--address`** / **`--peer`**.
- **Explicit mode** — `up … <iface> <addr/mask> [<peer/mask>]` sets Address and optional peer on the wire.

**Flags:** `--address=A` (IPv4), `--peer=P` (optional).

Requires **systemd-networkd** and root.

```bash
z-panel xray-tun help
```
