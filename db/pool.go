package db

import (
	"github.com/maxjkfc/cocola/db/config"
	"github.com/maxjkfc/cocola/db/mongo"
	"github.com/maxjkfc/cocola/db/redis"
	"github.com/maxjkfc/cocola/errors"
)

var Pools Pool

type Pool interface {
	NewConnect(config.Config) errors.Error
	Mgo(string) (mongo.Mgo, errors.Error)
	Redis(string) (redis.Redis, errors.Error)
	Status()
}

type pool struct {
	mgo     map[string]mongo.Mgo
	mariadb map[string]interface{}
	redis   map[string]redis.Redis
}

func init() {
	Pools = &pool{
		mgo: make(map[string]mongo.Mgo),
	}
}

func (p *pool) NewConnect(c config.Config) errors.Error {

	switch c.DBtype {

	case Mongo_Driver:
		m, err := mongo.Dial(c)
		if err != nil {
			return err
		}
		p.mgo[c.Tag] = m

	case Mariadb_Driver, Mysql_Driver:
		return nil

	case Redis_Driver:
		r, err := redis.Dial(c)
		if err != nil {
			return err
		}
		p.redis[c.Tag] = r
	}
	return nil
}

func (p *pool) Mgo(name string) (mongo.Mgo, errors.Error) {
	if x, ok := p.mgo[name]; !ok {
		return nil, errors.ErrorPoolNotFound
	} else {
		return x, nil
	}
}

func (p *pool) Redis(name string) (redis.Redis, errors.Error) {
	if x, ok := p.redis[name]; !ok {
		return nil, errors.ErrorPoolNotFound
	} else {
		return x, nil
	}
}

func (p *pool) Status() {
	for _, v := range p.mgo {
		v.Status()
	}

}
