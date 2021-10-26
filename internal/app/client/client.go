package main

import (
	"context"
	"fmt"
	"log"

	"github.com/vitorbgouveia/clean-effective-go/api/protocols/sale/salepb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	c := salepb.NewSaleServiceClient(conn)
	doUnary(c)
}

func doUnary(c salepb.SaleServiceClient) {
	fmt.Println("Starting to do a unary RPC...")
	req := &salepb.SumSaleValuesRequest{
		Cpf: "10640463479",
	}
	res, err := c.SumSaleValues(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Greet RPC: %v", err)
	}
	log.Printf("Response from Greet: %v", res)
}
