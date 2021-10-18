EXECUTABLE=centctl
WINDOWS=$(EXECUTABLE)_windows_amd64.exe
LINUX=$(EXECUTABLE)_linux_amd64

LASTTAG := $(shell git describe --tags --abbrev=0)
MODULE := $(shell env GO111MODULE=on go list -m)
REVISION := $(shell git rev-parse HEAD)

all: clean build ## Build and run tests

build: windows linux darwin ## Build binaries

windows: $(WINDOWS) ## Build for Windows

linux: $(LINUX) ## Build for Linux

darwin: $(DARWIN) ## Build for Darwin (macOS)

$(WINDOWS):
	env GOOS=windows GOARCH=amd64 go build -ldflags '-X $(MODULE)/cmd.lastGitTag=$(LASTTAG) -X $(MODULE)/cmd.lastGitCommit=$(REVISION)' -o $(WINDOWS) main.go

$(LINUX):
	env GOOS=linux GOARCH=amd64 go build -ldflags '-X $(MODULE)/cmd.lastGitTag=$(LASTTAG) -X $(MODULE)/cmd.lastGitCommit=$(REVISION)' -o $(LINUX) main.go

$(DARWIN):
	env GOOS=darwin GOARCH=amd64 go build -ldflags '-X $(MODULE)/cmd.lastGitTag=$(LASTTAG) -X $(MODULE)/cmd.lastGitCommit=$(REVISION)' -o $(DARWIN) main.go

clean: ## Remove previous build
	rm -f $(WINDOWS) $(LINUX) $(DARWIN)
