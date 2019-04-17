package main

import "fmt"

func main() {
	// 切片和数组的区别
	// 数组 [] 里面的长度是一个固定的常量，数组不能修改长度，len 和 cap 都是 5
	a := [5]int{}
	fmt.Printf("a: len = %d, cap = %d\n", len(a), cap(a))

	// 切片， [] 里面为空，或者为 ... ,切片的长度和容量可以不固定
	b := []int{}
	fmt.Printf("b: len = %d, cap = %d\n", len(b), cap(b))

	b = append(b, 22) // 给切片 b 末尾追加一个成员 22
	fmt.Println("b = ", b)
	fmt.Printf("b: len = %d, cap = %d\n", len(b), cap(b))

	// 自动推导类型，同时初始化
	c1 := []int{1, 2, 3, 4, 5}
	fmt.Println("c1 = ", c1)

	// 借助 make 函数，格式： make(切片类型，长度，容量)
	c2 := make([]int, 5, 11)
	fmt.Printf("c2: len = %d, cap = %d\n", len(c2), cap(c2))

	// 如果不指定容量，容量和长度一样
	c3 := make([]int, 6)
	fmt.Printf("c3: len = %d, cap = %d\n", len(c3), cap(c3))

}
