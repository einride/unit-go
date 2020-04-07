PRETTIER_VERSION := 2.0.1
PRETTIER_DIR := $(abspath $(dir $(lastword $(MAKEFILE_LIST))))/bin/$(PRETTIER_VERSION)
PRETTIER := $(PRETTIER_DIR)/node_modules/.bin/prettier

$(PRETTIER):
	npm install --no-save --no-audit --prefix $(PRETTIER_DIR) prettier@$(PRETTIER_VERSION)
	chmod +x $@
	touch $@
