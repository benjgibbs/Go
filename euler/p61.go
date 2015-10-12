package main

import (
	"fmt"
)

/**
 * Takes a 4 digit in and returns the first two
 */
func front(i int) int {
	return i / 100
}

/**
 * Takes a 4 digit int and returns the last two
 */
func back(i int) int {
	return i % 100
}

func seq(min, max int, next func(n int) int) []int {
	var res []int
	last := 0
	for n := 1; last < max; n++ {
		last = next(n)
		if last >= min && last <= max {
			res = append(res, last)
		}
	}
	return res
}

type Number struct {
	value int
	front int
	back  int
	class int
}

func newNumber(x, c int) *Number {
	result := Number{
		value: x,
		front: front(x),
		back:  back(x),
		class: c}
	return &result
}

func populate(xs []int, c int, f, b *map[int][]Number) {
	for _, x := range xs {
		n := newNumber(x, c)
		(*f)[(*n).front] = append((*f)[(*n).front], *n)
		(*b)[(*n).back] = append((*b)[(*n).back], *n)
	}
}

func main() {
	fronts := make(map[int][]Number)
	backs := make(map[int][]Number)

	trias := seq(1000, 9999, func(n int) int { return n * (n - 1) / 2 })
	squas := seq(1000, 9999, func(n int) int { return n * n })
	pents := seq(1000, 9999, func(n int) int { return n * (3*n - 1) / 2 })
	hexes := seq(1000, 9999, func(n int) int { return n * (2*n - 1) })
	hepts := seq(1000, 9999, func(n int) int { return n * (5*n - 3) / 2 })
	octos := seq(1000, 9999, func(n int) int { return n * (3*n - 2) })

	fmt.Printf("Sizes: trias=%d, squas=%d, pents=%d, hexes=%d, hepts=%d, octos=%d\n",
		len(trias), len(squas), len(pents), len(hexes), len(hepts), len(octos))
	/*
		populate(trias, 3, &fronts, &backs)
		populate(squas, 4, &fronts, &backs)
		populate(pents, 5, &fronts, &backs)
		populate(hexes, 6, &fronts, &backs)
		populate(hepts, 7, &fronts, &backs)
		populate(octos, 8, &fronts, &backs)
		const CYCLE_LEN = 6
	*/
	populate([]int{1122}, 1, &fronts, &backs)
	populate([]int{2244}, 2, &fronts, &backs)
	populate([]int{2233}, 2, &fronts, &backs)
	populate([]int{3344}, 3, &fronts, &backs)
	populate([]int{3311}, 3, &fronts, &backs)
	const CYCLE_LEN = 3

	// fmt.Printf("Fronts: %d\n", fronts)
	// fmt.Printf("Backs: %d\n", backs)
	// fmt.Println("")

	for _, front := range fronts {
		for _, start := range front {

			visited := []Number{}
			visited = recurse(start, visited, fronts)

			if len(visited) == CYCLE_LEN && visited[0].front == visited[CYCLE_LEN-1].back {
				var sum int
				for _, x := range visited {
					sum += x.value
				}
				fmt.Printf("Found it: %d (%d)\n", sum, visited)
				return
			}
			if true || len(visited) == CYCLE_LEN {
				fmt.Printf("Not found. start=%d, len=%d, route=%d\n", start, len(visited), visited)
				//fmt.Println()
			}
		}
	}
}

func recurse(v Number, visited []Number, fronts map[int][]Number) []Number {
	visited = append(visited, v)
	//fmt.Printf("Recursing. v=%d, visitied=%d\n", v, visited)
	next := fronts[v.back]
	for _, f := range next {
		if eligible(visited, f) {
			return recurse(f, visited, fronts)
		}
	}
	return visited
}

func eligible(seen []Number, t Number) bool {
	for _, x := range seen {
		if x.class == t.class || x.value == t.value {
			//fmt.Printf("Not eligbile. %d/%d %d/%d\n", x.class, t.class, x.value, t.value)
			return false
		}
	}
	//fmt.Printf("Eligible %d\n", t)
	return true
}

//[{7140 40 71 6}
//    {4371 71 43 3}
//    {4371 71 43 3}
//       {6943 43 69 7}]
