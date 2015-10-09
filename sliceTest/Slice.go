package main

import (
	"fmt"
)

func byValue(f []string) {
	f[1] = "value"
}

func byPointer(f *[]string) {
	(*f)[0] = "pointer"
}

func main() {
	vals := []string{"one", "two", "three"}
	fmt.Println(vals)
	byValue(vals)
	fmt.Println(vals)
	byPointer(&vals)
	fmt.Println(vals)
}
