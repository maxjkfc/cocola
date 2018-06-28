package db

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/globalsign/mgo"
	"github.com/maxjkfc/cocola/errors"
)

type Mgo interface {
	New() *mgo.Session
	Status() string
}

type MgoDB interface {
	DB(string) *mgo.Database
	C(string) *mgo.Collection
	Close()
}

type mgodriver struct {
	session *mgo.Session
	config  *mgo.DialInfo
	status  Status
	err     error
}

func init() {
	mgo.SetStats(true)

}

// mongoDial - Mongo Driver Connect
func mongoDial(config Config) (Mgo, errors.Error) {
	return new(mgodriver).dial(config)
}

// New - new the mongo client instance
func (m *mgodriver) New() *mgo.Session {
	return m.session.Copy()
}

// dial - connect the mongo server
func (m *mgodriver) dial(config Config) (Mgo, errors.Error) {

	m.setStats(config)

	m.session, m.err = mgo.DialWithInfo(m.config)
	if m.err != nil {
		return nil, ErrorConnectFailed
	}

	return m, nil
}

func (m *mgodriver) setStats(config Config) {

	m.config = &mgo.DialInfo{
		Addrs:    strings.Split(config.Host, ","),
		Username: config.Account,
		Password: config.Password,
		Database: config.Database,
	}

	if config.ConnectTimeOut < 0 {
		m.config.Timeout = 20 * time.Second
	} else {
		m.config.Timeout = time.Duration(config.ConnectTimeOut) * time.Second
	}
	//
	m.status.Tag = config.Tag
	m.status.Username = config.Account
	m.status.Host = config.Host
	m.status.ConnectTime = time.Now()
}

// Status - Query the mongo client and server status
func (m *mgodriver) Status() string {
	if err := m.session.Ping(); err != nil {
		m.status.Ping = false
		m.status.Msg = err.Error()
	} else {
		m.status.Ping = true
		m.status.Msg = "Sussess"
	}
	//
	snapshot := mgo.GetStats()
	spew.Dump(snapshot)

	//
	subTime := time.Now().Sub(m.status.ConnectTime)
	m.status.TotalTime = strconv.FormatFloat(subTime.Hours(), 'g', 4, 64)

	ans, _ := json.Marshal(m.status)

	return string(ans)
}
