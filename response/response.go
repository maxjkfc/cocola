package response

import (
	"time"

	"github.com/maxjkfc/cocola/errors"
)

type Response struct {
	Data   interface{} `json:"data"`
	Status Status      `json:"status"`
}

type Status struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Time string `json:"time"`
	Unix int64  `json:unix`
}

func Error(err errors.Error) Response {
	return response(nil, err)
}

func Resp(data interface{}) Response {
	return response(data, errors.NotError)

}

func response(data interface{}, err errors.Error) Response {

	r := &Response{
		Data: data,
	}

	r.Status.Code = err.GetC()
	r.Status.Msg = err.Error()

	t := time.Now()

	r.Status.Time = t.Format(time.RFC3339)
	r.Status.Unix = t.Unix()

	return r
}
