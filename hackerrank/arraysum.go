package main

import "fmt"

func main() {
	var a, b, c uint32
	fmt.Scanf("%v\n", &a)
	for i := uint32(0); i < a; i++ {
		fmt.Scanf("%v", &b)
		c += b
	}
	fmt.Println(c)
}
