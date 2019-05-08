package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func RecvFile(fileName string, conn net.Conn) {
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println("os.Create err: ", err)
		return
	}

	buf := make([]byte, 1024*4)
	// 接收文件内容，接收多少，写多少
	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("文件接收完毕")
			} else {
				fmt.Println("conn.Read err: ", err)
			}
			return
		}
		f.Write(buf[:n]) // 往文件写入内容

	}
}

func main() {
	// 监听
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("net.Listen err: ", err)
		return
	}

	defer listener.Close()

	// 阻塞，等待用户连接
	conn, err1 := listener.Accept()
	if err1 != nil {
		fmt.Println("listener.Accept err: ", err1)
		return
	}

	defer conn.Close()

	var n int
	buf := make([]byte, 1024) // 读取对方发送的文件名
	n, err = conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read err: ", err)
		return
	}

	fileName := string(buf[:n]) // 读取文件名

	// 回复"ok"
	conn.Write([]byte("ok"))

	// 读取文件内容
	RecvFile(fileName, conn)

}
