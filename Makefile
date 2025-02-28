# Copyright (c) The Kowabunga Project
# Apache License, Version 2.0 (see LICENSE or https://www.apache.org/licenses/LICENSE-2.0.txt)
# SPDX-License-Identifier: Apache-2.0

.DEFAULT_GOAL := all

BIN_DIR = bin
DOCS_DIR = docs
SDK_DIR = sdk
PACKAGE_NAME = kowabunga

include Make.api.mk
include Make.go.mk
include Make.typescript.mk

V = 0
Q = $(if $(filter 1,$V),,@)
M = $(shell printf "\033[34;1m▶\033[0m")

ifeq ($(V), 1)
  OUT = ""
else
  OUT = ">/dev/null"
endif

.PHONY: all
all: pre sdk post ; @
	$Q echo "done"

# Updates all go modules
.PHONY: update
update: ; $(info $(M) updating modules…) @
	$Q go get -u ./...
	$Q go mod tidy

.PHONY: pre
pre: ; $(info $(M) [Git] cleaning up auto-generated code…) @
	$Q git rm -rf --quiet $(DOCS_DIR) $(SDK_DIR) || true

.PHONY: post
post: ; $(info $(M) [Git] adding auto-generated code…) @
	$Q git add $(DOCS_DIR) $(SDK_DIR)

.PHONY: bin
bin: ; $(info $(M) [Misc] create local bin directory) @
	$Q mkdir -p $(BIN_DIR)

.PHONY: sdk
sdk: api go-sdk typescript-sdk; @

.PHONY: clean
clean: clean-api clean-go clean-typescript ; $(info $(M) cleaning…) @
	$Q rm -rf $(BIN_DIR)
	$Q rm -rf $(SDK_DIR)
