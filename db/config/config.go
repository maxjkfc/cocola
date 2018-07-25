package config

import (
	"encoding/json"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type Config struct {
	DBtype         string `json:"dbtype"`
	Account        string `json:"account"`
	Password       string `json:"password"`
	Database       string `json:"database"`
	Host           string `json:"host"`
	Protocol       string `json:"protocol"`
	ConnectTimeOut int64  `json:"connect_time_out"`
	Tag            string `json:"tag"`
}

// Format - Trime the space form string
func (c *Config) Format() {

	t := reflect.TypeOf(*c)
	v := reflect.ValueOf(c)

	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).Type.Kind() == reflect.String {

			v1 := v.Elem().Field(i)
			v1.SetString(strings.TrimSpace(v1.String()))
		}
	}
}

type Status struct {
	Ping        bool      `json:"ping"`
	Username    string    `json:"username"`
	Host        string    `json:"host"`
	Database    string    `json:"database"`
	ConnectTime time.Time `json:"connect_time"`
	TotalTime   string    `json:"total_time"`
	Connecting  int       `json:"connecting"`
	DBType      string    `json:"db_type"`
	Tag         string    `json:"tag"`
	Msg         string    `json:"msg"`
}

func (s *Status) GetTotalTime() {
	subTime := time.Now().Sub(s.ConnectTime)

	if subTime.Hours() >= 1 {
		s.TotalTime = strconv.FormatFloat(subTime.Hours(), 'f', 0, 64) + "h "
	}

	if subTime.Minutes() >= 1 {
		s.TotalTime = s.TotalTime + strconv.FormatFloat(subTime.Minutes(), 'f', 0, 64) + "m "
	}

	if subTime.Seconds() > 0 {
		s.TotalTime = s.TotalTime + strconv.FormatFloat(subTime.Seconds(), 'f', -1, 64) + "s "
	}

}

func (s *Status) Json() string {

	ans, err := json.Marshal(s)

	if err != nil {
		panic(err)
	}

	return string(ans)
}

func (s *Status) Set(c Config) {
	s.Tag = c.Tag
	s.Username = c.Account
	s.Host = c.Host
	s.ConnectTime = time.Now()
	s.DBType = c.DBtype
	s.Database = c.Database
}
