package db

import (
	"strings"
	"time"

	"github.com/globalsign/mgo"
	"github.com/maxjkfc/cocola/errors"
)

type Mgo interface {
	New() *mgo.Session
}

type MgoDB interface {
	DB(string) *mgo.Database
	C(string) *mgo.Collection
	Close()
}

type mgodb struct {
	s   *mgo.Session
	db  *mgo.Database
	c   *mgo.Collection
	err error
}

type mgodriver struct {
	session *mgo.Session
	db      *mgo.Database
	c       *mgo.DialInfo
	err     error
}

func mongoDial(config Config) (Mgo, errors.Error) {
	return new(mgodriver).dial(config)
}

func (m *mgodriver) dial(config Config) (Mgo, errors.Error) {

	m.c = &mgo.DialInfo{
		Addrs:    strings.Split(config.Host, ","),
		Username: config.Account,
		Password: config.Password,
		Database: config.Database,
	}

	if config.ConnectTimeOut < 0 {
		m.c.Timeout = 20 * time.Second
	} else {
		m.c.Timeout = time.Duration(config.ConnectTimeOut) * time.Second
	}

	m.session, m.err = mgo.DialWithInfo(m.c)
	if m.err != nil {
		return nil, ErrorConnectFailed
	}

	return m, nil
}

func (m *mgodriver) New() *mgo.Session {
	return m.session.Copy()
}

func (m *mgodriver) Status() {

}
