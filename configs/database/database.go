package database

import (
	"context"
	"fmt"
	"log"

	"github.com/vitorbgouveia/clean-effective-go/configs/env"
	"github.com/vitorbgouveia/clean-effective-go/tools"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var clientDB *mongo.Client
var ctxDB context.Context

// OptionsDatabase list options functions to databse
type Databaser interface {
	Connect()
	Client() (*mongo.Client, context.Context)
}

// Database struct
type database struct{}

func DB() Databaser {
	var d Databaser = database{}
	return d
}

// Connect load connection with database and create context database
func (d database) Connect() {
	var c = env.Envs.Database
	str := fmt.Sprintf("mongodb://%s:%s@%s:%d", c.User, c.Password, c.Host, c.Port)

	ctxDB = context.Background()
	client, err := mongo.Connect(ctxDB, options.Client().ApplyURI(str))
	tools.FatalError(err, "Error connecting database")
	log.Println("Database connected")

	clientDB = client
}

// ClientDB return client database and context database
func (d database) Client() (*mongo.Client, context.Context) {
	return clientDB, ctxDB
}
