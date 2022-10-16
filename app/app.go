package app

import (
	"fmt"
	"lovoo/calculator/api"
	"lovoo/calculator/repository"
	"lovoo/calculator/service"
	"net"

	"google.golang.org/grpc"
)

type Application struct {
}

func (a Application) Start() {
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		panic(err)
	}

	calculatorRepository := repository.NewCalculatorRepository()
	calculatorService := service.NewCalculatorService(calculatorRepository)
	calculatorServer := api.NewCalculatorServer(calculatorService)

	server := grpc.NewServer()
	api.RegisterCalculatorServiceServer(server, calculatorServer)

	fmt.Println("Server is starting...")
	if err := server.Serve(listener); err != nil {
		panic(err)
	}
}
