package main

import (
	"fmt"
	"os"
	"strconv"
)

var houses []int

func main() {
	target, _ := strconv.Atoi(os.Args[1])

	limit := 1 + target/10
	houses = make([]int, limit)

	for e := 1; ; e++ {
		for h := e; h < limit; h += e {
			houses[h] += e * 10
		}
		if houses[e] >= target {
			fmt.Printf("%d has %d presents\n", e, houses[e])
			//fmt.Println(e, houses)
			break
		}
	}
}
