package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(3 * time.Second)
	isOk := timer.Reset(1 * time.Second) // 重新设定时间为1秒
	fmt.Println(isOk)

	go func() {
		<-timer.C // 延时
		fmt.Println("子协程开始执行")
	}()

	//timer.Stop() // 停止定时器

	for {

	}

}
