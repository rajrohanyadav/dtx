PROJECT_NAME=dtx
BIN="bin"
SRC=$(shell find . -name "*.go")
TARGET?=github.com/rajrohanyadav/dtx

BINARY ?= ${PROJECT_NAME}

VERSION = 1.0.0
COMMIT = $(shell git rev-parse --short=7 HEAD)
BRANCH = $(shell git rev-parse --abbrev-ref HEAD)

LDFLAGS = -ldflags "-X main.VERSION=${VERSION} -X main.commit=${COMMIT} -X main.BRANCH=${BRANCH} -X main.BUILD_NUMBER=${BUILD_NUMBER}"

ifeq (, $(shell which golangci-lint))
$(warning "could not find golangci-lint in $(PATH), run: curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh")
endif

.PHONY: all build fmt vet lint test install_deps clean

default: all

all: vet fmt test build

build: clean windows darwin linux

linux: install_deps
	$(info ******************** building for linux amd64 ********************)
	@GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o $(BIN)/${BINARY}-linux-amd64 ${TARGET}; 
	$(info ******************** building for linux arm64 ********************)
	@GOOS=linux GOARCH=arm64 go build ${LDFLAGS} -o $(BIN)/${BINARY}-linux-arm64 ${TARGET}; 

darwin: install_deps
	$(info ******************** building for darwin amd64 ********************)
	@GOOS=darwin GOARCH=amd64 go build ${LDFLAGS} -o $(BIN)/${BINARY}-darwin-amd64 ${TARGET}; 
	$(info ******************** building for darwin arm64 ********************)
	@GOOS=darwin GOARCH=arm64 go build ${LDFLAGS} -o $(BIN)/${BINARY}-darwin-arm64 ${TARGET}; 

windows: install_deps
	$(info ******************** building for windows ********************)
	@GOOS=windows GOARCH=${GOARCH} go build ${LDFLAGS} -o $(BIN)/${BINARY}-windows.exe ${TARGET}; 

fmt:
	$(info ******************** checking formatting ********************)
	go fmt $$(go list ./... | grep -v /vendor/);
	goimports -local github.com/golangci/golangci-lint -w $$(find . -type f -iname \*.go)

lint:
	$(info ******************** running lint tools ********************)
	golangci-lint run -v

vet:
	$(info ******************** vetting ********************)
	go vet $$(go list ./... | grep -v /vendor/)

test: install_deps
	$(info ******************** running tests ********************)
	go clean -testcache
	go test -v ./...

install_deps:
	$(info ******************** downloading dependencies ********************)
	go get -v ./...

clean:
	rm -rf $(BIN)
