package errcode

import (
	"fmt"
	"net/http"
)

type Error struct {
	Code    int      `json:"code"`
	Msg     string   `json:"msxg"`
	Details []string `json:"details"`
}

var codes = map[int]string{}

func NewError(code int, msg string) *Error {
	if _, ok := codes[code]; ok {
		s := fmt.Sprintf("错误码%d已经存在，请更换一个", code)
		panic(any(s))
	}

	codes[code] = msg
	return &Error{Code: code, Msg: msg}
}

func (e *Error) GetCode() int {
	return e.Code
}

func (e *Error) GetMsg() string {
	return e.Msg
}

func (e *Error) GetDetail() []string {
	return e.Details
}

func (e *Error) WithDetails(details ...string) *Error {
	newError := *e

	newError.Details = []string{}
	for _, d := range details {
		newError.Details = append(newError.Details, d)
	}
	return &newError
}

func (e *Error) StatusCode() int {
	switch e.Code {
	case Success.Code:
		return http.StatusOK
	case ServiceError.Code:
		return http.StatusInternalServerError
	}
	return http.StatusInternalServerError
}
