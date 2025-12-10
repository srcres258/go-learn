package main

import (
	"net/http"
	"fmt"
	"io"
)

func main() {
	response, err := http.Get("http://127.0.0.1:11451")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	byteData, _ := io.ReadAll(response.Body)
	fmt.Println(string(byteData))
}

