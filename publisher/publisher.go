package publisher

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func Ping() {
	fmt.Println("Pong From Publisher!")
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func Publish(ch *amqp.Channel, q amqp.Queue) {

	for i := 0; i < 10; i++ {
		// Message to be sent
		body := fmt.Sprintf("Hello, Message From RabbitMQ! ::: %v", i)

		// Publish the message to the queue
		err := ch.Publish(
			"",     // Exchange
			q.Name, // Routing key
			false,  // Mandatory
			false,  // Immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			})
		failOnError(err, "Failed to publish a message")

	}

}
