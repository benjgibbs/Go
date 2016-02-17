package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func find(cs []int, total int) (bool, [][]int) {
	if total == 0 {
		return true, [][]int{}
	}
	if total < 0 {
		return false, [][]int{}
	}
	res := [][]int{}
	for i := 0; i < len(cs); i++ {
		c := cs[i]
		if t, subRes := find(cs[:i], (total - c)); t {
			if len(subRes) == 0 && c == total {
				res = append(res, []int{c})
			} else {
				for j := 0; j < len(subRes); j++ {
					cp := append(subRes[j], c)
					res = append(res, cp)
				}
			}
		}
	}
	return true, res
}

func main() {
	total, _ := strconv.Atoi(os.Args[1])
	containers := []int{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		containers = append(containers, num)
	}
	fmt.Println(containers)
	_, b := find(containers, total)
	for _, bi := range b {
		fmt.Println(bi)
	}
	fmt.Println(len(b))
}
