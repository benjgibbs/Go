package main

import (
	"fmt"
)

func binomial(n, k int64) int64 {
	var num, den int64 = 1, 1
	for i := int64(1); i <= k; i++ {
		num *= n + 1 - i
		den *= i
	}
	return num / den
}

func abs(x int64) int64 {
	if x >= 0 {
		return int64(x)
	}
	return int64(-x)
}

func p85() {
	const want int64 = 2000000
	var (
		maxLen  int64 = 80
		closest int64 = want
	)

	for a := int64(1); a < maxLen; a++ {
		for b := a + 1; b < maxLen; b++ {
			numRects := binomial(int64(a+1), 2) * binomial(int64(b+1), 2)
			distance := abs(numRects - want)
			if distance < closest {
				fmt.Printf("New closer numRects=%d, (%d,%d) with area %d\n", numRects, a, b, a*b)
				closest = distance
			}
		}
	}
}
