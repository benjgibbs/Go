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
