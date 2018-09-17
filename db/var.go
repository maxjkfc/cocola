package db

import "github.com/globalsign/mgo"

const (
	Mongo_Driver   = "mongo"
	Mariadb_Driver = "mariadb"
	Mysql_Driver   = "mysql"
	Redis_Driver   = "redis"
	Mongo_Protocol = "mongodb"
)

func NotFound(err error) bool {
	if err == mgo.ErrNotFound {
		return true
	}

	return false
}
