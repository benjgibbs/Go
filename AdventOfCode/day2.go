package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	sum := 0
	ribbon := 0
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "x")
		dims := []int{}
		for _, t := range parts {
			p, _ := strconv.Atoi(t)
			dims = append(dims, p)
		}
		lxh := dims[0] * dims[1]
		lxw := dims[0] * dims[2]
		hxw := dims[1] * dims[2]
		sum += 2*(lxh+lxw+hxw) + min(lxh, min(lxw, hxw))

		ribbon += dims[0] * dims[1] * dims[2]
		ms := min2(dims)
		ribbon += 2 * (ms[0] + ms[1])
	}
	fmt.Printf("Wrapping paper required: %d ft\n", sum)
	fmt.Printf("Ribbon required: %d ft\n", ribbon)
}

func min2(dims []int) []int {
	sort.Ints(dims)
	return dims[:2]
}
