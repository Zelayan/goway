package response

import (
	"encoding/json"
	"github.com/pkg/errors"
)

const (
	CreateHttpRequestFailed = 1000
	ParseHttpResponseFailed = 1001
	ProxyUrlNotFound        = 1004
	ProxyError              = 1502
	LimitTimeOutError       = 1504
)

type RestResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func NewError(code int, data interface{}) error {
	ret := &RestResponse{
		code,
		StatusText(code),
		data,
	}
	v, _ := json.Marshal(ret)
	return errors.New(string(v))
}
func StatusText(code int) string {
	return respStatus[code]
}

var respStatus = map[int]string{
	CreateHttpRequestFailed: "create HTTP Request failed",
	ParseHttpResponseFailed: "parse HTTP Response Body failed",
	ProxyUrlNotFound:        "API Not Found",
	ProxyError:              "Bad Gateway, Failed to request the background service ",
	LimitTimeOutError:       "API rate limit, timeout",
}
