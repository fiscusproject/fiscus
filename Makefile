help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

build:
	go build -ldflags="-w -s" -o dist/server ./cmd/app

deps:
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@v1.31.1

generate:
	go generate ./...
