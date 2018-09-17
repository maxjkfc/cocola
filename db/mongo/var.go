package mongo

import "github.com/globalsign/mgo"

const (
	ConnectTimeOut = 60
	StatusSuccess  = "Success"
)

func NotFound(err error) bool {
	if err == mgo.ErrNotFound {
		return true
	}

	return false
}
