package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
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

func (m *MongoCollection) updateRecords(ts *TS) {
	var lateReason *int
	if ts.LateReason != nil {
		lateReason = &ts.LateReason.Reason
	}

	updates := []Update{}
	for _, loc := range ts.Locations {
		if loc.Pass != nil || loc.Wtp != nil {
			//Ignore passing points
			continue
		}

		jsps := []JourneyStopPoint{}
		iter := m.col.Find(Key(ts.Rid, loc.Tpl)).Iter()
		iter.All(&jsps)
		err := iter.Close()
		failIf(err)
		var jsp JourneyStopPoint
		switch len(jsps) {
		case 0:
			jsp = JourneyStopPoint{}
		case 1:
			jsp = jsps[0]
		default:
			log.Fatalf("Should only get one result for %s/%s.  Got %d.\n %+v`n",
				ts.Rid, loc.Tpl, len(jsps), jsps)
		}
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
		updates = append(updates, Update{id, jsp})
	}
	log.Printf("Updating %d stops for RID: %s.\n", len(updates), ts.Rid)
	for _, upd := range updates {
		log.Printf("%+v\n", upd)
		_, err := m.col.Upsert(Key(upd.Id.Rid, upd.Id.Station), upd.jsp)
		failIf(err)
	}
}

func (m *MongoCollection) SaveStream(feed *NREUpdates) {
	log.Println("Saving stats to Mongo DB running on:", MongoHost)

	for xml := range *feed {
		update := ParsePportXml(xml)
		if ur := update.Ur; ur != nil {
			if ts := ur.Ts; ts != nil {
				log.Println("Updating", ts.Rid)
				m.updateRecords(ts)
			}
		}
	}
}

func Key(rid, station string) bson.M {
	return bson.M{"_id": bson.M{"rid": rid, "station": station}}
}
