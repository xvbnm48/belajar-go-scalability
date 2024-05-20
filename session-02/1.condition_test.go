package session02

import (
	"fmt"
	"testing"
)

func TestIfElse(t *testing.T) {
	var currentYear = 2024

	if age := currentYear - 2001; age < 17 {
		t.Logf("You are %d years old, you are not allowed to drive", age)
	} else {
		t.Logf("You are %d years old, you are allowed to drive", age)
	}
}

func TestSwitchCondition(T *testing.T) {
	score := 70
	switch {
	case score >= 90:
		T.Log("A")
	case score >= 80:
		T.Log("B")
	case score >= 70:
		T.Log("C")
	case score >= 60:
		T.Log("D")
	default:
		T.Log("E")
	}
}

func TestSwitcConditionhWithRelationalOperator(t *testing.T) {
	score := 70
	switch {
	case score == 100:
		t.Log("Perfect")
	case (score < 100) && (score >= 90):
		t.Log("A")
	default:
		t.Log("Not Perfect")
	}
}

func TestSwitchFallthrough(t *testing.T) {
	score := 6
	switch {
	case score == 8:
		fmt.Println("perfect")
	case (score <= 8) && (score > 3):
		fmt.Println("You are good")
		fallthrough
	case score < 5:
		fmt.Println("You need to learn more")
	default:
		fmt.Println("You are not good")
	}
}

func TestNestedCondition(t *testing.T) {
	var score = 20
	if score > 7 {
		switch score {
		case 10:
			t.Log("Perfect")
		default:
			t.Log("Good")
		}
	} else {
		if score == 5 {
			t.Log("Not Bad")
		} else if score == 3 {
			t.Log("Bad")
		} else {
			t.Log("Very Bad")
			if score == 0 {
				t.Log("You are not good")
			}
		}
	}
}
