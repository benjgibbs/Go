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
		codeLen += 2 // add quotes
		for i := 0; i < len(line); i++ {
			switch line[i] {
			case '"':
				codeLen += 2
			case '\\':
				if line[i+1] == 'x' {
					i += 3
					codeLen += 5
				} else if line[i+1] == '\\' {
					i++
					codeLen += 4
				} else {
					codeLen += 2
				}
			default:
				codeLen++
			}
		}
		strLen += len(line)
	}
	fmt.Printf("code length %d, string length: %d, result = %d\n", codeLen, strLen, (codeLen - strLen))
}
