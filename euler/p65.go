package main

import (
	"fmt"
	"math/big"
)

// Assume that 2 is the zero'th term so that the k works out
func gen(term int) *big.Int {
	switch term % 3 {
	case 0, 1:
		return big.NewInt(1)
	}
	next := 2 * (term + 1) / 3
	return big.NewInt(int64(next))
}

func main() {
	h1, h2, k1, k2 := big.NewInt(3), big.NewInt(2), big.NewInt(1), big.NewInt(1)
	for term := 2; term < 100; term++ {
		a := gen(term)

		h, k := big.NewInt(0), big.NewInt(0)

		h.Mul(a, h1)
		h.Add(h, h2)

		k.Mul(a, k1)
		k.Add(k, k2)

		k2, k1 = k1, k
		h2, h1 = h1, h
	}
	fmt.Printf("%d / %d", h1, k1)
	fmt.Printf(" : sum is %d\n", sumDigits(h1))
}

func sumDigits(toSum *big.Int) int64 {
	sum := big.NewInt(0)
	ten := big.NewInt(10)
	add := big.NewInt(0)
	zero := big.NewInt(0)
	for toSum.Cmp(zero) > 0 {
		add.Mod(toSum, ten)
		sum.Add(sum, add)
		toSum.Div(toSum, ten)
	}
	return sum.Int64()
}
