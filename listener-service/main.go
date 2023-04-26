package main

import (
	"fmt"
	"listener/event"
	"log"
	"math"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {

	rabbitConn, err := connect()
	if err != nil {
		log.Fatal(err)
	}

	defer rabbitConn.Close()

	log.Println("Listening for consuming RabbitMQ messages...")

	consumer, err := event.NewConsumer(rabbitConn)
	if err != nil {
		log.Fatal(err)
	}

	err = consumer.Listen([]string{"log.INFO", "log.WARNING", "log.ERROR"})
	if err != nil {
		log.Fatal(err)
	}

}

func connect() (*amqp.Connection, error) {

	var counts int64
	var backOff = 1 * time.Second
	var connection *amqp.Connection

	for {
		c, err := amqp.Dial("amqp://rabbitmq:password@rabbitmq:5672/")

		if err != nil {

			fmt.Println("RabbitMQ not yet ready...")
			counts++

		} else {

			log.Println("Connected to RabbitMQ")

			connection = c
			break
		}

		if counts > 5 {
			fmt.Println(err)
			return nil, err
		}

		backOff = time.Duration(math.Pow(float64(counts), 2)) * time.Second
		log.Println("backing off for", backOff)
		time.Sleep(backOff)

		continue
	}

	return connection, nil
}
