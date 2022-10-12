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

	resp, err := client.Addition(ctx, &api.AdditionCalculationRequest{
		FirstNumber:  1,
		SecondNumber: 2,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)

	resp, err = client.Division(ctx, &api.DivisionCalculationRequest{
		FirstNumber:  6,
		SecondNumber: 3,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)

	resp, err = client.Multiplication(ctx, &api.MultiplicationCalculationRequest{
		FirstNumber:  5,
		SecondNumber: 2,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)

	resp, err = client.Subtraction(ctx, &api.SubtractionCalculationRequest{
		FirstNumber:  6,
		SecondNumber: 5,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)

	respx, err := client.AllCalculations(ctx, &api.AllCalculationRequest{
		FirstNumber:  10,
		SecondNumber: 2,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(respx)
}
