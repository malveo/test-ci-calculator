package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
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
	fmt.Println("Versione 1.0.1")
	if len(os.Args) != 4 {
		fmt.Println("Usage: calculator <oeration> <num1> <num2>")
		fmt.Println("Operations: add, sub, mul, div")
		fmt.Println("Example: calculator add 10 5")
		os.Exit(1)
	}

	operation := os.Args[1]
	a, err := strconv.ParseFloat(os.Args[2], 64)
	if err != nil {
		fmt.Printf("Error: invalid number '%s'\n", os.Args[2])
		os.Exit(1)
	}

	b, err := strconv.ParseFloat(os.Args[3], 64)
	if err != nil {
		fmt.Printf("Error: invalid number '%s'\n", os.Args[3])
		os.Exit(1)
	}

	var result float64
	switch operation {
	case "add":
		result = Add(a, b)
		fmt.Printf("%.2f + %.2f = %.2f\n", a, b, result)
	case "sub":
		result = Subtract(a, b)
		fmt.Printf("%.2f - %.2f = %.2f\n", a, b, result)
	case "mul":
		result = Multiply(a, b)
		fmt.Printf("%.2f * %.2f = %.2f\n", a, b, result)
	case "div":
		result, err := Divide(a, b)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("%.2f / %.2f = %.2f\n", a, b, result)
	default:
		fmt.Printf("Error: unknown operation '%s'\n", operation)
		fmt.Println("Valid operations: add, sub, mul, div")
		os.Exit(1)
	}
}
