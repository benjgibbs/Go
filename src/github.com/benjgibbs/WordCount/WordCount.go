package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

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

func main() {

	fmt.Println("Starting")

	dat, err := ioutil.ReadFile("warpeace.txt")
	if err != nil {
		panic(err)
	}

	counts := make(map[string]int)
	parts := strings.FieldsFunc(string(dat), func(r rune) bool {
		switch r {
		case ' ', '.', '\n', '\r', ':', '*', ';', '(', ')', '%', '£', '\t', ',',
			'"', '-', '\'', '?', '/', '!':
			return true
		}
		return false
	})

	for i := range parts {
		part := parts[i]
		part = strings.ToLower(part)
		counts[part]++
	}
	frequencies := []WordFreq{}
	wordCount := 0
	for k, v := range counts {
		freq := WordFreq{
			Word: k,
			Freq: v,
		}
		frequencies = append(frequencies, freq)
		wordCount += v
	}
	sort.Sort(ByFreq(frequencies))
	for i := range frequencies {
		freq := frequencies[i]
		fmt.Printf("%s => %d\n", freq.Word, freq.Freq)
	}
	fmt.Printf("Total Words: %d\n", wordCount)
}
