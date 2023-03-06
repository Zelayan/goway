package goway_context

import (
	"github.com/Zelayan/goway/gateway/response"
	"math"
	"net/http"
)

const abortIndex int = math.MaxInt >> 1

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

func (ctx *GoWayContext) Next() {
	ctx.index++
	s := len(ctx.filters)
	for ; ctx.index < s; ctx.index++ {
		ctx.filters[ctx.index](ctx)
	}
}

func (ctx *GoWayContext) Abort() {
	ctx.index = abortIndex
}

func (ctx *GoWayContext) AddFilter(filter func(ctx *GoWayContext)) {
	ctx.filters = append(ctx.filters, filter)
}

func (ctx *GoWayContext) AbortWithStatus(httpCode int, errorCode int, err error) {
	ctx.ErrorHandler(httpCode, errorCode, err)
	ctx.Abort()
}

func (ctx *GoWayContext) ErrorHandler(httpCode int, errorCode int, err error) {
	if ctx.ResponseWriter != nil {
		ctx.ResponseWriter.Header().Set("Content-Type", "application/json;charset=UTF-8")
		ctx.ResponseWriter.WriteHeader(httpCode)
		ctx.ResponseWriter.Write([]byte(response.NewError(errorCode, err.Error()).Error()))
	}
}
