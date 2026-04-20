# `z-panel install-shell`

Устанавливает **bash completion** для `z-panel`. По умолчанию с root скрипт попадает в системный путь (например **`/usr/share/bash-completion/completions/z-panel`**). С **`--user` / `-u`** — в **`XDG_DATA_HOME`** или **`~/.local/share/bash-completion/completions/`** без root.

```bash
sudo z-panel install-shell
z-panel install-shell --user
```

После установки удобно дополнять подкоманды и `help` по Tab.

```bash
z-panel install-shell help
```
