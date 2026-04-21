# Zetep Panel — التوثيق (العربية)

## نظرة عامة

**Zetep Panel** (`z-panel`) أداة سطر أوامر لنظام Linux تضبط **التوجيه السياسي** بأسلوب يشبه **wg-quick**: المسار الافتراضي عبر واجهة مختارة، ووسم اختياري عبر **cgroup**، وقواعد جدار ناري لتقليل التسرب، ومساعدات لـ **UFW** و**systemd-networkd** (عناوين TUN).

## المتطلبات

- Linux مع أداة `ip`؛ معظم العمليات تحتاج `sudo`
- **Bash 4+** لإكمال التبويب؛ يُنصح بـ **bash-completion**
- اختياري: `nft`، `iptables`/`ip6tables`، `systemd-networkd`، `ufw`

## البناء

```bash
cd z-panel
go build -o z-panel .
```

يُشترط **Go 1.21+**.

## التثبيت

محليًا:

```bash
sudo z-panel install
```

عن بُعد (مع `scp` و`ssh` على العميل):

```bash
z-panel --ssh=user@host install
```

المسارات الافتراضية: الثنائي `/usr/local/bin/z-panel`، الإعداد `/etc/z-panel/config.toml`.

## الإعداد

- **`language`**: `auto` أو أحد `en`، `zh`، `hi`، `es`، `fr`، `ar`، `bn`، `pt`، `ru`، `ur`
- **`schema_version`** تديره البرنامج؛ الإعدادات القديمة تحتاج **`config migrate`**
- التفاصيل: `sudo z-panel config init`

## فهرس الأوامر

| الأمر | التوثيق |
|-------|---------|
| `version` | [cmd/version.md](cmd/version.md) |
| `install` | [cmd/install.md](cmd/install.md) |
| `install-shell` | [cmd/install-shell.md](cmd/install-shell.md) |
| `config` | [cmd/config.md](cmd/config.md) |
| `xray-redirect` | [cmd/xray-redirect.md](cmd/xray-redirect.md) |
| `ufw` | [cmd/ufw.md](cmd/ufw.md) |
| `xray-tun` | [cmd/xray-tun.md](cmd/xray-tun.md) |

التوثيق الكامل بالإنجليزية: [../en/README.md](../en/README.md).

## الترخيص

**رخصة GNU العمومية الإصدار 2 أو أي إصدار لاحق** (GPL-2.0-or-later).
