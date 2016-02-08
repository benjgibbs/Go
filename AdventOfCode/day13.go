package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

type Neighbour struct {
	p1, p2 string
	v      int
}

type Neighbours []Neighbour

var nbs = Neighbours{}
var ppl = []string{}

func main() {
	parseInput()
	max := math.MinInt32
	for _, perm := range perms(ppl) {
		happiness := 0
		happiness += findHappiness(perm[len(perm)-1], perm[0])
		for i := 0; i < len(perm)-1; i++ {
			happiness += findHappiness(perm[i], perm[i+1])
		}
		if happiness > max {
			fmt.Println("Better configuration: ", perm)
			fmt.Println("Happiness now: ", happiness)
			max = happiness
		}

	}
}

func findHappiness(p1, p2 string) int {
	result := 0
	for _, nb := range nbs {
		if nb.p1 == p1 && nb.p2 == p2 {
			result += nb.v
		} else if nb.p1 == p2 && nb.p2 == p1 {
			result += nb.v
		}
	}
	return result
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

func parseInput() {
	scanner := bufio.NewScanner(os.Stdin)
	matcher := regexp.MustCompile(`(\w+) would (gain|lose) (\w+) happiness units by sitting next to (\w+).`)

	for scanner.Scan() {
		match := matcher.FindStringSubmatch(scanner.Text())
		if len(match) == 5 {
			p1 := match[1]
			p2 := match[4]
			sign := 1
			if match[2] == "lose" {
				sign = -1
			}
			w, err := strconv.Atoi(match[3])
			if err != nil {
				panic("Bad weight: " + match[3])
			}
			w = w * sign
			addPerson(p1)
			addPerson(p2)
			nbs = append(nbs, Neighbour{p1, p2, w})
		}
	}
	fmt.Println("Num People: ", len(ppl))
	fmt.Println("People: ", ppl)
	fmt.Println("Neighbours: ", nbs)
}

func addPerson(s string) {
	for _, p := range ppl {
		if p == s {
			return
		}
	}
	ppl = append(ppl, s)
}
