package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 设置种子，只需一次
	rand.Seed(time.Now().UnixNano())

	var a [10]int
	n := len(a)
	for i := 0; i < n; i++ {
		a[i] = rand.Intn(100)
		fmt.Printf("a[%d] = %d\n", i, a[i])
	}

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}

	fmt.Printf("\n冒泡排序，升序\n")

	for i := 0; i < n; i++ {
		fmt.Println(a[i])
	}

	fmt.Println(a)

}
