package main

import (
	"fmt"
	"time"
)

// 创建一个全局变量channel
var ch = make(chan int)

func main() {
	// 两个人同时调用打印机
	go Person1()
	go Person2()

	for { // 死循环的额主协程

	}
}

// Person1执行完成后，才会执行Person2
func Person1() {
	Printer("hello")
	ch <- 666 // 给通道写入数据
}

func Person2() {
	<-ch // 读取通道的数据，如果通道没有数据，它就会阻塞
	Printer("world")
}

// 打印单个字符的打印机
func Printer(str string) {
	for _, data := range str {
		fmt.Printf("%c", data)
		time.Sleep(time.Second)
	}
	fmt.Printf("\n")
}
