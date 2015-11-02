package main

import (
	"log"
	"strconv"
)

func fact(n uint64) uint64 {
	var res uint64 = 1
	for n > 0 {
		res *= n
		n--
	}
	return res
}

func atoi(s string) int {
	i, e := strconv.Atoi(s)
	if e != nil {
		log.Fatal(e)
	}
	return i
}

func containsInt(x int, xs []int) (bool, int) {
	for i, v := range xs {
		if v == x {
			return true, i
		}
	}
	return false, -1
}

func contains(x uint64, ys []uint64) (bool, int) {
	for i, y := range ys {
		if x == y {
			return true, i
		}
	}
	return false, -1
}

func Heap(xs []int) [][]int {
	var result [][]int
	heap(len(xs), xs, &result)
	return result
}

func heap(n int, xs []int, res *[][]int) {
	if n == 1 {
		out := make([]int, len(xs))
		copy(out, xs)
		*res = append(*res, out)
	} else {
		for i := 0; i < n-1; i += 1 {
			heap(n-1, xs, res)
			if n%2 == 0 {
				xs[i], xs[n-1] = xs[n-1], xs[i]
			} else {
				xs[0], xs[n-1] = xs[n-1], xs[0]
			}
		}
		heap(n-1, xs, res)
	}
}
