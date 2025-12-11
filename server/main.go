package main

import (
	"net"
	"fmt"
	"io"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:11451")

	listen, err := net.ListenTCP("tcp", addr)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("tcp server bind addr:", addr.String())
	for {
		// 等待连接
		conn, err := listen.AcceptTCP()
		if err != nil {
			fmt.Println("Accept err:", err)
			break
		}
		// 获取连接的客户端地址
		fmt.Println(conn.RemoteAddr().String(), " connect success")
		// 读取客户端发送的数据
		for {
			var buf []byte = make([]byte, 1024)
			n, err := conn.Read(buf)
			// 客户端退出
			if err == io.EOF {
				fmt.Println(conn.RemoteAddr().String(), " client exit")
				break
			}
			fmt.Println(string(buf[0:n]))
			conn.Write([]byte("你说的是:" + string(buf[0:n])))
		}
	}
}

