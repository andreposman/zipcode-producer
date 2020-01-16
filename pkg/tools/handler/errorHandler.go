package tools

import "log"

//FailOnError is a simple error handling function
func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
