package main

import (
	"fmt"
	"lovoo/calculator/api"
	"lovoo/calculator/service"
	"net"

	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		panic(err)
	}

	calculatorService := service.NewCalculatorService()
	calculatorServer := api.NewCalculatorServer(calculatorService)

	server := grpc.NewServer()
	api.RegisterCalculatorServiceServer(server, calculatorServer)

	fmt.Println("Server is starting...")
	if err := server.Serve(listener); err != nil {
		panic(err)
	}
}
