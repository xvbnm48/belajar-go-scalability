package session02

import (
	"fmt"
	"testing"
)

func TestLoopingFirstWay(t *testing.T) {
	for i := 0; i < 3; i++ {
		t.Log(i)
	}
}

func TestLoopingSecondWay(t *testing.T) {
	i := 0

	for i < 3 {
		fmt.Println(i)
		i++
	}
}

func TestLoopingThirdWay(t *testing.T) {
	i := 0
	for {
		fmt.Println("angka", i)
		i++
		if i == 3 {
			break
		}
	}
}

func TestLoopingFourthWay(t *testing.T) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}

		if i > 8 {
			break

		}

		fmt.Println("angka", i)
	}
}

func TestNestedLooping(t *testing.T) {
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}
}

func TestNestedLoopingLabel(t *testing.T) {
outerloop:
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			if j == 3 {
				break outerloop
			}
			fmt.Print(j, " ")
		}
		fmt.Println()
	}
}
