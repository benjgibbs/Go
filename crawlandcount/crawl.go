package main

import (
	"bytes"
	"gopkg.in/xmlpath.v2"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"
	"sync/atomic"
)

const MAX_DEPTH = 3
const POOL_SIZE = 16

type UrlQuery struct {
	urls  []string
	depth int
}

type ParseParams struct {
	depth int
	base  string
	url   string
}

type WordCount struct {
	word  string
	count int
}

type WordCounts map[string]int

type CountQueue chan WordCounts
type PageParseQueue chan ParseParams
type QueryQueue chan UrlQuery

type WorkerFunc func(a ParseParams)

type WordCountList []WordCount

func (wcl WordCountList) Len() int {
	return len(wcl)
}

func (wcl WordCountList) Less(i, j int) bool {
	return wcl[i].count > wcl[j].count
}

func (wcl WordCountList) Swap(i, j int) {
	wcl[i], wcl[j] = wcl[j], wcl[i]
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

func parsePage(depth int, base, url string, results CountQueue, queryCount *int64, urlQuery QueryQueue) {
	log.Println("Crawling:", url, "at depth:", depth)
	reader := fromUrl(url)
	root, err := xmlpath.ParseHTML(reader)
	failWith(err)
	words := countWords(root)
	log.Println("Sending:", url)
	results <- words

	if depth < MAX_DEPTH {
		links := findLinks(base, root)
		numLinks := int64(len(links))
		log.Println("Found", numLinks, "links in", url)
		atomic.AddInt64(queryCount, numLinks)
		urlQuery <- UrlQuery{links, depth}
	}
	count := atomic.AddInt64(queryCount, -1)
	log.Println("Current count:", count)
	if count == 0 {
		close(results)
	}
}

func createUrlQuery(base string, queryCount *int64, workers *PageParseQueue) QueryQueue {
	urlq := make(QueryQueue)
	go func() {
		seenUrls := []string{}
		for q := range urlq {
			for _, url := range q.urls {
				if contains(seenUrls, url) {
					log.Println("Skipping", url, " Already exists")
					atomic.AddInt64(queryCount, -1)
				} else {
					seenUrls = append(seenUrls, url)
					go func() {
						*workers <- ParseParams{q.depth + 1, base, url}
					}()
				}
			}
		}
	}()
	return urlq
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
	words := collectCounts(results, urlQuery)
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

func collectCounts(results CountQueue, urlQuery QueryQueue) WordCounts {
	words := WordCounts{}
	pageVisits := 0
	for update := range results {
		pageVisits += 1
		words = combine(words, update)
	}
	log.Println("Page visits:", pageVisits)
	return words
}

func createPool(sz int, f WorkerFunc) PageParseQueue {
	queue := make(PageParseQueue)
	for i := 0; i < sz; i++ {
		go func(i int) {
			for a := range queue {
				log.Println("Starting", a, "on worker", i)
				f(a)
				log.Println("Worker", i, "Finished")
			}
		}(i)
	}
	return queue
}

func combine(a, b WordCounts) WordCounts {
	for k, v := range a {
		b[k] += v
	}
	return b
}

func contains(ss []string, s string) bool {
	for _, x := range ss {
		if x == s {
			return true
		}
	}
	return false
}

func isInterestingLink(link string) bool {
	return !strings.HasPrefix(link, "mailto:") &&
		!strings.HasPrefix(link, "whatsapp:") &&
		!strings.HasPrefix(link, "http://") &&
		!strings.HasPrefix(link, "https://") &&
		!strings.HasPrefix(link, "#")
}

func findLinks(base string, root *xmlpath.Node) []string {
	path := xmlpath.MustCompile("//a/@href")
	result := []string{}
	iter := path.Iter(root)
	for iter.Next() {
		node := iter.Node()
		link := node.String()
		if isInterestingLink(link) {
			result = append(result, base+link)
		}
	}
	return result
}

func countWords(root *xmlpath.Node) WordCounts {
	path := xmlpath.MustCompile("//p")
	words := make(WordCounts)
	iter := path.Iter(root)
	for iter.Next() {
		node := iter.Node()
		for _, w := range strings.Split(node.String(), " ") {
			w = strings.ToLower(w)
			w = strings.TrimFunc(w, func(r rune) bool {
				return !(r >= 'a' && r <= 'z')
			})
			if len(w) > 0 {
				words[w] += 1
			}
		}
	}
	return words
}
