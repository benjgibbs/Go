package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func ReadFromFile(fileName string, msBetweenUpdates int) NREUpdates {
	f, err := os.Open(fileName)
	failIf(err)
	r := bufio.NewReader(f)
	result := make(NREUpdates)

	delim := byte('\n')
	fmt.Println("Using delim: ", delim)
	go func() {
		for {
			line, err := r.ReadBytes(delim)
			if err != nil || len(line) == 0 {
				break
			}
			result <- line[:len(line)-1]
			time.Sleep(time.Duration(msBetweenUpdates) * time.Millisecond)
		}
		close(result)
	}()
	return result
}
