package main

import (
	"fmt"
	"time"
)

func main() {
	// 创建一个长度为3有缓存的channel
	ch := make(chan int, 3)

	fmt.Printf("len(ch) = %d, cap(ch) = %d\n", len(ch), cap(ch))

	go func() {
		for i := 0; i < 10; i++ { // 超过定义的通道长度便会阻塞
			ch <- i // 只有channel里面的数据读完才能往里面写
			fmt.Printf("子协程[%d]: len(ch) = %d, cap(ch) = %d\n", i, len(ch), cap(ch))
		}
	}()

	time.Sleep(2 * time.Second)

	for i := 0; i < 10; i++ {
		num := <-ch
		fmt.Printf("num = %d\n", num)
	}

}
