package repository

type CalculatorRepository interface {
	Add(float32, float32) (float32, error)
	Sub(float32, float32) (float32, error)
	Div(float32, float32) (float32, error)
	Multi(float32, float32) (float32, error)
}
type calculatorRepository struct {
}

func (s calculatorRepository) Add(n1, n2 float32) (float32, error) {
	return n1 + n2, nil
}
func (s calculatorRepository) Sub(n1, n2 float32) (float32, error) {
	return n1 - n2, nil
}
func (s calculatorRepository) Div(n1, n2 float32) (float32, error) {
	return n1 / n2, nil
}
func (s calculatorRepository) Multi(n1, n2 float32) (float32, error) {
	return n1 * n2, nil
}

func NewCalculatorRepository() CalculatorRepository {
	return calculatorRepository{}
}
