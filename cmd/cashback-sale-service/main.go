package main

import (
	"github.com/vitorbgouveia/clean-effective-go/configs/database"
	"github.com/vitorbgouveia/clean-effective-go/configs/stream"
	"github.com/vitorbgouveia/clean-effective-go/internal/app/server"
)

func main() {
	database.DB().Connect()
	server.AppServer().Start()
	stream.Stream().Start()
}
