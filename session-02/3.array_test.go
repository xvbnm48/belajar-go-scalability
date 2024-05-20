package session02

import (
	"fmt"
	"strings"
	"testing"
)

func TestArray(t *testing.T) {
	numbers := [4]int{1, 2, 3, 4}
	strings := [2]string{"hello", "world"}

	fmt.Printf("%v\n", numbers)
	fmt.Printf("%v\n", strings)
}

func TestArrayModify(t *testing.T) {
	fruits := [3]string{"apple", "banana", "mango"}
	fruits[0] = "grape"
	fruits[1] = "melon"
	fruits[2] = "strawberry"
	fmt.Printf("%v\n", fruits)
}

func TestArrayLoopThrough(t *testing.T) {
	fruits := [3]string{"apple", "banana", "mango"}
	for i, v := range fruits {
		fmt.Printf("index: %d, value: %s\n", i, v)
	}

	fmt.Println(strings.Repeat("#", 10))
	for i := 0; i < len(fruits); i++ {
		fmt.Printf("index: %d, value: %s\n", i, fruits[i])
	}
}

func TestArrayMultiDimensional(t *testing.T) {
	balances := [2][3]int{{1, 2, 3}, {4, 5, 6}}

	for _, arr := range balances {
		for _, v := range arr {
			fmt.Printf("%d ", v)
		}
		fmt.Println()
	}
}
