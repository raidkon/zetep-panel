# `z-panel xray-tun`

Записывает drop-in `.network` для **systemd-networkd**, чтобы у TUN-интерфейса был стабильный IPv4-адрес (и при необходимости маршрут к пиру), затем перезагружает **networkd**.

## Синтаксис

```bash
sudo z-panel xray-tun up [флаги] <interfaceName> ip
sudo z-panel xray-tun up [флаги] <interfaceName> <address[/mask]> [<peer[/mask]>]
sudo z-panel xray-tun down <interfaceName> ip
```

- **Режим `ip`** — после `up … <iface> ip` адрес и пир берутся из значений по умолчанию в конфиге, если не заданы **`--address`** / **`--peer`**.
- **Явный режим** — `up … <iface> <addr/mask> [<peer/mask>]` задаёт Address и при необходимости peer.

**Флаги:** `--address=A` (IPv4), `--peer=P` (необязательно).

Нужны **systemd-networkd** и root.

```bash
z-panel xray-tun help
```
