package main

import (
	"log"
)

type WordCounts map[string]int

func collectCounts(results CountQueue) WordCounts {
	words := WordCounts{}
	pageVisits := 0
	for update := range results {
		pageVisits += 1
		words = combine(words, update)
	}
	log.Println("Page visits:", pageVisits)
	return words
}

func combine(a, b WordCounts) WordCounts {
	for k, v := range a {
		b[k] += v
	}
	return b
}
