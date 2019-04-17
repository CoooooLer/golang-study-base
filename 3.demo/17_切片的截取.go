package main

import "fmt"

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// s1 := array
	// fmt.Println("s1 = ", s1)
	// fmt.Printf("len = %d, cap = %d\n", len(s1), len(s1))

	// low:high:max 取下标从 low 开始的元素，len = high-low, cap = max - low
	s2 := array[:] // 不指定长度和容量，取所有
	fmt.Println("s2 = ", s2)
	fmt.Printf("len = %d, cap = %d\n", len(s2), cap(s2))

	// 操作某个元素，和操作数组一样
	data := array[3]
	fmt.Println("data = ", data)

	// array[1], array[2], array[3]  len = 4-1 = 3, cap = 5-1 = 4
	s3 := array[1:4:5]
	fmt.Println("s3 = ", s3)
	fmt.Printf("len = %d, cap = %d\n", len(s3), cap(s3))

	s4 := array[:3] // 从 0 开始，取3个元素，容量？
	fmt.Println("s4 = ", s4)
	fmt.Printf("len = %d, cap = %d\n", len(s4), cap(s4))

	s5 := array[3:] // 从 3 开始，到末尾
	fmt.Println("s5 = ", s5)
	fmt.Printf("len = %d, cap = %d\n", len(s5), cap(s5))
}
