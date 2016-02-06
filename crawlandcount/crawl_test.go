package main

import (
	"sync"
	"testing"
	"time"
)

func Test_Pool(t *testing.T) {
	queue := make(chan int)
	wg := sync.WaitGroup{}
	for p := 0; p < 3; p++ {
		wg.Add(1)
		go func(p int) {
			for m := range queue {
				t.Log("pool: ", p, " message: ", m)
				time.Sleep(10 * time.Millisecond)
			}
			wg.Done()
		}(p)
	}
	for i := 0; i < 20; i++ {
		queue <- i
	}
	close(queue)
	wg.Wait()
}
