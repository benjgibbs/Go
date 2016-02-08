package main

import (
	"log"
	"sync/atomic"
)

type UrlQuery struct {
	urls  []string
	depth int
}

type QueryQueue chan UrlQuery

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

func contains(ss []string, s string) bool {
	for _, x := range ss {
		if x == s {
			return true
		}
	}
	return false
}
