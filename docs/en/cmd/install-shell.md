# `z-panel install-shell`

Installs **bash completion** for `z-panel`. By default (with root) it writes to the system path (e.g. **`/usr/share/bash-completion/completions/z-panel`**). With **`--user` / `-u`**, installs under **`XDG_DATA_HOME`** or **`~/.local/share/bash-completion/completions/`** without root.

```bash
sudo z-panel install-shell
z-panel install-shell --user
```

Typical use: tab-complete subcommands and `help` after installation.

```bash
z-panel install-shell help
```
