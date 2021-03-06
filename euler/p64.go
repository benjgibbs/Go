package main

import (
	"fmt"
	"math"
)

type iteration struct {
	m, d, a int
}

func equal(i, j iteration) bool {
	return i.a == j.a && i.d == j.d && i.m == j.m
}

func contFracLen(s int) int {

	a0 := int(math.Floor(math.Sqrt(float64(s))))
	m, d, a := 0, 1, a0
	previous := []iteration{}

	for {
		m = d*a - m
		d = (s - (m * m)) / d
		a = (a0 + m) / d
		it := iteration{m, d, a}
		for i, x := range previous {
			if equal(x, it) {
				return len(previous) - i
			}
		}
		previous = append(previous, it)
	}
}

func countOddsUpTo(bound int) int {
	squares := map[int]struct{}{}
	for i := 1; i*i <= bound; i++ {
		squares[i*i] = struct{}{}
	}

	oddCount := 0
	for i := 2; i <= bound; i++ {
		_, sq := squares[i]
		if !sq {
			if contFracLen(i)%2 == 1 {
				oddCount++
			}
		}
	}
	return oddCount
}

func p64() {
	fmt.Printf("Odd fractions up to %d number %d\n", 13, countOddsUpTo(13))
	fmt.Printf("Odd fractions up to %d number %d\n", 10000, countOddsUpTo(10000))
}
