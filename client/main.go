package main

import (
	"context"
	"fmt"
	"lovoo/calculator/api"

	"google.golang.org/grpc"
)

func main() {
	serverAddress := "localhost:8000"

	connection, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := api.NewCalculatorServiceClient(connection)

	ctx := context.Background()

	//resp, err := client.Addition(ctx, &api.AdditionCalculationRequest{})
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(resp)

	resp, err := client.Division(ctx, &api.DivisionCalculationRequest{})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
