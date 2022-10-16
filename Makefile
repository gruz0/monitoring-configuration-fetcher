export ROOT=$(realpath $(dir $(lastword $(MAKEFILE_LIST))))
export DEBUG=true
export APP=monitoring-configuration-fetcher
export LDFLAGS="-w -s"
export MONITORING_CONFIGURATION_SERVICE_URL=http://localhost:8080
export OUTPUT_DIR=./public

all: build test

build:
	go build -race -o $(APP) .

build-static:
	CGO_ENABLED=0 go build -race -v -o $(APP) -a -installsuffix cgo -ldflags $(LDFLAGS) .

run:
	go run -race .

test:
	go test -v -race ./...

docker-build:
	docker build --rm -t gruz0/monitoring-configuration-fetcher .

docker-run:
	docker run --rm -it -e "MONITORING_CONFIGURATION_SERVICE_URL=$(MONITORING_CONFIGURATION_SERVICE_URL)" -e "OUTPUT_DIR=$(OUTPUT_DIR)" gruz0/monitoring-configuration-fetcher

.PHONY: build run build-static test build-container run-container
