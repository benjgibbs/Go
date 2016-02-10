package main

import "log"

func failIf(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
