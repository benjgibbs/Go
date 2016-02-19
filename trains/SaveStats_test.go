package main

import (
	"gopkg.in/mgo.v2/bson"
	"testing"
)

type Person struct {
	Name  string
	Phone string
}

func TestExmaple(t *testing.T) {

	m := Connect("trains_test", "test_people")
	//	defer m.Close()
	//	defer m.db.DropDatabase()
	//	defer m.col.DropCollection()

	err := m.col.Insert(&Person{"Ale", "+55 53 8116 9639"}, &Person{"Cla", "+55 53 8402 8510"})
	failIf(err)

	result := Person{}
	err = m.col.Find(bson.M{"name": "Ale"}).One(&result)
	failIf(err)
	if result.Phone != "+55 53 8116 9639" {
		t.Error("Wrong phone number for Ale")
	}
}

func TestTrainUpdate(t *testing.T) {
	m := Connect("trains_test", "test_journies")
	defer m.Close()
	m.db.DropDatabase()
	m.col.DropCollection()
	updates := ReadFromFile("/Users/ben/Git/GoCode/var/one.dat", 0)
	m.SaveStream(updates)
	result := JourneyStopPoint{}
	err := m.col.Find(Key("201602142632334", "ELGH")).One(&result)
	if err != nil {
		t.Error("Got error: ", err)
	}
}
