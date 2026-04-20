# `z-panel install`

Устанавливает бинарник **z-panel** и при необходимости запускает настройку конфигурации.

## Локальная установка (на той же машине)

```bash
sudo z-panel install
```

- Копирует **текущий запущенный** исполняемый файл в `/usr/local/bin/z-panel` (путь по умолчанию).
- Если нет `/etc/z-panel/config.toml`, запускается интерактивный **`z-panel config init`**.
- Если конфиг есть, но схема старше встроенной, выполняется **`z-panel config migrate`** (с подтверждением).

## Удалённая установка

```bash
z-panel install user@host
```

- Через `scp` копирует локальный бинарник `z-panel` на удалённый хост (тот же путь по умолчанию).
- Через `ssh -t user@host` выполняет там `sudo z-panel install` (возможны `config init` / `config migrate`).

На клиенте в `PATH` должны быть `scp` и `ssh`.

## Справка

```bash
z-panel install help
```
