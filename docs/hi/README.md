# Zetep Panel — दस्तावेज़ (हिन्दी)

## परिचय

**Zetep Panel** (`z-panel`) Linux के लिए एक कमांड-लाइन टूल है जो **wg-quick** जैसे **policy routing** को सेट करता है: चुने इंटरफ़ेस पर डिफ़ॉल्ट रूट, वैकल्पिक **cgroup** मार्किंग, लीक-रोधी फ़ायरवॉल नियम, और **UFW** तथा **systemd-networkd** (TUN पता) सहायक कमांड।

## आवश्यकताएँ

- `ip` वाला Linux; अधिकांश कार्यों के लिए `sudo`
- टैब-पूर्णता के लिए **Bash 4+**; **bash-completion** अनुशंसित
- वैकल्पिक: `nft`, `iptables`/`ip6tables`, `systemd-networkd`, `ufw`

## बिल्ड

```bash
cd z-panel
go build -o z-panel .
```

**Go 1.21+** आवश्यक।

## इंस्टॉल

स्थानीय:

```bash
sudo z-panel install
```

दूरस्थ (क्लाइंट पर `scp` / `ssh`):

```bash
z-panel --ssh=user@host install
```

डिफ़ॉल्ट पथ: बाइनरी `/usr/local/bin/z-panel`, कॉन्फ़िग `/etc/z-panel/config.toml`।

## कॉन्फ़िगरेशन

- **`language`**: `auto` या `en`, `zh`, `hi`, `es`, `fr`, `ar`, `bn`, `pt`, `ru`, `ur` में से एक
- **`schema_version`** प्रोग्राम द्वारा प्रबंधित; पुराने कॉन्फ़िग के लिए **`config migrate`**
- विवरण: `sudo z-panel config init`

## कमांड सूची

| कमांड | दस्तावेज़ |
|--------|-----------|
| `version` | [cmd/version.md](cmd/version.md) |
| `install` | [cmd/install.md](cmd/install.md) |
| `install-shell` | [cmd/install-shell.md](cmd/install-shell.md) |
| `config` | [cmd/config.md](cmd/config.md) |
| `xray-redirect` | [cmd/xray-redirect.md](cmd/xray-redirect.md) |
| `ufw` | [cmd/ufw.md](cmd/ufw.md) |
| `xray-tun` | [cmd/xray-tun.md](cmd/xray-tun.md) |

पूर्ण अंग्रेज़ी दस्तावेज़: [../en/README.md](../en/README.md)।

## लाइसेंस

**GNU General Public License संस्करण 2 या बाद का** (GPL-2.0-or-later)।
