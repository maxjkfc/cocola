package errors

import (
	"encoding/json"
	"errors"
	"fmt"
	"sort"
)

var (
	errlist       *errorList
	ErrorNotExist = errors.New("Error Code Not Exist")
	ErrorExist    = errors.New("Error Code have Exist")
)

type errorList struct {
	errs map[int]Error // error list
	list []int         // error code list
}

type err struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
}

type Error interface {
	Error() string
	Json() string
}

func init() {
	if errlist == nil {
		errlist = &errorList{
			errs: make(map[int]Error),
			list: make([]int, 0),
		}
	}
}

func New(code int, msg string) Error {

	if _, ok := errlist.errs[code]; ok {
		panic(ErrorExist.Error())
	}

	errlist.errs[code] = &err{
		Code: code,
		Msg:  msg,
	}
	errlist.list = append(errlist.list, code)

	return errlist.errs[code]
}

func Err(code int) Error {
	if x, ok := errlist.errs[code]; ok {
		return x
	} else {
		panic(ErrorNotExist.Error())
	}

}

func List() {
	sort.Ints(errlist.list)

	for _, v := range errlist.list {
		fmt.Printf("ErrorCode: %d \t ErrorMsg: %s\n", v, errlist.errs[v])
	}
}

func Keys() []int {
	sort.Ints(errlist.list)
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
