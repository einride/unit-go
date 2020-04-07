GIT_DIR := $(abspath $(dir $(lastword $(MAKEFILE_LIST))))

# git-verify-nodiff: verify that there is no differences between the staging area and the working directory.
.PHONY: git-verify-nodiff
git-verify-nodiff:
	$(GIT_DIR)/git-verify-nodiff.sh
