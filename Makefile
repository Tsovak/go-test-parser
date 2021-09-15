export GOPATH ?= $(shell go env GOPATH)
export GO111MODULE ?= on
GOBASE=$(shell pwd)
export GOBIN=$(GOBASE)/bin
BIN_DIR = bin
APPNAME = app
LDFLAGS ?=

#.DEFAULT_GOAL := all

.PHONY: all
all: build

.PHONY: mod
mod:
	go mod download

.PHONY: clean
clean: ## run all cleanup tasks
	go clean ./...
	rm -rf $(BIN_DIR)
	cd cmd/go-runner && packr2 clean
	cd cmd/go-test-parser && packr2 clean

golangci: ## install golangci-linter
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b ${BIN_DIR} v1.21.0

install_packr: ## install packr2
	go install github.com/gobuffalo/packr/v2/packr2

.PHONY: install_deps
install_deps: golangci install_packr ## install necessary dependencies

.PHONY: build
build:  ## build all applications
	go build -ldflags "$(LDFLAGS)" -o $(BIN_DIR)/go-test-parser cmd/go-test-parser/*.go
	go build -ldflags "$(LDFLAGS)" -o $(BIN_DIR)/go-runner cmd/go-runner/*.go

.PHONY: packr
packr:  ##
	cd cmd/go-runner && $(GOBIN)/packr2
	cd cmd/go-test-parser && $(GOBIN)/packr2


.PHONY: unit
unit:  ## run unit tests
	go test -v ./... -count 10 -race

.PHONY: unit-ci
unit-ci:  ## run unit tests with json output
	go test -v ./... -count 10 -race -json

.PHONY: test-with-coverage
test-with-coverage: ## run tests with coverage mode
	go test -v ./... -count 1 -race -coverprofile=coverage.out


.PHONY: lint
lint: golangci ## run linter
	${BIN_DIR}/golangci-lint --color=always run ./... -v --timeout 5m

.PHONY: help
help: ## display help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
	
