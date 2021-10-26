package consumers

import (
	"log"

	"github.com/streadway/amqp"
)

type Starter interface {
	Start(conn *amqp.Connection)
}

type Consumer struct {}

// Start consumers
func (c Consumer) Start(conn *amqp.Connection) {
	chanConsumers := make(chan bool)

	go handlerEffectedSaleConsumer(conn, chanConsumers)

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-chanConsumers
	log.Printf("Stoped wait for messages")
}
