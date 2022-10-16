package app

import (
	"fmt"
	"lovoo/calculator/api"
	"lovoo/calculator/repository"
	"lovoo/calculator/service"
	"net"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

type Application struct {
}

func (a Application) Start() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		panic("Server Port env is empty!")
	}

	listener, err := net.Listen("tcp", ":"+serverPort)
	if err != nil {
		panic(err)
	}

	calculatorRepository := repository.NewCalculatorRepository()
	calculatorService := service.NewCalculatorService(calculatorRepository)
	calculatorServer := api.NewCalculatorServer(calculatorService)

	server := grpc.NewServer()
	api.RegisterCalculatorServiceServer(server, calculatorServer)

	fmt.Printf("Server is starting on port:%s\n", serverPort)
	if err := server.Serve(listener); err != nil {
		panic(err)
	}
}

func NewApplication() Application {
	return Application{}
}
