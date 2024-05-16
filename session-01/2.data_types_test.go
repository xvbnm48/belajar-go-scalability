package session01

import (
	"fmt"
	"testing"
)

func TestNumber(t *testing.T) {
	var first uint8 = 48
	var second int8 = -48
	var third float32 = 3.65

	// %T is used to print the type of the variable
	fmt.Printf("%d is %T\n", first, first)
	fmt.Printf("%d is %T\n", second, second)
	fmt.Printf("%.2f is %T\n", third, third)
}

func TestBool(t *testing.T) {
	var condition bool = true
	fmt.Printf("is it permitted ? %t\n", condition)
}

func TestString(t *testing.T) {
	var message string = "Hello, World!"
	var anotherMessage string = " Hira Dazzle"
	fmt.Println(message + anotherMessage)
	fmt.Println(message)
}
