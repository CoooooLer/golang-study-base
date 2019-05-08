package main

import (
	"fmt"
	"net"
)

func main() {
	// 监听
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	defer listener.Close()

	// 阻塞等待用户链接
	conn, err1 := listener.Accept()
	if err1 != nil {
		fmt.Println("err1 = ", err1)
		return
	}

	// 接收用户的请求
	buf := make([]byte, 1024) // 1024大小的缓冲区
	n, err2 := conn.Read(buf)
	if err2 != nil {
		fmt.Println("err2 = ", err2)
		return
	}
	fmt.Println("buf = ", string(buf[:n])) // 切记转换为字符串

	defer conn.Close()
}
