# Zetep Panel — Documentation (français)

## Présentation

**Zetep Panel** (`z-panel`) est un outil en ligne de commande pour Linux qui configure le **routage par politique** à la manière de **wg-quick** : route par défaut via une interface choisie, marquage optionnel via **cgroup**, règles pare-feu anti-fuites, et aides pour **UFW** et **systemd-networkd** (adressage TUN).

## Prérequis

- Linux avec `ip` ; la plupart des opérations nécessitent `sudo`
- **Bash 4+** pour la complétion ; **bash-completion** recommandé
- Optionnel : `nft`, `iptables`/`ip6tables`, `systemd-networkd`, `ufw`

## Compilation

```bash
cd z-panel
go build -o z-panel .
```

**Go 1.21+** requis.

## Installation

Locale :

```bash
sudo z-panel install
```

À distance (`scp` / `ssh` sur la machine cliente) :

```bash
z-panel install user@host
```

Chemins par défaut : binaire `/usr/local/bin/z-panel`, configuration `/etc/z-panel/config.toml`.

## Configuration

- **`language`** : `auto` ou l’un de `en`, `zh`, `hi`, `es`, `fr`, `ar`, `bn`, `pt`, `ru`, `ur`
- **`schema_version`** géré par le programme ; les anciens fichiers nécessitent **`config migrate`**
- Détails : `sudo z-panel config init`

## Référence des commandes

| Commande | Documentation |
|----------|----------------|
| `version` | [cmd/version.md](cmd/version.md) |
| `install` | [cmd/install.md](cmd/install.md) |
| `install-shell` | [cmd/install-shell.md](cmd/install-shell.md) |
| `config` | [cmd/config.md](cmd/config.md) |
| `xray-redirect` | [cmd/xray-redirect.md](cmd/xray-redirect.md) |
| `ufw` | [cmd/ufw.md](cmd/ufw.md) |
| `xray-tun` | [cmd/xray-tun.md](cmd/xray-tun.md) |

Documentation complète en anglais : [../en/README.md](../en/README.md).

## Licence

**GNU General Public License version 2 ou ultérieure** (GPL-2.0-or-later).
