# `z-panel ufw`

Inspects **UFW** and prints guidance or template commands so your tunnel rules coexist with UFW.

## `check`

Typical workflow: run the suggested checks, then apply the printed rules if they match your policy.

```bash
sudo z-panel ufw check
z-panel ufw check --lan-cidr=192.168.0.0/22 xray2tun
```

## `masq-check`

Checks **`iptables-save -t nat`** for **POSTROUTING** rules with **MASQUERADE** or **SNAT** and outbound interface **`-o <interface>`** (e.g. `ppp0`, `wan1`, `xray2tun`). If none are found, prints a line to add under a **`*nat`** block in **`/etc/ufw/before.rules`** (next to your other MASQUERADE rules), then **`sudo ufw reload`**.

```bash
sudo z-panel ufw masq-check xray2tun
sudo z-panel ufw masq-check --lan-cidr=192.168.0.0/22 xray2tun
```

LAN CIDR in the suggested rule defaults from **`config.toml`** (`default_lan_cidr`).

```bash
z-panel ufw help
```

Reading **`iptables-save`** usually requires **root**.
