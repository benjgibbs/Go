package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	input := os.Args[1]
	t1, _ := strconv.Atoi(os.Args[2])
	fmt.Printf("Input: %s, Iterations: %d\n", input, t1)
	for i := 0; i < t1; i++ {
		next := ""
		c := input[0]
		lastChange := 0
		for j := 0; j < len(input); j++ {
			if c != input[j] {
				cnt := j - lastChange
				next = next + strconv.Itoa(cnt) + string(c)
				c = input[j]
				lastChange = j
			}
		}
		cnt := len(input) - lastChange
		next = next + strconv.Itoa(cnt) + string(c)
		input = next
	}
	l1 := len(input)
	fmt.Println("Length: ", l1)
	if len(input) < 80 {
		fmt.Println(input)
	}

}
