package mongo

import (
	"testing"

	"github.com/globalsign/mgo/bson"
	"github.com/maxjkfc/cocola/db/config"
)

var (
	m   Mgo
	err error
)

type user struct {
	Id   bson.ObjectId `bson:"_id,omitempty"`
	Name string
	Age  int
}

func Test_MgoDial(t *testing.T) {

	c := config.Config{
		DBtype: "Mongo",
		Host:   "localhost:27017",
		//Account: "admin",
		Tag: "Mongo",
	}

	m, err = Dial(c)
	if err != nil {
		t.Error(err)
	} else {
		t.Log("Connect Sussess")
	}
}
func Test_GetMgo(t *testing.T) {
	x := user{
		Id:   bson.NewObjectId(),
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
	c := m.New().DB("test").C("user")
	u := new(user)
	q := c.Find(bson.M{"name": "Admin"})
	if err := q.One(u); err != nil {
		t.Error(err)
	} else {
		t.Log("Success")
	}

}

func Test_Status(t *testing.T) {
	t.Log(m.Status())

}
