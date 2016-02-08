package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	strLen := 0
	codeLen := 0
	for scanner.Scan() {
		line := scanner.Text()
		for i := 0; i < len(line); i++ {
			switch line[i] {
			case '"':
				//do nothing
			case '\\':
				if line[i+1] == 'x' {
					i += 3
					strLen++
				} else {
					i++
					strLen++
				}
			default:
				strLen++
			}
		}
		codeLen += len(line)
	}
	fmt.Printf("string length: %d, code length %d, result = %d\n", strLen, codeLen, (codeLen - strLen))
}
