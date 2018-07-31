package db

import (
	"github.com/maxjkfc/cocola/db/config"
	"github.com/maxjkfc/cocola/db/mongo"
	"github.com/maxjkfc/cocola/db/redis"
	"github.com/maxjkfc/cocola/errors"
)

var p *pool

type pool struct {
	mgosession mongo.MgoSession
	mariadb    map[string]interface{}
	redis      map[string]redis.Redis
	list       []Mod
}

type Mod struct {
	DType string
	Tag   string
}

func init() {
	p = &pool{
		redis: make(map[string]redis.Redis, 0),
		list:  make([]Mod, 0),
	}
}

func NewConnect(c config.Config) errors.Error {

	switch c.DBtype {

	case Mongo_Driver:
		m, err := mongo.Dial(c)
		if err != nil {
			return err
		}
		p.mgosession = m

	case Mariadb_Driver, Mysql_Driver:
		return nil

	case Redis_Driver:
		r, err := redis.Dial(c)
		if err != nil {
			return err
		}
		p.redis[c.Tag] = r
	default:
		return errors.ErrorDataBaseType
	}

	p.list = append(p.list, Mod{
		DType: c.DBtype,
		Tag:   c.Tag,
	})

	return nil
}

func Mgo() mongo.MgoSession {
	return p.mgosession
}

func Redis(name string) (redis.Redis, errors.Error) {
	if x, ok := p.redis[name]; !ok {
		return nil, errors.ErrorPoolNotFound
	} else {
		return x, nil
	}
}

func List() []Mod {
	return p.list
}

func Status() {
	p.mgosession.Status()
}
