package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	quit := make(chan bool)

	// 新建一个子协程
	go func() {
		for {
			select {
			// 每个 case 语句都是一个IO操作
			case num := <-ch:
				fmt.Println("num = ", num)
			case <-time.After(3 * time.Second):
				fmt.Println("超时")
				quit <- true // 往通道写入数据
			}
		}
	}()

	for i := 0; i < 5; i++ {
		ch <- i
		time.Sleep(time.Second)
	}

	<-quit // 读取通道的数据，如果没有数据，就阻塞
	fmt.Println("程序结束")
}
