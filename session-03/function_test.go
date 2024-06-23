package session_03

import (
	"fmt"
	"math"
	"strings"
	"testing"
)

func TestFunction(t *testing.T) {
	greet("fariz", "jaksel")
	var names = []string{"Fariz,", "vini"}
	printMsg := greetReturn("hai", names)
	fmt.Println(printMsg)
}

func greet(name, address string) {
	fmt.Println("Hello there, my name is ", name)
	fmt.Println("I Live in ", address)
}

func greetReturn(msg string, names []string) string {
	var joinStr = strings.Join(names, " ")
	var result = fmt.Sprintf("%s %s", msg, joinStr)
	return result
}

func TestFunctionMultipleReturn(t *testing.T) {
	var diameter float64 = 15
	var area, circumference = calculate(diameter)
	fmt.Println("area: ", area)
	fmt.Println("circumference: ", circumference)
}

func calculate(d float64) (area float64, circumference float64) {
	area = math.Pi * math.Pow(d/2, 2)
	circumference = math.Pi * d
	return
}
