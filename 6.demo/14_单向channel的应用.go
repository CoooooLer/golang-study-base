package main

import (
	"fmt"
)

func main() {
	// 创建一个双向无缓存的channel
	ch := make(chan int)

	// 生产者，生产数字，写入channel，新开一个goroutine
	go producter(ch)
	// 消费者，从channel读取数据并打印
	consumer(ch)

}

// 只能写，不能读
func producter(write chan<- int) {
	for i := 0; i < 10; i++ {
		write <- i
	}
	close(write)
}

// 只能读，不能写
func consumer(read <-chan int) {
	for num := range read {
		fmt.Println("num = ", num)
	}
}
