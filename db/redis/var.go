package redis

import "time"

const (
	TCP                 = "tcp"
	ConnectTimeOut      = 60
	MaxRedisDatabase    = 16
	MaxRedisIdleSize    = 100
	MaxRedisActiveSize  = 0
	MaxRedisIdleTimeout = 60 * time.Second
	StatusSuccess       = "Success"
	StatusFailed        = "Failed"
)

// Redis Command Const
const (
	Redis_PING = "PING"
)
