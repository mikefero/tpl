.DEFAULT_GOAL := build

# Ensure golang is available
ifeq (, $(shell which go 2> /dev/null))
$(error "'go' is not installed or available in PATH")
endif

APP_NAME := tpl
APP_DIR := $(patsubst %/,%,$(dir $(abspath $(lastword $(MAKEFILE_LIST)))))
APP_WORKDIR := $(shell pwd)

include mk/build.mk
include mk/run.mk
include mk/test.mk
include mk/tools.mk

# Catch all for add-command which takes arguments
# Ensures make doesn't fail for unknown job when executing add-command
%:
	@:
