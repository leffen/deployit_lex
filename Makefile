# Author: Leif Terje Fonnes, desember 2019
#!make

include .env

APP=deployit_lex
GOOS=linux
GOARCH=amd64
GOBUILDFLAGS ?= CGO_ENABLED=$(CGO_ENABLED) GOOS=$(GOOS) GOARCH=$(GOARCH)

BuildTime=$(shell date +%Y-%m-%d.%H:%M:%S)

# Check whether the git repository is dirty or not.
GitDirty=$(shell git status --porcelain --untracked-files=no)
GitCommit=$(shell git rev-parse --short HEAD)
ifneq ($(GitDirty),)
	GitCommit:= $(GitCommit)-dirty
endif

SOURCES=$(APP).go ops.go

all: build
.PHONY: build

bump:
	bump_version patch deployit_lex.go

clean:
	rm -rf bin/*

compile: test clean
	$(GOBUILDFLAGS) go build -ldflags "-s -X main.CommitHash=$(GitCommit) -X main.BuildTime=$(BuildTime)"  -o bin/$(APP) .

release: test bump compile
	serverless deploy

test: vet
	@# this target should always be listed first so "make" runs the tests.
	go test -cover -race -v

vet: $(MEGACHECK)
	@# We can't vet the vendor directory, it fails.
	go list ./... | grep -v vendor | xargs go vet


