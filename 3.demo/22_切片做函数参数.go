package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 切片作为函数参数传入到函数时，在函数中改变的值就是切片的值
func main() {
	n := 10
	// s := []int{}
	// 使用 make 创建一个切片
	s := make([]int, n, n)

	InitData(s)
	fmt.Printf("初始化：%v\n", s)
	BubbleSort(s)
	fmt.Printf("排序后：%v\n", s)
}

// 产生随机数
func InitData(s []int) {
	n := len(s)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < n; i++ {
		s[i] = rand.Intn(100)
	}
}

// 冒泡排序
func BubbleSort(s []int) {
	n := len(s)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
}
