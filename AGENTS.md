# AGENTS.md

## Cursor Cloud specific instructions

### Overview

`z-panel` is a single-binary Go CLI tool for Linux TUN policy routing. There are no databases, Docker containers, or external services — all tests are pure Go unit tests with mocked system commands.

### Key commands

| Action | Command |
|--------|---------|
| Install deps | `go mod download` |
| Lint | `go vet ./...` |
| Test | `make test` or `go test ./...` |
| Test (CI-style) | `go test ./... -count=1 -race -shuffle=on` |
| Coverage | `make cover` |
| Build | `go build -o z-panel .` |

### Non-obvious notes

- The binary requires **root** and Linux networking stack for most runtime commands (`xray-redirect up`, `config init`, etc.), but **all tests run without root** — they mock system calls.
- `config init` is **interactive** (TTY prompts); avoid running it in non-interactive shells.
- Tests use `Z_PANEL_LANG=en` internally; no env var setup is needed to run them.
- Go 1.21+ is required (`go.mod`); the VM ships Go 1.22.2 which satisfies this.
- The only external Go dependency is `github.com/pelletier/go-toml/v2`.
