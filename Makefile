# all: all tasks required for a complete build
.PHONY: all
all: \
	circleci-config-validate \
	markdown-lint \
	mod-tidy \
	go-generate \
	go-lint \
	go-review \
	go-test \
	git-verify-submodules \
	git-verify-nodiff

export GO111MODULE := on

# clean: remove generated build files
.PHONY: clean
clean:
	rm -rf \
		internal/common/funnel \
		vendor \
		build \
		test/mocks
	find -name '*wire_gen.go' -exec rm {} \+

.PHONY: build
build:
	@git submodule update --init --recursive $@

include build/rules.mk
build/rules.mk: build
	@# Included in submodule: build

.PHONY: mod-tidy
mod-tidy:
	go mod tidy

.PHONY: go-lint
go-lint: $(GOLANGCI_LINT)
	$(GOLANGCI_LINT) run --enable-all

ifeq ($(CI),true)
# prevent OOM error when CircleCI erroneously reports 32 available VCPUs
GO_BUILD_FLAGS := -p 2
else
GO_BUILD_FLAGS :=
endif

# go-generate: (re-)generate Go code using go generate
.PHONY: go-generate
go-generate: $(STRINGER) $(WIRE)
	go generate $(GO_BUILD_FLAGS) ./...

# go-test: run Go test suite
.PHONY: go-test
go-test:
	go test $(GO_BUILD_FLAGS) -count 1 -race -cover ./...

# markdown-lint: lint Markdown files with markdownlint
.PHONY: markdown-lint
markdown-lint: $(MARKDOWNLINT)
	$(MARKDOWNLINT) --ignore build --ignore vendor --ignore waysure-log-collector-sidecar .

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
