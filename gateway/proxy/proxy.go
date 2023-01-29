package proxy

import (
	"net/http"
	"net/http/httputil"
)

type GoWayProxy struct {
}

// Dispatch all request
func (g *GoWayProxy) Dispatch(writer http.ResponseWriter, request *http.Request) {
	remoteUrl := "http://localhost:7070/hello/test"
	proxy := httputil.ReverseProxy{
		Director: func(req *http.Request) {
			req.Header.Add("X-Forwarded-Host", req.Host)
			req.Header.Add("X-Origin-Host", remoteUrl)
			req.URL.Scheme = "http"
			req.URL.Host = "localhost:7070"
		},
		Transport:      nil,
		FlushInterval:  0,
		ErrorLog:       nil,
		BufferPool:     nil,
		ModifyResponse: nil,
		ErrorHandler:   nil,
	}
	proxy.ServeHTTP(writer, request)
}

func NewGoWayProxy() *GoWayProxy {
	return &GoWayProxy{}
}
