package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"time"
)

const (
	MongoHost = "127.0.0.1"
)

type MongoCollection struct {
	session *mgo.Session
	db      *mgo.Database
	col     *mgo.Collection
}

func Connect(db, col string) *MongoCollection {
	m := MongoCollection{}
	var err error
	m.session, err = mgo.Dial(MongoHost)
	failIf(err)
	m.db = m.session.DB(db)
	m.col = m.db.C(col)
	return &m
}

func (m *MongoCollection) Close() {
	m.session.Close()
}

type JourneyStopPoint struct {
	rid, station, platform    string
	lateReason                *int
	schedArrive, actualArrive time.Time
	schedDepart, actualDepart time.Time
}

func (m *MongoCollection) updateRecords(ts *TS) {
	var lateReason *int
	if ts.LateReason != nil {
		lateReason = &ts.LateReason.Reason
	}

	updates := []JourneyStopPoint{}
	for _, loc := range ts.Locations {
		if loc.Pass != nil || loc.Wtp != nil {
			//Ignore passing points
			continue
		}

		jsps := []JourneyStopPoint{}
		iter := m.col.Find(bson.M{"rid": ts.Rid, "station": loc.Tpl}).Iter()
		iter.All(&jsps)
		err := iter.Close()
		failIf(err)
		var jsp JourneyStopPoint
		switch len(jsps) {
		case 0:
			jsp = JourneyStopPoint{rid: ts.Rid, station: loc.Tpl}
		case 1:
			jsp = jsps[0]
		default:
			log.Fatal("Should only get one result for", ts.Rid, loc.Tpl,
				"Got", len(jsps))
		}
		if lateReason != nil {
			jsp.lateReason = lateReason
		}
		if loc.Plat != nil {
			jsp.platform = loc.Plat.Plat
		}
		if loc.Pta != nil {
			jsp.schedArrive = ParseTime(*loc.Pta)
		}
		if loc.Ptd != nil {
			jsp.schedDepart = ParseTime(*loc.Ptd)
		}
		if loc.Wta != nil {
			jsp.actualArrive = ParseTime(*loc.Wta)
		}
		if loc.Wtd != nil {
			jsp.actualDepart = ParseTime(*loc.Wtd)
		}
		updates = append(updates, jsp)
	}
	log.Printf("Updating %d stops for RID: %s.\n", len(updates), ts.Rid)
	err := m.col.Insert(updates)
	failIf(err)
}

func (m *MongoCollection) SaveStream(feed NREUpdates) {
	log.Println("Saving stats to Mongo DB running on:", MongoHost)

	for xml := range feed {
		update := ParsePportXml(xml)
		if ur := update.Ur; ur != nil {
			if ts := ur.Ts; ts != nil {
				log.Println("Updating", ts.Rid)
				m.updateRecords(ts)
			}
		}
	}

}
