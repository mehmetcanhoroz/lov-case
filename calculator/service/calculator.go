package service

import (
	"fmt"
	"lovoo/calculator/repository"
)

type CalculatorService interface {
	Add(float32, float32) (float32, error)
	Sub(float32, float32) (float32, error)
	Div(float32, float32) (float32, error)
	Multi(float32, float32) (float32, error)
}
type calculatorService struct {
	calculatorRepository repository.CalculatorRepository
}

func (s calculatorService) Add(n1, n2 float32) (float32, error) {
	return s.calculatorRepository.Add(n1, n2)
}
func (s calculatorService) Sub(n1, n2 float32) (float32, error) {
	return s.calculatorRepository.Sub(n1, n2)
}
func (s calculatorService) Div(n1, n2 float32) (float32, error) {
	if n2 == 0 {
		return 0, fmt.Errorf("second number cannot be zero for division operation")
	}

	return s.calculatorRepository.Div(n1, n2)
}
func (s calculatorService) Multi(n1, n2 float32) (float32, error) {
	return s.calculatorRepository.Multi(n1, n2)
}

func NewCalculatorService(calculatorRepository repository.CalculatorRepository) CalculatorService {
	return calculatorService{
		calculatorRepository: calculatorRepository,
	}
}
