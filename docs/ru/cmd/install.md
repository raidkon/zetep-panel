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
z-panel --ssh=user@host install
```

- С вашей машины выполняется **`ssh -t user@host sudo z-panel install`**; на удалённой стороне те же шаги, что и при локальной установке (при необходимости `config init` / `config migrate`).

На клиенте в `PATH` должен быть `ssh`.

## Справка

```bash
z-panel install help
```
