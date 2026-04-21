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
z-panel --ssh=user@host install
```

- Runs **`ssh -t user@host sudo z-panel install`** from your machine; the remote side performs the same steps as a local install (including `config init` / `config migrate` when needed).

Requires `ssh` in `PATH` on the client.

## Help

```bash
z-panel install help
```
