package redis

import (
	"strconv"

	"github.com/gomodule/redigo/redis"
)

//Cmd - Redis cmds
type Cmd interface {
	// SELECT
	SELECT(int) Cmd
	// SET Key Value
	SET(key, value interface{}) Cmd
	// DELETE Key
	DELETE(key string) Cmd
	// GET Key
	GET(key string) interface{}
	// GETByInt64
	GETByInt64(key string) int64
	// GETSET Key
	GETSET(key, value string) interface{}
	// SETTEX  Key , value , ttl
	SETEX(key, value string, ttl int64) Cmd
	// EXPIRE key , ttl
	EXPIRE(key string, ttl int64) Cmd
	//
	EXISTS(key string) bool
	// TTL key
	TTL(key string) int64
	// HSET key , filed , value
	HSET(key, filed, value string) Cmd
	// HMSET key , kv
	HMSET(kv []interface{}) Cmd
	// HGET key ,filed
	HGET(key, filed string) interface{}
	//
	HGETALL(key string) map[string]string
	//
	HDEL(key, filed string) Cmd
	//
	HLEN(key string) int64
	//
	KEYS(pattern string) []string
	//
	SCAN(match string, count int) []string
	//
	INCR(key string) int64
	//
	DECR(key string) int64
	//
	INCRBy(key string, amount int64) int64
	//
	DECRBy(key string, amount int64) int64
	//
	Error() error
	//
	Close()
	//
	ToString(interface{}) string
}

type cmds struct {
	db     redis.Conn
	err    error
	result interface{}
}

// SET - Set the Key , value
func (r *cmds) SET(key, value interface{}) Cmd {
	_, r.err = r.db.Do("SET", key, value)
	return r
}

// DELETE - Delete the Key , value
func (r *cmds) DELETE(key string) Cmd {
	_, r.err = r.db.Do("DEL", key)
	return r
}

// GET - Get the Key , value
func (r *cmds) GET(key string) interface{} {
	r.result, r.err = r.db.Do("GET", key)
	return r.result
}

// GETByInt64 - Get the Key , value
func (r *cmds) GETByInt64(key string) int64 {
	var temp int64
	temp, r.err = redis.Int64(r.db.Do("GET", key))
	return temp
}

// GETSET - Get the Old Value , Set the New Value
func (r *cmds) GETSET(key, value string) interface{} {
	r.result, r.err = r.db.Do("GETSET", key, value)
	return r.result
}

// SETEX -
func (r *cmds) SETEX(key, value string, t int64) Cmd {
	r.result, r.err = r.db.Do("SETEX", key, t, value)
	return r
}

// EXPIRE - EXPIRE  the key life
func (r *cmds) EXPIRE(key string, ttl int64) Cmd {
	r.result, r.err = r.db.Do("EXPIRE", key, ttl)
	return r
}

// TTL - Query the key lifetime
func (r *cmds) TTL(key string) int64 {
	r.result, r.err = redis.Int64(r.db.Do("TTL", key))
	return r.result.(int64)
}

// EXISTS - Check the key
func (r *cmds) EXISTS(key string) bool {
	if r.result, r.err = r.db.Do("EXISTS", key); r.err != nil {
		return false
	} else if r.result.(int64) != 1 {
		return false
	}
	return true
}

// SELECT - switch the redis database
func (r *cmds) SELECT(to int) Cmd {
	_, r.err = r.db.Do("SELECT", to)
	return r
}

// HSET -
func (r *cmds) HSET(key, filed, value string) Cmd {
	r.result, r.err = r.db.Do("HSET", key, filed, value)
	return r
}

// HGET -
func (r *cmds) HGET(key, filed string) interface{} {
	r.result, r.err = r.db.Do("HGET", key, filed)
	return r.result
}

// HMSET -
func (r *cmds) HMSET(kv []interface{}) Cmd {
	r.result, r.err = r.db.Do("HMSET", kv...)
	return r
}

// HGETALL -
func (r *cmds) HGETALL(key string) map[string]string {
	var temp map[string]string
	temp, r.err = redis.StringMap(r.db.Do("HGETALL", key))
	return temp
}

// HDEL -
func (r *cmds) HDEL(key, filed string) Cmd {
	r.result, r.err = r.db.Do("HDEL", key, filed)
	return r
}

// HLEN -
func (r *cmds) HLEN(key string) (x int64) {
	x, r.err = redis.Int64(r.db.Do("HLEN", key))
	return x
}

func (r *cmds) SCAN(match string, count int) []string {
	var (
		result = make([]string, 0)
		loop   = true
	)

	if count < 1 {
		count = 10
	}

	args := []interface{}{
		0,
		"MATCH",
		match,
		"COUNT",
		count,
	}

	for loop {
		var ans []interface{}
		ans, r.err = redis.Values(r.db.Do("SCAN", args...))
		if r.err != nil {
			return nil
		}

		args[0], r.err = strconv.Atoi(string(ans[0].([]uint8)))
		if r.err != nil {
			return nil
		}

		if args[0] == 0 {
			loop = false
		}
		x, _ := redis.Strings(ans[1], nil)
		result = append(result, x...)
	}
	return result
}

func (r *cmds) KEYS(pattern string) []string {
	temp, err := redis.Strings(r.db.Do("KEYS", pattern))
	r.err = err
	return temp
}

func (r *cmds) INCR(key string) int64 {
	var x int64
	x, r.err = redis.Int64(r.db.Do("INCR", key))
	return x
}

func (r *cmds) DECR(key string) int64 {
	var x int64
	x, r.err = redis.Int64(r.db.Do("DECR", key))
	return x
}

func (r *cmds) INCRBy(key string, amount int64) int64 {
	var x int64
	x, r.err = redis.Int64(r.db.Do("INCRBY", key, amount))
	return x
}

func (r *cmds) DECRBy(key string, amount int64) int64 {
	var x int64
	x, r.err = redis.Int64(r.db.Do("DECRBY", key, amount))
	return x
}

// Close - close the db connect
func (r *cmds) Close() {
	r.db.Close()
}
func (r *cmds) Error() error {
	return r.err
}

func (r *cmds) Int64(arg interface{}) int64 {
	switch arg.(type) {
	case int64:
		return arg.(int64)
	case int:
		return int64(arg.(int))

	case string:
		var x int64
		x, r.err = strconv.ParseInt(arg.(string), 10, 64)
		return x
	case []uint8:
		var x int64
		x, r.err = strconv.ParseInt(string(arg.([]uint8)), 10, 64)
		return x
	default:
		return 0
	}
}

func (r *cmds) ToString(str interface{}) string {
	switch str.(type) {
	case []uint8:
		return string(str.([]uint8))
	case int64:
		return strconv.FormatInt(str.(int64), 64)
	case int:
		return strconv.Itoa(str.(int))
	case string:
		return str.(string)
	default:
		return ""
	}

}
