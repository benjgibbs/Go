package main

import (
	"gopkg.in/gcfg.v1"
)

const server = "datafeeds.nationalrail.co.uk"

type Cfg struct {
	Network struct {
		User, Pass string
	}
	Nre struct {
		User, Pass, Queue string
	}
}

func main() {
	cfg := Cfg{}
	err := gcfg.ReadFileInto(&cfg, "etc/trains.gcfg")
	failIf(err)
	fromNre(cfg)
}
