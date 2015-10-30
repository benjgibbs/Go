package main

import (
	"fmt"
)

/*
	p1 can do anything - p2 has to use perfect strategy
	p1(x,x,x) is the state that p1 leaves the piles in
	X(1,2,3)
		p1(1,2,2), p2(0,2,2), p1(0,0,2), p2(0,0,0) p1 loses
		p1(1,2,1), p2(1,0,1), p1(0,0,1), p2(0,0,0) p1 loses
		p1(1,2,0), p2(1,1,0), p1(1,0,0), p2(0,0,0) p1 loses
		p1(1,1,3), p2(1,1,0), p1(1,0,0), p2(0,0,0) p1 loses
		p1(1,0,3), p2(1,0,1), p1(1,0,0), p2(0,0,0) p1 loses
		p1(0,2,3), p2(0,2,2), p1(0,2,0), p2(0,0,0) p1 loses
		p1(0,2,3), p2(0,2,2), p1(0,2,1), p2(0,1,1), p1(0,0,1), p2(0,0,0) p1 loses

	X(2,4,6)
		p1(2,4,5), p2(2,4,1), p1(2,4,0), p2(2,2,0), p1(2,0,0), p2(0,0,0) p1 loses
		p1(2,2,6), p2(2,2,0), p1(2,0,0), p2(0,0,0) p1 loses
		p1(2,3,6), p2(2,2,6), p1(2,0,0), p2(0,0,0) p1 loses

	Working backwards
		- want two even piles
		- if there are 2 piles of differing sizes we can reduce one to the same size
		- if there are 3 piles, 2 of same size and one other we can remove the odd one
		- if there are 3 piles, all of different sizes,
				if we reduce one pile to the size of another our openent can can then remove the odd one and win
				if we reduce one pile to a size different from either of the others then our openent is in our position
				we want to force our openent to leave two piles the same size, so we make a difference of one

		given (1,3,4) => (1,2,4), (1,3,2) but now openent can (1,2,3), (
			from 1,2,3 = if we reduce 3, or 2 our oponent wins
			so 0,2,3 - openent plays 0, 1, 3, we play 0, 1, 2, oponent plays 0,0,2, we win

			If we (1) are given (1,1,2):
				Give 2(1,1,0),1(1,0,0),2(0,0,0) - 2  loses
				Give 2(1,0,2),1(1,0,1),2(1,0,0),1(0,0,0) 1 loses
				Give 2(1,0,2),1(1,0,0),2(0,0,0) - 2 loses

		(1,3,5)
			p(1,3,0) q(1,2,0), p(1,1,0) => p wins
			p(1,3,0) q(1,1,0) p(
*/

func numZeros(xs []uint64) int {
	sum := 0
	for _, x := range xs {
		if x == 0 {
			sum++
		}
	}
	return sum
}

func numEvenPiles(xs []uint64) int {
	numEvens := 0
	for _, x := range xs {
		if x%2 == 0 {
			numEvens++
		}
	}
	return numEvens
}

func x(xs []uint64) int {
	switch numZeros(xs) {
	case 3:
		return 0 // LOST - shouldn't get here really
	case 2:
		return 1 // WIN - can remove the whole last pile and win
	case 1:
		if numEvenPiles(xs) == 2 {
			// LOST - we eiter make one pile odd => oppenent wins
			// 				or we remove the whole pile => lost
			return 0
		} else {
			return 1
		}
	case 0:
		switch numEvenPiles(xs) {
		case 0, 2, 3:
			return 1
		}
	}
	return 0
}

func p301() {
	fmt.Println("Hello from 301")

}
