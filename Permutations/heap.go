package permutations

func Heap(xs []int) [][]int {
	var result [][]int
	heap(len(xs), xs, &result)
	return result
}

func heap(n int, xs []int, res *[][]int) {
	if n == 1 {
		out := make([]int, len(xs))
		copy(out, xs)
		*res = append(*res, out)
	} else {
		for i := 0; i < n-1; i += 1 {
			heap(n-1, xs, res)
			if n%2 == 0 {
				xs[i], xs[n-1] = xs[n-1], xs[i]
			} else {
				xs[0], xs[n-1] = xs[n-1], xs[0]
			}
		}
		heap(n-1, xs, res)
	}
}
