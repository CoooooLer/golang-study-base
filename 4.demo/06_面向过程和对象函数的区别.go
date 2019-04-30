package main

import "fmt"

func main() {
	result := add01(1, 2)
	fmt.Println("result = ", result)

	var a long = 2
	r := a.Add02(3) // 调用方法格式：变量名.函数(所需参数)
	fmt.Println("r = ", r)
}

// 面向过程
func add01(a int, b int) int {
	return a + b
}

// 面向对象，方法：给某个类型绑定一个函数
type long int

// tmp 叫接收者，接收者就是传递的一个参数
func (tmp long) Add02(other long) long {
	return tmp + other
}
