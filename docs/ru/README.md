# Zetep Panel — документация (русский)

## Обзор

**Zetep Panel** (`z-panel`) — утилита командной строки для Linux, которая настраивает **policy routing** в духе **wg-quick**: маршрут по умолчанию через выбранный интерфейс, при необходимости обход через **cgroup** и метки пакетов, правила файрвола против утечек, а также вспомогательные команды для **UFW** и **systemd-networkd** (адресация TUN).

## Требования

- Linux с утилитой `ip`; для большинства операций нужен root (`sudo`)
- **Bash 4+** для скриптов автодополнения; желательно пакет **bash-completion**
- По желанию: `nft`, `iptables`/`ip6tables`, `systemd-networkd`, `ufw`

## Сборка

```bash
cd z-panel
go build -o z-panel .
```

Нужен **Go 1.21+**.

## Установка бинарника

**Локально** (копирует текущий исполняемый файл в `/usr/local/bin/z-panel`, при необходимости запускает мастер конфигурации):

```bash
sudo z-panel install
```

**Удалённо** (с вашей машины разработки; используются `scp` и один сеанс `ssh -t`):

```bash
z-panel --ssh=user@host install
```

Пути по умолчанию заданы в коде (`internal/config`): бинарник `/usr/local/bin/z-panel`, конфиг `/etc/z-panel/config.toml`.

## Конфигурация

Файл: **`/etc/z-panel/config.toml`** (создаётся через `z-panel config init`).

Важные поля:

- **`language`** — язык: `auto` или один из `en`, `zh`, … При **`auto`** используются только стандартные переменные локали: **`LANGUAGE`**, **`LC_ALL`**, **`LC_MESSAGES`**, **`LANG`** (без `Z_PANEL_*`).
- **`no_banner`**, **`ssh_no_tty`**, **`ssh_no_multiplex`** — необязательные флаги; см. комментарий в шапке файла при сохранении
- **`schema_version`** — ведётся программой; при устаревшей схеме потребуется **`z-panel config migrate`** (или миграция при `install`)
- Таблицы маршрутизации, fwmark, шаблоны UFW, параметры Xray TUN — см. интерактивный `config init` (один `config.toml`; снимок `xray-redirect` тоже в нём)

Настройки z-panel читаются **только из `config.toml`**; переменных окружения вида `Z_PANEL_*` для конфигурации и SSH нет.

## Справка по командам

| Команда | Документация |
|---------|--------------|
| `version` | [cmd/version.md](cmd/version.md) |
| `install` | [cmd/install.md](cmd/install.md) |
| `install-shell` | [cmd/install-shell.md](cmd/install-shell.md) |
| `config` | [cmd/config.md](cmd/config.md) |
| `xray-redirect` | [cmd/xray-redirect.md](cmd/xray-redirect.md) |
| `ufw` | [cmd/ufw.md](cmd/ufw.md) |
| `xray-tun` | [cmd/xray-tun.md](cmd/xray-tun.md) |

Общая справка:

```bash
z-panel help
z-panel <команда> help
```

## Лицензия

Программа распространяется на условиях **GNU General Public License версии 2 или любой более поздней**.
