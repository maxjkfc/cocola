package mongo

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/maxjkfc/cocola/errors"
)

type MgoCmd interface {
	Insert(args interface{}) error
	Update(condition interface{}, args interface{}) error
	UpdateByID(id string, args interface{}) error
	UpdateAll(condition interface{}, args interface{}) (*mgo.ChangeInfo, error)
	Upsert(condition interface{}, args interface{}) (*mgo.ChangeInfo, error)
	FindByID(id string) Query
	Find(interface{}) Query
	Delete(condition interface{}) error
	DeleteAll(condition interface{}) (*mgo.ChangeInfo, error)
	DeleteByID(id string) error
}

type mgodb struct {
	s    *mgo.Session
	c    *mgo.Collection
	q    *mgo.Query
	err  error
	errs errors.Error
}

func newMgo(session *mgo.Session, db, c string) MgoCmd {
	return &mgodb{
		s: session,
		c: session.DB(db).C(c),
	}
}

// Insert - 寫入資料
func (m *mgodb) Insert(args interface{}) error {
	defer m.close()
	return m.c.Insert(args)
}

// UpdateByID - 更新資料
func (m *mgodb) Update(condition interface{}, args interface{}) error {
	defer m.close()
	return m.c.Update(condition, args)
}

// UpdateByID - 更新資料
func (m *mgodb) UpdateByID(id string, args interface{}) error {
	defer m.close()
	return m.c.UpdateId(bson.ObjectIdHex(id), args)
}

// UpdateByID - 更新資料
func (m *mgodb) UpdateAll(condition interface{}, args interface{}) (*mgo.ChangeInfo, error) {
	defer m.close()
	return m.c.UpdateAll(condition, args)
}

// Upsert
func (m *mgodb) Upsert(condition interface{}, args interface{}) (*mgo.ChangeInfo, error) {
	defer m.close()
	return m.c.Upsert(condition, args)
}

// FindByID - 查詢資料
func (m *mgodb) FindByID(id string) Query {
	m.q = m.c.FindId(bson.ObjectIdHex(id))
	return m
}

// MgoQuery -
func (m *mgodb) Find(condition interface{}) Query {
	m.q = m.c.Find(condition)
	return m
}

// Delete -
func (m *mgodb) Delete(condition interface{}) error {
	defer m.close()
	return m.c.Remove(condition)
}

// DeleteAll - Delete all the document from the specific condition
func (m *mgodb) DeleteAll(condition interface{}) (*mgo.ChangeInfo, error) {
	defer m.close()
	return m.c.RemoveAll(condition)
}

// DeleteByID - Delete the document by ID
func (m *mgodb) DeleteByID(id string) error {
	defer m.close()
	return m.c.RemoveId(bson.ObjectIdHex(id))
}

func (m *mgodb) close() {
	m.s.Close()
}
