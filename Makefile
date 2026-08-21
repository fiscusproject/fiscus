help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo dev)
LDFLAGS := -w -s -X github.com/fiscusproject/fiscus/internal/core/commons.Version=$(VERSION)

build:
	go build -ldflags="$(LDFLAGS)" -o dist/server ./cmd/app

deps:
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@v1.31.1

generate:
	go generate ./...
