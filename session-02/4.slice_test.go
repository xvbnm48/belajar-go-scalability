package session02

import (
	"fmt"
	"testing"
)

func TestSliceFunction(t *testing.T) {
	// slice
	fruits := make([]string, 3)
	_ = fruits
	fmt.Printf("%v\n", fruits)
}

func TestSliceAppend(t *testing.T) {
	fruits := make([]string, 3)
	fruits = append(fruits, "apple")
	fmt.Printf("%v\n", fruits)
}

func TestSliceAppendFunctionWithEllipsis(t *testing.T) {
	fruits := []string{"apple", "banana", "mango"}
	fruits2 := []string{"grape", "melon", "strawberry"}
	fruits = append(fruits, fruits2...)

	fmt.Printf("%v\n", fruits)
}
