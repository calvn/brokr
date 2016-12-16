TEST?=$(shell go list ./... | grep -v /vendor/)

# Get git commit information
GIT_COMMIT=$(shell git rev-parse --short HEAD)
GIT_DIRTY=$(shell test -n "`git status --porcelain`" && echo "+CHANGES" || true)

default: test

test: generate
	@echo " ==> Running tests..."
	@go list $(TEST) \
		| grep -v "/vendor/" \
		| xargs -n1 go test -v -timeout=60s $(TESTARGS)
.PHONY: test

generate:
	@echo " ==> Generating..."
	@find . -type f -name '.DS_Store' -delete
	@go list ./... \
		| grep -v "/vendor/" \
		| xargs -n1 go generate $(PACKAGES)
.PHONY: generate


build: generate
	@echo " ==> Cleaning up old directory..."
	@rm -rf bin && mkdir -p bin
	@echo " ==> Building..."
	@go build -ldflags "-X github.com/calvn/brokr/cmd.GitCommit=${GIT_COMMIT}${GIT_DIRTY}" -o bin/brokr .
	@echo " ==> Installing..."
	@cp bin/brokr $(GOPATH)/bin
.PHONY: build

build-linux: create-build-image remove-dangling build-native
.PHONY: build-linux
