package main

import (
	"fmt"
	"gopkg.in/xmlpath.v2"
	"log"
	//"os"
	"http"
	"strings"
)

func main() {
	file, err := os.Open("bbc.html")
	if err != nil {
		log.Fatal(err)
	}

	root, err := xmlpath.ParseHTML(file)
	if err != nil {
		log.Fatal(err)
	}

	toVisit := []string{}
	visited := []string{}

	fmt.Println("Words:", countWords(root))
	toVisit = addLinks(findLinks(root), visited, toVisit)
	fmt.Println("ToVisit:", toVisit)
}

func contains(s string, ss []string) bool {
	for _, x := range ss {
		if x == s {
			return true
		}
	}
	return false
}

func addLinks(links, visited, toVisit []string) []string {
	for _, link := range links {
		link = strings.TrimSpace(link)
		if !strings.HasPrefix(link, "mailto:") &&
			!strings.HasPrefix(link, "whatsapp:") &&
			!strings.HasPrefix(link, "http://") &&
			!strings.HasPrefix(link, "https://") &&
			!strings.HasPrefix(link, "#") &&
			!contains(link, toVisit) &&
			!contains(link, visited) {
			toVisit = append(toVisit, link)
		}
	}
	return toVisit
}

func findLinks(root *xmlpath.Node) []string {
	path := xmlpath.MustCompile("//a/@href")
	result := []string{}
	iter := path.Iter(root)
	for iter.Next() {
		node := iter.Node()
		result = append(result, node.String())
	}
	return result
}

func countWords(root *xmlpath.Node) map[string]int {
	path := xmlpath.MustCompile("//p")
	words := make(map[string]int)
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
