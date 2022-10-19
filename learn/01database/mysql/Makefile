GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
BINARY_NAME=cmd
DOCKER=docker
DOCKERBUILD=$(DOCKER) build
PROJECT_PATH=$(shell pwd)

all: docker
	@echo "all"
build: clean
	@echo "build clean"
	cd $(BINARY_NAME) && $(GOBUILD) -o service
build-proto:
	@echo "build-proto"
docker: build
	@echo "docker build"
	tar -cvf - ./cmd/service configset | $(DOCKERBUILD) -t $(BINARY_NAME) -f $(PROJECT_PATH)/Dockerfile .
clean:
	@echo "clean"
run:
	@echo "run"
local:
	@echo "local"
