package goway_context

import (
	"net/http"
)

type Filter func(c *GoWayContext)

var globalFilters []Filter

type GoWayContext struct {
	ResponseWriter http.ResponseWriter
	StatusCode     int
	Request        *http.Request
	// middleware
	filters []Filter
	index   int
}

func Use(f Filter) {
	globalFilters = append(globalFilters, f)
}

func NewGoWayContext(response http.ResponseWriter, request *http.Request) *GoWayContext {
	return &GoWayContext{
		ResponseWriter: response,
		Request:        request,
		index:          -1,
		filters:        globalFilters,
	}
}

func (c *GoWayContext) Next() {
	c.index++
	s := len(c.filters)
	for ; c.index < s; c.index++ {
		c.filters[c.index](c)
	}
}

func (c *GoWayContext) AddFilter(filter func(ctx *GoWayContext)) {
	c.filters = append(c.filters, filter)
}
