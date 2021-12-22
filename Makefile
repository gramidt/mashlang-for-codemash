include ./Makefile.Common

TOOLS_MOD_DIR := ./internal/tools

.DEFAULT_GOAL := all

.PHONY: all
all: install-tools impi lint misspell test build

.PHONY: build
build:
	GO111MODULE=on CGO_ENABLED=0 go build -o ./bin/mash_$(GOOS)_$(GOARCH)$(EXTENSION) .

.PHONY: run
run:
	GO111MODULE=on CGO_ENABLED=0 go run main.go

.PHONY: install-tools
install-tools:
	cd $(TOOLS_MOD_DIR) && go install github.com/client9/misspell/cmd/misspell
	cd $(TOOLS_MOD_DIR) && go install github.com/pavius/impi/cmd/impi
	cd $(TOOLS_MOD_DIR) && go install github.com/golangci/golangci-lint/cmd/golangci-lint
