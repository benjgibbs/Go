package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

func check(what, line string, expect int) int {
	matcher := regexp.MustCompile(what + ": (\\d+)")
	match := matcher.FindStringSubmatch(line)
	if len(match) == 0 {
		return 0
	}
	valStr := match[1]
	val, _ := strconv.Atoi(valStr)
	return (expect - val) * (expect - val)
}

func main() {
	minDistance := math.MaxInt32
	scanner := bufio.NewScanner(os.Stdin)

	idRe := regexp.MustCompile(`Sue (\d+):`)

	for scanner.Scan() {
		line := scanner.Text()
		sue := idRe.FindStringSubmatch(line)
		dist := 0
		dist += check("children", line, 3)
		dist += check("cats", line, 7)
		dist += check("samoyeds", line, 2)
		dist += check("pomeranians", line, 3)
		dist += check("akitas", line, 0)
		dist += check("vizslas", line, 0)
		dist += check("goldfish", line, 5)
		dist += check("trees", line, 3)
		dist += check("cars", line, 2)
		dist += check("perfumes", line, 1)
		if dist < minDistance {
			fmt.Println("New minimum:", dist, " Sue:", sue)
			minDistance = dist
		}
	}
}
