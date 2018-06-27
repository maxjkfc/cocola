package db

import "time"

type Config struct {
	DBtype         string
	Account        string
	Password       string
	Database       string
	Host           string
	Protocol       string
	ConnectTimeOut int
	Tag            string
}

type Status struct {
	Ping        bool      `json:"ping"`
	Username    string    `json:"username"`
	Host        string    `json:"host"`
	ConnectTime time.Time `json:"connect_time"`
	UpdateTime  time.Time `json:"update_time"`
	PoolSize    int       `json:"pool_size"`
	Connecting  int       `json:"connecting"`
	DBType      string    `json:"db_type"`
	Tag         string    `json:"tag"`
}
