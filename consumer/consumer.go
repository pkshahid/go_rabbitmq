package consumer

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func Ping() {
	fmt.Println("Pong From Consumer")
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func Consume(ch *amqp.Channel, q amqp.Queue) {
	// Consume messages from the queue
	msgs, err := ch.Consume(
		q.Name, // Queue name
		"",     // Consumer
		true,   // Auto acknowledge
		false,  // Exclusive
		false,  // No-local
		false,  // No-wait
		nil,    // Arguments
	)
	failOnError(err, "Failed to Consume a message")

	for d := range msgs {
		fmt.Printf("Received a message: %s\n", d.Body)
	}

}
