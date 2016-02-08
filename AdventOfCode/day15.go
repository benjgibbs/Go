package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Ingredient struct {
	name                                             string
	capacity, durability, flavour, texture, calories int
}

type Ingredients []Ingredient

var ingredients Ingredients = Ingredients{}

func perms(of, count int) [][]int {
	if count == 1 {
		return [][]int{{of}}
	}
	res := [][]int{}
	for i := 0; i <= of; i++ {
		part := perms(of-i, count-1)
		for j := 0; j < len(part); j++ {
			part[j] = append(part[j], i)
		}
		res = append(res, part...)
	}

	return res
}

func main() {
	parseInput()
	calcMax()
}

func calcMax() {
	max := uint64(0)

	ni := len(ingredients)

	for _, weight := range perms(100, ni) {

		c, d, f, t := 0, 0, 0, 0

		for i := 0; i < ni; i++ {
			c += ingredients[i].capacity * weight[i]
			d += ingredients[i].durability * weight[i]
			f += ingredients[i].flavour * weight[i]
			t += ingredients[i].texture * weight[i]
		}

		if c > 0 && d > 0 && f > 0 && t > 0 {
			p := uint64(c) * uint64(d) * uint64(f) * uint64(t)
			if p > max {
				fmt.Println("New sum:", p, "Weights:", weight)
				max = p
			}
		}
	}
	//Sugar: capacity 3, durability 0, flavor 0, texture -3, calories 2
}

func parseInput() {
	pattern := regexp.MustCompile(`(\w+): \w+ (-?\d+), \w+ (-?\d+), \w+ (-?\d+), \w+ (-?\d+), \w+ (-?\d+)`)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		match := pattern.FindStringSubmatch(line)
		c, _ := strconv.Atoi(match[2])
		d, _ := strconv.Atoi(match[3])
		f, _ := strconv.Atoi(match[4])
		t, _ := strconv.Atoi(match[5])
		cal, _ := strconv.Atoi(match[6])
		ingredients = append(ingredients, Ingredient{match[1], c, d, f, t, cal})
	}
	fmt.Println(ingredients)
}
