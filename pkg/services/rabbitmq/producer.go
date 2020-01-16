package rabbitmq

import (
	"encoding/json"
	"log"

	tools "github.com/andreposman/zipcode-producer/pkg/tools/handler"
	"github.com/streadway/amqp"
)

// PublishMessage sends the message to a queue
func PublishMessage(zipCodes []string) {

	connString := "amqp://guest:guest@localhost:5672/"

	conn, err := amqp.Dial(connString)
	tools.FailOnError(err, "Failed to connect do RabbitMQ")
	defer conn.Close()

	channel, err := conn.Channel()
	tools.FailOnError(err, "Failed to open a channel")
	defer channel.Close()

	queue, err := channel.QueueDeclare(
		"zipcodes", // name
		false,      // durable
		false,      // delete when unused
		false,      // exclusive
		false,      // no-wait
		nil,        // arguments
	)
	tools.FailOnError(err, "Faile do declare a queue")

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
	tools.FailOnError(err, "Failed to publish a message")
}
