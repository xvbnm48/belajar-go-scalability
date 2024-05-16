package session01

import (
	"fmt"
	"testing"
)

func TestVariable(t *testing.T) {
	var name string
	name = "vini"
	age := 21
	fmt.Printf("Name: %s, Age: %d\n", name, age)
}

func TestMultipleVariabel(t *testing.T) {
	var name, age, country = "vini", 21, "Indonesia"
	first, second, third := "first", "second", "third"
	fmt.Println(name, age, country)
	fmt.Println(first, second, third)
}

func TestUnderScoreVariable(t *testing.T) {
	var firstVariable string
	var name, age, country = "vini", 21, "Indonesia"
	_, _, _, _ = firstVariable, name, age, country
}

func TestPrintf(t *testing.T) {
	var name = "vini"
	age := 21
	country := "indonesia"
	fmt.Printf("Name: %s, Age: %d, Country: %s\n", name, age, country)
	// string using %s, integer using %d, float using %f
}
