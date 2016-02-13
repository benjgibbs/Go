package main

import (
	"fmt"
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

func Find(locs *[]Location, name string) (bool, *Location) {
	for _, loc := range *locs {
		if loc.Tpl == name {
			return true, &loc
		}
	}
	return false, nil
}

func main() {
	cfg := Cfg{}
	err := gcfg.ReadFileInto(&cfg, "etc/trains.gcfg")
	failIf(err)
	feed := NREFeed{}
	updates := feed.Subscribe(cfg)
	defer feed.Unsubscribe()
	for update := range updates {
		if found, loc := Find(&update.UR.TS.Locations, "WCHSTR"); found {
			fmt.Printf("WINCHESTER: %+v\n", loc)
			fmt.Printf("WINCHESTER: %+v\n", update)
		} else if found, loc := Find(&update.UR.TS.Locations, "WATRLMN"); found {
			fmt.Printf("WATERLOO: %+v\n", loc)
			fmt.Printf("WATERLOO: %+v\n", update)
		}
	}
}
