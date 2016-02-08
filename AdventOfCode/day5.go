package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func countVowels(line string) int {
	count := 0
	for _, c := range line {
		switch c {
		case 'a', 'e', 'i', 'o', 'u':
			count++
		}
	}
	return count
}
func hasRepeat(line string) bool {
	for i := 1; i < len(line); i++ {
		if line[i-1] == line[i] {
			return true
		}
	}
	return false
}
func hasBannedStrings(line string) bool {
	return strings.Contains(line, "ab") || strings.Contains(line, "cd") ||
		strings.Contains(line, "pq") || strings.Contains(line, "xy")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		if !hasBannedStrings(line) && countVowels(line) > 2 && hasRepeat(line) {
			sum++
		}
	}
	fmt.Println("Nice string count: ", sum)
}
