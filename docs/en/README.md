# Zetep Panel — Documentation (English)

## Overview

**Zetep Panel** helps you route traffic through a tunnel interface using Linux policy routing (ip rules, routing tables, optional nftables/iptables), similar in spirit to **wg-quick** but oriented toward **Xray** / proxy stacks. It can:

- Bring up **xray-redirect** (full tunnel + cgroup marking + anti-leak)
- Install **bash completion** generated from registered commands
- Manage **config.toml** interactively, including schema migration after upgrades
- Check **UFW** rules and print template commands
- Write **systemd-networkd** `.network` drop-ins for TUN addresses (**xray-tun**)

## Requirements

- Linux with `ip`, root for most operations (`sudo`)
- **Bash 4+** for completion scripts; **bash-completion** package recommended
- Optional: `nft`, `iptables`/`ip6tables`, `systemd-networkd`, `ufw`

## Build

```bash
cd z-panel
go build -o z-panel .
```

Requires **Go 1.21+**.

## Install binary

**Local** (copies running binary to `/usr/local/bin/z-panel`, then config wizard if needed):

```bash
sudo z-panel install
```

**Remote** (from your dev machine; uses `scp` and one `ssh -t` session):

```bash
z-panel --ssh=user@host install
```

Paths are defined in code (`internal/config`): default install path `/usr/local/bin/z-panel`, config `/etc/z-panel/config.toml`.

## Configuration

Configuration file: **`/etc/z-panel/config.toml`** (created by `z-panel config init`).

Important fields include:

- **`language`** — UI language: `auto` or one of `en`, `zh`, `hi`, `es`, `fr`, `ar`, `bn`, `pt`, `ru`, `ur`
- **`schema_version`** — managed by the program; older configs trigger **`z-panel config migrate`** (or migration during `install`)
- Routing table / fwmark, state directory, UFW templates, Xray TUN defaults — see interactive `config init`

Environment overrides for language when `language = auto`: `Z_PANEL_LANG`, `LANGUAGE`, `LC_ALL`, `LC_MESSAGES`, `LANG`.

## Command reference

| Command | Documentation |
|---------|----------------|
| `version` | [cmd/version.md](cmd/version.md) |
| `install` | [cmd/install.md](cmd/install.md) |
| `install-shell` | [cmd/install-shell.md](cmd/install-shell.md) |
| `config` | [cmd/config.md](cmd/config.md) |
| `xray-redirect` | [cmd/xray-redirect.md](cmd/xray-redirect.md) |
| `ufw` | [cmd/ufw.md](cmd/ufw.md) |
| `xray-tun` | [cmd/xray-tun.md](cmd/xray-tun.md) |

Global help:

```bash
z-panel help
z-panel <command> help
```

## License

Licensed under the **GNU General Public License v2.0 or later**.
