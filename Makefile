.DEFAULT_GOAL := all

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

run-consumer:
	./tmp/consumer

run-sender:
	./tmp/sender

swagger:
	swag init -g cmd/sender/main.go --parseDependency --parseInternal

dev:
	make swagger
	air server

all: swagger dev
