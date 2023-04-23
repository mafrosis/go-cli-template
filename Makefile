IMPORT:=github.com/mafrosis/go-cli-template
BIN:=go-cli-template
UNAME:=$(shell uname -s)

# Default to 64bit linux binaries
GOOS?=linux
GOARCH?=amd64

CGO_ENABLED=0

ifndef DATE
	DATE:=$(shell date -u '+%Y%m%d')
endif

ifndef SHA
	SHA:=$(shell git rev-parse --short HEAD)
endif

ifndef VERSION
	ifneq ($(RELEASE_TAG),)
		VERSION?=$(subst v,,$(RELEASE_TAG))
	else
		VERSION?=$(SHA)
	endif
endif

# Embed the version and git commit into the built binary
LDFLAGS+=-s -w -extldflags "-static" -X "$(IMPORT)/pkg/version.String=$(VERSION)" -X "$(IMPORT)/pkg/version.Revision=$(SHA)" -X "$(IMPORT)/pkg/version.Date=$(DATE)"

# Find all go source files
SOURCES?=$(shell find . -name "*.go" -type f)

.PHONY: all
all: build

.PHONY: clean
clean:
	go clean -i ./...
	rm -rf bin dist

.PHONY: build
build: bin/$(BIN)

$(addprefix bin/,$(BIN)): $(SOURCES)
	 go build -v -tags netgo -ldflags '$(LDFLAGS)' -o $@ main.go
