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

	if err := NewConnect(c); err != nil {
		t.Error(err)
	}

	if Mgo() == nil {

		t.Error("Not Find the Database")
	} else {

		t.Log("Status:", Mgo().Status())
	}

}

func Test_RedisConnect(t *testing.T) {
	c := config.Config{
		DBtype: Redis_Driver,
		Host:   "localhost:6379",
		Tag:    "r",
	}
	if err := NewConnect(c); err != nil {
		t.Error(err)
	}
	if r, err := Redis(c.Tag); err != nil {
		t.Error("Not Find the Database")
	} else {
		t.Log("Status:", r.Status())
	}
}
