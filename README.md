# Zetep Panel

**Zetep Panel** (`z-panel`) is a command-line tool for Linux that sets up **TUN policy routing** in a **wg-quick–style** way: default routes through a chosen interface, optional cgroup-based bypass marking, firewall anti-leak rules, and helpers for **UFW** and **systemd-networkd** TUN addressing.

## Documentation by language

| Language | Documentation |
|----------|----------------|
| English | [docs/en/README.md](docs/en/README.md) |
| Русский | [docs/ru/README.md](docs/ru/README.md) |
| 简体中文 | [docs/zh/README.md](docs/zh/README.md) |
| हिन्दी | [docs/hi/README.md](docs/hi/README.md) |
| Español | [docs/es/README.md](docs/es/README.md) |
| Français | [docs/fr/README.md](docs/fr/README.md) |
| العربية | [docs/ar/README.md](docs/ar/README.md) |
| বাংলা | [docs/bn/README.md](docs/bn/README.md) |
| Português | [docs/pt/README.md](docs/pt/README.md) |
| اردو | [docs/ur/README.md](docs/ur/README.md) |

## Quick start

```bash
go build -o z-panel .
sudo cp z-panel /usr/local/bin/z-panel   # or: sudo z-panel install
sudo z-panel config init
```

See [docs/en/README.md](docs/en/README.md) for configuration, i18n, and command reference.

## License

Licensed under the terms of **GNU General Public License Version 2 or later**.

See [https://www.gnu.org/licenses/old-licenses/gpl-2.0.html](https://www.gnu.org/licenses/old-licenses/gpl-2.0.html) and [https://www.gnu.org/licenses/gpl-3.0.html](https://www.gnu.org/licenses/gpl-3.0.html).
