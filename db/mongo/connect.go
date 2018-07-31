package mongo

import (
	"strings"
	"time"

	"github.com/globalsign/mgo"
	"github.com/maxjkfc/cocola/db/config"
	"github.com/maxjkfc/cocola/errors"
)

type Mgo interface {
	New() *mgo.Session
	Status() string
	Close()
}

type mgodriver struct {
	session *mgo.Session
	c       config.Config
	s       config.Status
	err     error
	errs    errors.Error
}

func init() {
	mgo.SetStats(true)
}

// mongoDial - Mongo Driver Connect
func Dial(c config.Config) (Mgo, errors.Error) {
	return new(mgodriver).dial(c)
}

// New - new the mongo client instance
func (m *mgodriver) New() *mgo.Session {
	return m.session.Copy()
}

// dial - connect the mongo server
func (m *mgodriver) dial(c config.Config) (Mgo, errors.Error) {
	dconfig, err := m.setConfig(c)
	if err != nil {
		return nil, err
	}

	m.setStatus()
	m.session, m.err = mgo.DialWithInfo(dconfig.(*mgo.DialInfo))
	if m.err != nil {
		return nil, errors.ErrorConnectFailed
	}

	return m, nil
}

func (m *mgodriver) setConfig(c config.Config) (interface{}, errors.Error) {
	c.Format()
	m.c = c

	mconfig := &mgo.DialInfo{
		Addrs:    strings.Split(m.c.Host, ","),
		Username: c.Account,
		Password: c.Password,
		Database: c.Database,
		AppName:  c.Tag,
	}

	if m.c.ConnectTimeOut <= 0 {
		m.c.ConnectTimeOut = ConnectTimeOut
	}

	mconfig.Timeout = (time.Duration(m.c.ConnectTimeOut) * time.Second)

	return mconfig, nil

}

// setStats - use the set the Status
func (m *mgodriver) setStatus() {
	m.s.Set(m.c)
}

// Status - Query the mongo client and server status
func (m *mgodriver) Status() string {
	if err := m.session.Ping(); err != nil {
		m.s.Ping = false
		m.s.Msg = err.Error()
	} else {
		m.s.Ping = true
		m.s.Msg = StatusSuccess
	}

	//
	snapshot := mgo.GetStats()
	m.s.Connecting = snapshot.SocketsAlive
	m.s.GetTotalTime()

	return m.s.Json()
}

func (m *mgodriver) Close() {
	m.session.Close()
}
