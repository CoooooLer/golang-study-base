package main

import "fmt"

func main() {
	// 数组做函数参数，它是值传递
	// 实参数组的每个元素都给形参数组拷贝一份
	// 形参数组是实参数组的复制品
	a := [5]int{1, 2, 3, 4, 5}
	Modify(a)
	fmt.Println("a = ", a)

}

func Modify(a [5]int) {
	a[0] = 666
	fmt.Println("a = ", a)
}
