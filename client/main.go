package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"lovoo/calculator/api"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

var (
	calculation string
	n1, n2      int
)

func main() {
	fmt.Printf("%v \n", os.Args[1])
	if err := run(os.Args, os.Stdout); err != nil {
		_, err := fmt.Fprintf(os.Stderr, "%s\n", err)
		if err != nil {
			panic(err)
		}
		os.Exit(1)
	}
}

func run(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet(args[0], flag.ExitOnError)
	flags.StringVar(&calculation, "method", "all", "-method add")
	flags.IntVar(&n1, "a", 1, "-a 0")
	flags.IntVar(&n2, "b", 1, "-b 0")

	fmt.Println(flags)
	fmt.Println(calculation)
	fmt.Println(n1)
	fmt.Println(n2)

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	serverHost := os.Getenv("SERVER_HOST")
	if serverHost == "" {
		panic("Server Host env is empty!")
	}

	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		panic("Server Port env is empty!")
	}

	serverAddress := fmt.Sprintf("%s:%s", serverHost, serverPort)

	connection, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := api.NewCalculatorServiceClient(connection)

	ctx := context.Background()

	switch calculation {
	case "add":
		resp, err := client.Addition(ctx, &api.AdditionCalculationRequest{
			FirstNumber:  float32(n1),
			SecondNumber: float32(n2),
		})
		if err != nil {
			panic(err)
		}
		fmt.Println(resp)
	case "sub":
		resp, err := client.Subtraction(ctx, &api.SubtractionCalculationRequest{
			FirstNumber:  float32(n1),
			SecondNumber: float32(n2),
		})
		if err != nil {
			panic(err)
		}
		fmt.Println(resp)
	case "div":
		resp, err := client.Division(ctx, &api.DivisionCalculationRequest{
			FirstNumber:  float32(n1),
			SecondNumber: float32(n2),
		})
		if err != nil {
			panic(err)
		}
		fmt.Println(resp)
	case "multi":
		resp, err := client.Multiplication(ctx, &api.MultiplicationCalculationRequest{
			FirstNumber:  float32(n1),
			SecondNumber: float32(n2),
		})
		if err != nil {
			panic(err)
		}
		fmt.Println(resp)
	default:
		resp, err := client.AllCalculations(ctx, &api.AllCalculationRequest{
			FirstNumber:  float32(n1),
			SecondNumber: float32(n2),
		})
		if err != nil {
			panic(err)
		}
		fmt.Println(resp)
	}
	return nil
}
