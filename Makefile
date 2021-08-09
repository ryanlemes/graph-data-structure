# Global vars
PROJECT_NAME := $(shell basename "$(PWD)")

## fmt: Run formatter
fmt:
	go fmt ./...

## build: Build Docker image
docker-build:
	docker build -t ${PROJECT_NAME} .

## run: Run Docker image
docker-run:
	docker run -it ${PROJECT_NAME}

## build: Build server binary
build:
	go build -o ${PROJECT_NAME} ./cmd/cli

## run: Run server
run:
	go run ./cmd/cli