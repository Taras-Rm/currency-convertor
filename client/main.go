package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Taras-Rm/currency-convertor/grpc_api/api"
	"google.golang.org/grpc"
)

func main()  {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}

	client := api.NewCurrencyConvertorClient(conn)

	resp, err := client.ExchangeRate(context.Background(), &api.ExchangeRateRequest{

	})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp)
}