package main

import (
	"fmt"
	"lunit-test/calc"
)

func main() {
	num0, num1 := 1.0, 2.0

	fmt.Printf("%.2f+%.2f=%.2f\n", num0, num1, calc.Add(num0, num1))
	fmt.Printf("%.2f-%.2f=%.2f\n", num0, num1, calc.Subtract(num0, num1))
	val, err := calc.Divide(num0, num1)
	if err != nil {
	}
	fmt.Printf("%.2f/%.2f=%.2f\n", num0, num1, val)
	fmt.Printf("%.2f*%.2f=%.2f\n", num0, num1, calc.Multiply(num0, num1))
}
