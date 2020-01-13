package reader

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

// OpenFile as the name suggets, open and returns an csv file.
func openFile() *os.File {
	file := "../data/zipcodes.csv"

	csvData, err := os.Open(file)
	if err != nil {
		log.Fatalln(" - ❌ Error:", err)
	}
	// defer csvData.Close()

	return csvData
}

// ReadFile reads the csv file and returns an array of strings
func ReadFile() []string {

	// ZipCodes := new(models.ZipCode)
	// var ZipCode = &ZipCodes
	var ZipCode []string
	csvData := openFile()
	reader := csv.NewReader(bufio.NewReader(csvData))

	lines, err := reader.Read()

	if err != nil {
		fmt.Println(err)
	}

	for _, line := range lines {
		ZipCode = append(ZipCode, line)
	}

	return ZipCode
}
