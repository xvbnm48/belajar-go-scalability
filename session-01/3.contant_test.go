package session01

import (
	"fmt"
	"testing"
)

func TestConstant(t *testing.T) {
	const name string = "vini"
	fmt.Printf("Name: %s\n", name)
}

func TestArithmeticOperator(t *testing.T) {
	const firstNumber = 10
	const secondNumber = 5
	const result = firstNumber + secondNumber
	var age = 21
	var salary = 1.5
	// convert float to int
	resultConvert := age * int(salary)
	fmt.Printf("Result: %d\n", resultConvert)
	fmt.Printf("Result: %d\n", result)
}

func TestRasional(t *testing.T) {
	var first bool
	first = 2 < 3
	var second bool = "vini" == "vini"
	var third bool = 10 != 2.3
	var fourth bool = 21 <= 11

	fmt.Printf("First: %t\n", first)
	fmt.Printf("Second: %t\n", second)
	fmt.Printf("Third: %t\n", third)
	fmt.Printf("Fourth: %t\n", fourth)
}

func TestLogicalOperator(t *testing.T) {
	var first bool
	first = true && false
	var second bool = true || false
	var third bool = !true

	fmt.Printf("First: %t\n", first)
	fmt.Printf("Second: %t\n", second)
	fmt.Printf("Third: %t\n", third)
}
