package main

import (
	"fmt"
	"testing"
)

func TestSlicing(t *testing.T) {
	bytes := []byte{1, 2, 3, 4}
	fmt.Println(bytes)
	fmt.Println(bytes[:len(bytes)-1])
	fmt.Println(bytes[:1])
	fmt.Println(bytes[1:])
}
