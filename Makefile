MAKEFLAGS += -j2
.DEFAULT_GOAL := run

sender:
	go run cmd/sender/main.go

consumer:
	go run cmd/consumer/main.go

env:
	cp .env.example .env

build:
	make swagger
	go build -ldflags="-s -w" -o tmp/sender cmd/sender/main.go
	go build -ldflags="-s -w" -o tmp/consumer cmd/consumer/main.go

target: run-consumer run-sender

run:
	make build
	make target

run-consumer:
	./tmp/consumer

run-sender:
	./tmp/sender

swagger:
	swag init -g cmd/sender/main.go --parseDependency --parseInternal

dev:
	go get -u github.com/automation-co/husky
	go get -u github.com/golangci/golangci-lint/cmd/golangci-lint
	go get -u github.com/swaggo/swag/cmd/swag
	make swagger
	air server

all: swagger dev
