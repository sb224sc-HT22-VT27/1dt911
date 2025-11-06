package calc_test

import (
	"lunit-test/calc"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testcases_devide = []struct {
	name        string
	expected    float64
	numerator   float64
	denominator float64
	wantErr     bool
}{
	{"division", 5.0, 10.0, 2.0, false},
	{"division w neg num", -5.0, -10.0, 2.0, false},
	{"division w ng dnum", -5.0, 10.0, -2.0, false},
	{"division by zero", 0.0, 5.0, 0.0, true},
}

func TestAdd(t *testing.T) {
	// Arrange
	num0, num1 := 5.0, 6.0
	expected := 11.0

	// Act
	got := calc.Add(num0, num1)

	// Assert
	if got != expected {
		t.Errorf("Expected %.2f got %.2f", expected, got)
	}
}

func TestSubtract(t *testing.T) {
	// Arrange
	num0, num1 := 5.0, 6.0
	expected := -1.0

	// Act
	got := calc.Subtract(num0, num1)

	// Assert
	if got != expected {
		t.Errorf("Expected %.2f got %.2f", expected, got)
	}
}

func TestDivide(t *testing.T) {
	// Arrange
	num0, num1 := 3.0, 6.0
	expected := 0.5

	// Act
	got, err := calc.Divide(num0, num1)

	if err != nil {
	}

	// Assert
	if got != expected {
		t.Errorf("Expected %.2f got %.2f", expected, got)
	}

	for _, tc := range testcases_devide {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)
			got, err := calc.Divide(tc.numerator, tc.denominator)

			if tc.wantErr {
				assert.Error(err)
			}

			assert.Equal(tc.expected, got)
		})
	}
}

func TestMultiply(t *testing.T) {
	// Arrange
	num0, num1 := 5.0, 6.0
	expected := 30.0

	// Act
	got := calc.Multiply(num0, num1)

	// Assert
	if got != expected {
		t.Errorf("Expected %.2f got %.2f", expected, got)
	}
}
