package main

import (
	"fmt"
	"net/http"
)

const (
	port = 8080
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte(fmt.Sprintf("hello: %d", port)))
	})
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
