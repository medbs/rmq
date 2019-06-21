package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

// The flag was used in the declaration because the connections start
// with amqp.Dial(), and it requires a command line argument or environment variable.
var (
	amqpURI = flag.String("amqp", "amqp://guest:guest@localhost:5672/", "AMQP URI")
)

var conn *amqp.Connection
var ch *amqp.Channel
var replies <-chan amqp.Delivery
var message string

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}
func initAmqp() {
	var err error
	var q amqp.Queue

	conn, err = amqp.Dial(*amqpURI)
	failOnError(err, "connection to RabbitMQ failed")

	log.Printf("got Connection, getting Channel...")

	ch, err = conn.Channel()
	failOnError(err, "Failed to open a channel")

	log.Printf("got Channel, declaring Exchange (%s)", "rmq-exchange")

	err = ch.ExchangeDeclare(
		"rmq-exchange", // name of the exchange
		"direct",       // type
		true,           // durable
		false,          // delete when complete
		false,          // internal
		false,          // noWait
		nil,            // arguments
	)
	failOnError(err, "Failed to declare the Exchange")

	log.Printf("declared Exchange, declaring Queue (%s)", "rmq-exchange")

	q, err = ch.QueueDeclare(
		"rmq.queue", // name, leave empty to generate a unique name
		true,        // durable
		false,       // delete when usused
		false,       // exclusive
		false,       // noWait
		nil,         // arguments
	)
	failOnError(err, "Error declaring the Queue")

	log.Printf("declared Queue (%q %d messages, %d consumers), binding to Exchange (key %q)",
		q.Name, q.Messages, q.Consumers, "rmq-routingkey")

	err = ch.QueueBind(
		q.Name,           // name of the queue
		"rmq-routingkey", // bindingKey
		"rmq-exchange",   // sourceExchange
		false,            // noWait
		nil,              // arguments
	)
	failOnError(err, "Error binding to the Queue")

	log.Printf("Queue bound to Exchange, starting Consume (consumer tag %q)", "go-amqp-example")

	replies, err = ch.Consume(
		q.Name,         // queue
		"rmq-consumer", // consumer
		true,           // auto-ack
		false,          // exclusive
		false,          // no-local
		false,          // no-wait
		nil,            // args
	)
	failOnError(err, "Error consuming the Queue")
}

func init() {
	flag.Parse()
	initAmqp()
}

func main() {
	log.Println("Start consuming the Queue...")
	var count int = 1

	for r := range replies {
		log.Printf("Consuming reply number %d", count)
		json.Unmarshal(r.Body, &message)
		fmt.Printf(message)
		count++
	}
}
