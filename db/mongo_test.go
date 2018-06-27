package db

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

var (
	m   Mgo
	err error
)

type user struct {
	Name string
	Age  int
}

func Test_MgoDial(t *testing.T) {

	c := Config{
		DBtype: Mongo,
		Host:   "localhost:27017",
		//Account: "admin",
		Tag: "Mongo",
	}

	m, err = mongoDial(c)
	if err != nil {
		spew.Dump(err)
		t.Error(err)
	} else {
		t.Log("Connect Sussess")
	}
}
func Test_GetMgo(t *testing.T) {
	x := user{
		Name: "Admin",
		Age:  13,
	}

	c := m.New().DB("test").C("user")
	err := c.Insert(x)
	if err != nil {
		panic(err)
	}

}

func Test_FindMgo(t *testing.T) {

}
