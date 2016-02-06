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
	"sync"
	"sync/atomic"
)

const MAX_DEPTH = 3
const POOL_SIZE = 16

type CrawlFnArg struct {
	depth   int
	url     string
	workers ArgQueue
}

type Result struct {
	url   string
	words WordCounts
}

type WordCount struct {
	word  string
	count int
}

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

type WordCounts map[string]int
type ResultQueue chan Result
type ArgQueue chan CrawlFnArg
type CrawlFn func(a CrawlFnArg)

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

func crawl(depth int, url string, workers ArgQueue, results ResultQueue, count *int64) {
	log.Println("Crawling:", url, "at depth:", depth)
	reader := fromUrl(url)
	root, err := xmlpath.ParseHTML(reader)
	failWith(err)
	words := countWords(root)
	log.Println("Sending:", url)
	results <- Result{url, words}
	if depth < MAX_DEPTH {
		links := findLinks(root)
		log.Println("Found", len(links), "links in", url)
		for _, link := range links {
			atomic.AddInt64(count, 1)
			go func() {
				workers <- CrawlFnArg{depth + 1, url + link, workers}
			}()
		}
	}
	count2 := atomic.AddInt64(count, -1)
	log.Println("Current count:", count2)
	if count2 == 0 {
		close(results)
		close(workers)
	}
}

func main() {
	var count int64 = 1
	results := make(ResultQueue)

	workers, _ := createPool(POOL_SIZE, func(a CrawlFnArg) {
		crawl(a.depth, a.url, a.workers, results, &count)
	})

	workers <- CrawlFnArg{1, "http://bbc.co.uk/", workers}
	words := collect(results)
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

func collect(results ResultQueue) WordCounts {
	urls := []string{}
	words := WordCounts{}
	uniqueVisits := 0
	pageVisits := 0
	for page := range results {
		log.Println("Update on:", page.url)
		pageVisits += 1
		if !contains(page.url, urls) {
			uniqueVisits += 1
			urls = append(urls, page.url)
			words = combine(words, page.words)
		} else {
			log.Println("Already seen:", page.url)
		}
	}
	log.Println("Unique visits:", uniqueVisits, "Total visits:", pageVisits)
	return words
}

func createPool(sz int, f CrawlFn) (ArgQueue, sync.WaitGroup) {
	queue := make(ArgQueue)
	wg := sync.WaitGroup{}
	for i := 0; i < sz; i++ {
		wg.Add(1)
		go func(i int) {
			for a := range queue {
				log.Println("Starting", a, "on worker", i)
				f(a)
				log.Println("Worker", i, "Finished")
			}
			wg.Done()
		}(i)
	}
	return queue, wg
}

func combine(a, b WordCounts) WordCounts {
	for k, v := range a {
		b[k] += v
	}
	return b
}

func contains(s string, ss []string) bool {
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

func findLinks(root *xmlpath.Node) []string {
	path := xmlpath.MustCompile("//a/@href")
	result := []string{}
	iter := path.Iter(root)
	for iter.Next() {
		node := iter.Node()
		link := node.String()
		if isInterestingLink(link) {
			result = append(result, link)
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
				return !(r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
			})
			words[w] += 1
		}
	}
	return words
}
