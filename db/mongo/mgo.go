package mongo

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/maxjkfc/cocola/errors"
)

type MgoCmd interface {
	Insert(args interface{}) error
	Update(selector interface{}, args interface{}) error
	UpdateByID(id string, args interface{}) error
	UpdateAll(selector interface{}, args interface{}) (*mgo.ChangeInfo, error)
	Upsert(selector interface{}, args interface{}) (*mgo.ChangeInfo, error)
	FindByID(id string, value interface{}) error
	Find(selector interface{}, value interface{}) error
	FindAll(selector interface{}, value interface{}) error
	FindAllBySort(selector interface{}, value interface{}, sort ...string) error
	FindSkipLimit(selector, value interface{}, skip, limit int, sort ...string) error
	Delete(selector interface{}) error
	DeleteAll(selector interface{}) (*mgo.ChangeInfo, error)
	DeleteByID(id string) error
	Count(selector interface{}) (int, error)
}

type mgodb struct {
	s    *mgo.Session
	c    *mgo.Collection
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
func (m *mgodb) Update(selector interface{}, args interface{}) error {
	defer m.close()
	return m.c.Update(selector, args)
}

// UpdateByID - 更新資料
func (m *mgodb) UpdateByID(id string, args interface{}) error {
	defer m.close()
	return m.c.UpdateId(bson.ObjectIdHex(id), args)
}

// UpdateByID - 更新資料
func (m *mgodb) UpdateAll(selector interface{}, args interface{}) (*mgo.ChangeInfo, error) {
	defer m.close()
	return m.c.UpdateAll(selector, args)
}

// Upsert
func (m *mgodb) Upsert(selector interface{}, args interface{}) (*mgo.ChangeInfo, error) {
	defer m.close()
	return m.c.Upsert(selector, args)
}

// FindByID - 查詢資料
func (m *mgodb) FindByID(id string, value interface{}) error {
	defer m.close()
	return m.c.FindId(bson.ObjectIdHex(id)).One(value)
}

// MgoQuery -
func (m *mgodb) Find(selector interface{}, value interface{}) error {
	defer m.close()
	return m.c.Find(selector).One(value)
}

// FindAll - find all document from the specific condition
func (m *mgodb) FindAll(selector interface{}, value interface{}) error {
	defer m.close()
	return m.c.Find(selector).All(value)
}

// FindAll - find all document from the specific condition
func (m *mgodb) FindAllBySort(selector interface{}, value interface{}, sort ...string) error {
	defer m.close()
	return m.c.Find(selector).Sort(sort...).All(value)
}

// FindSkipLimit -{}
func (m *mgodb) FindSkipLimit(selector, value interface{}, skip, limit int, sort ...string) error {
	defer m.close()
	return m.c.Find(selector).Skip(skip).Limit(limit).Sort(sort...).All(value)
}

// Delete -
func (m *mgodb) Delete(selector interface{}) error {
	defer m.close()
	return m.c.Remove(selector)
}

// DeleteAll - Delete all the document from the specific condition
func (m *mgodb) DeleteAll(selector interface{}) (*mgo.ChangeInfo, error) {
	defer m.close()
	return m.c.RemoveAll(selector)
}

// DeleteByID - Delete the document by ID
func (m *mgodb) DeleteByID(id string) error {
	defer m.close()
	return m.c.RemoveId(bson.ObjectIdHex(id))
}

func (m *mgodb) Count(selector interface{}) (int, error) {
	defer m.close()
	return m.c.Find(selector).Count()
}

func (m *mgodb) close() {
	m.s.Close()
}
