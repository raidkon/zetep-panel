# `z-panel install`

Installs the **z-panel** binary and runs configuration setup when needed.

## Local install (on the machine where you run the command)

```bash
sudo z-panel install
```

- Copies the **currently running** executable to `/usr/local/bin/z-panel` (default path).
- If `/etc/z-panel/config.toml` is missing, runs **`z-panel config init`** interactively.
- If the config exists but is older than the built-in schema, runs **`z-panel config migrate`** (with confirmation).

## Remote install (from your machine: upload this binary, then install on the host)

```bash
z-panel --ssh=user@host install
```

- Copies **this** `z-panel` binary to the host with **`scp`**, then runs **`ssh -t user@host`** with `chmod` and *
  *`sudo /tmp/… install`**, so the same steps run on the **remote** system as a local `sudo z-panel install` (including
  `config init` / `config migrate` when needed).
- Requires **`ssh`** and **`scp`** in `PATH` on the client.

If the program is **already** installed on the server and you are not replacing it from the client, use:

```bash
z-panel --ssh-connect=user@host install
```

(runs the **remote** `z-panel` on the host — it copies that binary to the install path, not your local build.)

## Help

```bash
z-panel install help
```
