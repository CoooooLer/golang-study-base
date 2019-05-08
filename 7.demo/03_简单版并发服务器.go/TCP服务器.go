package main

import (
	"fmt"
	"net"
	"strings"
)

func HandleConn(conn net.Conn) {
	// 函数调用完毕，自动关闭conn
	defer conn.Close()

	// 获取客户端的网络地址信息
	addr := conn.RemoteAddr().String()
	fmt.Println(addr, " connect successful")

	buf := make([]byte, 2048)

	// 读取用户数据
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read err: ", err)
			return
		}

		fmt.Println("len = ", len(string(buf[:n])))

		// 输出用户的地址
		fmt.Printf("[%s]: %s\n", addr, string(buf[:n]))
		// if "exit" == string(buf[:n-1]) { // nc测试
		if "exit" == string(buf[:n-2]) { // client 测试
			fmt.Println(addr, " exit")
			return
		}

		// 把数据转换为大写，再发送给用户
		_, err1 := conn.Write([]byte(strings.ToUpper(string(buf[:n]))))
		if err1 != nil {
			println("write err: ", err.Error())
			return
		}
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

	// 接收多个用户
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept err: ", err)
			return
		}

		// 处理用户请求，每一个用户新建一个协程
		go HandleConn(conn)
	}

}
