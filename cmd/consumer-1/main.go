package main

import (
	"log"

	"github.com/vkunssec/rabbit-mq-simple/pkg/repository"
)

func main() {
	err := repository.ReceiveMessageFromQueue1()
	if err != nil {
		log.Printf("Error receiving message from queue 1: %v", err)
	}
}
