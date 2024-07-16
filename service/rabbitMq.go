package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"

	amqp "github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		fmt.Printf("%s: %s", msg, err)
	}
}

func SendMq(queueName string, payload interface{}) error {
	amqpUrl := "amqp://guest:guest@localhost:5672/"
	if os.Getenv("AMQP_URL") != "" {
		amqpUrl = os.Getenv("AMQP_URL")
	}
	conn, err := amqp.Dial(amqpUrl)
	failOnError(err, "Failed to conn ect to RabbitMQ")
	if err != nil {
		return errors.New("Failed to conn ect to RabbitMQ")
	}
	defer conn.Close()
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	if err != nil {
		return errors.New("Failed to open a channel")
	}
	defer ch.Close()
	q, err := ch.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	failOnError(err, "Failed to declare a queue")
	if err != nil {
		return errors.New("Failed to declare a queue")
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	if err := ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        b,
		}); err != nil {
		failOnError(err, "Failed to publish a message")
		return errors.New("Failed to publish a message")
	}
	log.Printf(queueName)
	log.Printf(" [x] Sent %s\n", b)
	return nil
}
