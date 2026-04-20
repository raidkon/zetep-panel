# Zetep Panel — ডকুমেন্টেশন (বাংলা)

## সংক্ষিপ্ত বিবরণ

**Zetep Panel** (`z-panel`) হল Linux-এর জন্য একটি কমান্ড-লাইন টুল যা **wg-quick**-সদৃশ **policy routing** সেটআপ করে: নির্বাচিত ইন্টারফেসে ডিফল্ট রুট, ঐচ্ছিক **cgroup** মার্কিং, লিক-প্রতিরোধী ফায়ারওয়াল নিয়ম, এবং **UFW** ও **systemd-networkd** (TUN ঠিকানা) সহায়ক কমান্ড।

## প্রয়োজনীয়তা

- `ip` সহ Linux; বেশিরভাগ কাজের জন্য `sudo`
- ট্যাব সম্পূর্ণতার জন্য **Bash 4+**; **bash-completion** সুপারিশকৃত
- ঐচ্ছিক: `nft`, `iptables`/`ip6tables`, `systemd-networkd`, `ufw`

## বিল্ড

```bash
cd z-panel
go build -o z-panel .
```

**Go 1.21+** প্রয়োজন।

## ইনস্টল

স্থানীয়:

```bash
sudo z-panel install
```

রিমোট (ক্লায়েন্টে `scp` / `ssh`):

```bash
z-panel install user@host
```

ডিফল্ট পথ: বাইনারি `/usr/local/bin/z-panel`, কনফিগ `/etc/z-panel/config.toml`।

## কনফিগারেশন

- **`language`**: `auto` অথবা `en`, `zh`, `hi`, `es`, `fr`, `ar`, `bn`, `pt`, `ru`, `ur` এর একটি
- **`schema_version`** প্রোগ্রাম দ্বারা পরিচালিত; পুরনো কনফিগের জন্য **`config migrate`**
- বিস্তারিত: `sudo z-panel config init`

## কমান্ড সূচি

| কমান্ড | ডকুমেন্ট |
|--------|----------|
| `version` | [cmd/version.md](cmd/version.md) |
| `install` | [cmd/install.md](cmd/install.md) |
| `install-shell` | [cmd/install-shell.md](cmd/install-shell.md) |
| `config` | [cmd/config.md](cmd/config.md) |
| `xray-redirect` | [cmd/xray-redirect.md](cmd/xray-redirect.md) |
| `ufw` | [cmd/ufw.md](cmd/ufw.md) |
| `xray-tun` | [cmd/xray-tun.md](cmd/xray-tun.md) |

সম্পূর্ণ ইংরেজি ডকুমেন্ট: [../en/README.md](../en/README.md)।

## লাইসেন্স

**GNU General Public License সংস্করণ 2 বা তার পরের** (GPL-2.0-or-later)।
