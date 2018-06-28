package db

import (
	"testing"
)

func Test_MgoConnect(t *testing.T) {
	c := Config{
		DBtype: Mongo,
		Host:   "localhost:27017",
		//Account: "admin",
		Tag: "Mongo_Pool",
	}

	if err := Pools.NewConnect(c); err != nil {
		t.Error(err)
	}
	if m, err := Pools.Mgo(c.Tag); err != nil {
		t.Error("Not Find the Database")
	} else {
		t.Log("Database Name:", m.New().DB("test").Name)
	}

}
