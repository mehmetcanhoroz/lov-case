package api

import (
	"context"
	"lovoo/calculator/service"
)

type CalculatorServer struct {
	UnimplementedCalculatorServiceServer
	calculatorService service.CalculatorService
}

func (s CalculatorServer) Addition(ctx context.Context, in *AdditionCalculationRequest) (*CalculationResponse, error) {
	result, err := s.calculatorService.Add(in.FirstNumber, in.SecondNumber)
	if err != nil {
		return nil, err
	}
	response := CalculationResponse{
		Result: result,
	}
	return &response, nil
}
func (s CalculatorServer) Subtraction(ctx context.Context, in *SubtractionCalculationRequest) (*CalculationResponse, error) {
	result, err := s.calculatorService.Sub(in.FirstNumber, in.SecondNumber)
	if err != nil {
		return nil, err
	}
	response := CalculationResponse{
		Result: result,
	}
	return &response, nil
}
func (s CalculatorServer) Division(ctx context.Context, in *DivisionCalculationRequest) (*CalculationResponse, error) {
	result, err := s.calculatorService.Div(in.FirstNumber, in.SecondNumber)
	if err != nil {
		return nil, err
	}
	response := CalculationResponse{
		Result: result,
	}
	return &response, nil
}
func (s CalculatorServer) Multiplication(ctx context.Context, in *MultiplicationCalculationRequest) (*CalculationResponse, error) {
	result, err := s.calculatorService.Multi(in.FirstNumber, in.SecondNumber)
	if err != nil {
		return nil, err
	}
	response := CalculationResponse{
		Result: result,
	}
	return &response, nil
}
func (s CalculatorServer) AllCalculations(ctx context.Context, in *AllCalculationRequest) (*AllCalculationsResponse, error) {
	return nil, nil
}
func NewCalculatorServer(calculatorService service.CalculatorService) CalculatorServer {
	return CalculatorServer{
		calculatorService: calculatorService,
	}
}
