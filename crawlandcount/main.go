package main

import (
	"bytes"
	"flag"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sort"
)

type CountQueue chan WordCounts

var maxDepth *int
var poolSize *int
var site *string

func init() {
	maxDepth = flag.Int("maxdepth", 3, "The number of pages to descend from the front page")
	poolSize = flag.Int("poolsize", 16, "The number of threads to use")
	site = flag.String("site", "http://www.bbc.co.uk", "The site to visit")
	flag.Parse()
	log.Println("maxdepth: ", *maxDepth)
	log.Println("poolsize: ", *poolSize)
	log.Println("site: ", *site)
}

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
	results := make(CountQueue)
	var workers PageParseQueue
	var currentlyCrawlingCount int64 = 1 // Start off at one which is the root page
	urlQuery := createUrlQuery(*site, &currentlyCrawlingCount, &workers)
	workers = createPool(*poolSize, func(a ParseParams) {
		parsePage(a.depth, a.base, a.url, results, &currentlyCrawlingCount, urlQuery)
	})

	workers <- ParseParams{1, *site, *site}
	words := collectCounts(results)
	printTopN(words, 20)
}

func printTopN(words WordCounts, n int) {
	wcl := WordCountList{}
	for word, count := range words {
		wcl = append(wcl, WordCount{word, count})
	}
	sort.Sort(wcl)
	log.Println("Top", n, "words.")
	for i := 0; i < n && i < len(wcl); i++ {
		log.Println(wcl[i])
	}
}
