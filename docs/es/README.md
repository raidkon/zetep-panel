# Zetep Panel — Documentación (español)

## Descripción

**Zetep Panel** (`z-panel`) es una herramienta de línea de comandos para Linux que configura **enrutamiento por políticas** al estilo **wg-quick**: ruta predeterminada por una interfaz elegida, marcado opcional vía **cgroup**, reglas de cortafuegos anti-fugas, y utilidades para **UFW** y **systemd-networkd** (direccionamiento TUN).

## Requisitos

- Linux con `ip`; la mayoría de operaciones requieren `sudo`
- **Bash 4+** para el autocompletado; se recomienda **bash-completion**
- Opcional: `nft`, `iptables`/`ip6tables`, `systemd-networkd`, `ufw`

## Compilar

```bash
cd z-panel
go build -o z-panel .
```

Se necesita **Go 1.21+**.

## Instalación

Local:

```bash
sudo z-panel install
```

Remota (requiere `scp` y `ssh` en el cliente):

```bash
z-panel install user@host
```

Rutas por defecto: binario `/usr/local/bin/z-panel`, configuración `/etc/z-panel/config.toml`.

## Configuración

- **`language`**: `auto` o uno de `en`, `zh`, `hi`, `es`, `fr`, `ar`, `bn`, `pt`, `ru`, `ur`
- **`schema_version`** lo gestiona el programa; configuraciones antiguas requieren **`config migrate`**
- Más campos: `sudo z-panel config init`

## Índice de comandos

| Comando | Documentación |
|---------|----------------|
| `version` | [cmd/version.md](cmd/version.md) |
| `install` | [cmd/install.md](cmd/install.md) |
| `install-shell` | [cmd/install-shell.md](cmd/install-shell.md) |
| `config` | [cmd/config.md](cmd/config.md) |
| `xray-redirect` | [cmd/xray-redirect.md](cmd/xray-redirect.md) |
| `ufw` | [cmd/ufw.md](cmd/ufw.md) |
| `xray-tun` | [cmd/xray-tun.md](cmd/xray-tun.md) |

Documentación completa en inglés: [../en/README.md](../en/README.md).

## Licencia

**GNU General Public License versión 2 o posterior** (GPL-2.0-or-later).
