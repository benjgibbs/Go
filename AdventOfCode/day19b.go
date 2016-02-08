package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

type Replacements []string

func main() {
	reps := make(map[string]Replacements)
	matcher := regexp.MustCompile(`(\w+) => (\w+)`)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		match := matcher.FindStringSubmatch(line)
		if len(match) == 3 {
			reps[match[1]] = append(reps[match[1]], match[2])
		} else if len(line) > 0 {
			fmt.Println(len(unique))
		}
	}
}
