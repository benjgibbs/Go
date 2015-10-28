package main

import (
	"fmt"
)

func isBouncy(x uint64) bool {
	last := x % 10
	var firstPass = true
	var increasing = false
	for x > 0 {
		cur := x % 10
		//fmt.Printf("x=%d, last=%d, cur=%d, firstPass=%t, increasing=%t\n", x, last, cur, firstPass, increasing)
		if firstPass {
			if last != cur {
				firstPass = false
				increasing = last < cur
			}
		} else {
			if cur < last && increasing {
				return true
			} else if cur > last && !increasing {
				return true
			}
		}
		last = cur
		x = x / 10
	}
	return false
}

func p112() {
	var count uint64 = 0

	var frac float64 = 0
	const limit = 0.99
	var i = 0
	for frac < limit {
		i++
		if isBouncy(uint64(i)) {
			count++
		}
		frac = float64(count) / float64(i)
	}
	fmt.Printf("limit=%f, i=%d, count=%d, frac=%f\n", limit, i, count, frac)
}

func testUnder1000() {
	const limit = 1000
	var count uint64 = 0
	for i := 0; i < limit; i++ {
		if isBouncy(uint64(i)) {
			count++
		}
	}
	fmt.Printf("%d bouncy numbers below %d\n", count, limit)
}

func tests() {
	checkBouncy(123456789, false)
	checkBouncy(123454789, true)
	checkBouncy(987654321, false)
	checkBouncy(987674321, true)
	checkBouncy(997674321, true)
	checkBouncy(111111111, false)
	checkBouncy(134468, false)
	checkBouncy(66420, false)
	checkBouncy(155349, true)
}

func checkBouncy(n uint64, expect bool) {
	is := "is not"
	result := isBouncy(n)
	if result {
		is = "is"
	}
	err := "OK"
	if expect != result {
		err = "ERROR"
	}
	fmt.Printf("%s: %d %s Bouncy\n", err, n, is)
}
