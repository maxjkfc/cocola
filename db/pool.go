package db

import (
	"github.com/maxjkfc/cocola/errors"
)

var Pools Pool

type Pool interface {
	NewConnect(Config) errors.Error
	Mgo(string) (Mgo, errors.Error)
}

type pool struct {
	num     int
	mgo     map[string]Mgo
	mariadb map[string]interface{}
	redis   map[string]interface{}
}

func init() {
	Pools = &pool{
		mgo: make(map[string]Mgo),
	}
}

func (p *pool) NewConnect(config Config) errors.Error {
	switch config.DBtype {
	case Mongo:
		m, err := mongoDial(config)
		if err != nil {
			return err
		}
		p.mgo[config.Tag] = m
	case Mariadb, Mysql:
	case Redis:

	}
	return nil
}

func (p *pool) Mgo(name string) (Mgo, errors.Error) {
	if x, ok := p.mgo[name]; !ok {
		return nil, ErrorPoolNotFound
	} else {
		return x, nil
	}
}
