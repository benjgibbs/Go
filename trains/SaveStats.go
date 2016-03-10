package main

import (
	"time"
)

type Id struct {
	Rid     string
	Station string
}

type JourneyStopPoint struct {
	Platform     string
	LateReason   int
	SchedArrive  time.Time
	ActualArrive time.Time
	SchedDepart  time.Time
	ActualDepart time.Time
}

type Update struct {
	Id  Id
	jsp JourneyStopPoint
}
