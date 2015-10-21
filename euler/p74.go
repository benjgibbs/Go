package main

import (
	"fmt"
)

type Calculator struct {
	cache map[uint64]uint64
}

func (c *Calculator) fact(i uint64) uint64 {
	result, found := c.cache[i]
	if found {
		return result
	}
	result = uint64(1)
	fact := i
	for fact > 1 {
		result *= fact
		fact--
	}
	c.cache[i] = result
	return result
}

func sdf(i uint64, c *Calculator) uint64 {
	result := uint64(0)
	for i > 0 {
		result += c.fact(i % 10)
		i /= 10
	}
	return result
}

func contains(x uint64, ys []uint64) (bool, int) {
	for i, y := range ys {
		if x == y {
			return true, i
		}
	}
	return false, 0
}

func cycleLen(i uint64, c *Calculator) int {
	seen := []uint64{i}
	for {
		i = sdf(i, c)
		//fmt.Printf("i=%d, seen=%d\n", i, seen)
		t, _ := contains(i, seen)
		if t {
			//fmt.Printf("Found: p=%d, seen=%d\n", p, seen)
			return len(seen)
		}
		seen = append(seen, i)
	}
}

func p74() {
	c := Calculator{cache: make(map[uint64]uint64)}
	const (
		iters = 1000000
	)
	fmt.Printf("Cycle len 145: %d\n", cycleLen(145, &c))
	fmt.Printf("Cycle len 169: %d\n", cycleLen(169, &c))
	fmt.Printf("Cycle len 69: %d\n", cycleLen(69, &c))
	max, num60s := 0, 0
	for i := 1; i < iters; i++ {
		cl := cycleLen(uint64(i), &c)
		if cl == 60 {
			num60s++
		}
		if cl > max {
			max = cl
		}
	}
	fmt.Printf("Max: %d, Num with 60: %d\n", max, num60s)

}
