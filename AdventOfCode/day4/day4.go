package main

import (
	"crypto/md5"
	"fmt"
	"os"
	"strconv"
)

func main() {
	key := os.Args[1]
	for num := 0; ; num++ {
		sum := key + strconv.Itoa(num)
		out := md5.Sum([]byte(sum))
		if 0 == out[0] && 0 == out[1] && 0 == (out[2]&0xF0) {
			md5string := fmt.Sprintf("%x", out)
			fmt.Printf("\n5 num: %d, result: %s, %d\n", num, md5string, out)
		}
		if 0 == out[0] && 0 == out[1] && 0 == (out[2]&0xFF) {
			md5string := fmt.Sprintf("%x", out)
			fmt.Printf("\n6 num: %d, result: %s, %d\n", num, md5string, out)
			return
		}
	}
}
