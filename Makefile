sender:
	go run cmd/sender/main.go

consumer:
	go run cmd/consumer/main.go

env:
	cp .env.example .env

build:
	go build -ldflags="-s -w" -o tmp/sender cmd/sender/main.go
	go build -ldflags="-s -w" -o tmp/consumer cmd/consumer/main.go

run-consumer:
	./tmp/consumer

run-sender:
	./tmp/sender
