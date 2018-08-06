package response

import (
	"time"

	"github.com/maxjkfc/cocola/errors"
)

type R struct {
	Data   interface{} `json:"data"`
	Status Status      `json:"status"`
}

type Status struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Time string `json:"time"`
	Unix int64  `json:"unix"`
}

func Error(err errors.Error) R {
	return response(nil, err)
}

func Resp(data interface{}) R {
	return response(data, errors.NotError)
}

func response(data interface{}, err errors.Error) R {
	t := time.Now()
	return R{
		Data: data,
		Status: Status{
			Code: err.GetC(),
			Msg:  err.Error(),
			Time: t.Format(time.RFC3339),
			Unix: t.Unix(),
		},
	}
}
