package proxy

import (
	"github.com/Zelayan/goway/gateway/router"
	"net/http"
	"net/http/httputil"
)

type GoWayProxy struct {
}

// Dispatch all request
func (g *GoWayProxy) Dispatch(writer http.ResponseWriter, request *http.Request) {

	match, _, err := router.Match(request.URL.Path)
	if err != nil {

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
	proxy.ServeHTTP(writer, request)
}

func NewGoWayProxy() *GoWayProxy {
	return &GoWayProxy{}
}
