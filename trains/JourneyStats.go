package main

import (
	"fmt"
	"time"
)

func Find(locs *[]Location, name string) (bool, *Location) {
	for _, loc := range *locs {
		if loc.Tpl == name {
			return true, &loc
		}
	}
	return false, nil
}

type Journey struct {
	id           string
	src, dest    string
	lateReason   int
	schedDepart  time.Time
	actualDepart time.Time
	schedArrive  time.Time
	actualArrive time.Time
}

func SaveStatsFor(place string, feed NREUpdates) {
	fmt.Println("Saving stats...")

	for xml := range feed {
		update := XmlToStructs(xml)
		locations := &update.UR.TS.Locations
		if found, loc := Find(locations, place); found {
			fmt.Printf("%s -> %s -> %s\n", (*locations)[0].Tpl,
				loc.Tpl, (*locations)[len(*locations)-1].Tpl)
		} else {
			//fmt.Println("No match")
		}
	}
}
