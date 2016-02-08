package main

import (
	"bufio"
	"fmt"
	"os"
)

func hasRepeatPair(line string) bool {
	for i := 0; i < len(line)-2; i++ {
		for j := i + 2; j < len(line)-1; j++ {
			if line[i] == line[j] && line[i+1] == line[j+1] {
				return true
			}
		}
	}
	fmt.Println("no pair")
	return false
}

func hasGapRepeat(line string) bool {
	for i := 0; i < len(line)-3; i++ {
		if line[i] == line[i+2] {
			return true
		}
	}
	fmt.Println("no gap")
	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	nice := 0
	for scanner.Scan() {
		line := scanner.Text()
		if hasRepeatPair(line) && hasGapRepeat(line) {
			nice++
		}
	}
	fmt.Println("Nice string count: ", nice)
}
