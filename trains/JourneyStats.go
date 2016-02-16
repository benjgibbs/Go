package main

import (
	"fmt"
	"time"
)

const HoursMinsSecs = "15:04:05"
const HoursMins = "15:04"

func Find(locs *[]Location, name string) (bool, *Location) {
	for _, loc := range *locs {
		if loc.Tpl == name {
			return true, &loc
		}
	}
	return false, nil
}

type Times struct {
	schedArrive  time.Time
	actualArrive time.Time
	schedDepart  time.Time
	actualDepart time.Time
}

type Stops map[string]*Times

type Journey struct {
	id         string
	stops      Stops
	lateReason *int
}

func PrintJourney(j Journey) {
	fmt.Println("Journey - Rid:", j.id)
	if j.lateReason != nil {
		fmt.Println("Late Reason:", j.lateReason)
	}
	for station, times := range j.stops {
		fmt.Printf("%-10sArr: %s / %s. Dep: %s / %s. %.0f/%.0f seconds late\n", station,
			times.schedArrive.Format(HoursMinsSecs),
			times.actualArrive.Format(HoursMinsSecs),
			times.schedDepart.Format(HoursMinsSecs),
			times.actualDepart.Format(HoursMinsSecs),
			times.actualArrive.Sub(times.schedArrive).Seconds(),
			times.actualDepart.Sub(times.schedDepart).Seconds())
	}
}

func ParseTime(s string) time.Time {
	result, err := time.Parse(HoursMins, s)
	if err != nil {
		result, err = time.Parse(HoursMinsSecs, s)
	}
	failIf(err)
	return result
}

func (j Journey) Update(ts *TS) Journey {
	if ts.LateReason != nil {
		j.lateReason = &ts.LateReason.Reason
	}
	if j.stops == nil {
		j.stops = make(Stops)
	}
	for _, loc := range ts.Locations {
		if loc.Pass != nil || loc.Wtp != nil {
			continue
		}
		if _, ok := j.stops[loc.Tpl]; !ok {
			j.stops[loc.Tpl] = &Times{}
		}
		times := j.stops[loc.Tpl]

		if loc.Pta != nil {
			times.schedArrive = ParseTime(*loc.Pta)
		}
		if loc.Ptd != nil {
			times.schedDepart = ParseTime(*loc.Ptd)
		}
		if loc.Wta != nil {
			times.actualArrive = ParseTime(*loc.Wta)
		}
		if loc.Wtd != nil {
			times.actualDepart = ParseTime(*loc.Wtd)
		}
	}
	return j
}

func CreateJourney(ts *TS) Journey {
	result := Journey{id: ts.Rid}
	result.Update(ts)
	return result
}

func SaveStatsFor(place string, feed NREUpdates) {
	fmt.Println("Saving stats...")
	journies := make(map[string]Journey)
	for xml := range feed {
		update := XmlToStructs(xml)

		if ur := update.Ur; ur != nil {
			if deact := ur.Deactivated; deact != nil {
				if _, ok := journies[deact.Rid]; ok {
					PrintJourney(journies[deact.Rid])
					delete(journies, deact.Rid)
				}
			} else if ts := ur.Ts; ts != nil {
				if _, ok := journies[ts.Rid]; ok {
					journies[ts.Rid] = journies[ts.Rid].Update(ts)
				} else {
					journies[ts.Rid] = CreateJourney(ts)
				}
			}
		}
	}
}
