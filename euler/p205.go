package main

import (
	"fmt"
	"math"
	"math/rand"
)

func throw(numSides, n, cnt int, acc map[int]int) {
	if n == 0 {
		acc[cnt] = acc[cnt] + 1
	} else {
		for i := 1; i <= numSides; i++ {
			throw(numSides, n-1, cnt+i, acc)
		}
	}
}

func p205() {
	// p can score from 9 to 36
	pScoresSum := make(map[int]int)
	throw(4, 9, 0, pScoresSum)
	pScoresProb := make(map[int]float64)
	pTot := math.Pow(4.0, 9.0)
	for id, cnt := range pScoresSum {
		pScoresProb[id] = float64(cnt) / pTot
	}
	cScoresSum := make(map[int]int)
	throw(6, 6, 0, cScoresSum)
	cTot := math.Pow(6.0, 6.0)
	cScoresProb := make(map[int]float64)
	for id, cnt := range cScoresSum {
		cScoresProb[id] = float64(cnt) / cTot
	}

	// P(p wins)
	// = P( p has x ) * P( c has less than x)
	// = P(p has x) * (P(c has x-1) + P(c has x-2) + ... + P(c has 1))

	pWinProb := 0.0
	for pScore := 9; pScore <= 36; pScore++ {
		probCLower := 0.0
		for cScore := 6; cScore < pScore; cScore++ {
			probCLower += cScoresProb[cScore]
		}
		pWinProb += pScoresProb[pScore] * probCLower
	}

	//bruteForce()
	fmt.Printf("Prob P Wins %.7f\n", pWinProb)
}

func bruteForce() {
	const samples = 100000
	pWins := 0
	for i := 0; i < samples; i++ {
		pTot := 0
		for p := 0; p < 9; p++ {
			pTot += (rand.Intn(4) + 1)
		}
		cTot := 0
		for c := 0; c < 6; c++ {
			cTot += (rand.Intn(6) + 1)
		}
		if pTot > cTot {
			pWins++
		}
	}
	fmt.Printf("By Force: Prob P wins: %.7f\n", float64(pWins)/float64(samples))
}
