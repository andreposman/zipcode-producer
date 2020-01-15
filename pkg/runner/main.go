package main

import (
	"github.com/andreposman/zipcode-producer/pkg/amqp"
)

func main() {

	amqp.ReceiveMessage()
}
