package controllers

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/andreposman/zipcode-producer/pkg/services/rabbitmq"
)

// OpenFile as the name suggets, open and returns an csv file.
func openFile() *os.File {
	file := "../data/zipcodes.csv"

	csvData, err := os.Open(file)
	if err != nil {
		log.Fatalln(" - ❌ Error:", err)
	}

	return csvData
}

// ReadFile reads the csv file and returns an array of strings
func ReadFile() []string {

	var ZipCodes []string
	csvData := openFile()
	reader := csv.NewReader(bufio.NewReader(csvData))

	lines, err := reader.Read()

	if err != nil {
		fmt.Println(err)
	}

	for _, line := range lines {
		ZipCodes = append(ZipCodes, line)
	}

	return ZipCodes
}

// ProduceMessage receives an array of strings and publish on a rabbitmq queue
func ProduceMessage() {
	zipcodes := ReadFile()
	rabbitmq.PublishMessage(zipcodes)
}
