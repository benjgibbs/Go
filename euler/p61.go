package main

import (
	"fmt"
)

const CYCLE_LEN = 6

func front(i int) int {
	return i / 100
}

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

	populate(trias, 3, &fronts, &backs)
	populate(squas, 4, &fronts, &backs)
	populate(pents, 5, &fronts, &backs)
	populate(hexes, 6, &fronts, &backs)
	populate(hepts, 7, &fronts, &backs)
	populate(octos, 8, &fronts, &backs)

	for _, front := range fronts {
		for _, start := range front {
			var found bool
			visited := []Number{}
			visited, found = recurse(start, visited, fronts)
			if found {
				var sum int
				for _, x := range visited {
					sum += x.value
				}
				fmt.Printf("Found it: %d (%d)\n", sum, visited)
				return
			}
		}
	}
}

func recurse(v Number, visited []Number, fronts map[int][]Number) ([]Number, bool) {
	visited = append(visited, v)
	if len(visited) == CYCLE_LEN && visited[0].front == visited[CYCLE_LEN-1].back {
		return visited, true
	}

	next := fronts[v.back]
	for _, f := range next {
		if eligible(visited, f) {
			res, found := recurse(f, visited, fronts)
			if found {
				return res, true
			}
		}
	}
	return visited, false
}

func eligible(seen []Number, t Number) bool {
	for _, x := range seen {
		if x.class == t.class || x.value == t.value {
			return false
		}
	}
	return true
}
