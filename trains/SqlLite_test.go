package main

import (
	"testing"
	"time"
)

func TestCreationOfNotStored(t *testing.T) {
	db := NewDB("./trains.db")
	defer cleanUp(db)
	result := db.getStop("1234", "WAT")
	if result.Id.Rid != "1234" {
		t.Error("Wrong rid", result.Id.Rid)
	}
	if result.Id.Station != "WAT" {
		t.Error("Wrong station", result.Id.Station)
	}
}

func TestDbRoundTrip(t *testing.T) {
	db := NewDB("./trains.db")
	defer cleanUp(db)
	update := Update{
		Id{"1234", "WAT"},
		JourneyStopPoint{"1a", 99,
			time.Date(1, 1, 1, 12, 0, 0, 0, time.UTC),
			time.Date(1, 1, 1, 12, 0, 1, 0, time.UTC),
			time.Date(1, 1, 1, 12, 0, 2, 0, time.UTC),
			time.Date(1, 1, 1, 12, 0, 3, 0, time.UTC)}}

	db.updateStopPoint(&update)

	result := db.getStop("1234", "WAT")

	if result.Id.Rid != "1234" {
		t.Error("Wrong rid", result.Id.Rid)
	}
	if result.Id.Station != "WAT" {
		t.Error("Wrong station", result.Id.Station)
	}

	if result.jsp.Platform != "1a" {
		t.Error("Wrong platform: ", result.jsp.Platform)
	}

	if result.jsp.LateReason != 99 {
		t.Error("Wrong late reason: ", result.jsp.LateReason)
	}

	if result.jsp.SchedArrive != time.Date(1, 1, 1, 12, 0, 0, 0, time.UTC) {
		t.Error("Failed to get the correct arrival time", result)
	}

	if result.jsp.ActualArrive != time.Date(1, 1, 1, 12, 0, 1, 0, time.UTC) {
		t.Error("Failed to get the correct arrival time", result)
	}

	if result.jsp.SchedDepart != time.Date(1, 1, 1, 12, 0, 2, 0, time.UTC) {
		t.Error("Failed to get the correct depart time", result)
	}

	if result.jsp.ActualDepart != time.Date(1, 1, 1, 12, 0, 3, 0, time.UTC) {
		t.Error("Failed to get the correct depart time", result)
	}
}

func TestTrainUpdateIntoSqlite(t *testing.T) {
	db := NewDB("./trains2.db")
	//defer cleanUp(&db)
	updates := ReadFromFile("/Users/ben/Git/GoCode/var/one.dat", 0)
	db.SaveStream(updates)
}

func cleanUp(db *DB) {
	db.dropStops()
}
