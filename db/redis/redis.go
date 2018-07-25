package redis

import (
	"strconv"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/maxjkfc/cocola/db/config"
	"github.com/maxjkfc/cocola/errors"
)

type Redis interface {
	New() redis.Conn
	Status() string
	Close()
}

type redisdriver struct {
	p    *redis.Pool
	s    config.Status
	c    config.Config
	err  error
	errs errors.Error
}

func Dial(c config.Config) (Redis, errors.Error) {
	return new(redisdriver).dial(c)
}

func (r *redisdriver) New() redis.Conn {
	return r.p.Get()
}

func (r *redisdriver) dial(c config.Config) (Redis, errors.Error) {

	opt, err := r.setConfig(c)
	if err != nil {
		return nil, err
	}

	r.setStatus()

	r.p = &redis.Pool{
		MaxIdle:     MaxRedisIdleSize,
		MaxActive:   MaxRedisActiveSize,
		IdleTimeout: MaxRedisIdleTimeout,
		Dial: func() (redis.Conn, error) {
			return redis.Dial(TCP, r.c.Host, opt.([]redis.DialOption)...)
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) (err error) {
			_, err = c.Do(Redis_PING)
			return err
		},
	}

	if err := r.p.TestOnBorrow(r.p.Get(), time.Now()); err != nil {
		return nil, errors.ErrorConnectFailed
	} else {
		return r, nil
	}
}

func (r *redisdriver) setConfig(c config.Config) (interface{}, errors.Error) {

	c.Format()
	r.c = c

	// new the redis dial option
	option := make([]redis.DialOption, 0, 0)

	// set the ConnectTimeOut , if ConnectTimeOut < 0 , default = 20
	if c.ConnectTimeOut <= 0 {
		r.c.ConnectTimeOut = ConnectTimeOut
	}

	option = append(option, redis.DialConnectTimeout(time.Duration(r.c.ConnectTimeOut)*time.Second))

	// set the password
	if r.c.Password != "" {
		option = append(option, redis.DialPassword(r.c.Password))
	}

	// set the database

	if r.c.Database == "" {
		r.c.Database = "0"
	}

	if dbn, err := strconv.Atoi(r.c.Database); err == nil && dbn <= MaxRedisDatabase {
		option = append(option, redis.DialDatabase(dbn))
	} else {
		return nil, errors.ErrorConfigSet
	}

	return option, nil
}

func (r *redisdriver) setStatus() {
	r.s.Set(r.c)
}

func (r *redisdriver) Status() string {
	r.s.GetTotalTime()
	r.s.Connecting = r.p.ActiveCount()
	if r.p.TestOnBorrow(r.p.Get(), time.Now()) != nil {
		r.s.Ping = false
		r.s.Msg = StatusFailed
	} else {
		r.s.Ping = true
		r.s.Msg = StatusSuccess
	}
	return r.s.Json()
}

func (r *redisdriver) Close() {
	r.p.Close()
}
