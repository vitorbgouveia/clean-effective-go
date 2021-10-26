package services

import "google.golang.org/grpc"

type StarterRPC interface {
	StartRPC(s *grpc.Server)
}

type rpc struct{}

func RPC() StarterRPC {
	var s StarterRPC = rpc{}
	return s
}

// StartRPC start rpc Services
func (r rpc) StartRPC(server *grpc.Server) {
	startSaleService(server)
}
