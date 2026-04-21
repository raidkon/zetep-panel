# Zetep Panel — 文档（简体中文）

## 概述

**Zetep Panel**（`z-panel`）是面向 Linux 的命令行工具，用于以类似 **wg-quick** 的方式配置 **策略路由**：默认路由走选定接口，可选 cgroup 旁路与标记、防泄漏防火墙规则，并提供 **UFW** 与 **systemd-networkd**（TUN 地址）相关辅助命令。

## 环境要求

- 带 `ip` 的 Linux；多数操作需 `sudo`
- 自动补全需 **Bash 4+**；建议安装 **bash-completion**
- 可选：`nft`、`iptables`/`ip6tables`、`systemd-networkd`、`ufw`

## 构建

```bash
cd z-panel
go build -o z-panel .
```

需要 **Go 1.21+**。

## 安装

本地：

```bash
sudo z-panel install
```

远程（需本机 `scp` / `ssh`）：

```bash
z-panel --ssh=user@host install
```

默认路径：`/usr/local/bin/z-panel`，配置 `/etc/z-panel/config.toml`。

## 配置

- **`language`**：`auto` 或 `en`、`zh`、`hi`、`es`、`fr`、`ar`、`bn`、`pt`、`ru`、`ur`
- **`schema_version`** 由程序维护；旧配置需 **`config migrate`**
- 详细字段见 `sudo z-panel config init`

## 命令索引

| 命令 | 文档 |
|------|------|
| `version` | [cmd/version.md](cmd/version.md) |
| `install` | [cmd/install.md](cmd/install.md) |
| `install-shell` | [cmd/install-shell.md](cmd/install-shell.md) |
| `config` | [cmd/config.md](cmd/config.md) |
| `xray-redirect` | [cmd/xray-redirect.md](cmd/xray-redirect.md) |
| `ufw` | [cmd/ufw.md](cmd/ufw.md) |
| `xray-tun` | [cmd/xray-tun.md](cmd/xray-tun.md) |

完整英文说明见 [../en/README.md](../en/README.md)。

## 许可证

**GNU 通用公共许可证第 2 版或更高版本**（GPL-2.0-or-later）。
