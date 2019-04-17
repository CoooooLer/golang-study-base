package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 设置种子，只需一次
	// 如果种子参数一样，每次运行程序产生的随机数都一样
	// rand.Seed(555)
	rand.Seed(time.Now().UnixNano()) //
	for i := 0; i < 5; i++ {

		// 产生随机数
		// fmt.Println("rand = ", rand.Int()) // 产水很大的随机数
		fmt.Println("rand = ", rand.Intn(10)) // 限制在10以内的数
	}
}
