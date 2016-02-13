package main

import (
	"fmt"
)

func Find(locs *[]Location, name string) (bool, *Location) {
	for _, loc := range *locs {
		if loc.Tpl == name {
			return true, &loc
		}
	}
	return false, nil
}

func SaveStatsFor(place string, feed NREUpdates) {
	fmt.Println("Saving stats...")
	for xml := range feed {
		update := XmlToStructs(xml)
		if found, loc := Find(&update.UR.TS.Locations, place); found {
			fmt.Printf("WINCHESTER: %+v\n", loc)
			fmt.Printf("WINCHESTER: %+v\n", update)
		} else {
			//fmt.Println("No match")
		}
	}
}
