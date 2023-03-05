package main

import (
	"github.com/Zelayan/goway/gateway/go_way"
	_ "net/http/pprof"
)

func main() {
	goWay := go_way.NewServer()
	goWay.Start()
}
