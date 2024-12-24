MAKEFLAGS += -j2
.DEFAULT_GOAL := run

sender:
	go run cmd/sender/main.go

consumer-1:
	go run cmd/consumer-1/main.go

consumer-2:
	go run cmd/consumer-2/main.go

consumers: consumer-1 consumer-2

env:
	cp .env.example .env

build:
	make swagger
	go build -ldflags="-s -w" -o tmp/sender cmd/sender/main.go
	go build -ldflags="-s -w" -o tmp/consumer-1 cmd/consumer-1/main.go
	go build -ldflags="-s -w" -o tmp/consumer-2 cmd/consumer-2/main.go

target: run-consumer-1 run-consumer-2 run-sender

run:
	make build
	make target

run-consumer-1:
	./tmp/consumer-1

run-consumer-2:
	./tmp/consumer-2

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
