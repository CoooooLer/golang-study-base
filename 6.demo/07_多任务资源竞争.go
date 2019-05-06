package main

import (
	"fmt"
	"time"
)

func main() {

	// 新建两个协程，代表两个人，同时使用打印机
	go Person1()
	go Person2()

	for { // 定义一个死循环主协程  主协程一定要写在子协程后面

	}

}

func Person1() {
	Printer("hello")

}

func Person2() {
	Printer("world")

}

// 定义一个打印机，参数为字符串，打印每一个字符
// 打印机属于公共资源
func Printer(str string) {
	for _, data := range str {
		// fmt.Print(data)
		fmt.Printf("%c", data)
		time.Sleep(time.Second)
	}
	fmt.Printf("\n")
}
