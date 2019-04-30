package main

import (
	"fmt"
	"time"
)

func main() {
	// 在函数调用前添加 go 关键字，创建并发执行单元
	go test() // 新建一个携程，新建一个任务

	for {
		fmt.Println("this is main goroutine")
		time.Sleep(time.Second) // 延时执行1秒
	}

}

func test() {
	for {
		fmt.Println("this is test goroutine")
		time.Sleep(time.Second)
	}
}
