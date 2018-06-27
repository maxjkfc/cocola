package db

import (
	"github.com/maxjkfc/cocola/errors"
)

var (
	ErrorConnectFailed = errors.New(201, "Connect Failed.")
	ErrorPoolNotFound  = errors.New(210, "Not Find the DB in the Pool With the Tag Name")
)
