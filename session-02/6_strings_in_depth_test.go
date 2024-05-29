package session02

import (
	"fmt"
	"testing"
)

func TestLoopingOverString(T *testing.T) {
	city := "tokyo"
	for i := 0; i < len(city); i++ {
		fmt.Printf("index: %d", "value: %d", i, city[i])
	}
	jakarta := []byte{
		74, 97, 107, 97, 114, 116, 97,
	}
	for i := 0; i < len(jakarta); i++ {
		fmt.Printf("index: %d", "value: %d", i, jakarta[i])
	}
}
