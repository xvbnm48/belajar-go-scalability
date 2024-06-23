package session_03

import (
	"fmt"
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
	var result string = fmt.Sprintf("%s %s", msg, joinStr)
	return result
}
