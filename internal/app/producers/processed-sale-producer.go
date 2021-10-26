package producers

import (
	"fmt"

	"github.com/streadway/amqp"
	"github.com/vitorbgouveia/clean-effective-go/internal/presentation/dtos"
	"github.com/vitorbgouveia/clean-effective-go/tools"
)

// ProcessedSaleProducer publish message to quee
func ProcessedSaleProducer(conn *amqp.Connection, sale dtos.OutputSale) {
	ch, err := conn.Channel()
	tools.FatalError(err, "Failed to open a channel")
	defer ch.Close()

	const queueName = "processed-sale"
	q, err := ch.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when unused
		true,      // exclusive
		false,     // no-wait
		nil,       // arguments
	)

	strSale := fmt.Sprint(sale)
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(strSale),
		})
	tools.FatalError(err, "Failed to publish a message")
}
