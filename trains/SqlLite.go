package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"time"
)

type DB struct {
	dbName      string
	db          *sql.DB
	checkStmt   *sql.Stmt
	upsertStmt  *sql.Stmt
	getStopStmt *sql.Stmt
}

func NewDB(dbName string) *DB {
	db := createDb(dbName)
	check := `select count(*) from stops where rid like ? and station like ?`
	checkStmt, err := db.Prepare(check)
	failIf(err)

	//TODO: Make this more of an upsert than an insert
	upsert := `insert into stops values(?, ?, ?, ?, ?, ?, ?, ?)`
	upsertStmt, err := db.Prepare(upsert)
	failIf(err)

	getStop := `select 
								platform, 
								late_reason, 
								scheduled_arrival, 
								scheduled_departure, 
								actual_arrival, 
								actual_departure 
							from stops 
							where rid like ? and station like ?`

	getStopStmt, err := db.Prepare(getStop)
	failIf(err)

	return &DB{dbName, db, checkStmt, upsertStmt, getStopStmt}
}

func createDb(name string) *sql.DB {
	log.Println("Creating", name)
	db, err := sql.Open("sqlite3", name)
	failIf(err)

	create := `create table if not exists stops(
					rid text,
					station text,
					platform text,
					late_reason integer,
					scheduled_arrival timestamp,
					scheduled_departure timestamp,
					actual_arrival timestamp,
					actual_departure timestamp,
					primary key(rid, station))`

	_, err = db.Exec(create)
	failIf(err)
	return db
}

func (db *DB) updateStopPoint(update *Update) {
	db.upsertStmt.Exec(
		update.Id.Rid,
		update.Id.Station,
		update.jsp.Platform,
		update.jsp.LateReason,
		update.jsp.SchedArrive,
		update.jsp.SchedDepart,
		update.jsp.ActualArrive,
		update.jsp.ActualDepart,
	)
}

func (db *DB) dropStops() {
	drop := `drop table if exists stops`
	_, err := db.db.Exec(drop)
	failIf(err)
}

func (db *DB) getStop(rid, station string) *Update {
	result := Update{}
	result.Id = Id{rid, station}

	row := db.getStopStmt.QueryRow(rid, station)

	if row == nil {
		log.Printf("Could not find %s/%s so creating a new update",
			rid, station)
		return &result
	}

	var platform string
	var lateReason int
	var sa, sd, aa, ad time.Time
	err := row.Scan(&platform, &lateReason, &sa, &sd, &aa, &ad)
	if err != nil {
		log.Printf("Could not scan %s/%s so creating a new update: %s",
			rid, station, err)
		return &result
	}
	log.Println("Got Dates:", sa, aa, sd, ad)
	result.jsp = JourneyStopPoint{
		platform, lateReason, sa, aa, sd, ad}

	return &result
}

func (db *DB) updateRecords(ts *TS) {
	var lateReason *int
	if ts.LateReason != nil {
		lateReason = &ts.LateReason.Reason
	}
	for _, loc := range ts.Locations {
		jsp := new(JourneyStopPoint)
		if lateReason != nil {
			jsp.LateReason = *lateReason
		}
		if loc.Plat != nil {
			jsp.Platform = loc.Plat.Plat
		}
		if loc.Pta != nil {
			jsp.SchedArrive = ParseTime(*loc.Pta)
		}
		if loc.Ptd != nil {
			jsp.SchedDepart = ParseTime(*loc.Ptd)
		}
		if loc.Wta != nil {
			jsp.ActualArrive = ParseTime(*loc.Wta)
		}
		if loc.Wtd != nil {
			jsp.ActualDepart = ParseTime(*loc.Wtd)
		}
		id := Id{ts.Rid, loc.Tpl}
		db.updateStopPoint(&Update{id, *jsp})
	}
}

func (db *DB) SaveStream(feed NREUpdates) {
	log.Println("Saving stats to Sqlite running on:", MongoHost)

	for xml := range feed {
		update := ParsePportXml(xml)
		if ur := update.Ur; ur != nil {
			if ts := ur.Ts; ts != nil {
				log.Println("Updating", ts.Rid)
				db.updateRecords(ts)
			}
		}
	}
}
