package collections

import (
	"container/list"
	"fmt"
	"testing"
)

const iters = 10000000

func TestAppendingToASlice(t *testing.T) {
	var res []int
	for i := 0; i < iters; i++ {
		res = append(res, i)
	}
	sum := 0
	for i := 0; i < iters; i++ {
		sum += res[i]
	}
	fmt.Printf("Sum is: %d\n", sum)
}

func BenchmarkAppendingToASlice(b *testing.B) {
	var res []int
	for i := 0; i < iters; i++ {
		res = append(res, i)
	}
}

func BenchmarkAppendingToAMap(b *testing.B) {
	res := make(map[int]int)
	for i := 0; i < iters; i++ {
		res[i] = i
	}
}

func BenchmarkAppendingToAList(b *testing.B) {
	res := list.New()
	for i := 0; i < iters; i++ {
		res.PushBack(i)
	}
}
