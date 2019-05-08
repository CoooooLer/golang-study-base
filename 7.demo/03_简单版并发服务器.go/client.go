package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// 主动连接服务器
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("net.listener err: ", err)
		return
	}

	// main调用完毕,关闭连接
	defer conn.Close()

	go func() {
		// 获取从键盘输入的内容，发送给服务器
		str := make([]byte, 1024)
		for {
			n, err := os.Stdin.Read(str) // 获取键盘输入的内容，存到str
			if err != nil {
				fmt.Println("os.Stdin err: ", err)
				return
			}

			// 把输入的内容发送给服务器
			conn.Write(str[:n])
		}

	}()

	// 接收服务器返回的数据
	// 切片缓冲
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf) // 接收服务器的请求
		if err != nil {
			fmt.Println("conn.Read err: ", err)
			return
		}
		fmt.Println(string(buf[:n])) // 打印接收到的内容，转换为字符串再打印
		// fmt.Println("len(buf[:n])", len(string(buf[:n])))
	}
}
