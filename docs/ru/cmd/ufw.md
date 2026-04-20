# `z-panel ufw`

Проверяет **UFW** и выводит подсказки или шаблоны команд, чтобы правила туннеля согласовывались с UFW.

## `check`

```bash
sudo z-panel ufw check
z-panel ufw check --lan-cidr=192.168.0.0/22 xray2tun
```

## `masq-check`

Смотрит вывод **`iptables-save -t nat`**: есть ли в **POSTROUTING** правила **MASQUERADE** или **SNAT** с **`-o <интерфейс>`**. Если нет — печатает строку для блока **`*nat`** в **`/etc/ufw/before.rules`** и напоминает **`sudo ufw reload`**.

```bash
sudo z-panel ufw masq-check xray2tun
sudo z-panel ufw masq-check --lan-cidr=192.168.0.0/22 xray2tun
```

CIDR LAN в подсказке по умолчанию берётся из **`config.toml`** (`default_lan_cidr`).

```bash
z-panel ufw help
```

Обычно нужен **root** для чтения таблицы **nat**.
