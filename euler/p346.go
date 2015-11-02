package main

import (
	"fmt"
)

func isBaseNRepUnit(n, num uint) bool {
	//fmt.Printf("num=%d\n", num)
	cont := num - 1
	if cont == 0 {
		return true
	}
	if cont%n > 0 {
		//fmt.Printf("Fail: %d\n", cont%6)
		return false
	}
	return isBaseNRepUnit(n, cont/n)
}

func repUnit(base, limit uint) []uint {
	var res []uint
	for i := 1; i < limit; i = i + base + 1 {
		res = append(res, i)
	}
	return res
}

func p346() {
	var num uint = 0
	for i := uint(0); i < 8; i++ {
		num = num | 1<<i
		for j := uint(3); j < num; j++ {
			if isBaseNRepUnit(j, num) {
				fmt.Println("Strong repunit: ", num)
			}
		}
	}

}
