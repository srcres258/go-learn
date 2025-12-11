package main

import (
	"net"
	"fmt"
	"io"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:11451")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	for {
		var byteData = make([]byte, 1024)
		n, err := conn.Read(byteData)
		if err == io.EOF {
			break
		}
		fmt.Println(string(byteData[:n]))
	}
}

