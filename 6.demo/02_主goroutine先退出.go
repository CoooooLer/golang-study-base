package main

import (
	"fmt"
	"time"
)

// 主协程退出，子协程也退出
func main() {

	go func() {
		i := 0
		for {
			i++
			fmt.Println("子协程: i = ", i)
			time.Sleep(time.Second)
		}
	}()

	i := 0
	for {
		i++
		fmt.Println("main i = ", i)
		time.Sleep(time.Second)
		if i > 3 {
			break
		}
	}
}
