package reader

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
)

// OpenFile as the name suggets, open and returns an csv file.
func OpenFile() *os.File {

	csvData, err := os.Open("../../data/zipcodes.csv")

	if err != nil {
		fmt.Println(err)
	}

	return csvData
}

func ReadFile() []string {

	csvData := OpenFile()
	reader := csv.NewReader(bufio.NewReader(csvData))

	ZipCode := new ZipCode
}
