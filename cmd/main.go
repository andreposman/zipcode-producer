package main

import (
	"github.com/andreposman/zipcode-producer/database"
)

func main() {
	// rabbitmq.ReceiveMessage()

	database.Connect()
}
