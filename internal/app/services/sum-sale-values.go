package services

import (
	"context"

	"github.com/vitorbgouveia/clean-effective-go/api/protocols/sale/salepb"
	"github.com/vitorbgouveia/clean-effective-go/internal/presentation/dtos"
	"github.com/vitorbgouveia/clean-effective-go/internal/usecases"
	"google.golang.org/grpc"
)

type server struct {
	salepb.UnimplementedSaleServiceServer
}

func (*server) SumSaleValues(ctx context.Context, req *salepb.SumSaleValuesRequest) (*salepb.SumSaleValuesResponse, error) {
	cpf := req.GetCpf()

	p := dtos.ParamsGetSumSaleValues{
		Cpf: cpf,
	}
	s := usecases.SumSaleValues(p)
	res := &salepb.SumSaleValuesResponse{
		Cpf:           cpf,
		TotalValue:    s.TotalValue,
		TotalCashback: s.TotalCashback,
	}

	return res, nil
}

func startSaleService(s *grpc.Server) {
	salepb.RegisterSaleServiceServer(s, &server{})
}
