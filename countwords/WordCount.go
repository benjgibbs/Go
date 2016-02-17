package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
	"time"
)

const THRESHOLD int = 10000
const NUM_THREADS int = 4

type WordFreqs struct {
	Freqs map[string]int
	Count int
}

func main() {
	fmt.Println("ReadingFile...")
	dat, err := ioutil.ReadFile("/Users/ben/Books/warpeace.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Naive...")
	start := time.Now()
	freqs := naive(dat)
	fmt.Printf("Took %d ms.\n", time.Now().Sub(start)/1000000)
	printResult(freqs)

	fmt.Println("Parallel...")
	start = time.Now()
	freqs = parallel(dat)
	fmt.Printf("Took %d ms.\n", time.Now().Sub(start)/1000000)
	printResult(freqs)
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func parallel(dat []byte) WordFreqs {
	messages := make(chan WordFreqs)
	chunk := len(dat) / NUM_THREADS
	for i := 0; i < NUM_THREADS; i++ {
		start := i * chunk
		end := min(start+chunk, len(dat))
		go func() { messages <- naive(dat[start:end]) }()
	}

	var result WordFreqs
	result.Freqs = make(map[string]int)
	for i := 0; i < NUM_THREADS; i++ {
		part := <-messages
		for k, v := range part.Freqs {
			result.Freqs[k] += v
		}
		result.Count += part.Count
	}
	return result
}

func naive(dat []byte) WordFreqs {
	parts := strings.FieldsFunc(string(dat), func(r rune) bool {
		switch r {
		case ' ', '.', '\n', '\r', ':', '*', ';', '(', ')', '%', '£', '\t', ',',
			'"', '-', '\'', '?', '/', '!':
			return true
		}
		return false
	})
	counts := make(map[string]int)
	for i := range parts {
		part := parts[i]
		part = strings.ToLower(part)
		counts[part]++
	}
	wordCount := 0
	for _, v := range counts {
		wordCount += v
	}
	return WordFreqs{Freqs: counts, Count: wordCount}
}

type WordFreq struct {
	Word string
	Freq int
}

type ByFreq []WordFreq

func (wf ByFreq) Len() int {
	return len(wf)
}

func (wf ByFreq) Swap(i, j int) {
	wf[i], wf[j] = wf[j], wf[i]
}

func (wf ByFreq) Less(i, j int) bool {
	return wf[i].Freq < wf[j].Freq
}

func printResult(words WordFreqs) {
	freqs := make([]WordFreq, len(words.Freqs))
	i := 0
	for k, v := range words.Freqs {
		freqs[i] = WordFreq{k, v}
		i++
	}
	sort.Sort(ByFreq(freqs))
	for i := range freqs {
		freq := freqs[i]
		if freq.Freq > THRESHOLD {
			fmt.Printf("%s => %d\n", freq.Word, freq.Freq)
		}
	}
	fmt.Printf("Total Words: %d.\n", words.Count)
}
