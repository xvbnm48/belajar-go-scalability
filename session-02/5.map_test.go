package session02

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	idol := map[string]string{}
	idol["berlian"] = "hira dazzle"
	idol["edm_idol"] = "shoujo complex"
	idol["jpop_idol"] = "twenty nine teens"

	fmt.Println(idol)
	fmt.Println("idol local:", idol["berlian"], idol["edm_idol"], idol["jpop_idol"])
}

func TestMapLooping(t *testing.T) {
	var person = map[string]string{
		"name":    "vini",
		"age":     "18",
		"address": "tangerang",
	}

	for key, value := range person {
		fmt.Println(key, ":", value)
	}
}

func TestDeletingMap(t *testing.T) {
	var person = map[string]string{
		"name":    "vini",
		"age":     "18",
		"address": "tangerang",
	}
	fmt.Println("before delete:", person)
	delete(person, "age")
	fmt.Println("after delete:", person)
}

func TestMapDetectingValue(t *testing.T) {
	var person = map[string]string{
		"name":    "vini",
		"age":     "18",
		"address": "tangerang",
	}

	value, exist := person["age"]
	if exist {
		fmt.Println("value:", value, "exist:", exist)
	} else {
		fmt.Println("value does not exist")
	}

	delete(person, "age")
	value, exist = person["age"]
	if exist {
		fmt.Println("value:", value, "exist:", exist)
	} else {
		fmt.Println("value does not exist")
	}
}

func TestCombiningSliceWithMap(t *testing.T) {
	people := []map[string]string{
		{"name": "vini", "age": "18"},
		{"name": "meia", "age": "20"},
		{"name": "riichan", "age": "20"},
	}
	for i, person := range people {
		fmt.Printf("person %d: %s\n", i, person["name"])
	}
}
