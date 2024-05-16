package session02

import "testing"

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
