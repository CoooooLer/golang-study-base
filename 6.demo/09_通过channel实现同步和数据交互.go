package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	defer fmt.Println("主协程即将结束")

	go func() {
		defer fmt.Println("子协程调用完毕")

		for i := 0; i < 2; i++ {
			fmt.Printf("子协程: i = %d\n", i)
			time.Sleep(time.Second)
		}
		ch <- "6666"
	}()

	str := <-ch // 通道没有数据阻塞
	fmt.Println("str = ", str)

}
