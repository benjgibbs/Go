package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

const numPiles = 3

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	parcels := []int{}
	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		num, e := strconv.Atoi(line)
		if e != nil {
			panic(e)
		}
		parcels = append(parcels, num)
		total += num
	}
	target := total / numPiles

	for i :-
}


func prod(pile []int) int {
	res := 1
	for i := 0; i < len(pile); i++ {
		res *= pile[i]
	}
	return res
}

func sum(pile []int) int {
	res := 0
	for _, x := range pile {
		res += x
	}
	return res
}
