package main

import "github.com/Zelayan/goway/gateway/go_way"

func main() {
	goWay := go_way.NewServer()
	goWay.Start()
}
