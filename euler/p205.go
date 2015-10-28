package main

import (
	"fmt"
	"math/rand"
)

func p205() {
	// p can score from 9 to 36
	tot := 0
	pScoresSum := make(map[int]int)
	for p1 := 1; p1 <= 4; p1++ {
		for p2 := 1; p2 <= 4; p2++ {
			for p3 := 1; p3 <= 4; p3++ {
				for p4 := 1; p4 <= 4; p4++ {
					for p5 := 1; p5 <= 4; p5++ {
						for p6 := 1; p6 <= 4; p6++ {
							for p7 := 1; p7 <= 4; p7++ {
								for p8 := 1; p8 <= 4; p8++ {
									for p9 := 1; p9 <= 4; p9++ {
										idx := p1 + p2 + p3 + p4 + p5 + p6 + p7 + p8 + p9
										pScoresSum[idx] = pScoresSum[idx] + 1
										tot += 1
									}
								}
							}
						}
					}
				}
			}
		}
	}
	pScoresProb := make(map[int]float64)
	for id, cnt := range pScoresSum {
		pScoresProb[id] = float64(cnt) / float64(tot)
	}
	tot = 0
	cScoresTot := make(map[int]int)
	for c1 := 1; c1 <= 6; c1++ {
		for c2 := 1; c2 <= 6; c2++ {
			for c3 := 1; c3 <= 6; c3++ {
				for c4 := 1; c4 <= 6; c4++ {
					for c5 := 1; c5 <= 6; c5++ {
						for c6 := 1; c6 <= 6; c6++ {
							idx := c1 + c2 + c3 + c4 + c5 + c6
							cScoresTot[idx] = cScoresTot[idx] + 1
							tot++
						}
					}
				}
			}
		}
	}
	cScoresProb := make(map[int]float64)
	for id, cnt := range cScoresTot {
		cScoresProb[id] = float64(cnt) / float64(tot)
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
