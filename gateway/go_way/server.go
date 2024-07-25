package go_way

import (
	"fmt"
	goway_context "github.com/Zelayan/goway/gateway/context"
	"github.com/Zelayan/goway/gateway/limit"
	"github.com/Zelayan/goway/gateway/log"
	"github.com/Zelayan/goway/gateway/proxy"
	"github.com/Zelayan/goway/gateway/router"
	"go.uber.org/zap"
	"net/http"
	_ "net/http/pprof"
	"time"
)

type Server struct {
}

func (s *Server) Start() error {
	wayProxy := proxy.NewGoWayProxy()
	err := router.InitRouter()
	goway_context.Use(log.Logger())
	goway_context.Use(limit.RateLimit(time.Millisecond, 1))
	initPprof()
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

func initPprof() {
	go func() {
		http.ListenAndServe("0.0.0.0:8899", nil)
	}()
}
