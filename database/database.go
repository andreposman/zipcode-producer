package database

import (
	"fmt"
	"os"
	"strings"
)

//SaveToFile receives the messages and saves to file
func SaveToFile(messages []string) {
	path := "consumedData.json"
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}

	file.WriteString(strings.Join(messages, "\n"))
}
