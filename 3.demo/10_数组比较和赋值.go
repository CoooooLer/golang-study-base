package main

import "fmt"

func main() {
	// 数组比较，只支持 == 或 ！= ，比较是不是每一个元素都一样，两个数组的类型必须一样
	a := [5]int{1, 2, 3, 4, 5}
	b := [5]int{1, 2, 3, 4, 5}
	c := [5]int{1, 2, 3}
	fmt.Println("a == b", a == b)
	fmt.Println("a == c", a == c)

	// 同类型的数组可以赋值
	d := a
	fmt.Println("d = ", d)
}
