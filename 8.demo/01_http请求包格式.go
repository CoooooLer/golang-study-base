package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Println("net.Listen err: ", err)
		return
	}

	defer listener.Close()

	conn, err1 := listener.Accept()
	if err1 != nil {
		fmt.Println("listener.Accept err: ", err1)
		return
	}

	buf := make([]byte, 1024*4)
	n, err2 := conn.Read(buf)
	if n == 0 {
		fmt.Println("conn.Read err: ", err2)
		return
	}

	fmt.Printf("#%v#", string(buf[:n]))

	// 先运行此文件，再在浏览器输入 127.0.0.1:8000 测试

}
