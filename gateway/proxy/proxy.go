package proxy

import (
	goway_context "github.com/Zelayan/goway/gateway/context"
	"github.com/Zelayan/goway/gateway/response"
	"github.com/Zelayan/goway/gateway/router"
	"go.uber.org/zap"
	"net/http"
	"net/http/httputil"
)

type GoWayProxy struct {
}

// Dispatch all request
func (g *GoWayProxy) Dispatch(writer http.ResponseWriter, request *http.Request) {

	ctx := goway_context.NewGoWayContext(writer, request)
	ctx.AddFilter(g.Filter)
	defer func() {
		if err := recover(); err != nil {
			zap.L().Error("dispatch failed", zap.Any("err", err))
			g.globalRecover(ctx, err)
			return
		}
	}()
	ctx.Next()
}

func (g *GoWayProxy) Filter(ctx *goway_context.GoWayContext) {
	match, _, err := router.Match(ctx.Request.URL.Path)
	if err != nil {
		zap.L().Error("match router failed", zap.Error(err))
		ctx.ErrorHandler(http.StatusBadGateway, response.ProxyError, err)
		return
	}
	proxy := httputil.ReverseProxy{
		Director: func(req *http.Request) {
			req.Header.Add("X-Forwarded-Host", req.Host)
			req.Header.Add("X-Origin-Host", match.Host)
			req.URL.Scheme = match.Scheme
			req.URL.Host = match.Host
			req.URL.Path = match.Path
		},
	}
	proxy.ServeHTTP(ctx.ResponseWriter, ctx.Request)
}

func (g *GoWayProxy) globalRecover(ctx *goway_context.GoWayContext, errMsg interface{}) {
	if ctx.ResponseWriter != nil {
		ctx.ResponseWriter.Header().Set("Content-Type", "application/json;charset=UTF-8")
		ctx.ResponseWriter.WriteHeader(http.StatusBadGateway)
		ctx.ResponseWriter.Write([]byte(response.NewError(response.ProxyError, errMsg).Error()))
	}
}

func NewGoWayProxy() *GoWayProxy {
	return &GoWayProxy{}
}
