package src

import (
	"github.com/streadway/amqp"
)

type Consumer struct {
	MessageChannel chan amqp.Delivery
	RabbitMQ       *RabbitMQ
}

func NewConsumer(messageChannel chan amqp.Delivery, rabbitMQ *RabbitMQ) *Consumer {
	return &Consumer{
		MessageChannel: messageChannel,
		RabbitMQ:       rabbitMQ,
	}
}

func (c *Consumer) Start() {
	for message := range c.MessageChannel {
		body := string(message.Body)

		if body == "valid" {
			message.Ack(false)

			c.RabbitMQ.Notify("Finalizou com sucesso!", "text/html", "json-parser-success", "")
		} else {
			message.Nack(false, false)
		}
	}
}
