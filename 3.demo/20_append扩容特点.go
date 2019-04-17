package main

import "fmt"

func main() {
	// 切片如果超过原来的容量，通常以2倍容量扩容
	s := make([]int, 0, 0) // 用make函数声明一个切片
	// fmt.Println("s = ", s)
	oldCap := cap(s) // 计算容量
	for i := 0; i < 9; i++ {
		s = append(s, i)
		if newCap := cap(s); newCap > oldCap {
			fmt.Printf("newCap = %d\n", newCap)
			oldCap = newCap
		}
	}
}
