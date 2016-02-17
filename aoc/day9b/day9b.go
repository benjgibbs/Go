package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

type Path struct {
	e1, e2 string
	d      int
}

type Paths []Path
type Countries []string

var paths Paths = Paths{}
var countries Countries = Countries{}

func hasCounry(cs Countries, c string) bool {
	for i := 0; i < len(cs); i++ {
		if cs[i] == c {
			return true
		}
	}
	return false
}

func addCountry(c string) {
	if !hasCounry(countries, c) {
		countries = append(countries, c)
	}
}

func findPath(e1, e2 string) (bool, Path) {
	for _, p := range paths {
		if (p.e1 == e1 && p.e2 == e2) || (p.e1 == e2 && p.e2 == e1) {
			return true, p
		}
	}
	return false, Path{}
}

func parseInput() {
	scanner := bufio.NewScanner(os.Stdin)
	pattern := regexp.MustCompile(`(.*) to (.*) = (\d+)`)
	for scanner.Scan() {
		line := scanner.Text()
		if match := pattern.FindStringSubmatch(line); len(match) > 0 {
			distance, err := strconv.Atoi(match[3])
			if err != nil {
				panic(fmt.Sprintf("bad distances %s on line %s", match[3], line))
			}
			paths = append(paths, Path{match[1], match[2], distance})
			addCountry(match[1])
			addCountry(match[2])
		}
	}
}

func countRoutes() {
	counts := map[string]int{}
	for _, p := range paths {
		counts[p.e1]++
		counts[p.e2]++
	}
	fmt.Println(counts)
}

func perms(of []string) [][]string {
	if len(of) == 1 {
		return [][]string{of}
	}
	a := of[0]
	b := of[1:]
	res := [][]string{}
	for _, ps := range perms(b) {
		for psi := 0; psi < len(ps)+1; psi++ {
			perm := make([]string, len(ps)+1)
			copy(perm[0:], ps[:psi])
			copy(perm[psi:], []string{a})
			copy(perm[psi+1:], ps[psi:])
			res = append(res, perm)
		}
	}
	return res
}

func main() {
	parseInput()
	countRoutes()
	fmt.Println(len(countries))
	max := math.MinInt32
	route := []string{}
	for _, perm := range perms(countries) {
		dist := 0
		for ci := 0; ci < len(perm)-1; ci++ {
			if e, p := findPath(perm[ci], perm[ci+1]); e {
				dist += p.d
			} else {
				dist = math.MinInt32
			}
		}
		if dist > max {
			max = dist
			route = perm
		}
	}
	fmt.Println("Max: ", max)
	fmt.Println("Route: ", route)
}
