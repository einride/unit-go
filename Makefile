# all: all tasks required for a complete build
.PHONY: all
all: \
	circleci-config-validate \
	markdown-lint \
	go-mod-tidy \
	go-lint \
	go-review \
	go-test \
	git-verify-submodules \
	git-verify-nodiff

export GO111MODULE := on

# clean: remove generated build files
.PHONY: clean
clean:
	rm -rf build

.PHONY: build
build:
	@git submodule update --init --recursive $@

include build/rules.mk
build/rules.mk: build
	@# included in submodule: build

.PHONY: go-mod-tidy
go-mod-tidy:
	go mod tidy -v

.PHONY: go-lint
go-lint: $(GOLANGCI_LINT)
	$(GOLANGCI_LINT) run --enable-all

# go-test: run Go test suite
.PHONY: go-test
go-test:
	go test -count 1 -race -cover ./...

# markdown-lint: lint Markdown files with markdownlint
.PHONY: markdown-lint
markdown-lint: $(MARKDOWNLINT)
	$(MARKDOWNLINT) --ignore build .

# docker-lint: lint Dockerfiles with Hadolint
.PHONY: docker-lint
docker-lint: $(HADOLINT)
	git ls-files --exclude='Dockerfile*' --ignored | xargs -L 1 $(HADOLINT)

# circleci-config-validate: validate the CircleCI build config
.PHONY: circleci-config-validate
circleci-config-validate: $(CIRCLECI)
	$(CIRCLECI) config validate

# go-review: review Go code with goreview
.PHONY: go-review
go-review: $(GOREVIEW)
	$(GOREVIEW) -c 1 ./...
