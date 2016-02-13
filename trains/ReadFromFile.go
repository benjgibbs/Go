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

	fmt.Println("Reading from", fileName)
	scanner := bufio.NewScanner(r)
	go func() {
		for scanner.Scan() {
			line := scanner.Bytes()
			result <- line
			time.Sleep(time.Duration(msBetweenUpdates) * time.Millisecond)
		}
		close(result)
		failIf(scanner.Err())
	}()
	return result
}
