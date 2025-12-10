package main

import (
	"net/http"
	"fmt"
)

func Index(writer http.ResponseWriter, request *http.Request) {
	fmt.Println(request.URL.Path, request.UserAgent())
	writer.Write([]byte("Hello World"))
}

func main() {
	http.HandleFunc("/", Index)
	var addr = "127.0.0.1:11451"
	fmt.Println("Starting server at", addr)
	http.ListenAndServe(addr, nil)
}

