package main

import (
	"log"
)

type WorkerFunc func(a ParseParams)

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
