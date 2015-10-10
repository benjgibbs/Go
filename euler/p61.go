package main

import (
	"fmt"
)

/**
 * Takes a 4 digit in and returns the first two
 */
func front(i int) int {
	return i % 100
}

/**
 * Takes a 4 digit int and returns the last two
 */
func back(i int) int {
	return i / 100
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
	back int
	class int
}

func newNumber(x,c int) *Number {
	result := Number {
		value: x,
		front: front(x),
		back: back(x),
		class: c }
	return &result
}

func populate(xs []int, c int, f,b *map[int][]*Number) {
	for _,x := range xs {
		n := newNumber(x,c)
		f[(*n).front] = append(f[(*n).front],n)
		b[(*n).back] = append(f[(*n).back],n)
	}
}

func contains(xs []int, t int) bool {
	for _,x := range xs {
		if x == t {
			return true
		}
	}
	return false
}

func main() {
	trias := seq(1000, 9999, func(n int) int { return n * (n - 1) / 2 })
	squas := seq(1000, 9999, func(n int) int { return n * n })
	pents := seq(1000, 9999, func(n int) int { return n * (3*n - 1) / 2 })
	hexes := seq(1000, 9999, func(n int) int { return n * (2*n - 1) })
	hepts := seq(1000, 9999, func(n int) int { return n * (5*n - 3) / 2 })
	octos := seq(1000, 9999, func(n int) int { return n * (3*n - 2) })
	fmt.Printf("Sizes: trias=%d, squas=%d, pents=%d, hexes=%d, hepts=%d, octos=%d\n",
		len(trias), len(squas), len(pents), len(hexes), len(hepts), len(octos))

	fronts := make(map[int][]*Number)
	backs  := make(map[int][]*Number)
	populate(trias, 3, &fronts, &backs)
	populate(squas, 4, &fronts, &backs)
	populate(pents, 5, &fronts, &backs)
	populate(hexes, 6, &fronts, &backs)
	populate(hepts, 7, &fronts, &backs)
	populate(octos, 8, &fronts, &backs)

	for _,v := backs {
		visited := []Number{}
		recurse(v,&visited,&fronts,&backs)
		if len(visited) == 6 {
			fmt.Printf("Found it: %d", visited)
		}
	}
}

func recurse(v, visited []*Number,fronts, backs *map[int][]*Number) {
	for _,b := range v {
		visited := append(visited,b.class)
		for _, f := range fronts[b.back]{
			if !contains(visited, f) {
				recurse(b,visited,fronts,backs)
			}
		}
	}
}


//	for _, ord := range permutations.heap([]int{3, 4, 5, 6, 7, 8}) {
//		for i = 0; i < len(ord)-1; i++ {
//			x := ord[i]
//			y := ord[i+1]
//			var ends *Ends
//			switch x {
//			case 3:
//				ends = &triaEnds
//			case 4:
//				ends = &squaEnds
//			case 5:
//				ends = &pentEnds
//			case 6:
//				ends = &hexeEnds
//			case 7:
//				ends = &heptEnds
//			case 8:
//				ends = &octaEnds
//			}
//			for _,starts := ends.back
//		}
