package amqp

import (
	"encoding/json"
	"log"

	"github.com/streadway/amqp"
)

// failOnError helper function to check the return value for each amqp call:
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s:", msg, err)
	}
}

// SendMessage sends the message to a queue
func SendMessage(zipCodes []string) {

	connString := "amqp://guest:guest@localhost:5672/"

	conn, err := amqp.Dial(connString)
	failOnError(err, "Failed to connect do RabbitMQ")
	defer conn.Close()

	channel, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer channel.Close()

	queue, err := channel.QueueDeclare(
		"zipcodes", // name
		false,      // durable
		false,      // delete when unused
		false,      // exclusive
		false,      // no-wait
		nil,        // arguments
	)
	failOnError(err, "Faile do declare a queue")

	message, _ := json.Marshal(zipCodes)

	err = channel.Publish(
		"",         // exchange
		queue.Name, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	log.Printf(" [x] Sent: %s", message)
	failOnError(err, "Failed to publish a message")
}
