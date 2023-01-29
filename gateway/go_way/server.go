package go_way

import (
	"github.com/Zelayan/goway/gateway/proxy"
	"net/http"
)

type Server struct {
}

func (s *Server) Start() {
	wayProxy := proxy.NewGoWayProxy()

	http.HandleFunc("/", wayProxy.Dispatch)
	err := http.ListenAndServe(":8888", nil)

	if err != nil {
		panic(err)
	}
}

func NewServer() *Server {
	return &Server{}
}
