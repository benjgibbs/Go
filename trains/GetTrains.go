package main

import (
	"flag"
	"gopkg.in/gcfg.v1"
	"log"
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

const (
	WRITE_TO_FILE  = 1
	READ_FROM_FILE = 2
	READ_FROM_NRE  = 3
	CFG_FILE       = "etc/trains.gcfg"
	DAT_FILE       = "var/darwin.dat"
	DAT_FILE_SIZE  = 1000
	MS_GAP         = 1
)

var mode int

func main() {
	cfg := Cfg{}
	err := gcfg.ReadFileInto(&cfg, CFG_FILE)
	failIf(err)
	switch mode {
	case WRITE_TO_FILE:
		log.Printf("Writing to file %s.", DAT_FILE)
		feed := NREFeed{}
		updates := feed.Subscribe(cfg)
		defer feed.Unsubscribe()
		WriteToFile(DAT_FILE_SIZE, DAT_FILE, updates)
	case READ_FROM_FILE:
		log.Printf("Reading from file %s.\n", DAT_FILE)
		updates := ReadFromFile(DAT_FILE, MS_GAP)
		ProcessFeed(updates)
	case READ_FROM_NRE:
		log.Printf("Writing to file %s.", DAT_FILE)
		feed := new(NREFeed)
		updates := feed.Subscribe(cfg)
		//defer feed.Unsubscribe()
		ProcessFeed(updates)
	}
}

func ProcessFeed(feed NREUpdates) {
	db := NewDB("var/trains.db")
	db.SaveStream(feed)
}

func init() {
	read := flag.Bool("read", false, "Read darwin from file")
	write := flag.Bool("write", false, "Subscribe to darwin and write to the file")
	flag.Parse()
	if *read {
		mode = READ_FROM_FILE
	} else if *write {
		mode = WRITE_TO_FILE
	} else {
		log.Println("Reading from NRE")
		mode = READ_FROM_NRE
	}
}
