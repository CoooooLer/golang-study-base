package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		for {
			fmt.Print("child goroutine")
			time.Sleep(time.Second)
		}
	}() // 调用匿名函数
}
