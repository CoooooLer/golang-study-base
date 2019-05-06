package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(2 * time.Second)
	<-timer.C
	fmt.Println("时间到")

	time.Sleep(2 * time.Second)
	fmt.Println("时间到02")

	<-time.After(2 * time.Second) // 定时两秒，阻塞两秒，两秒后产生一个事件，往channel里面写内容
	fmt.Println("时间到03")
}
