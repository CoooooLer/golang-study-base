package main

import (
	"fmt"
)

func main() {
	ch := make(chan int) // 创建一个无缓存channel

	// 新建一个goroutine
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i // 往通道写数据
		}

		close(ch) // 不需要往通道写入数据时，关闭通道
		//ch <- 666 // 关闭channel后，无法再往通道写入数据，但是可以读取数据
	}()

	// 通过range实现遍历
	for num := range ch {
		fmt.Println("num = ", num)
	}

	// for {
	// 	// 如果ok为true，说明通道没有关闭
	// 	if num, ok := <-ch; ok == true {
	// 		fmt.Println("num = ", num)
	// 	} else { // 通道关闭
	// 		break
	// 	}
	// }
}
