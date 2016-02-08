package main

import (
	"gopkg.in/xmlpath.v2"
	"log"
	"strings"
	"sync/atomic"
)

type ParseParams struct {
	depth int
	base  string
	url   string
}

type PageParseQueue chan ParseParams

func parsePage(depth int, base, url string, results CountQueue, queryCount *int64, urlQuery QueryQueue) {
	log.Println("Crawling:", url, "at depth:", depth)
	reader := fromUrl(url)
	root, err := xmlpath.ParseHTML(reader)
	failWith(err)
	words := countWords(root)
	log.Println("Sending:", url)
	results <- words

	if depth < *maxDepth {
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
		link = strings.TrimSpace(link)
		if isInterestingLink(link) {
			result = append(result, base+link)
		} else if strings.HasPrefix(link, base) {
			log.Println("Using full link:", link)
			result = append(result, link)
		} else {
			log.Println("Ignoring:", link)
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
