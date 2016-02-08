package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sort"
)

const MAX_DEPTH = 3
const POOL_SIZE = 16

type CountQueue chan WordCounts

func failWith(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func fromUrl(url string) io.Reader {
	resp, err := http.Get(url)
	failWith(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	failWith(err)
	return bytes.NewReader(body)
}

func fromFile() io.Reader {
	file, err := os.Open("bbc.html")
	failWith(err)
	return file
}

func main() {
	var count int64 = 1
	results := make(CountQueue)
	var workers PageParseQueue
	url := "http://bbc.co.uk/"
	urlQuery := createUrlQuery(url, &count, &workers)
	workers = createPool(POOL_SIZE, func(a ParseParams) {
		parsePage(a.depth, a.base, a.url, results, &count, urlQuery)
	})

	workers <- ParseParams{1, url, url}
	words := collectCounts(results)
	printTopN(words, 20)
}

func printTopN(words WordCounts, n int) {
	wcl := WordCountList{}
	for word, count := range words {
		wcl = append(wcl, WordCount{word, count})
	}
	sort.Sort(wcl)
	log.Print("Top", n, "words.")
	for i := 0; i < n && i < len(wcl); i++ {
		log.Println(wcl[i])
	}
}
