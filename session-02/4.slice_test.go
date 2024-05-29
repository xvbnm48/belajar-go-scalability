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

func TestSliceCopyFunction(t *testing.T) {
	idol_local := []string{"hira dazzle", "shoujo complex", "twenty nine teens"}
	idol_japan := []string{"nogizaka46", "keyakizaka46", "hinatazaka46"}

	mix := copy(idol_local, idol_japan)
	fmt.Println("idol local:", idol_local)
	fmt.Println("idol japan:", idol_japan)
	fmt.Println("copied data:", mix)
	// print result after copy

}

func TestSliceSlicing(t *testing.T) {
	idol := []string{"hira dazzle", "shoujo complex", "twenty nine teens", "nogizaka46", "keyakizaka46", "hinatazaka46"}
	idol_local := idol[:3]
	idol_japan := idol[3:]

	fmt.Println("idol local:", idol_local)
	fmt.Println("idol japan:", idol_japan)
}

func TestSliceCombiningSliceAndAppend(t *testing.T) {
	idol := []string{"hira dazzle", "shoujo complex", "twenty nine teens", "nogizaka46", "keyakizaka46", "hinatazaka46"}
	idol2 := idol[1:4] // shoujo complex, twenty nine teens, nogizaka46
	idol2[0] = "amai monogatari"
	fmt.Println("idol2:", idol2)
	fmt.Println("idol:", idol)
}
