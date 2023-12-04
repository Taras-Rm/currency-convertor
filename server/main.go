package main

import (
	"context"
	"log"
	"net"

	"github.com/Taras-Rm/currency-convertor/grpc_api/api"
	"github.com/Taras-Rm/currency-convertor/server/currency"
	"google.golang.org/grpc"
)

type currencyConvertorServer struct {
	api.UnimplementedCurrencyConvertorServer
}

func (s *currencyConvertorServer) ExchangeRate(ctx context.Context, req *api.ExchangeRateRequest) (*api.ExchangeRateResponse, error) {
	resp := api.ExchangeRateResponse{}

	rate := currency.ExchangeCurrency(req.CurrFrom, req.CurrTo)

	resp.Rate = rate

	return &resp, nil
}

func main()  {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}

	grpcServer := grpc.NewServer()

	api.RegisterCurrencyConvertorServer(grpcServer, &currencyConvertorServer{})

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Println(err)
	}
}