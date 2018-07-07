SHELL = /bin/bash

CURRENT_DIR = $(shell pwd)
#GOPATH := $(CURRENT_DIR)
VET_REPORT = $(GOPATH)/bin/vet.report
LINT_REPORT = $(GOPATH)/bin/lint.report
TEST_REPORT = $(GOPATH)/bin/tests.xml
GOARCH = amd64

GO ?= go

TAG ?= $(USER)
VERSION ?= 1.0.0
COMMIT = $(shell git rev-parse HEAD)
BRANCH = $(shell git rev-parse --abbrev-ref HEAD)

PKG := salat-timekeeper

VCS_REF = $(shell git rev-parse HEAD)
VCS_URL = $(shell git config --get remote.origin.url)

IMAGE ?= $(PKG)
TAG ?= latest
VERSION ?= $(0.0.0)


all: clean pretest test compile

test: compile
	$(GO) test ./... 

clean:
	-rm -f $(VET_REPORT)
	-rm -f $(GOPATH)/bin/*-amd64

pretest:
	dep ensure ; \
	$(GO) fmt ./... ; \
	$(GO) vet ./... 2>&1 | tee $(VET_REPORT) ; \
	gometalinter ./... --vendor 2>&1 | tee $(LINT_REPORT) ; \

compile: darwin linux windows

darwin:
	GOOS=darwin GOARCH=amd64 $(GO) build -o $(GOPATH)/bin/$(PKG)-$(VERSION)-darwin-amd64

linux:
	GOOS=linux GOARCH=amd64 $(GO) build -o $(GOPATH)/bin/$(PKG)-$(VERSION)-linux-amd64

windows:
	GOOS=windows GOARCH=amd64 $(GO) build -o $(GOPATH)/bin/$(PKG)-$(VERSION)-windows-amd64

.PHONY: pretest test clean linux darwin windows 
