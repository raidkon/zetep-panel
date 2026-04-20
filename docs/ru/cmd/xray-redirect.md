# `z-panel xray-redirect`

Управляет стеком **полного туннеля / redirect**: policy routing (`ip rule` / таблица маршрутизации), `suppress_prefixlength`, маршрут по умолчанию через туннель, sysctl `src_valid_mark`, правила nft против утечек и при необходимости маркировка исходящего трафика через **cgroup v2** (iptables `cgroup`), в духе **wg-quick**.

## Подкоманды

```bash
sudo z-panel xray-redirect up [флаги] <interface>
sudo z-panel xray-redirect down <interface>
```

**Флаги `up`** (до имени интерфейса), в том числе:

- `--table=N` — таблица маршрутизации и fwmark (по умолчанию из конфига)
- `--no-mark` — отключить путь маркировки cgroup
- `--ipv6` — IPv6 default `::/0` и правила для IPv6
- `--bypass-unit=auto|…` — systemd-юнит, чей cgroup используется для обхода
- `--bypass-cgroup=path` — явный путь cgroup v2 от корня иерархии

**Примеры:**

```bash
sudo z-panel xray-redirect up xray2tun
sudo z-panel xray-redirect up --bypass-unit=sing-box xray2tun
sudo z-panel xray-redirect up --table=51844 --ipv6 xray2tun
sudo z-panel xray-redirect down xray2tun
```

Для `up` / `down` нужен root. Полный список флагов: `z-panel xray-redirect help`.
