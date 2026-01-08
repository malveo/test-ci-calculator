package main

import (
	"errors"
	"fmt"
)

// Add returns the sum of two numbers
func Add(a, b float64) float64 {
	return a + b
}

// Subtract returns the difference between two numbers
func Subtract(a, b float64) float64 {
	return a - b
}

// Multiply returns the product of two numbers
func Multiply(a, b float64) float64 {
	return a * b
}

// Divide returns the quotient of two numbers
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func main() {
	fmt.Println("Calculator Demo")
	fmt.Printf("10 + 5 = %.2f\n", Add(10, 5))
	fmt.Printf("10 - 5 = %.2f\n", Subtract(10, 5))
	fmt.Printf("10 * 5 = %.2f\n", Multiply(10, 5))

	result, err := Divide(10, 5)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("10 / 5 = %.2f\n", result)
	}

	_, err = Divide(10, 0)
	if err != nil {
		fmt.Printf("10 / 0 = Error: %v\n", err)
	}
}
