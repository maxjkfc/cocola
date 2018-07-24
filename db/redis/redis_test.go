package redis

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/maxjkfc/cocola/db/config"
)

var (
	r   Redis
	err error
)

func Test_RedisPool(t *testing.T) {
	c := config.Config{
		Account:  "",
		Password: "",
		Host:     "localhost:6379",
		Tag:      "Redis-Test",
		DBtype:   "Redis",
	}

	r, err = Dial(c)

	if err != nil {
		t.Fatal(err)
	}

	ans := r.Status()
	spew.Dump(ans)
}
