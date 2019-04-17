package main

import (
	"fmt"
)

// map 存储的是键值对
func main() {
	// 定义一个变量，类型为 map[int]string
	var m1 map[int]string
	fmt.Println("m1 = ", m1)
	// 对于 map,只有 len, 没有 cap
	fmt.Println("len = ", len(m1))

	// 通过 make 创建
	m2 := make(map[int]string)
	fmt.Println("m2 = ", m2)
	fmt.Println("len = ", len(m2))

	// 通过 make 创建，可以指定长度，只是指定了容量，但是里面没有数据
	m3 := make(map[int]string, 3)
	m3[0] = "first" // 元素赋值
	m3[1] = "second"
	m3[2] = "third"
	fmt.Printf("m3 = %v\n", m3) // map 打印的值是无序的
	fmt.Println("len = ", len(m3))

	// 初始化，键值是唯一的
	m4 := map[string]string{"name": "小明", "number": "123456"}
	fmt.Println("m4 = ", m4)
	fmt.Println("len = ", len(m4))
}
