package main

import (
	"log"

	"github.com/vkunssec/rabbit-mq-simple/pkg/repository"
)

func main() {
	err := repository.ReceiveMessageFromQueue2()
	if err != nil {
		log.Printf("Error receiving message from queue 2: %v", err)
	}
}
