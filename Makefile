GO_FLAGS   ?=
NAME       := k13s
OUTPUT_BIN ?= execs/${NAME}
PACKAGE    := github.com/kswapd/$(NAME)
GIT_REV    ?= $(shell git rev-parse --short HEAD)
SOURCE_DATE_EPOCH ?= $(shell gdate +%s)
DATE       ?= $(shell gdate -u -d @${SOURCE_DATE_EPOCH} +"%Y-%m-%dT%H:%M:%SZ")
VERSION    ?= v1.0.3
IMG_NAME   := kswapd/${NAME}
IMAGE      := ${IMG_NAME}:${VERSION}

default: help

test:   ## Run all tests
	@go clean --testcache && go test ./...

cover:  ## Run test coverage suite
	@go test ./... --coverprofile=cov.out
	@go tool cover --html=cov.out

build:  ## Builds the CLI
	@go build ${GO_FLAGS} \
	-ldflags "-w -s -X ${PACKAGE}/cmd.version=${VERSION} -X ${PACKAGE}/cmd.commit=${GIT_REV} -X ${PACKAGE}/cmd.date=${DATE}" \
	-a -tags netgo -o ${OUTPUT_BIN} main.go

kubectl-stable-version:  ## Get kubectl latest stable version
	@curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt

img:    ## Build Docker Image
	@docker build --rm -t ${IMAGE} .

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":[^:]*?## "}; {printf "\033[38;5;69m%-30s\033[38;5;38m %s\033[0m\n", $$1, $$2}'
