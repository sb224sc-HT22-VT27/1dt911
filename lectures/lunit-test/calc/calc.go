package calc

import "errors"

// Capital first letter in function names specify public function

func Add(num0, num1 float64) float64 {
	return num0 + num1
}

func Subtract(num0, num1 float64) float64 {
	return num0 - num1
}

func Divide(num0, num1 float64) (float64, error) {
	if num1 == 0.0 {
		return 0.0, errors.New("division by zero")
	}
	return num0 / num1, nil
}

func Multiply(num0, num1 float64) float64 {
	return num0 * num1
}
