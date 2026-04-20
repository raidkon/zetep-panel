# Zetep Panel — Documentação (português)

## Visão geral

**Zetep Panel** (`z-panel`) é uma ferramenta de linha de comando para Linux que configura **roteamento por política** no estilo **wg-quick**: rota padrão por uma interface escolhida, marcação opcional via **cgroup**, regras de firewall anti-vazamento, e utilitários para **UFW** e **systemd-networkd** (endereçamento TUN).

## Requisitos

- Linux com `ip`; a maioria das operações exige `sudo`
- **Bash 4+** para conclusão de comandos; recomenda-se **bash-completion**
- Opcional: `nft`, `iptables`/`ip6tables`, `systemd-networkd`, `ufw`

## Compilar

```bash
cd z-panel
go build -o z-panel .
```

É necessário **Go 1.21+**.

## Instalação

Local:

```bash
sudo z-panel install
```

Remota (cliente com `scp` e `ssh`):

```bash
z-panel install user@host
```

Caminhos padrão: binário `/usr/local/bin/z-panel`, configuração `/etc/z-panel/config.toml`.

## Configuração

- **`language`**: `auto` ou um de `en`, `zh`, `hi`, `es`, `fr`, `ar`, `bn`, `pt`, `ru`, `ur`
- **`schema_version`** é mantido pelo programa; configs antigas precisam de **`config migrate`**
- Detalhes: `sudo z-panel config init`

## Índice de comandos

| Comando | Documentação |
|---------|---------------|
| `version` | [cmd/version.md](cmd/version.md) |
| `install` | [cmd/install.md](cmd/install.md) |
| `install-shell` | [cmd/install-shell.md](cmd/install-shell.md) |
| `config` | [cmd/config.md](cmd/config.md) |
| `xray-redirect` | [cmd/xray-redirect.md](cmd/xray-redirect.md) |
| `ufw` | [cmd/ufw.md](cmd/ufw.md) |
| `xray-tun` | [cmd/xray-tun.md](cmd/xray-tun.md) |

Documentação completa em inglês: [../en/README.md](../en/README.md).

## Licença

**GNU General Public License versão 2 ou posterior** (GPL-2.0-or-later).
