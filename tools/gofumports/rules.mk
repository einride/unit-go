GOFUMPORTS_VERSION := aaa7156f4122b1055c466e26e77812fa32bac1d9
GOFUMPORTS_dir := $(abspath $(dir $(lastword $(MAKEFILE_LIST))))
GOFUMPORTS := $(GOFUMPORTS_dir)/bin/gofumports

$(GOFUMPORTS):$(GOFUMPORTS_dir)/go.mod
	cd $(GOFUMPORTS_dir) && go build -o $@ mvdan.cc/gofumpt/gofumports
