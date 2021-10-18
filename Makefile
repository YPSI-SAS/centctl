# MIT License

# Copyright (c)  2020-2021 YPSI SAS

# Permission is hereby granted, free of charge, to any person obtaining a copy
# of this software and associated documentation files (the "Software"), to deal
# in the Software without restriction, including without limitation the rights
# to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
# copies of the Software, and to permit persons to whom the Software is
# furnished to do so, subject to the following conditions:

# The above copyright notice and this permission notice shall be included in all
# copies or substantial portions of the Software.

# THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
# IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
# FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
# AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
# LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
# OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE

LASTTAG := $(shell git describe --tags --abbrev=0)
MODULE := $(shell env GO111MODULE=on go list -m)
REVISION := $(shell git rev-parse HEAD)
WINDOWS := $(MODULE)_windows_amd64.exe
LINUX := $(MODULE)_linux_amd64

# Defines the default target that `make` will to try to make,
# or in the case of a phony target, execute the specified commands
# This target is executed whenever we just type `make`
.DEFAULT_GOAL = help

.PHONY: help all build windows linux darwin clean clean-docker docker

help: ## Print help on Makefile
	@echo "Please use 'make <target>' where <target> is one of"
	@echo ""
	@grep -hE '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-17s\033[0m %s\n", $$1, $$2}'
	@echo ""
	@echo "Check the Makefile to know exactly what each target is doing."

all: clean build ## Build and run tests

build: windows linux darwin ## Build binaries

windows: $(WINDOWS) ## Build for Windows

linux: $(LINUX) ## Build for Linux

darwin: $(DARWIN) ## Build for Darwin (macOS)

$(WINDOWS):
	env GOOS=windows GOARCH=amd64 go build \
	    -ldflags '-w -s -X $(MODULE)/cmd.lastGitTag=$(LASTTAG) -X $(MODULE)/cmd.lastGitCommit=$(REVISION)' \
		-o $(WINDOWS) main.go

$(LINUX):
	env GOOS=linux GOARCH=amd64 go build \
	    -ldflags '-w -s -X $(MODULE)/cmd.lastGitTag=$(LASTTAG) -X $(MODULE)/cmd.lastGitCommit=$(REVISION)' \
		-o $(LINUX) main.go

$(DARWIN):
	env GOOS=darwin GOARCH=amd64 go build \
	    -ldflags '-w -s -X $(MODULE)/cmd.lastGitTag=$(LASTTAG) -X $(MODULE)/cmd.lastGitCommit=$(REVISION)' \
		-o $(DARWIN) main.go

clean: ## Remove previous build
	rm -f $(WINDOWS) $(LINUX) $(DARWIN)

clean-docker: ## Remove Docker image
	docker rmi $(MODULE)

docker: clean-docker ## Build Docker container
	docker build -t $(MODULE) --build-arg lastTag=$(LASTTAG) --build-arg lastCommit=$(REVISION) --no-cache .