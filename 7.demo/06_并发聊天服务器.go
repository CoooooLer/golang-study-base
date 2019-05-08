package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

type Client struct {
	C    chan string // 用户发送数据的管道
	Name string      // 用户民
	Addr string      // 网络地址
}

var message = make(chan string)

// 保存在线用户   cliAddr ==> Client
var onlineMap map[string]Client

// 新开一个协程用于转发消息。只要有消息来了，遍历map，给每个map成员都发送此消息
func Manager() {
	// 给map分配空间
	onlineMap = make(map[string]Client)

	for {
		msg := <-message // 没有消息前，这里会阻塞

		// 遍历map，给map每个成员发送消息
		for _, cli := range onlineMap {
			cli.C <- msg
		}
	}
}

// 新建一个协程，专门给客户端发送消息
func WriteMsgToClient(cli Client, conn net.Conn) { // 注意类型
	for msg := range cli.C {
		conn.Write([]byte(msg + "\n"))
	}

}

// 拼接广播的消息
func MakeMsg(cli Client, msg string) (buf string) {
	buf = "[" + cli.Addr + "]" + cli.Name + ": " + msg
	return
}

// 处理用户连接
func HandleConn(conn net.Conn) {
	defer conn.Close()

	// 获取客户端的网络地址
	cliAddr := conn.RemoteAddr().String()

	// 创建一个结构体，默认用户名和网络地址一样
	cli := Client{make(chan string), cliAddr, cliAddr}

	// 把结构体添加到onlineMap
	onlineMap[cliAddr] = cli

	// 新建一个协程，专门给客户端发送消息
	go WriteMsgToClient(cli, conn)

	// 广播某个用户上线
	//message <- "[" + cli.Addr + "]" + cli.Name + ": login"
	message <- MakeMsg(cli, "login")

	//提示,我是谁
	cli.C <- MakeMsg(cli, "i am here")

	// 用户退出的通道
	isQuit := make(chan bool)

	// 判断用户是否超时
	hasData := make(chan bool)

	// 新建一个协程，接收用户发过来的信息
	go func() {
		buf := make([]byte, 2048)
		for {
			n, err := conn.Read(buf)
			if n == 0 { // 对方断开，或者出问题
				isQuit <- true
				// fmt.Println("isQuit = ", <-isQuit)
				// time.Sleep(time.Second)
				fmt.Println("conn.Read err: ", err)
				return
			}
			msg := string(buf[:n-1]) // windows nc 测试，多一个换行符

			// 通过输入 who ,显示所有在线用户
			if len(msg) == 3 && msg == "who" {
				conn.Write([]byte("user list: \n"))
				for _, tmp := range onlineMap {
					msg = tmp.Addr + ":" + tmp.Name + "\n"
					conn.Write([]byte(msg))
				}
				// 通过输入 rename | newName , 修改用户名
			} else if len(msg) >= 8 && msg[:6] == "rename" { // rename | goooo
				name := strings.Split(msg, "|")[1] // 截取用户输入的新用户名
				cli.Name = name
				onlineMap[cliAddr] = cli
				conn.Write([]byte("rename successful \n"))
			} else {
				// 转发内容
				message <- MakeMsg(cli, msg)
				hasData <- true
			}
		}
	}()

	for {
		// 通过select检测channel的流动
		select {
		case <-isQuit:
			delete(onlineMap, cliAddr)           // 从map删除退出的用户
			message <- MakeMsg(cli, "logout \n") // 广播用户退出
			return
		case <-hasData:

		case <-time.After(time.Second * 20):
			delete(onlineMap, cliAddr)
			message <- MakeMsg(cli, "time out leave \n")
			return
		}
	}
}

func main() {
	// 监听
	listener, err := net.Listen("tcp", "0.0.0.0:8000")
	if err != nil {
		fmt.Println("net.Listen err: ", err)
		return
	}

	defer listener.Close()

	// 新开一个协程用于转发消息。只要有消息来了，遍历map，给每个map成员都发送此消息
	go Manager()

	// 主协程，循环阻塞等待用户连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept err: ", err)
			continue
		}
		go HandleConn(conn) // 处理用户连接
	}

}
