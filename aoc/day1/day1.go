package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		for i, c := range line {
			switch c {
			case '(':
				sum++
			case ')':
				sum--
			}
			if sum == -1 {
				fmt.Println("Entered the basement at:", i+1)
			}
		}
		fmt.Printf("Sum: %d\n", sum)
	}
}
