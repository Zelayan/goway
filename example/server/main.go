package main

import (
	"flag"
	"fmt"
	"net/http"
)

func main() {
	port := "1"
	flag.StringVar(&port, "port", "", "")
	flag.Parse()
	server(port)
}

func server(port string) {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte(fmt.Sprintf("hello: %s", port)))
	})
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
