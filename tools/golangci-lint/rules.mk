golangci_lint_dir := $(abspath $(dir $(lastword $(MAKEFILE_LIST))))
GOLANGCI_LINT := $(golangci_lint_dir)/bin/golangci-lint

$(GOLANGCI_LINT): $(golangci_lint_dir)/go.mod
	cd $(golangci_lint_dir) && go build -o $@ github.com/golangci/golangci-lint/cmd/golangci-lint
