package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)    // 数字通信
	quit := make(chan bool) // 程序是否结束

	// 消费者，从channel读取内容
	go func() { // 子协程一定要写在主协程的前面
		for i := 0; i < 8; i++ {
			num := <-ch
			fmt.Println(num)
		}
		quit <- true
	}()

	// 生产者，生产数字并写入到channel
	fibonacci(ch, quit)

}

// ch只写，quit只读
func fibonacci(ch chan<- int, quit <-chan bool) {
	x, y := 1, 1
	for {
		select { // 监听数据流动
		// 每个 case 语句都是一个IO操作
		case ch <- x:
			x, y = y, x+y
		case flag := <-quit:
			fmt.Println("flag = ", flag)
			return
		}
	}

}
