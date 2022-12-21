package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/maurodesouza/dead-letter-exchange-go-example/src"
	"github.com/streadway/amqp"
)

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func main() {
	messageChannel := make(chan amqp.Delivery)

	rabbitMQ := src.NewRabbitMQ("json-parser", "app-parser", "json-parser-dlx")
	ch := rabbitMQ.Connect()

	defer ch.Close()

	rabbitMQ.Consume(messageChannel)

	consumer := src.NewConsumer(messageChannel, rabbitMQ)
	consumer.Start()
}
