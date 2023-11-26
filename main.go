package main

import (
	consumer "go_rabbitmq/consumer"
	publisher "go_rabbitmq/publisher"
	"log"

	"github.com/streadway/amqp"
)

func main() {
	// Just ping
	consumer.Ping()
	publisher.Ping()

	// Connect to RabbitMQ server
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// Create a channel
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// Declare a queue
	q, err := ch.QueueDeclare(
		"hello", // Queue name
		false,   // Durable
		false,   // Delete when unused
		false,   // Exclusive
		false,   // No-wait
		nil,     // Arguments
	)
	failOnError(err, "Failed to declare a queue")

	publisher.Publish(ch, q)
	consumer.Consume(ch, q)
	log.Print("End of Program.")
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
