package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"testing"
)

type Person struct {
	Name  string
	Phone string
}

func TestExmaple(t *testing.T) {
	session, err := mgo.Dial("127.0.0.1")
	failIf(err)
	defer session.Close()

	db := session.DB("trains_test")
	defer db.DropDatabase()
	c := db.C("test_people")
	defer c.DropCollection()

	err = c.Insert(&Person{"Ale", "+55 53 8116 9639"}, &Person{"Cla", "+55 53 8402 8510"})
	failIf(err)

	result := Person{}
	err = c.Find(bson.M{"name": "Ale"}).One(&result)
	failIf(err)
	if result.Phone != "+55 53 8116 9639" {
		t.Error("Wrong phone number for Ale")
	}
}

func TestTrainUpdate(t *testing.T) {

}
