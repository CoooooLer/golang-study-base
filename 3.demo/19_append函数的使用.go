package main

import "fmt"

func main() {
	s := []int{}

	fmt.Println("s = ", s)
	fmt.Printf("len = %d, cap = %d\n", len(s), cap(s))

	// 在原切片的末尾追加元素
	s = append(s, 1)
	s = append(s, 2)
	s = append(s, 3)
	fmt.Println("s = ", s)
	fmt.Printf("len = %d, cap = %d\n", len(s), cap(s))

	s1 := []int{1, 1, 1}
	fmt.Println("s1 = ", s1)
	s1 = append(s1, 2)
	s1 = append(s1, 22)
	s1 = append(s1, 2222)
	s1 = append(s1, 22222)
	fmt.Println("s1 = ", s1)
	fmt.Printf("len = %d, cap = %d\n", len(s1), cap(s1))

}
