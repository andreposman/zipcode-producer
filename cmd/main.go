package main

import (
	"github.com/andreposman/zipcode-producer/pkg/controllers"
	"github.com/andreposman/zipcode-producer/pkg/services/rabbitmq"
)

func main() {

	zipCodes := controllers.ReadFile()
	rabbitmq.ReceiveMessage()
	// rabbitmq.SendMessage(zipCodes)

}
