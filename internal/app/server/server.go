package server

import (
	"fmt"
	"log"
	"net"

	"github.com/vitorbgouveia/clean-effective-go/configs/env"
	"github.com/vitorbgouveia/clean-effective-go/internal/app/services"
	"github.com/vitorbgouveia/clean-effective-go/tools"
	"google.golang.org/grpc"
)

// OptionsServer list functions to server
type Server interface {
	Start()
}

// Server struct
type appServer struct{}

func AppServer() Server {
	var s Server = appServer{}
	return s
}

// Start server application
func (sv appServer) Start() {
	port := fmt.Sprint(env.Envs.App.Port)

	go func() {
		lis, err := net.Listen("tcp", ":"+port)
		tools.FatalError(err, "Failed to listen: %v")

		s := grpc.NewServer()
		services.RPC().StartRPC(s)

		log.Println("Server started on port: " + port)
		err = s.Serve(lis)
		tools.FatalError(err, "Failed to serve: %v")
	}()
}
