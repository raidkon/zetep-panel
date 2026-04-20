# `z-panel config`

Управляет файлом **`/etc/z-panel/config.toml`**.

## `config init`

Интерактивный мастер создания конфигурации (при необходимости создаёт каталог `/etc/z-panel`).

```bash
sudo z-panel config init
```

## `config migrate`

Обновляет существующий конфиг до текущего **`schema_version`**, по возможности сохраняя пользовательские значения.

```bash
sudo z-panel config migrate
```

Могут запрашиваться подтверждения при неоднозначных или разрушительных изменениях.

```bash
z-panel config help
```
