package mongo

import (
	"testing"

	"github.com/globalsign/mgo/bson"
	"github.com/maxjkfc/cocola/db/config"
)

var (
	m   MgoSession
	err error
	u   *user
)

type user struct {
	Id   bson.ObjectId `bson:"_id,omitempty"`
	Name string
	Age  int
}

func init() {
	u = &user{
		Name: "kk",
		Age:  10,
	}

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

	err := m.DB("test", "user").Insert(x)
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

func Benchmark_NewSession(b *testing.B) {
	u.Name = "NewSession"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		x := m.New()
		x.DB("test").C("info").Insert(u)
		defer x.Close()
	}
}

func Benchmark_NewDB(b *testing.B) {
	u.Name = "NewDB"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.DB("test", "info").Insert(u)
	}
}
