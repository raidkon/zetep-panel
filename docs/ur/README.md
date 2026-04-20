# Zetep Panel — دستاویزات (اردو)

## تعارف

**Zetep Panel** (`z-panel`) Linux کے لیے کمانڈ لائن ٹول ہے جو **wg-quick** جیسا **policy routing** سیٹ اپ کرتا ہے: منتخب انٹرفیس پر ڈیفالٹ روٹ، اختیاری **cgroup** مارکنگ، لیک روکنے والے فائر وال اصول، اور **UFW** و **systemd-networkd** (TUN ایڈریسنگ) معاون کمانڈز۔

## ضروریات

- `ip` والا Linux؛ زیادہ تر عملیات کے لیے `sudo`
- ٹیب مکملی کے لیے **Bash 4+**؛ **bash-completion** تجویز کردہ
- اختیاری: `nft`، `iptables`/`ip6tables`، `systemd-networkd`، `ufw`

## بلڈ

```bash
cd z-panel
go build -o z-panel .
```

**Go 1.21+** درکار۔

## انسٹال

مقامی:

```bash
sudo z-panel install
```

ریموٹ (کلائنٹ پر `scp` / `ssh`):

```bash
z-panel install user@host
```

ڈیفالٹ راستے: بائنری `/usr/local/bin/z-panel`، کنفیگ `/etc/z-panel/config.toml`۔

## کنفیگریشن

- **`language`**: `auto` یا `en`، `zh`، `hi`، `es`، `fr`، `ar`، `bn`، `pt`، `ru`، `ur` میں سے ایک
- **`schema_version`** پروگرام سنبھالتا ہے؛ پرانے کنفیگ کے لیے **`config migrate`**
- تفصیل: `sudo z-panel config init`

## کمانڈ فہرست

| کمانڈ | دستاویز |
|--------|---------|
| `version` | [cmd/version.md](cmd/version.md) |
| `install` | [cmd/install.md](cmd/install.md) |
| `install-shell` | [cmd/install-shell.md](cmd/install-shell.md) |
| `config` | [cmd/config.md](cmd/config.md) |
| `xray-redirect` | [cmd/xray-redirect.md](cmd/xray-redirect.md) |
| `ufw` | [cmd/ufw.md](cmd/ufw.md) |
| `xray-tun` | [cmd/xray-tun.md](cmd/xray-tun.md) |

مکمل انگریزی دستاویز: [../en/README.md](../en/README.md)۔

## لائسنس

**GNU General Public License ورژن 2 یا اس کے بعد** (GPL-2.0-or-later)۔
