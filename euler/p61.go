package main

import (
	"fmt"
	"github.com/benjgibbs/permutations"
)

func match(i, j int) bool {
	p1 := (i % 100)
	p2 := (j / 100)
	return p1 == p2
}

func isCircular(is []int) bool {
	for i := 0; i < len(is)-1; i++ {
		if !match(is[i], is[i+1]) {
			return false
		}
	}
	return match(is[len(is)-1], is[0])
}

func check(is []int) {
	fmt.Printf("%s: %t\n", is, isCircular(is))
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

func main() {
	trias := seq(1000, 9999, func(n int) int { return n * (n - 1) / 2 })
	squas := seq(1000, 9999, func(n int) int { return n * n })
	pents := seq(1000, 9999, func(n int) int { return n * (3*n - 1) / 2 })
	hexes := seq(1000, 9999, func(n int) int { return n * (2*n - 1) })
	hepts := seq(1000, 9999, func(n int) int { return n * (5*n - 3) / 2 })
	octos := seq(1000, 9999, func(n int) int { return n * (3*n - 2) })
	fmt.Printf("Sizes: trias=%d, squas=%d, pents=%d, hexes=%d, hepts=%d, octos=%d\n",
		len(trias), len(squas), len(pents), len(hexes), len(hepts), len(octos))

	for t, tv := range octos {
		for _, sv := range squas {
			if sv == tv {
				continue
			}
			for _, pv := range pents {
				if pv == sv || pv == tv {
					continue
				}
				for _, xv := range hexes {
					if xv == pv || xv == sv || xv == tv {
						continue
					}
					for _, hv := range hepts {
						if hv == xv || hv == pv || hv == sv || hv == tv {
							continue
						}
						for _, ov := range octos {
							if ov == hv || ov == xv || ov == pv || ov == sv || ov == tv {
								continue
							}
							parts := []int{tv, sv, pv, xv, hv, ov}
							if !hasMatchingEnds(parts) {
								continue
							}
							fmt.Printf("Matching ends found: %d (%f)\n", parts, float32(t+1)/float32(len(octos)))
							for _, perm := range permutations.Heap(parts) {
								//fmt.Printf("Checking: %d\n", perm)
								if isCircular(perm) {
									fmt.Printf("Found: %d\n", perm)
									return
								}
							}
						}
					}
				}
			}
		}
	}
}

func hasMatchingEnds(xs []int) bool {
	begins := make(map[int]int)
	ends := make(map[int]int)
	for _, x := range xs {
		b := x / 100
		begins[b]++
		e := x % 100
		ends[e]++
	}
	for b, c := range begins {
		if c != ends[b] {
			return false
		}
	}
	return true
}

func showGivens() {
	trias := seq(1, 70, func(n int) int { return n * (n - 1) / 2 })
	squas := seq(1, 70, func(n int) int { return n * n })
	pents := seq(1, 70, func(n int) int { return n * (3*n - 1) / 2 })
	hexes := seq(1, 70, func(n int) int { return n * (2*n - 1) })
	hepts := seq(1, 70, func(n int) int { return n * (5*n - 3) / 2 })
	octos := seq(1, 70, func(n int) int { return n * (3*n - 2) })
	fmt.Printf("trias=%d\n", trias)
	fmt.Printf("squas=%d\n", squas)
	fmt.Printf("pents=%d\n", pents)
	fmt.Printf("hexes=%d\n", hexes)
	fmt.Printf("hepts=%d\n", hepts)
	fmt.Printf("octos=%d\n", octos)
}
