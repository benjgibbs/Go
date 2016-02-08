package main

import (
	"fmt"
)

func perms(of []string) [][]string {
	if len(of) == 1 {
		return [][]string{of}
	}
	a := of[0]
	b := of[1:]
	res := [][]string{}
	for _, ps := range perms(b) {
		for psi := 0; psi < len(ps)+1; psi++ {
			perm := make([]string, len(ps)+1)
			copy(perm[0:], ps[:psi])
			copy(perm[psi:], []string{a})
			copy(perm[psi+1:], ps[psi:])
			res = append(res, perm)
		}
		fmt.Println("res=", res)
	}
	return res

}

func main() {
	fmt.Println(perms([]string{"a", "b", "c"}))
}
