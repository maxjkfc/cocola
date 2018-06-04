package errors

import (
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
	msg  string
	code int
}

type Error interface {
	Error() string
	Code() int
}

func init() {
	if errlist == nil {
		errlist = &errorList{
			errs: make(map[int]Error),
			list: make([]int, 0),
		}
	}
}

func New(code int, msg string) {

	if _, ok := errlist.errs[code]; ok {
		panic(ErrorExist.Error())
	}

	errlist.errs[code] = &err{
		code: code,
		msg:  msg,
	}
	errlist.list = append(errlist.list, code)
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
	return e.msg
}

func (e *err) Code() int {
	return e.code
}
