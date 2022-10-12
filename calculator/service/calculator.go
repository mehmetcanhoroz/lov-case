package service

import "fmt"

type CalculatorService interface {
	Add(float32, float32) (float32, error)
	Sub(float32, float32) (float32, error)
	Div(float32, float32) (float32, error)
	Multi(float32, float32) (float32, error)
}
type calculatorService struct {
}

func (s calculatorService) Add(n1, n2 float32) (float32, error) {
	return n1 + n2, nil
}
func (s calculatorService) Sub(n1, n2 float32) (float32, error) {
	return n1 - n2, nil
}
func (s calculatorService) Div(n1, n2 float32) (float32, error) {
	if n2 == 0 {
		return 0, fmt.Errorf("second number cannot be zero for division operation")
	}

	return n1 / n2, nil
}
func (s calculatorService) Multi(n1, n2 float32) (float32, error) {
	return n1 * n2, nil
}

func NewCalculatorService() CalculatorService {
	return calculatorService{}
}
