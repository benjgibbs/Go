package main

import (
	"fmt"
)

func nextRepUnit(base, last uint) uint {
	return last*base + 1
}

func p346() {
	const limit = 10
	for upto := uint(1); upto < limit; upto++ {
		for b1 := uint(1); b1 < upto; b1++ {
			for v1 := 1; v1 < upto; v1 = nextRepUnit(v1, b1) {
				lastRep := make(map[uint]uint)
				for  b2 := uint(1); b2 < upto; b2 = b2++ {
					for v2 := lastRep
						fmt.Printf("repunit: v1=%d, b1=%d\n", v1, b1)
				}
			}
		}
	}
}

func repUnit(base, limit uint) []uint {
	var res []uint
	for i := uint(1); i < limit; i = i*base + 1 {
		res = append(res, i)
	}
	return res
}

func memoryBound() {
	const limit = 1000
	var repUnits [][]uint
	repUnits = make([][]uint, limit)

	for i := 2; i < limit; i++ {
		repUnits[i] = repUnit(uint(i), limit)
	}
	sum := uint(0)
	count := 0
	for n := uint(1); n < limit; n++ {
		if check(repUnits, n) {
			sum += n
			count++
		}
	}
	fmt.Printf("sum=%d\n", sum)
	fmt.Printf("count=%d\n", count)
}

func check(repUnits [][]uint, n uint) bool {
	nCount := 0
	for i := 0; i < len(repUnits); i++ {
		//fmt.Printf("repUnits[%d]: %d\n", i, repUnits[i])
		for _, x := range repUnits[i] {
			if n == x {
				nCount++
			}
			if nCount == 2 {
				//fmt.Printf("RepUnit: %d\n", n)
				return true
			}
		}
	}
	return false
}
