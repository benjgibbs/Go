package main

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
