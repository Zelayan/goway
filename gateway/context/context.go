package goway_context

import "net/http"

type GoWayContext struct {
	ResponseWriter http.ResponseWriter
	Request        *http.Request

	index int
}

func NewGoWayContext(response http.ResponseWriter, request *http.Request) *GoWayContext {
	return &GoWayContext{
		ResponseWriter: response,
		Request:        request,
	}
}

func (c *GoWayContext) Next() {
	c.index++

}
