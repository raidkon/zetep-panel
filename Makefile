# Default: plain go build uses Version from internal/config/config.go.
# With git: embeds `git describe --tags --always --dirty` so each commit changes the reported version.
GIT := $(shell git describe --tags --always --dirty 2>/dev/null)
LDFLAGS := $(if $(GIT),-ldflags '-X z-panel/internal/config.Version=$(GIT)')

.PHONY: build test
build:
	go build $(LDFLAGS) -o z-panel .

test:
	go test ./...
