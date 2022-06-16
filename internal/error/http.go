package error

import (
	"fmt"
	"github.com/go-kratos/kratos/v2/errors"
)

type HTTPError struct {
	Code    int               `json:"ret_code"`
	MapData map[string]string `json:"ret_data"`
	Message string            `json:"ret_msg"`
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("ret_code: %d, message: %s", e.Code, e.Message)
}

func NewHTTPError(code int, mapData map[string]string, msg string) *HTTPError {
	return &HTTPError{
		Code:    code,
		MapData: mapData,
		Message: msg,
	}
}

// FromHTTPError 把 err 转换成自定义的 HTTPError
func FromHTTPError(err error) *HTTPError {
	if nil == err {
		return nil
	}
	// 判断是不是自定义的 HTTPError
	if t1err := new(HTTPError); errors.As(err, &t1err) {
		return t1err
	}
	// 判断是不是 kratos/v2/errors
	if t1err := new(errors.Error); errors.As(err, &t1err) {
		mapData := t1err.Metadata
		mapData["error-reason"] = t1err.Reason
		return NewHTTPError(int(t1err.Code), mapData, t1err.Message)
	}
	// 都不是就返回 500
	return NewHTTPError(500, map[string]string{}, "internal.error.http")
}
