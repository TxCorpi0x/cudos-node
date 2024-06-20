PACKAGES=$(shell go list ./... | grep -v '/simulation')

VERSION := $(shell echo $(shell git describe --tags))
COMMIT := $(shell git log -1 --format='%H')

GO_VERSION := $(shell cat go.mod | grep -E 'go [0-9].[0-9]+' | cut -d ' ' -f 2)
DOCKER := $(shell which docker)

BUILDDIR ?= $(CURDIR)/build

ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=cudos-node \
	-X github.com/cosmos/cosmos-sdk/version.AppName=cudos-noded \
	-X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
	-X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT)

BUILD_FLAGS := -ldflags '$(ldflags)'

all: install

install: export CGO_LDFLAGS=-Wl,-rpath,$$ORIGIN/../
install: go.sum
		@echo "--> Installing cudos-noded"
		@go install -mod=readonly $(BUILD_FLAGS) -tags "ledger" ./cmd/cudos-noded


build: export CGO_LDFLAGS=-Wl,-rpath,$$ORIGIN/../
build: go.sum
		@echo "--> Building cudos-noded"
		@go build -mod=readonly $(BUILD_FLAGS) -o $(BUILDDIR)/ -tags "ledger" ./cmd/cudos-noded

build-linux-amd64: go.sum
	mkdir -p $(BUILDDIR)
	$(DOCKER) buildx create --name cudobuilder || true
	$(DOCKER) buildx use cudobuilder
	$(DOCKER) buildx build \
		--build-arg GO_VERSION=$(GO_VERSION) \
		--build-arg GIT_VERSION=$(VERSION) \
		--build-arg GIT_COMMIT=$(COMMIT) \
		--build-arg RUNNER_IMAGE=alpine:3.18 \
		--platform linux/amd64 \
		-t cudo:local-amd64 \
		--load \
		-f Dockerfile .
	$(DOCKER) rm -f cudobinary || true
	$(DOCKER) create -ti --name cudobinary cudo:local-amd64
	$(DOCKER) cp cudobinary:/bin/cudos-noded $(BUILDDIR)/cudo-linux-amd64
	$(DOCKER) rm -f cudobinary

go.sum: go.mod
		@echo "--> Ensure dependencies have not been modified"
		GO111MODULE=on go mod verify

test:
	@go test -v -mod=readonly $(PACKAGES)
