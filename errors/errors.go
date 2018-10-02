package errors

import (
	"encoding/json"
	"errors"

	"github.com/maxjkfc/cocola/log"
)

var (
	errlist       *errorList
	errorNotExist = errors.New("Error Code Not Exist")
	errorExist    = errors.New("Error Code have Exist")
)

type errorList struct {
	errs map[int]Error // error list
	list []*err        // error code list
}

type err struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
}

type Error interface {
	Error() string
	Json() string
	GetC() int
}

func new() {
	if errlist == nil {
		errlist = &errorList{
			errs: make(map[int]Error),
			list: make([]*err, 0),
		}
	}

}

// New - the New Error in list
func New(code int, msg string) Error {
	new()
	if _, ok := errlist.errs[code]; ok {
		log.Error("Same Error", errorExist)
		panic(errorExist.Error())
	}

	e := &err{
		Code: code,
		Msg:  msg,
	}

	errlist.errs[code] = e

	errlist.list = append(errlist.list, e)

	return errlist.errs[code]
}

// T - temp error
func T(code int, msg string) Error {
	return &err{
		Code: code,
		Msg:  msg,
	}
}

func List() []*err {
	return errlist.list
}

func (e *err) Error() string {
	return e.Msg
}

func (e *err) Json() string {
	x, err := json.Marshal(e)
	if err != nil {
		panic(err)
	}
	return string(x)
}

func (e *err) GetC() int {
	return e.Code
}
