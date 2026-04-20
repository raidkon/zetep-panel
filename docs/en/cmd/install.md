# `z-panel install`

Installs the **z-panel** binary and runs configuration setup when needed.

## Local install (on the machine where you run the command)

```bash
sudo z-panel install
```

- Copies the **currently running** executable to `/usr/local/bin/z-panel` (default path).
- If `/etc/z-panel/config.toml` is missing, runs **`z-panel config init`** interactively.
- If the config exists but is older than the built-in schema, runs **`z-panel config migrate`** (with confirmation).

## Remote install

```bash
z-panel install user@host
```

- Uses `scp` to copy your local `z-panel` binary to the remote host (same default path).
- Uses `ssh -t user@host` to run `sudo z-panel install` there (which may trigger `config init` / `config migrate`).

Requires `scp` and `ssh` in `PATH` on the client.

## Help

```bash
z-panel install help
```
