HADOLINT_VERSION := 1.17.3
HADOLINT_DIR := $(abspath $(dir $(lastword $(MAKEFILE_LIST))))
HADOLINT := $(HADOLINT_DIR)/hadolint

$(HADOLINT):
	mkdir -p $(dir $@)
	curl -s -L -o $(HADOLINT_DIR)/hadolint \
		https://github.com/hadolint/hadolint/releases/download/v$(HADOLINT_VERSION)/hadolint-linux-x86_64
	chmod +x $@
	touch $@
