package main

import (
	"fmt"
	"time"
)

func main() {
	// 创建一个无缓冲区的channel
	ch := make(chan int, 0)

	// len(ch) 缓冲区剩余数据个数， cap(ch) 缓冲区大小
	fmt.Printf("len(ch) = %d, cap(ch) = %d\n", len(ch), cap(ch))

	// 新建一个协程
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println("子协程： i = ", i)
			ch <- i // 往通道写内容，如果写入的数据没有被读取，便会阻塞

			// fmt.Printf("len(ch) = %d, cap(ch) = %d\n", len(ch), cap(ch))
		}
	}()

	time.Sleep(2 * time.Second)

	for i := 0; i < 5; i++ {
		num := <-ch // 读取同道中的内容，如果没有内容就阻塞
		fmt.Println("num = ", num)
	}

}
