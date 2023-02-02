package go_way

import (
	"fmt"
	"github.com/Zelayan/goway/gateway/proxy"
	"github.com/Zelayan/goway/gateway/router"
	"go.uber.org/zap"
	"net/http"
)

type Server struct {
}

func (s *Server) Start() error {
	wayProxy := proxy.NewGoWayProxy()
	err := router.InitRouter()
	if err != nil {
		return fmt.Errorf("init router failed: %w", err)
	}
	http.HandleFunc("/", wayProxy.Dispatch)
	err = http.ListenAndServe(":8888", nil)
	if err != nil {
		return fmt.Errorf("http server start failed: %w", err)
	}
	return nil
}

func NewServer() *Server {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	defer func() {
		err = logger.Sync()
		if err != nil {
			fmt.Print(err)
		}
	}()

	undo := zap.ReplaceGlobals(logger)
	defer undo()
	return &Server{}
}
