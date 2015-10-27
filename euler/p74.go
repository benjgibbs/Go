package main

import (
	"fmt"
)

func sdf(i uint64) uint64 {
	result := uint64(0)
	for i > 0 {
		result += fact(i % 10)
		i /= 10
	}
	return result
}

func cycleLen(i uint64) int {
	seen := []uint64{i}
	for {
		i = sdf(i)
		c, _ := contains(i, seen)
		if c {
			return len(seen)
		}
		seen = append(seen, i)
	}
}

func p74() {
	const iters = 1000000
	max, num60s := 0, 0
	for i := 1; i < iters; i++ {
		cl := cycleLen(uint64(i))
		if cl == 60 {
			num60s++
		}
		if cl > max {
			max = cl
		}
	}
	fmt.Printf("Max: %d, Num with 60: %d\n", max, num60s)
}
