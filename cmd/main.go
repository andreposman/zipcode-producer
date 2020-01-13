package main

import (
	"github.com/andreposman/zipcode-producer/pkg/amqp"
	"github.com/andreposman/zipcode-producer/pkg/reader"
)

func main() {

	zipCodes := reader.ReadFile()
	amqp.SendMessage(zipCodes)

}
