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

	go func() {
		 for {
			 var buf []byte = make([]byte, 1024)
			 n, err := conn.Read(buf)
			 // 服务端退出
			 if err == io.EOF {
				 break
			 }
			 fmt.Println("收到服务器回复:", string(buf[0:n]))
		 }
	}()

	for {
		var text string
		fmt.Print("请输入你要发送的内容:")
		fmt.Scanln(&text)
		if text == "q" {
			conn.Close()
			break
		}
		conn.Write([]byte(text))
	}
}

