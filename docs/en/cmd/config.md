# `z-panel config`

Manages **`/etc/z-panel/config.toml`**.

## `config init`

Interactive wizard to create a new configuration (and directory `/etc/z-panel` if missing).

```bash
sudo z-panel config init
```

## `config migrate`

Upgrades an existing config to the current **`schema_version`**, preserving user choices where possible.

```bash
sudo z-panel config migrate
```

You may be prompted to confirm destructive or ambiguous changes.

```bash
z-panel config help
```
