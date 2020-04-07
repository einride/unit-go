goreview_dir := $(abspath $(dir $(lastword $(MAKEFILE_LIST))))
GOREVIEW := $(goreview_dir)/bin/goreview

$(GOREVIEW): $(goreview_dir)/go.mod
	cd $(goreview_dir) && go build -o $@ github.com/einride/goreview/cmd/goreview
