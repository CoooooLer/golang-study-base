package main

import "fmt"

func main() {
	var a [10]int
	var b [5]int
	fmt.Printf("len(a) = %d, len(b) = %d\n", len(a), len(b))

	// var n int
	// n = 10
	// var c [n]int // error 指定数组的个数必须是常量，不能是变量

	// 数组元素从0开始到 len()-1 结束，不对称元素，这个数字叫下标
	a[0] = 1
	i := 1
	a[i] = 2 // a[1] = 2

	// 给数组每个元素赋值
	for i := 0; i < len(a); i++ {
		a[i] = i + 100
	}

	// 打印数组每个元素的值
	for i, data := range a {
		fmt.Printf("a[%d] = %d\n", i, data)
	}

}
