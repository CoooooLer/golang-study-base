package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 0, 0}
	s := a[0:2:5] // 从下标 0 开始，到下标 2 结束，不包括下标 2 的值 （左闭右开）,
	fmt.Println("s = ", s)
	fmt.Println("len(s) = ", len(s)) // 数组的长度  2-0=2
	fmt.Println("cap(s) = ", cap(s)) // 数组的容量  5-0=5

	fmt.Println("===============================")

	s = a[1:3:5]
	fmt.Println("s = ", s)
	fmt.Println("len(s) = ", len(s))
	fmt.Println("cap(s) = ", cap(s))

}
