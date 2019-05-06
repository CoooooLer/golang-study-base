package main

import (
	"fmt"
	"time"
)

// 验证 time.NewTimer() 时间到了只会执行一次
func main() {
	timer := time.NewTimer(time.Second)

	for {
		<-timer.C
		fmt.Println("时间到")
	}
}

func main01() {
	// 创建一个定时器，设置时间为2s,两秒后往time通道写入当前时间
	timer := time.NewTimer(2 * time.Second)
	fmt.Println("当前时间为: ", time.Now())

	// 2秒后，往timer.C写数据，有数据后，就可以读取
	t := <-timer.C // channel没有数据前后阻塞
	fmt.Println("t = ", t)
}
