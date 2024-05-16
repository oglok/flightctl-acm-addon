#IMAGE_NAME = quay.io/flightctl/flightctl-acm-addon
IMAGE_NAME=localhost:5001/flightctl/flightctl-acm-addon

GOBASE=$(shell pwd)
GOBIN=$(GOBASE)/bin
GO_BUILD_FLAGS := ${GO_BUILD_FLAGS}

.PHONY: build

bin:
	rm -rf bin/

	go build -buildvcs=false $(GO_BUILD_FLAGS) -o $(GOBIN)/flightctl-acm-addon main.go

build:
	rm -rf bin/
	podman build -t $(IMAGE_NAME):latest -f Containerfile .
	podman push $(IMAGE_NAME):latest
