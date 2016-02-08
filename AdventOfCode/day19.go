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
			unique := make(map[string]bool)
			fmt.Println(reps)
			for i := 0; i < len(line); i++ {
				if ns, c := reps[string(line[i])]; c {
					for _, n := range ns {
						rep := line[0:i]
						rep += n
						rep += line[i+1:]
						unique[rep] = true
					}
				} else if i < len(line)-1 {
					m := line[i : i+2]
					if ns, c := reps[m]; c {
						for _, n := range ns {
							rep := line[0:i]
							rep += n
							rep += line[i+2:]
							unique[rep] = true
						}
						i++
					}
				}
			}
			fmt.Println(len(unique))
		}
	}
}
