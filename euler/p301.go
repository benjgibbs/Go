package main

import (
	"fmt"
)

func p301() {
	sum := 0
	// pp http://www.math.ucla.edu/~tom/Game_Theory/comb.pdf
	for i := 1; i <= 2<<29; i++ {
		if i^(2*i)^(3*i) == 0 {
			sum++
		}
	}
	fmt.Printf("sum=%d\n", sum)
}
