package db

import (
	"testing"

	"github.com/maxjkfc/cocola/db/config"
)

func Test_MgoConnect(t *testing.T) {
	c := config.Config{
		DBtype: Mongo_Driver,
		Host:   "localhost:27017",
		//Account: "admin",
		Tag: "m",
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

func Test_RedisConnect(t *testing.T) {
	c := config.Config{
		DBtype: Redis_Driver,
		Host:   "localhost:6379",
		Tag:    "r",
	}
	if err := Pools.NewConnect(c); err != nil {
		t.Error(err)
	}
	if r, err := Pools.Redis(c.Tag); err != nil {
		t.Error("Not Find the Database")
	} else {
		t.Log("Database Name:", r.New().Do("PING"))
	}
}

func Test_Statux(t *testing.T) {
}
