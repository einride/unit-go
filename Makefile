# all: all tasks required for a complete build
.PHONY: all
all: \
	markdown-lint \
	go-mod-tidy \
	go-lint \
	go-review \
	go-test \
	git-verify-nodiff

include tools/git-no-diff/rules.mk
include tools/prettier/rules.mk
include tools/gofumports/rules.mk
include tools/goreview/rules.mk
include tools/golangci-lint/rules.mk
include tools/hadolint/rules.mk

export GO111MODULE := on

.PHONY: go-mod-tidy
go-mod-tidy:
	go mod tidy -v

.PHONY: go-lint
go-lint: $(GOLANGCI_LINT)
	# dupl: Disabled due to duplication between units (TODO: code-generate)
	$(GOLANGCI_LINT) run --enable-all --disable dupl,gomnd,wsl,funlen

# go-test: run Go test suite
.PHONY: go-test
go-test:
	go test -count 1 -race -cover ./...

# markdown-lint: lint Markdown files with markdownlint
.PHONY: markdown-lint
markdown-lint: $(PRETTIER)
	$(PRETTIER) --parser markdown --check *.md

# docker-lint: lint Dockerfiles with Hadolint
.PHONY: docker-lint
docker-lint: $(HADOLINT)
	git ls-files --exclude='Dockerfile*' --ignored | xargs -L 1 $(HADOLINT)

# go-review: review Go code with goreview
.PHONY: go-review
go-review: $(GOREVIEW)
	$(GOREVIEW) -c 1 ./...
