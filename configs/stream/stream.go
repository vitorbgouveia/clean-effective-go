package stream

import (
	"fmt"

	"github.com/streadway/amqp"
	"github.com/vitorbgouveia/clean-effective-go/configs/env"
	"github.com/vitorbgouveia/clean-effective-go/internal/app/consumers"
	"github.com/vitorbgouveia/clean-effective-go/tools"
)

var connStream *amqp.Connection

// OptionsStream list options functions to Stream server
type Streamer interface {
	Start()
	connection()
}

// Stream struct
type stream struct{}

func Stream() Streamer {
	var s Streamer = stream{}
	return s
}

// Connection load connection stream
func (s stream) connection() {
	var c = env.Envs.RabbitMQ
	strConn := fmt.Sprintf("amqp://%s:%s@%s:%d/", c.User, c.Pasword, c.Host, c.Port)

	conn, err := amqp.Dial(strConn)
	tools.FatalError(err, "Failed to connect to RabbitMQ")

	connStream = conn
}

// Start stream services
func (s stream) Start() {
	var c consumers.Starter = consumers.Consumer{}

	s.connection()
	c.Start(connStream)
}
