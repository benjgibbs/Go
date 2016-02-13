package main

import (
	"bufio"
	"log"
	"os"
)

func WriteToFile(numRecords int, fileName string, updates NREUpdates) {
	f, err := os.Create(fileName)
	failIf(err)
	w := bufio.NewWriter(f)
	defer w.Flush()

	for update := range updates {
		if numRecords == 0 {
			break
		}
		numRecords--
		n, err := w.Write(update)
		failIf(err)
		log.Printf("Wrote %d bytes to %s. Records left %d", n, fileName, numRecords)
		w.Write([]byte("\n"))
	}
}
