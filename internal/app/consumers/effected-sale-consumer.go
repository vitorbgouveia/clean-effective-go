package consumers

import (
	"encoding/json"

	"github.com/streadway/amqp"
	"github.com/vitorbgouveia/clean-effective-go/internal/app/producers"
	"github.com/vitorbgouveia/clean-effective-go/internal/presentation/dtos"
	"github.com/vitorbgouveia/clean-effective-go/internal/usecases"
	"github.com/vitorbgouveia/clean-effective-go/tools"
)

func handlerEffectedSaleConsumer(conn *amqp.Connection, chanConsumer chan bool) {
	ch, err := conn.Channel()
	tools.FatalError(err, "Failed to open a channel")
	defer ch.Close()

	const queueName = "effected-sale"
	q, err := ch.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when unused
		true,      // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	tools.FatalError(err, "Failed to declare a queue "+queueName)

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	tools.FatalError(err, "Failed to register a consumer")

	go func() {
		for d := range msgs {
			s := new(dtos.EntranceSale)
			err := json.Unmarshal([]byte(d.Body), s)
			tools.DomainError(err, "Failed to read body")

			if err == nil {
				processedSale := usecases.ProcessSale(*s)
				producers.ProcessedSaleProducer(conn, processedSale)
			}
		}
	}()
}
